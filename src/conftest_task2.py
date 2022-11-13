import pytest
from utils.database.sql import Database
from .task2.factories import generate_factories, TypeFactories
from .task2.migrator import apply_migration, Migrations


@pytest.fixture
def load_task2_scheme(database: Database):
    apply_migration(database=database, path=Migrations.task2_1)

@pytest.fixture
def factories(database: Database):
    return generate_factories(database)