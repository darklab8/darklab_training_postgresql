import pytest
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .test_3_1 import base_test as base_test_3_1

@pytest.mark.performance
@pytest.mark.parametrize("amount_of_posts", (10,1000,5000,50000))
def test_task3_1_performance_no_index(database: Database, apply_task2_migrations, factories: TypeFactories, amount_of_posts):
    """
    measured_time=0.0008785724639892578 for 3_1, amount_of_posts=10
    .measured_time=0.0010502338409423828 for 3_1, amount_of_posts=1000
    .measured_time=0.002232789993286133 for 3_1, amount_of_posts=5000
    .measured_time=0.009729385375976562 for 3_1, amount_of_posts=50000
    """

    base_test_3_1(database=database, factories=factories, amount_of_posts=amount_of_posts)

# @pytest.mark.parametrize("amount_of_posts", (10,1000,5000,50000))
# def test_task3_1_performance_with_index(database: Database, apply_task2_migrations, factories: TypeFactories, amount_of_posts):
#     """
#     """

#     base_test_3_1(database=database, factories=factories, amount_of_posts=amount_of_posts)