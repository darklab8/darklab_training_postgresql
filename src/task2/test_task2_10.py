import pytest
from utils.database.sql import Database
from .factories import generate_factories, TypeFactories
from .migrator import apply_migration, Migrations


@pytest.fixture
def load_task2_scheme(database: Database):
    apply_migration(database=database, path=Migrations.task2_1)

@pytest.fixture
def factories(database: Database):
    return generate_factories(database)

def test_generate_data(database: Database, load_task2_scheme, factories: TypeFactories):
    class Consts:
        users_total_amount: int = 100
        users_per_post: int = 50
        posts_total_amount: int = users_total_amount * users_per_post

    factories.user.create_batch(
        [factories.user.template(id=i) for i in range(Consts.users_total_amount)])

    factories.post.create_batch(
        [
            factories.post.template(id=i,author_id=i % Consts.users_total_amount)
            for i in range(Consts.posts_total_amount)
        ])

    factories.post.create_one(factories.post.template())
    factories.post.create_batch([factories.post.template() for i in range(100)])