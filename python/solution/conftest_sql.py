import pytest
from python.utils.database.sql import Database
from python.shared import settings
import secrets
from typing import Generator
import psycopg2
from contextlib import contextmanager
from typing import TYPE_CHECKING
if TYPE_CHECKING:
    from psycopg2._psycopg import _Cursor

@contextmanager
def raw_connection(full_url: str) -> Generator["_Cursor", None, None]:
    conn = psycopg2.connect(full_url)
    try:
        conn.set_session(autocommit=True)
        with conn.cursor() as cur:
            yield cur
    finally:
        conn.close()

@pytest.fixture()
def database() -> Generator[Database, None, None]:
    system_db = Database(
        url=settings.DATABASE_URL,
        name="postgres",
        debug=settings.DATABASE_DEBUG,
    )

    database = Database(
        url=settings.DATABASE_URL,
        name=f"test_database_{secrets.token_hex(10)}",
        debug=settings.DATABASE_DEBUG,
    )

    with raw_connection(system_db.full_url) as cur:
        try:
            cur.execute(f"CREATE DATABASE {database.name}")
            yield database
        finally:
            cur.execute(f"SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '{database.name}'")
            cur.execute(f"DROP DATABASE {database.name}")

@pytest.fixture
def session(database: Database):
    
    with database.get_core_session() as session:
        yield session
