import pytest
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .test_3_1 import base_test

class Params:
    variables = "amount_of_posts"
    amounts = (10,1000,5000,50000)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_1_performance_no_index(database: Database, apply_task2_migrations, factories: TypeFactories, amount_of_posts):
    """
    measured_time=0.0008785724639892578 for 3_1, amount_of_posts=10
    .measured_time=0.0010502338409423828 for 3_1, amount_of_posts=1000
    .measured_time=0.002232789993286133 for 3_1, amount_of_posts=5000
    .measured_time=0.009729385375976562 for 3_1, amount_of_posts=50000
    """

    base_test(database=database, factories=factories, amount_of_posts=amount_of_posts)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_1_performance_with_index(database: Database, apply_task3_migrations, factories: TypeFactories, amount_of_posts):
    """
    measured_time=0.0009191036224365234 for 3_1, amount_of_posts=10
    .measured_time=0.0012400150299072266 for 3_1, amount_of_posts=1000
    .measured_time=0.0017991065979003906 for 3_1, amount_of_posts=5000
    .measured_time=0.0068323612213134766 for 3_1, amount_of_posts=50000
    """

    base_test(database=database, factories=factories, amount_of_posts=amount_of_posts)

"""
для amount_of_posts=50000
EXPLAIN ANALYZE VERBOSE

"Aggregate  (cost=2958.67..2958.68 rows=1 width=8) (actual time=10.348..10.349 rows=1 loops=1)"
"  Output: count(id)"
"  ->  Seq Scan on public.post  (cost=0.00..2833.00 rows=50267 width=4) (actual time=0.007..8.307 rows=50000 loops=1)"
"        Output: id"
"        Filter: (post.author_id = 0)"
"        Rows Removed by Filter: 50000"
"Planning Time: 0.039 ms"
"Execution Time: 10.367 ms"

"Aggregate  (cost=2322.69..2322.70 rows=1 width=8) (actual time=9.944..9.945 rows=1 loops=1)"
"  Output: count(id)"
"  ->  Index Scan using post_author_id_idx on public.post  (cost=0.29..2197.52 rows=50070 width=4) (actual time=0.014..7.635 rows=50000 loops=1)"
"        Output: id"
"        Index Cond: (post.author_id = 0)"
"Planning Time: 0.293 ms"
"Execution Time: 9.969 ms"
"""

"""
Итого: Мало значительное, но стабильное улучшение :thinking:.
Query стал искать используя Index судя по анализу.
"""