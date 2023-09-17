from python.utils.database.sql import Database
from sqlalchemy import text

def test_check(database: Database) -> None:
    with database.get_core_session() as session:
        session.execute(text("SELECT 1;"))
