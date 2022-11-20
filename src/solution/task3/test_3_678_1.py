import pytest
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .test_3_1 import base_test
from .reusable_code import measure_time

class Params:
    variables = "amount_of_posts"
    # amounts = (10,1000,5000,50000,150000)
    amounts = (1000000,)
    # amounts = (300000,)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_1_performance_no_index(database: Database, apply_task2_migrations, factories: TypeFactories, amount_of_posts):
    """
    measured_time=0.0013275146484375 for 3_1, amount_of_posts=10
    .measured_time=0.0011966228485107422 for 3_1, amount_of_posts=1000
    .measured_time=0.002332448959350586 for 3_1, amount_of_posts=5000
    .measured_time=0.014454364776611328 for 3_1, amount_of_posts=50000
    .measured_time=0.038999319076538086 for 3_1, amount_of_posts=150000
    """
    with measure_time(f"launch, {amount_of_posts=}"):
        base_test(database=database, factories=factories, amount_of_posts=amount_of_posts)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_1_performance_with_index(database: Database, apply_task3_migrations, factories: TypeFactories, amount_of_posts):
    """
    measured_time=0.0006289482116699219 for 3_1, amount_of_posts=10
    .measured_time=0.001100301742553711 for 3_1, amount_of_posts=1000
    .measured_time=0.0021598339080810547 for 3_1, amount_of_posts=5000
    .measured_time=0.013616323471069336 for 3_1, amount_of_posts=50000
    .measured_time=0.0391993522644043 for 3_1, amount_of_posts=150000
    """
    with measure_time(f"launch, {amount_of_posts=}"):
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

post_author_id_idx is Btree

"Aggregate  (cost=2322.69..2322.70 rows=1 width=8) (actual time=9.944..9.945 rows=1 loops=1)"
"  Output: count(id)"
"  ->  Index Scan using post_author_id_idx on public.post  (cost=0.29..2197.52 rows=50070 width=4) (actual time=0.014..7.635 rows=50000 loops=1)"
"        Output: id"
"        Index Cond: (post.author_id = 0)"
"Planning Time: 0.293 ms"
"Execution Time: 9.969 ms"
"""

"==================="
"1'000'000 posts"
"""
"Finalize Aggregate  (cost=45079.50..45079.51 rows=1 width=8) (actual time=78.042..79.100 rows=1 loops=1)"
"  ->  Gather  (cost=45079.29..45079.50 rows=2 width=8) (actual time=77.952..79.094 rows=3 loops=1)"
"        Workers Planned: 2"
"        Workers Launched: 2"
"        ->  Partial Aggregate  (cost=44079.29..44079.30 rows=1 width=8) (actual time=76.088..76.089 rows=1 loops=3)"
"              ->  Parallel Seq Scan on post  (cost=0.00..43041.65 rows=415055 width=4) (actual time=0.011..62.606 rows=333333 loops=3)"
"                    Filter: (author_id = 0)"
"                    Rows Removed by Filter: 333333"
"Planning Time: 0.063 ms"
"Execution Time: 79.122 ms"
"""

"""
"Finalize Aggregate  (cost=33292.83..33292.84 rows=1 width=8) (actual time=74.411..75.265 rows=1 loops=1)"
"  ->  Gather  (cost=33292.62..33292.83 rows=2 width=8) (actual time=74.351..75.260 rows=3 loops=1)"
"        Workers Planned: 2"
"        Workers Launched: 2"
"        ->  Partial Aggregate  (cost=32292.62..32292.63 rows=1 width=8) (actual time=72.394..72.395 rows=1 loops=3)"
"              ->  Parallel Index Scan using idx_post_author_id on post  (cost=0.43..31254.98 rows=415055 width=4) (actual time=0.017..57.841 rows=333333 loops=3)"
"                    Index Cond: (author_id = 0)"
"Planning Time: 0.161 ms"
"Execution Time: 75.282 ms"
"""

"""
Итого: Мало значительное, но стабильное улучшение :thinking:.
Query стал искать используя Index судя по анализу. И судя по планировщику на 20% быстрее (до 26% на миллионе данных)
Но база данных отлично кеширует/ускоряет результаты даже на таком колве данных, в связи с чем реальные результаты получить затратно.
"""