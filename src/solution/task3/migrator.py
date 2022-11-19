from pathlib import Path

# order does not matter
migrations = (Path(__file__).parent / "migrations").glob("*.sql")
