import pytest
from python.utils.database.sql import Database
from ..task2.factories import TypeFactories
from .test_3_4 import base_test

class Params:
    variables = "N"
    # amounts = (100,500,1000,2500,5000,10000,50000,100000)
    amounts = (2000000,)

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_4_performance_no_index(database: Database, apply_task2_migrations: None, factories: TypeFactories, N: int) -> None:
    """
    measured_time=0.0012657642364501953 for N=100
    .measured_time=0.0014224052429199219 for N=500
    .measured_time=0.00171661376953125 for N=1000
    .measured_time=0.0026314258575439453 for N=2500
    .measured_time=0.003918647766113281 for N=5000
    .measured_time=0.007543087005615234 for N=10000
    .measured_time=0.04165530204772949 for N=50000
    .measured_time=0.08265352249145508 for N=100000
    """
    base_test(database, factories, N)
    # breakpoint()

@pytest.mark.performance
@pytest.mark.parametrize(Params.variables, Params.amounts)
def test_task3_4_performance_with_index(database: Database, apply_task3_migrations: None, factories: TypeFactories, N: int) -> None:
    """
    measured_time=0.0011429786682128906 for N=100
    .measured_time=0.0015687942504882812 for N=500
    .measured_time=0.0019197463989257812 for N=1000
    .measured_time=0.0027549266815185547 for N=2500
    .measured_time=0.0040400028228759766 for N=5000
    .measured_time=0.006357431411743164 for N=10000
    .measured_time=0.02451038360595703 for N=50000
    """
    base_test(database, factories, N)
    breakpoint()

"""
Для N=1000000

"Limit  (cost=352709.23..352834.23 rows=50000 width=12) (actual time=1545.598..1551.826 rows=50000 loops=1)"
"  ->  Sort  (cost=352659.23..355662.37 rows=1201255 width=12) (actual time=1537.474..1543.450 rows=70000 loops=1)"
"        Sort Key: (max(GREATEST(post.created_at, post_edition.edited_at))) DESC"
"        Sort Method: external merge  Disk: 9264kB"
"        ->  GroupAggregate  (cost=110075.53..249981.56 rows=1201255 width=12) (actual time=357.505..1456.667 rows=362988 loops=1)"
"              Group Key: post.id"
"              ->  Merge Left Join  (cost=110075.53..228959.59 rows=1201255 width=20) (actual time=357.492..1374.238 rows=400000 loops=1)"
"                    Merge Cond: (post.id = post_edition.post_id)"
"                    Filter: (('target'::text = ANY ((post.tags)::text[])) OR ('target'::text = ANY ((post_edition.tags)::text[])))"
"                    Rows Removed by Filter: 1740394"
"                    ->  Index Scan using post_pkey on post  (cost=0.43..84571.45 rows=2000068 width=49) (actual time=0.013..405.041 rows=2000000 loops=1)"
"                    ->  Materialize  (cost=110074.76..113523.43 rows=689735 width=56) (actual time=357.456..530.276 rows=800000 loops=1)"
"                          ->  Sort  (cost=110074.76..111799.09 rows=689735 width=56) (actual time=357.453..461.004 rows=800000 loops=1)"
"                                Sort Key: post_edition.post_id"
"                                Sort Method: external merge  Disk: 50952kB"
"                                ->  Seq Scan on post_edition  (cost=0.00..19609.35 rows=689735 width=56) (actual time=0.011..96.519 rows=800000 loops=1)"
"Planning Time: 0.263 ms"
"JIT:"
"  Functions: 16"
"  Options: Inlining false, Optimization false, Expressions true, Deforming true"
"  Timing: Generation 1.020 ms, Inlining 0.000 ms, Optimization 0.304 ms, Emission 5.603 ms, Total 6.927 ms"
"Execution Time: 1562.773 ms"


EXPLAIN ANALYZE
SELECT
	post.id,
	max(GREATEST(post.created_at, post_edition.edited_at)) as latest_date
FROM post
LEFT JOIN post_edition on post.id = post_edition.post_id
WHERE 'target' = ANY(post.tags) OR 'target' = ANY(post_edition.tags)
GROUP BY post.id
ORDER BY latest_date DESC
LIMIT 50000 OFFSET 1 * 20000;

"Limit  (cost=310980.59..311105.59 rows=50000 width=12) (actual time=1604.886..1611.307 rows=50000 loops=1)"
"  ->  Sort  (cost=310930.59..313933.72 rows=1201255 width=12) (actual time=1597.029..1603.155 rows=70000 loops=1)"
"        Sort Key: (max(GREATEST(post.created_at, post_edition.edited_at))) DESC"
"        Sort Method: external merge  Disk: 9264kB"
"        ->  GroupAggregate  (cost=0.85..208252.91 rows=1201255 width=12) (actual time=0.034..1508.494 rows=362988 loops=1)"
"              Group Key: post.id"
"              ->  Merge Left Join  (cost=0.85..187230.95 rows=1201255 width=20) (actual time=0.025..1423.082 rows=400000 loops=1)"
"                    Merge Cond: (post.id = post_edition.post_id)"
"                    Filter: (('target'::text = ANY ((post.tags)::text[])) OR ('target'::text = ANY ((post_edition.tags)::text[])))"
"                    Rows Removed by Filter: 1740394"
"                    ->  Index Scan using idx_post_id on post  (cost=0.43..84571.45 rows=2000068 width=49) (actual time=0.012..417.610 rows=2000000 loops=1)"
"                    ->  Index Scan using idx_post_edition_post_id on post_edition  (cost=0.42..71794.27 rows=689735 width=56) (actual time=0.006..538.956 rows=800000 loops=1)"
"Planning Time: 0.206 ms"
"JIT:"
"  Functions: 14"
"  Options: Inlining false, Optimization false, Expressions true, Deforming true"
"  Timing: Generation 0.894 ms, Inlining 0.000 ms, Optimization 0.284 ms, Emission 5.272 ms, Total 6.449 ms"
"Execution Time: 1614.620 ms"

"""

