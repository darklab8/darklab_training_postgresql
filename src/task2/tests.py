from pathlib import Path
import pytest
from utils.database.sql import Database
from .script_fill_db import script_fill_db
from .reusable_code import apply_migration, Migrations


@pytest.fixture
def load_test2_scheme(database: Database):
    apply_migration(database=database, path=Migrations.task2_1)

def test_generate_data(database: Database, load_test2_scheme):
    script_fill_db(database)