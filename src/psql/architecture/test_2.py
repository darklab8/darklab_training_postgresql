from ..app.utils import run_raw
import os
from sqlalchemy.sql import text
from sqlalchemy.orm import Session


def test_check():
    assert True


def test_check_db_conn(engine):
    results = run_raw(
        engine,
        """
SELECT EXISTS (
SELECT FROM 
    pg_tables
WHERE 
    schemaname = 'public'
    );""",
    )


def test_automap_db(filled_db, engine):

    results = run_raw(engine, "SELECT * FROM users LIMIT 10")

    count = 0
    for row in results:
        print(row)
        count += 1

    assert count == 10
