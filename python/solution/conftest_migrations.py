import pytest
from python.utils.database.sql import Database
from .task2.factories import generate_factories, TypeFactories
from .task2.migrator import Migrations as Migrations2
from .task3.migrator import create_indexes
from pathlib import Path
from sqlalchemy import text
from .conftest_sql import raw_connection


def get_sql_code(path: Path) -> str:
    with open(str(path), "r") as file:
        schema_sql_code = file.read()
        return schema_sql_code
    raise Exception("failed get_sql_code")

def apply_migration(database: Database, path: Path) -> None:
    stmt = text(get_sql_code(path))
    with database.get_core_connection() as conn:
        conn = conn.execution_options(isolation_level='AUTOCOMMIT')
        conn.execute(stmt)

            

@pytest.fixture
def apply_task2_migrations(database: Database) -> None:
    apply_migration(database=database, path=Migrations2.task_2_1)
    apply_migration(database=database, path=Migrations2.disable_triggers)


@pytest.fixture
def apply_task3_migrations(database: Database, apply_task2_migrations: None) -> None:
    with raw_connection(database.full_url) as cur:
        sql = get_sql_code(create_indexes)
        lines = sql.split("\n")

        for line in lines:
            
            if line.startswith("--"):
                continue
            if line == "":
                continue

            cur.execute(line)
        

@pytest.fixture
def factories(database: Database) -> TypeFactories:
    return generate_factories(database)





