from sqlalchemy import create_engine
from sqlalchemy.ext.automap import automap_base
from sqlalchemy.orm import Session
import random
import pytest
from .app.utils import run_query
from dataclasses import dataclass



@pytest.fixture(scope="session")
def engine():
    engine = create_engine(
        "postgresql://postgres:postgres@localhost:5432/postgres"  # , echo=True
    )

    return engine


@pytest.fixture(scope="session")
def inited_db(engine):

    with open("architecture/task_2_architecture.sql", "r") as file_:
        query = file_.read()

    results = run_query(engine, query)


def random_DATE():
    return f"20{random.randint(10,22):02}-{random.randint(1,12):02}-{random.randint(1,28):02}"


@dataclass
class database_generating_consts:
    users_total_amount: int = 1000
    users_per_post: int = 2
    posts_total_amount: int = int(users_total_amount / users_per_post)
    post_editions_per_user = 4
    post_editions_total_amount = users_total_amount * post_editions_per_user
    post_approvals_total_amount = users_total_amount


Consts = database_generating_consts()


def filled_db_initer(inited_db, engine):
    Base = automap_base()
    Base.prepare(engine, reflect=True)

    session = Session(engine)

    User = Base.classes.users
    session.bulk_save_objects(
        [
            User(
                id=i,
                first_name=f"name_{i}",
                second_name=f"second_name_{i}",
                birth_date=f"{random_DATE()}",
                email=f"email_{i}",
                password=f"password_{i}",
                address=f"address_{i}",
            )
            for i in range(Consts.users_total_amount)
        ]
    )

    Post = Base.classes.posts
    session.bulk_save_objects(
        [
            Post(
                id=i,
                author_id=i % Consts.users_total_amount,
                title=f"title_{i}",
                content=f"content_{i}",
                created_at=f"{random_DATE()}",
                status=random.choice(["draft", "published", "archived"]),
                tags=[random.choice(["abc", "def", "ghi"])],
            )
            for i in range(Consts.posts_total_amount)
        ]
    )

    PostEdit = Base.classes.post_editions
    session.bulk_save_objects(
        [
            PostEdit(
                id=i,
                user_id=i % Consts.users_total_amount,
                post_id=random.randint(0,i % Consts.posts_total_amount),
                edited_at=f"{random_DATE()}",
            )
            for i in range(Consts.post_editions_total_amount)
        ]
    )

    PostApproval = Base.classes.post_approvals
    session.bulk_save_objects(
        [
            PostApproval(
                id=i,
                user_id=i % Consts.users_total_amount,
                post_id=i % Consts.posts_total_amount,
                change=random.choice([1, 1, -1]),
                created_at=f"{random_DATE()}",
            )
            for i in range(Consts.post_approvals_total_amount)
        ]
    )

    PostVisit = Base.classes.post_visits_per_day
    session.bulk_save_objects(
        [
            PostVisit(
                id=i,
                post_id=i % Consts.posts_total_amount,
                visits=random.randint(1, 20),
                day_date=f"{random_DATE()}",
            )
            for i in range(Consts.posts_total_amount)
        ]
    )

    session.commit()

@pytest.fixture(scope="session")
def filled_db(inited_db, engine):
    filled_db_initer(inited_db, engine)


