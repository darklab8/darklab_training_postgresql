from pathlib import Path
from sqlalchemy import text

def query(query_filename: str, params = None) -> text:
    if params is None:
        params = dict()

    with open(str(Path(__file__).parent / "queries" / query_filename), "r") as file:
        schema_sql_code = file.read()
    
    return text(schema_sql_code).bindparams(**params)