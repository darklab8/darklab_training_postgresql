from pathlib import Path
from python.shared.settings import sql_folder

# order does not matter
migrations = (sql_folder / "task3" / "migrations").glob("*.sql")
