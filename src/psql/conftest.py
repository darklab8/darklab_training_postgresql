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

    with open("architecture/task_2_architecture_without_data.sql", "r") as file_:
        query = file_.read()

    results = run_raw(engine, query)


@pytest.fixture
def filled_db(inited_db, engine):
    Base = automap_base()
    Base.prepare(engine, reflect=True)

    User = Base.classes.users

    session = Session(engine)

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
            for i in range(100)
        ]
    )
    session.commit()
