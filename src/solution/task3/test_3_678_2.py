import pytest
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .test_3_2 import base_test

class Params:
    variables = "N"
    amounts = (5,500,2500,10000)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_2_performance_no_index(
        database: Database,
        apply_task2_migrations,
        factories: TypeFactories,
        N,
    ):
    """
    measured_time=0.0009379386901855469 for 3_2, N=5
    .measured_time=0.0016803741455078125 for 3_2, N=500
    .measured_time=0.00477290153503418 for 3_2, N=2500
    .measured_time=0.01749396324157715 for 3_2, N=10000
    """
    base_test(database, factories, N=10000)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_2_performance_with_index(
        database: Database,
        apply_task3_migrations,
        factories: TypeFactories,
        N,
    ):
    """
    measured_time=0.0008466243743896484 for 3_2, N=5
    .measured_time=0.0015785694122314453 for 3_2, N=500
    .measured_time=0.004488229751586914 for 3_2, N=2500
    .measured_time=0.015027046203613281 for 3_2, N=10000
    """
    base_test(database, factories, N)

"""
для N=10000
EXPLAIN ANALYZE VERBOSE

BEFORE INDEX:
"Limit  (cost=1943.77..1968.77 rows=10000 width=91) (actual time=10.850..12.197 rows=10000 loops=1)"
"  Output: id, author_id, title, content, tags, status, created_at, rating"
"  ->  Sort  (cost=1943.77..1993.77 rows=20000 width=91) (actual time=8.911..9.879 rows=10000 loops=1)"
"        Output: id, author_id, title, content, tags, status, created_at, rating"
"        Sort Key: post.created_at DESC"
"        Sort Method: quicksort  Memory: 3398kB"
"        ->  Seq Scan on public.post  (cost=0.00..515.00 rows=20000 width=91) (actual time=0.006..3.007 rows=20000 loops=1)"
"              Output: id, author_id, title, content, tags, status, created_at, rating"
"Planning Time: 0.141 ms"
"Execution Time: 12.568 ms"

AFTER INDEX:
post_created_at_idx is a btree
"Limit  (cost=0.29..930.26 rows=10000 width=91) (actual time=0.020..5.385 rows=10000 loops=1)"
"  Output: id, author_id, title, content, tags, status, created_at, rating"
"  ->  Index Scan Backward using post_created_at_idx on public.post  (cost=0.29..1860.24 rows=20000 width=91) (actual time=0.019..4.973 rows=10000 loops=1)"
"        Output: id, author_id, title, content, tags, status, created_at, rating"
"Planning Time: 0.300 ms"
"Execution Time: 5.636 ms"

CONCLUSION:
Improvement more than twice in plan :) Index is working.
C * Log2(N) * N becomes, C * N potentially? If Limit is Const and not an O(N).
"""