"""
 SELECT i.relname "Table Name",indexrelname "Index Name",
 pg_size_pretty(pg_total_relation_size(relid)) As "Total Size",
 pg_size_pretty(pg_indexes_size(relid)) as "Total Size of all Indexes",
 pg_size_pretty(pg_relation_size(relid)) as "Table Size",
 pg_size_pretty(pg_relation_size(indexrelid)) "Index Size",
 reltuples::bigint "Estimated table row count"
 FROM pg_stat_all_indexes i JOIN pg_class c ON i.relid=c.oid 
 WHERE i.relname='post' OR i.relname='post_edition'

 "post"	"post_pkey"	"53 MB"	"28 MB"	"25 MB"	"4408 kB"	200000
"post"	"idx_post_id"	"53 MB"	"28 MB"	"25 MB"	"4408 kB"	200000
"post"	"idx_post_author_id"	"53 MB"	"28 MB"	"25 MB"	"4176 kB"	200000
"post"	"idx_post_created"	"53 MB"	"28 MB"	"25 MB"	"5048 kB"	200000
"post"	"idx_post_tags"	"53 MB"	"28 MB"	"25 MB"	"11 MB"	200000

"post_edition"	"post_edition_pkey"	"21 MB"	"11 MB"	"10176 kB"	"1768 kB"	80000
"post_edition"	"idx_post_edition_post_id"	"21 MB"	"11 MB"	"10176 kB"	"2336 kB"	80000
"post_edition"	"idx_post_edition_tags"	"21 MB"	"11 MB"	"10176 kB"	"5080 kB"	80000
"post_edition"	"idx_post_edition_edited"	"21 MB"	"11 MB"	"10176 kB"	"2352 kB"	80000
"""



