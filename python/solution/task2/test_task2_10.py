from python.utils.database.sql import Database
from .factories import TypeFactories
import random


def test_generate_data(database: Database, apply_task2_migrations: None, factories: TypeFactories) -> None:
    class Consts:
        users_total_amount: int = 100
        users_per_post: int = 50
        posts_total_amount: int = users_total_amount * users_per_post

    users = factories.user.create_batch(
        (factories.user.template(id=i) for i in range(Consts.users_total_amount)))

    factories.post.create_batch(
        (
            factories.post.template(id=i,author_id=i % Consts.users_total_amount)
            for i in range(Consts.posts_total_amount)
        ))

    factories.post.create_one(factories.post.template(id=999999,author_id=random.choice(users).id))
    factories.post.create_batch((factories.post.template(id=999999+i,author_id=random.choice(users).id) for i in range(1,100)))