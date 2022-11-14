# import pytest
# from utils.database.sql import Database
# from ..task2.factories import TypeFactories
# from .test_3_2 import base_test

# class Params:
#     variables = "N"
#     amounts = (10,1000,5000,50000)

# @pytest.mark.performance
# @pytest.mark.parametrize(Params.variables, Params.amounts)
# def test_task3_2_performance_no_index(database: Database, apply_task2_migrations, factories: TypeFactories, N):
#     base_test(database, factories, N)

# @pytest.mark.performance
# @pytest.mark.parametrize(Params.variables, Params.amounts)
# def test_task3_2_performance_with_index(database: Database, apply_task3_migrations, factories: TypeFactories, N):
#     base_test(database, factories, N)