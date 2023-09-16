import pytest
from sqlalchemy_utils import database_exists, create_database, drop_database # type: ignore
from python.utils.database.sql import Database
from python.shared import settings
import secrets
from typing import Generator


@pytest.fixture()
def database() -> Generator[Database, None, None]:
    database = Database(
        url=settings.DATABASE_URL,
        name=f"test_database_{secrets.token_hex(10)}",
        debug=settings.DATABASE_DEBUG,
    )

    try:
        if not database_exists(database.full_url):
            create_database(database.full_url)
        yield database
    finally:
        if database_exists(database.full_url):
            drop_database(database.full_url)


@pytest.fixture
def session(database: Database):
    
    with database.get_core_session() as session:
        yield session
