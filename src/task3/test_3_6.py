import pytest
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .test_3_1 import base_test as base_test_3_1

class Task3_1:
    variables = "amount_of_posts"
    amounts = (10,1000,5000,50000)

@pytest.mark.performance
@pytest.mark.parametrize(Task3_1.variables, Task3_1.amounts)
def test_task3_1_performance_no_index(database: Database, apply_task2_migrations, factories: TypeFactories, amount_of_posts):
    """
    measured_time=0.0008785724639892578 for 3_1, amount_of_posts=10
    .measured_time=0.0010502338409423828 for 3_1, amount_of_posts=1000
    .measured_time=0.002232789993286133 for 3_1, amount_of_posts=5000
    .measured_time=0.009729385375976562 for 3_1, amount_of_posts=50000
    """

    base_test_3_1(database=database, factories=factories, amount_of_posts=amount_of_posts)

@pytest.mark.performance
@pytest.mark.parametrize(Task3_1.variables, Task3_1.amounts)
def test_task3_1_performance_with_index(database: Database, apply_task3_migrations, factories: TypeFactories, amount_of_posts):
    """
    measured_time=0.0009191036224365234 for 3_1, amount_of_posts=10
    .measured_time=0.0012400150299072266 for 3_1, amount_of_posts=1000
    .measured_time=0.0017991065979003906 for 3_1, amount_of_posts=5000
    .measured_time=0.0068323612213134766 for 3_1, amount_of_posts=50000
    """

    base_test_3_1(database=database, factories=factories, amount_of_posts=amount_of_posts)

"""
Итого: Небольшое замедление от появление индексов. Вроде ожидаемый результат при count операции.
"""