import pytest
from utils.database.sql import Database
from .task2.factories import generate_factories
from .task2.migrator import Migrations
from pathlib import Path


def apply_migration(database: Database, path: Path):
    with open(str(path), "r") as file:
        schema_sql_code = file.read()

    with database.get_core_session() as session:
        session.execute(schema_sql_code)
        session.commit()


@pytest.fixture
def apply_task2_migrations(database: Database):
    apply_migration(database=database, path=Migrations.task2_1)


@pytest.fixture
def factories(database: Database):
    return generate_factories(database)





