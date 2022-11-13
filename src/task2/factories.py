"""
10. Написать скрипт для заполнения таблицы тестовыми данными.
В базе данных должно быть не менее 100000 постов для разных пользователей;
Для генерации тестовых данных можно воспользоваться функцией generate_series;

P.S. При создании 10000 юзеров с 50 постами каждого, скрипт выполняется за 14 секунд
"""
from sqlalchemy.ext.automap import automap_base
from utils.database.sql import Database
import random
import sys
from dataclasses import dataclass, field
from types import SimpleNamespace
from unittest.mock import MagicMock

def random_DATE():
    return f"20{random.randint(10,22):02}-{random.randint(1,12):02}-{random.randint(1,28):02}"

def random_int_generator():
    for i in range(sys.maxsize):
        yield i

rand = random_int_generator()

def rnd_int():
    return next(rand)

@dataclass
class UserTemplate:
    id: int = field(default_factory=rnd_int)
    first_name: str = field(default_factory=lambda: f"name_{rnd_int()}")
    second_name: str = field(default_factory=lambda: f"second_name_{rnd_int()}")
    birth_date: str = field(default_factory=random_DATE)
    email: str = field(default_factory=lambda: f"email_{rnd_int()}")
    password: str = field(default_factory=lambda: f"password_{rnd_int()}")
    address: str = field(default_factory=lambda: f"address_{rnd_int()}")
    rating: int = 0

@dataclass
class PostTemplateRaw:
    id: int = field(default_factory=rnd_int)
    author_id: int = field(default_factory=rnd_int)
    title: str = field(default_factory=lambda: f"title_{rnd_int()}")
    content: str = field(default_factory=lambda: f"content_{rnd_int()}")
    created_at: str = field(default_factory=random_DATE)
    status: str = field(default_factory=
            lambda: random.choice(["draft", "published", "archived"]))
    tags: str = field(default_factory=
            lambda: [random.choice(["abc", "def", "ghi"]), random.choice(["jkl", "mno", "pqr"])])
    rating: int = 0

class FactoryConveyor:
    def __init__(self, database: Database, db_model, template):
        self.database = database
        self.db_model = db_model
        self.template = template

    def create_one(self, template):
        with self.database.get_core_session() as session:
            session.add(
                self.db_model(
                    **(template.__dict__)
                )
            )
            session.commit()
            return template

    def create_batch(self, templates: list):
        with self.database.get_core_session() as session:
            session.bulk_save_objects(
                [
                    self.db_model(
                        **(template.__dict__)
                    ) for template in templates
                ]
            )
            session.commit()
            return templates

class TypeFactories:
    user = FactoryConveyor(MagicMock(),MagicMock(),template=UserTemplate)
    post = FactoryConveyor(MagicMock(),MagicMock(),template=PostTemplateRaw)

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

    return factories