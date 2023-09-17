from pathlib import Path
from sqlalchemy import text, TextClause
from contextlib import contextmanager
from python.shared.settings import sql_folder
from enum import Enum, auto
import time
from typing import Generator

class Task(Enum):
    task3 = auto()
    task4 = auto()
    task5 = auto()
    task6 = auto()
    task7 = auto()

def query(query_filename: str, task: Task, params: dict | None = None) -> TextClause:
    if params is None:
        params = dict()

    with open(str(sql_folder / task.name / "queries" / query_filename), "r") as file:
        schema_sql_code = file.read()
    
    return text(schema_sql_code).bindparams(**params)

@contextmanager
def measure_time(obj: str="uknown") -> Generator[None,None,None]:
    start = time.time()
    yield
    end = time.time()
    elapsed = end - start
    print(f"measured_time={elapsed} for {obj}")
