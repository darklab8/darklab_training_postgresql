import pytest
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .test_3_2 import base_test

class Params:
    variables = "N"
    amounts = (5,)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_2_performance_no_index(database: Database, apply_task2_migrations, factories: TypeFactories, N):
    """
    measured_time=0.0009379386901855469 for 3_2, N=5
    .measured_time=0.0016803741455078125 for 3_2, N=500
    .measured_time=0.00477290153503418 for 3_2, N=2500
    .measured_time=0.01749396324157715 for 3_2, N=10000
    """
    base_test(database, factories, N)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_2_performance_with_index(database: Database, apply_task3_migrations, factories: TypeFactories, N):
    """
    measured_time=0.001138925552368164 for 3_2, N=5
    .measured_time=0.0015490055084228516 for 3_2, N=500
    .measured_time=0.007289886474609375 for 3_2, N=2500
    .measured_time=0.01748371124267578 for 3_2, N=10000
    """
    base_test(database, factories, N)