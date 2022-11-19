import pytest
from utils.database.sql import Database
from .task2.factories import generate_factories
from .task2.migrator import Migrations as Migrations2
from .task3.migrator import migrations as migrations3
from pathlib import Path
from sqlalchemy import text
from sqlalchemy.engine import create_engine


def apply_migration(database: Database, path: Path):
    with open(str(path), "r") as file:
        schema_sql_code = file.read()


    stmt = text(schema_sql_code)
    with database.get_core_connection() as conn:
        conn = conn.execution_options(isolation_level='AUTOCOMMIT')
        conn.execute(stmt)

            

@pytest.fixture
def apply_task2_migrations(database: Database):
    apply_migration(database=database, path=Migrations2.task_2_1)
    apply_migration(database=database, path=Migrations2.disable_triggers)


@pytest.fixture
def apply_task3_migrations(database: Database, apply_task2_migrations):
    for migration_path in migrations3:
        apply_migration(database=database, path=migration_path)

@pytest.fixture
def factories(database: Database):
    return generate_factories(database)





