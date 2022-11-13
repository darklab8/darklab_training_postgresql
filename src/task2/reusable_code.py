from pathlib import Path
from utils.database.sql import Database

class Migrations:
    task2_1 = Path(__file__).parent / "migrations" / "task2_1.sql"

def apply_migration(database: Database, path: Path):
    with open(str(path), "r") as file:
        schema_sql_code = file.read()

    with database.get_core_session() as session:
        session.execute(schema_sql_code)
        session.commit()
