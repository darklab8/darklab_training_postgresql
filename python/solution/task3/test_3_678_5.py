import pytest
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .test_3_5 import base_test

class Params:
    variables = "N"
    # amounts = (500,1000,5000,50000)
    amounts = (50000,)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_5_performance_no_index(database: Database, apply_task2_migrations, factories: TypeFactories, N):
    base_test(database, factories, N)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_5_performance_with_index(database: Database, apply_task3_migrations, factories: TypeFactories, N):
    base_test(database, factories, N)

"""
"Limit  (cost=3231.61..3231.74 rows=50 width=16) (actual time=19.611..19.617 rows=50 loops=1)"
"  ->  Sort  (cost=3231.61..3238.43 rows=2725 width=16) (actual time=19.610..19.613 rows=50 loops=1)"
"        Sort Key: (COALESCE((pr.created_at)::timestamp with time zone, now())) DESC"
"        Sort Method: top-N heapsort  Memory: 28kB"
"        ->  Merge Join  (cost=1648.78..3141.09 rows=2725 width=16) (actual time=12.809..19.327 rows=2755 loops=1)"
"              Merge Cond: (post.id = pr.post_id)"
"              ->  Index Only Scan using post_pkey on post  (cost=0.29..1306.29 rows=50000 width=4) (actual time=0.008..3.785 rows=50000 loops=1)"
"                    Heap Fetches: 0"
"              ->  Sort  (cost=1648.49..1655.30 rows=2725 width=14) (actual time=12.796..13.034 rows=2755 loops=1)"
"                    Sort Key: pr.post_id"
"                    Sort Method: quicksort  Memory: 269kB"
"                    ->  Seq Scan on post_approval pr  (cost=0.00..1493.00 rows=2725 width=14) (actual time=0.008..12.228 rows=2755 loops=1)"
"                          Filter: ((created_at <= now()) AND (created_at >= (now() - '1 year'::interval)))"
"                          Rows Removed by Filter: 47245"
"Planning Time: 0.483 ms"
"Execution Time: 19.650 ms"

"""


"""
Without expresion indexes

"Limit  (cost=2521.04..2521.16 rows=50 width=16) (actual time=28.015..28.022 rows=50 loops=1)"
"  ->  Sort  (cost=2521.04..2528.16 rows=2848 width=16) (actual time=28.014..28.018 rows=50 loops=1)"
"        Sort Key: (COALESCE((pr.created_at)::timestamp with time zone, now())) DESC"
"        Sort Method: top-N heapsort  Memory: 28kB"
"        ->  Hash Join  (cost=904.60..2426.43 rows=2848 width=16) (actual time=14.815..27.589 rows=2904 loops=1)"
"              Hash Cond: (pr.post_id = post.id)"
"              ->  Seq Scan on post_approval pr  (cost=0.00..1493.00 rows=2848 width=14) (actual time=0.029..11.804 rows=2904 loops=1)"
"                    Filter: ((created_at <= now()) AND (created_at >= (now() - '1 year'::interval)))"
"                    Rows Removed by Filter: 47096"
"              ->  Hash  (cost=844.82..844.82 rows=4782 width=4) (actual time=14.774..14.775 rows=50000 loops=1)"
"                    Buckets: 65536 (originally 8192)  Batches: 1 (originally 1)  Memory Usage: 2270kB"
"                    ->  Seq Scan on post  (cost=0.00..844.82 rows=4782 width=4) (actual time=0.007..6.817 rows=50000 loops=1)"
"Planning Time: 0.196 ms"
"Execution Time: 28.152 ms"
"""


"""
"Limit  (cost=2506.80..2506.92 rows=50 width=16) (actual time=23.198..23.204 rows=50 loops=1)"
"  ->  Sort  (cost=2506.80..2513.92 rows=2848 width=16) (actual time=23.197..23.200 rows=50 loops=1)"
"        Sort Key: (COALESCE(pr.created_at, '2022-11-20 00:00:00'::timestamp without time zone)) DESC"
"        Sort Method: top-N heapsort  Memory: 28kB"
"        ->  Hash Join  (cost=904.60..2412.19 rows=2848 width=16) (actual time=10.229..22.783 rows=2904 loops=1)"
"              Hash Cond: (pr.post_id = post.id)"
"              ->  Seq Scan on post_approval pr  (cost=0.00..1493.00 rows=2848 width=14) (actual time=0.013..11.687 rows=2904 loops=1)"
"                    Filter: ((created_at <= now()) AND (created_at >= (now() - '1 year'::interval)))"
"                    Rows Removed by Filter: 47096"
"              ->  Hash  (cost=844.82..844.82 rows=4782 width=4) (actual time=10.206..10.207 rows=50000 loops=1)"
"                    Buckets: 65536 (originally 8192)  Batches: 1 (originally 1)  Memory Usage: 2270kB"
"                    ->  Seq Scan on post  (cost=0.00..844.82 rows=4782 width=4) (actual time=0.005..4.857 rows=50000 loops=1)"
"Planning Time: 0.207 ms"
"Execution Time: 23.233 ms"
"""

"""
Conclusion: Improvement more than twice
"""