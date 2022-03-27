from sqlalchemy import create_engine
from sqlalchemy.ext.automap import automap_base
from sqlalchemy.orm import Session
import random
import pytest
from .app.utils import run_raw


@pytest.fixture
def engine():
    engine = create_engine(
        "postgresql://postgres:postgres@localhost:5432/postgres", echo=True
    )

    return engine


@pytest.fixture
def inited_db(engine):

    with open("architecture/task_2_architecture.sql", "r") as file_:
        query = file_.read()

    results = run_raw(engine, query)


@pytest.fixture
def filled_db(inited_db, engine):
    Base = automap_base()
    Base.prepare(engine, reflect=True)

    session = Session(engine)

    users_amount: int = 1000
    posts_count: int = 2000

    User = Base.classes.users
    session.bulk_save_objects(
        [
            User(
                id=i,
                first_name=f"name_{i}",
                second_name=f"second_name_{i}",
                birth_date=f"20{random.randint(10,22):02}-{random.randint(1,12):02}-{random.randint(1,28):02}",
                email=f"email_{i}",
                password=f"password_{i}",
                address=f"address_{i}",
            )
            for i in range(users_amount)
        ]
    )

    Post = Base.classes.posts
    session.bulk_save_objects(
        [
            Post(
                id=i,
                author_id=i % users_amount,
                title=f"title_{i}",
                content=f"content_{i}",
                created_at=f"20{random.randint(10,22):02}-{random.randint(1,12):02}-{random.randint(1,28):02}",
                status=random.choice(["draft", "published", "archived"]),
            )
            for i in range(posts_count)
        ]
    )
    session.commit()
