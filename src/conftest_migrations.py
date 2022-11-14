import pytest
from utils.database.sql import Database
from .task2.factories import generate_factories
from .task2.migrator import Migrations as Migrations2
from .task3.migrator import Migrations as Migrations3
from pathlib import Path
from sqlalchemy import text
from sqlalchemy.engine import create_engine


def apply_migration(database: Database, path: Path):
    with open(str(path), "r") as file:
        schema_sql_code = file.read()

    engine = create_engine(
        database.full_url,
        pool_pre_ping=False,
        echo=True,
    )
    stmt = text(schema_sql_code)
    with engine.connect().execution_options(isolation_level='AUTOCOMMIT') as conn:
        conn.execute(stmt)

            

@pytest.fixture
def apply_task2_migrations(database: Database):
    apply_migration(database=database, path=Migrations2.task_2_1)

@pytest.fixture
def apply_task3_migrations(database: Database, apply_task2_migrations):
    apply_migration(database=database, path=Migrations3.task_3_7)

@pytest.fixture
def factories(database: Database):
    return generate_factories(database)





