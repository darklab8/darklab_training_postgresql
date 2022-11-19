import pytest
from sqlalchemy_utils import database_exists, create_database, drop_database
from utils.database.sql import Database
from src.shared.databases import DatabaseFactory
from src.shared import settings
import secrets


@pytest.fixture()
def database():
    database = DatabaseFactory(
        url=settings.DATABASE_URL,
        name=f"test_database_{secrets.token_hex(10)}",
    )

    if not database_exists(database.full_url):
        create_database(database.full_url)

    yield database

    if database_exists(database.full_url):
        drop_database(database.full_url)


@pytest.fixture
def session(database: DatabaseFactory):
    with database.manager_to_get_session() as session:
        yield session
