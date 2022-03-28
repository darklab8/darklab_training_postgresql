from ..app.utils import run_query
import os
from sqlalchemy.sql import text
from sqlalchemy.orm import Session


def test_check():
    assert True


def test_check_db_conn(engine):
    results = run_query(
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

    results = run_query(engine, "SELECT * FROM users LIMIT 10")

    count = 0
    for row in results:
        print(row)
        count += 1

    assert count == 10


def test_check_tags(filled_db, engine):

    first_result = run_query(
        engine, "SELECT tags FROM posts LIMIT 1", return_first=True
    )

    assert len(first_result) == 2
    assert first_result[0] in ["abc", "def", "ghi"]


def test_users_with_rating_exist(filled_db, engine):

    results = run_query(engine, "SELECT * FROM users WHERE rating > 0")

    assert results.rowcount != 0
