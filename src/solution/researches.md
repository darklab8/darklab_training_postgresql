# Researches

## 1. При создании таблиц максимально ограничить возможные значения каждого поля (unique, foreign key и т.д.)

## 2. Использовать индексы

- https://www.postgresql.org/docs/current/sql-createindex.html
- https://www.postgresql.org/docs/current/indexes-types.html
- https://devcenter.heroku.com/articles/postgresql-indexes#index-types
- https://www.datavail.com/blog/postgresql-indexing-types-how-when-where-should-they-be-used/
- https://pganalyze.com/blog/gin-index

### Index Types

- B-Tree default.
    - Подходит для `<   <=   =   >=   > BETWEEN, IN, IS NULL, IS NOT NULL, LIKE in string beginning "foo%"`
    - Небольшое ускорение для сортировки
    - for selective WHERE usually
    - works better when one value in a row
    - sort by asciending by default, has a point to have reverse order of creation `CREATE INDEX articles_published_at_index ON articles(published_at DESC NULLS LAST);`

- Hash
    - для сравнений `=` только
    - Не используй до 10 пострегса, они были транзакционно не безопасными

- Gist
    - spatial index
    - for `<<   &<   &>   >>   <<|   &<|   |&>   |>>   @>   <@   ~=   &&`
    - `SELECT * FROM places ORDER BY location <-> point '(101,456)' LIMIT 10;`
    - geometric data types + full-text search
    - For dynamic data, GiST is a better option compare to GIN.

- SP-GiST
    - `<<   >>   ~=   <@   <<|   |>>`

- GIN
    - invertex index
    - `<@   @>   =   &&`
    - works for arrays?
    - great when many values in one row
    - works for full-text search too
    - Good for JSONB data, Array, Range types, Full text search, and hstore

- BRIN
    - `<   <=   =   >=   >`
    - Хорош для ситуации когда дата упорядочена по умолчанию в таблице. Данные коррелированы.
    - Best for timestamp!

- Bloom
    - Good for WHERE equal AND/OR cases
    - Covering all attributes
    - CREATE INDEX bloom_idx_bar ON foo.bar USING bloom (id,dept_id,zipcode)
        WITH (length=80, col1=4, col2=2, col3=4); # number of bits per column. from 80 to 4096 size index allowed

- multi collumn indexes
    - `SELECT name FROM test2 WHERE major = constant AND minor = constant;`
    - `CREATE INDEX test2_mm_idx ON test2 (major, minor);`
    - `B-tree, GiST, GIN, and BRIN` работают только для этого

- partial indexes
    - save space and improve performance
    - CREATE INDEX articles_flagged_created_at_index ON articles(created_at) WHERE flagged IS TRUE;

- expression indexes
    - improve speed of functions
    - CREATE INDEX users_lower_email ON users(lower(email));
    - CREATE INDEX articles_day ON articles ( date(published_at) ) # for timestamp

- unique indexes speed up look up.
    - unique indexes aren't exactly equal to unique constraint, but quire close. unique index is lower implementation

### Index Sizes


```
SELECT
   relname  as table_name,
   pg_size_pretty(pg_total_relation_size(relid)) As "Total Size",
   pg_size_pretty(pg_indexes_size(relid)) as "Index Size",
   pg_size_pretty(pg_relation_size(relid)) as "Actual Size"
   FROM pg_catalog.pg_statio_user_tables 
ORDER BY pg_total_relation_size(relid) DESC;
```

```
 SELECT i.relname "Table Name",indexrelname "Index Name",
 pg_size_pretty(pg_total_relation_size(relid)) As "Total Size",
 pg_size_pretty(pg_indexes_size(relid)) as "Total Size of all Indexes",
 pg_size_pretty(pg_relation_size(relid)) as "Table Size",
 pg_size_pretty(pg_relation_size(indexrelid)) "Index Size",
 reltuples::bigint "Estimated table row count"
 FROM pg_stat_all_indexes i JOIN pg_class c ON i.relid=c.oid 
 WHERE i.relname='uploads'
```



https://stackoverflow.com/questions/46470030/postgresql-index-size-and-value-number

```
# querying index size
"SELECT pg_size_retty (pg_relation_size(`indexname`));"
```

- btree 100%
- hash 138%. Хмм для 20 миллионов записей, есть данные что он в два раза меньше чем Btree https://www.datavail.com/blog/postgresql-indexing-types-how-when-where-should-they-be-used/
- Gin is very small and compact?
- SP-GIST, supposedly super huge and for in memory usage
- BRIN - the most space efficient. It can be 2000 times smaller than B-tree
  - на 20 миллионах записей. 10% быстрее для таймстапмов чем бтрии. и экономит 99% дискового места.
- Bloom: 4 times smaller than Btree. Multicolumn.

### Index Comparison

`https://www.datavail.com/blog/postgresql-indexing-types-how-when-where-should-they-be-used/` very good article on big data checking


|   Index Type:   | B-Tree                                                        | Hash                                                                          | BRIN                                                                                                 | GIN                                                                                                                                   | GiST                                                                                                              | Bloom                                                              |
| :---------------: | --------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------- |
|   Use Cases:   | # Most of use cases.# =, >=, <=, >, <, IN, BETWEEN operations | # Equality (=) operations.<br/># High cardinality column as an indexed column | # Timestamped sensor data.<br/># Internet of Things (IoT)<br/># Sequentially lined up large data set | # Static data<br/># Index arrays, jsonb, and tsvector# Full text search (LIKE ‘%string%’)<br/># Efficient for <@, &&, @@@ operators | # Dynamic data<br/># Geometries, array<br/># Useful when using PostGIS<br/># Full text search (LIKE ‘%string%’) | Good for WHERE = AND/OR =<br />(like HASH but for multiple column) |
|   Index Size:   | Large compare to Hash and BRIN                                | Small compare to B-Tree                                                       | Very small in size                                                                                   | Large compare to GiST                                                                                                                 | Small compare to GIN                                                                                              | Small, efficient, multi                                            |
| Comparable size | 100%                                                          | 50%                                                                           | 1%                                                                                                   | 21%                                                                                                                                   | 10%?                                                                                                              | 25%(multic). Technically variable.                                 |
| is multicollumn | Yes                                                           | No                                                                            | Yes                                                                                                  | Yes                                                                                                                                   | Yes                                                                                                               | yes                                                                |
|      Cons      |                                                               |                                                                               |                                                                                                      | Expensive to insert and update, becomes especially slow when gin_pending_list_limit is reached, if default                            |                                                                                                                   | Slower than B-tree for small amount of taken columns with btree    |

## 3. Добавление новых столбцов в таблицу не должно блокировать ее и ее записи.

- Решено в Postgresql 11 автоматически
- https://www.depesz.com/2018/04/04/waiting-for-postgresql-11-fast-alter-table-add-column-with-a-non-null-default/

## 4. Добавление нового индекса в таблицу не должно блокировать ее и ее записи.

- Создавай `CONCURRENTLY`
- https://www.postgresql.org/docs/current/sql-createindex.html


## Research about transactional level:

`TODO add here`

## Research regarding views and materialized views?

`TODO perhaps mention them here too`