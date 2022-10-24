from pathlib import Path
from sqlalchemy.ext.automap import automap_base
import pytest
from sqlalchemy.ext.automap import automap_base
import random
from dataclasses import dataclass
from utils.database.sql import Database


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

@pytest.fixture
def load_test2_scheme(database: Database):

    with open(str(Path(__file__).parent / "solution" / "code.sql"), "r") as file:
        schema_sql_code = file.read()

    with database.get_core_session() as session:
        session.execute(schema_sql_code)
        session.commit()


def test_generate_data(database: Database, load_test2_scheme):
    Base = automap_base()
    Base.prepare(database.engine, reflect=True)

    with database.get_core_session() as session:

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
                    tags=[random.choice(["abc", "def", "ghi"]), random.choice(["jkl", "mno", "pqr"])],
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