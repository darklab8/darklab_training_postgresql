"""
10. Написать скрипт для заполнения таблицы тестовыми данными.
В базе данных должно быть не менее 100000 постов для разных пользователей;
Для генерации тестовых данных можно воспользоваться функцией generate_series;

P.S. При создании 10000 юзеров с 50 постами каждого, скрипт выполняется за 14 секунд
"""
from sqlalchemy.ext.automap import automap_base
from python.utils.database.sql import Database
import random
import sys
from dataclasses import dataclass, field
from types import SimpleNamespace
from random import randrange
from datetime import timedelta, datetime, date
from python.solution.task3.reusable_code import measure_time
from typing import Protocol, TypeVar, Generic, Type, Iterator
from typing import TYPE_CHECKING
if TYPE_CHECKING:
    from _typeshed import Incomplete

random_date_start = datetime.strptime("2020/01/01 16:30", "%Y/%m/%d %H:%M")
random_date_end = datetime.strptime("2022/01/01 16:30", "%Y/%m/%d %H:%M")

def random_date() -> datetime:
    """
    This function will return a random datetime between two datetime 
    objects.
    """
    delta = random_date_end - random_date_start
    int_delta = (delta.days * 24 * 60 * 60) + delta.seconds
    random_second = randrange(int_delta)
    return random_date_start + timedelta(seconds=random_second)

def random_int_generator() -> Iterator[int]:
    for i in range(sys.maxsize):
        yield i

rand = random_int_generator()

def rnd_int() -> int:
    return next(rand)

def increasing_date() -> datetime:
    "always increasing forward"
    return datetime.fromtimestamp(random_date_start.timestamp() + rnd_int())

@dataclass
class UserTemplate:
    id: int = field(default_factory=rnd_int)
    first_name: str = field(default_factory=lambda: f"name_{rnd_int()}")
    second_name: str = field(default_factory=lambda: f"second_name_{rnd_int()}")
    birth_date: datetime = field(default_factory=random_date)
    email: str = field(default_factory=lambda: f"email_{rnd_int()}")
    password: str = field(default_factory=lambda: f"password_{rnd_int()}")
    address: str = field(default_factory=lambda: f"address_{rnd_int()}")
    rating: int = 0

@dataclass
class PostTemplateRaw:
    id: int = field(default_factory=rnd_int)
    author_id: int = field(default_factory=rnd_int)
    status: str = field(default_factory=
            lambda: random.choice(["draft", "published", "archived"]))
    
    created_at: datetime = field(default_factory=random_date)
    rating: int = 0

@dataclass
class PostEditionTemplateRaw:
    id: int = field(default_factory=rnd_int)
    post_id: int = field(default_factory=rnd_int)
    user_id: int = field(default_factory=rnd_int)
    edited_at: datetime = field(default_factory=random_date)

    # new content
    title: str = field(default_factory=lambda: f"title_{rnd_int()}")
    content: str = field(default_factory=lambda: f"content_{rnd_int()}")
    tags: list[str] = field(default_factory=
            lambda: [random.choice(["abc", "def", "ghi"]), random.choice(["jkl", "mno", "pqr"])])


@dataclass
class PostApprovalTemplateRaw:
    id: int = field(default_factory=rnd_int)
    post_id: int = field(default_factory=rnd_int)
    user_id: int = field(default_factory=rnd_int)
    created_at: datetime = field(default_factory=random_date)
    change: int = field(default_factory=lambda: random.choice([-1, 1]))

@dataclass
class PostVisitsTemplateRaw:
    id: int = field(default_factory=rnd_int)
    post_id: int = field(default_factory=rnd_int)
    day_date: date = field(default_factory=random_date)
    visits: int = field(default_factory=lambda: random.randint(0,1000))

Template = TypeVar("Template")

class FactoryConveyor(Generic[Template]):
    def __init__(self, database: Database, db_model: "Incomplete", template: Type[Template]):
        self.database: Database = database
        self.db_model = db_model
        self.template = template

    def create_one(self, template: Template) -> Template:
        with self.database.get_core_session() as session:
            session.add(
                self.db_model(
                    **(template.__dict__)
                )
            )
            session.commit()
            return template

    def create_batch(self, templates: Iterator[Template]) -> list[Template]:
        list_templates = list(templates)
        with self.database.get_core_session() as session:
            session.bulk_save_objects(
                [
                    self.db_model(
                        **(template.__dict__)
                    ) for template in list_templates
                ]
            )
            session.commit()
        return list_templates

    def create_batch_in_chunks(self, templates: Iterator[Template]) -> Iterator[Template]:
        with self.database.get_core_session() as session:
            with measure_time(f"query"):
                try:
                    while True:
                        prepare_objects = []
                        for i in range(10000):
                            prepare_objects.append(self.db_model(**(next(templates).__dict__)))
                        session.bulk_save_objects(prepare_objects)
                except StopIteration:
                    session.bulk_save_objects(prepare_objects)
                session.commit()
            return templates

class TypeFactories(Protocol):
    user: FactoryConveyor[UserTemplate]
    post: FactoryConveyor[PostTemplateRaw]
    post_visits: FactoryConveyor[PostVisitsTemplateRaw]
    post_edition: FactoryConveyor[PostEditionTemplateRaw]
    post_approval: FactoryConveyor[PostApprovalTemplateRaw]

def generate_factories(database: Database) -> TypeFactories:
    Base = automap_base()
    Base.prepare(database.engine, reflect=True)
    factories = SimpleNamespace()

    factories.user = FactoryConveyor(
        database=database, db_model=Base.classes.user_, template=UserTemplate)

    @dataclass
    class PostTemplate(PostTemplateRaw):
        author_id: int = field(default_factory=
            lambda: factories.user.create_one(factories.user.template()).id)
    factories.post = FactoryConveyor(
        database=database, db_model=Base.classes.post, template=PostTemplate)

    @dataclass
    class PostEditionTemplate(PostEditionTemplateRaw):
        post_id: int = field(default_factory=
            lambda: factories.post.create_one(factories.post.template()).id)
        user_id: int = field(default_factory=
            lambda: factories.user.create_one(factories.user.template()).id)
    factories.post_edition = FactoryConveyor(
        database=database, db_model=Base.classes.post_edition, template=PostEditionTemplate)

    class PostApprovalTemplate(PostApprovalTemplateRaw):
        post_id: int = field(default_factory=
            lambda: factories.post.create_one(factories.post.template()).id)
        user_id: int = field(default_factory=
            lambda: factories.user.create_one(factories.user.template()).id)
    factories.post_approval = FactoryConveyor(
        database=database, db_model=Base.classes.post_approval, template=PostApprovalTemplate)

    factories.post_visits = FactoryConveyor(
        database=database, db_model=Base.classes.post_visits_per_day, template=PostVisitsTemplateRaw)

    return factories
