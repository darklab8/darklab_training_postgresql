from pathlib import Path
from sqlalchemy import text
from contextlib import contextmanager

def query(query_filename: str, params = None) -> text:
    if params is None:
        params = dict()

    with open(str(Path(__file__).parent / "queries" / query_filename), "r") as file:
        schema_sql_code = file.read()
    
    return text(schema_sql_code).bindparams(**params)

import time

@contextmanager
def measure_time(obj="uknown"):
    start = time.time()
    yield
    end = time.time()
    elapsed = end - start
    print(f"measured_time={elapsed} for {obj}")