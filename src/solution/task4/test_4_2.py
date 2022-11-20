
from utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from ..task3.reusable_code import query, measure_time


def test_task4_1(database: Database, apply_task2_migrations, factories: TypeFactories):
    "1. Найти N наиболее посещаемых постов за день/месяц/год.;"


