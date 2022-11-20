import pytest
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .test_3_3 import base_test

class Params:
    variables = "N"
    amounts = (10,500,1000,2500,5000,10000,50000,100000)
    # amounts = (100000,)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_3_performance_no_index(database: Database, apply_task2_migrations, factories: TypeFactories, N):
    """
    measured_time=0.0007317066192626953 for N=10
    .measured_time=0.001401662826538086 for N=500
    .measured_time=0.0019099712371826172 for N=1000
    .measured_time=0.003145933151245117 for N=2500
    .measured_time=0.005482196807861328 for N=5000
    .measured_time=0.01079249382019043 for N=10000
    .measured_time=0.05513310432434082 for N=50000
    .measured_time=0.11269259452819824 for N=100000
    """
    base_test(database, factories, N)


@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_3_performance_with_index(database: Database, apply_task3_migrations, factories: TypeFactories, N):
    """
    measured_time=0.0006661415100097656 for N=10
    .measured_time=0.0013499259948730469 for N=500
    .measured_time=0.0017962455749511719 for N=1000
    .measured_time=0.0038406848907470703 for N=2500
    .measured_time=0.00547027587890625 for N=5000
    .measured_time=0.0101470947265625 for N=10000
    .measured_time=0.054205894470214844 for N=50000
    .measured_time=0.0756826400756836 for N=100000
    """
    base_test(database, factories, N)

"""
для N=50000
EXPLAIN ANALYZE VERBOSE

"Limit  (cost=5184.56..5184.92 rows=145 width=1142) (actual time=78.489..86.306 rows=50000 loops=1)"
"  Output: id, author_id, title, content, tags, status, created_at, rating"
"  ->  Sort  (cost=5184.56..5184.92 rows=145 width=1142) (actual time=78.488..84.341 rows=50000 loops=1)"
"        Output: id, author_id, title, content, tags, status, created_at, rating"
"        Sort Key: post.created_at"
"        Sort Method: external merge  Disk: 10008kB"
"        ->  Seq Scan on public.post  (cost=0.00..5179.35 rows=145 width=1142) (actual time=0.006..36.155 rows=100000 loops=1)"
"              Output: id, author_id, title, content, tags, status, created_at, rating"
"              Filter: ((post.status)::text = 'draft'::text)"
"              Rows Removed by Filter: 200000"
"Planning Time: 0.120 ms"
"Execution Time: 88.721 ms"

post_created_at_idx is a btree
"Limit  (cost=0.42..14557.37 rows=50000 width=91) (actual time=0.024..79.988 rows=50000 loops=1)"
"  Output: id, author_id, title, content, tags, status, created_at, rating"
"  ->  Index Scan using post_created_at_idx on public.post  (cost=0.42..29082.29 rows=99890 width=91) (actual time=0.024..77.971 rows=50000 loops=1)"
"        Output: id, author_id, title, content, tags, status, created_at, rating"
"        Filter: ((post.status)::text = 'draft'::text)"
"        Rows Removed by Filter: 100354"
"Planning Time: 0.234 ms"
"Execution Time: 81.025 ms"

Conclusion:
One less performed loop. We got rid of Sort and Seq Scan, and had it replaced to Index Scan
But... estimated cost only increased twice.
"""