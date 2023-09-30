# Task # 3

## Requirements

Задание 3

(Обязательно к выполнению)

Написать следующие запросы к БД:

1. Посчитать количество постов для пользователя с заданным ID;
2. Выбрать N опубликованных постов, отсортированных в порядке убывания даты создания;
3. Выбрать N постов в статусе "ожидает публикации", отсортированных в порядке возрастания даты создания;
4. Найти N недавно обновленных постов определенного тэга для K страницы (в каждой странице L постов).
5. Найти N постов с наибольшим рейтингом за день/месяц/год.
6. Оценить время выполнения запросов (на достаточном количестве тестовых данных) и проанализировать план выполнения запросов.
7. Сократить время на выполнение запросов, используя подходящие индексы. Сравнить время выполнения и план запросов после создания индексов.
8. Оценить размер используемых индексов. При возможности - сократить размер созданных индексов.

## Solution

## 1. Посчитать количество постов для пользователя с заданным ID;

```sql
--8<-- "sql/task3/queries/query3_1.sql"
```

```go
--8<-- "golang/task3_1_test.go"
```

## 2. Выбрать N опубликованных постов, отсортированных в порядке убывания даты создания;

```sql
--8<-- "sql/task3/queries/query3_2.sql"
```

```go
--8<-- "golang/task3_2_test.go"
```

## 3. Выбрать N постов в статусе "ожидает публикации", отсортированных в порядке возрастания даты создания;

```sql
--8<-- "sql/task3/queries/query3_3.sql"
```

```go
--8<-- "golang/task3_3_test.go"
```

## 4. Найти N недавно обновленных постов определенного тэга для K страницы (в каждой странице L постов).

```sql
--8<-- "sql/task3/queries/query3_4.sql"
```

```go
--8<-- "golang/task3_4_test.go"
```

## 5. Найти N постов с наибольшим рейтингом за день/месяц/год.

```sql
--8<-- "sql/task3/queries/query3_5.sql"
```

```go
--8<-- "golang/task3_5_test.go"
```

`6.` Оценить время выполнения запросов (на достаточном количестве тестовых данных) и проанализировать план выполнения запросов.

Выполним вместе с 7м пунктом

`7.` Сократить время на выполнение запросов, используя подходящие индексы. Сравнить время выполнения и план запросов после создания индексов.

Выполним вместе с 8 пунктом.1.1 Древние индексы, версия первая

## 7.1 итерация первая, Древние индексы

- CREATE UNIQUE INDEX CONCURRENTLY idx_post_id ON post USING BTREE (id);
- CREATE INDEX CONCURRENTLY idx_post_author_id ON post USING BTREE (author_id);
- CREATE INDEX CONCURRENTLY idx_post_created ON post USING BTREE (created_at);
- CREATE INDEX CONCURRENTLY idx_post_edition_post_id ON post_edition USING BTREE (post_id);
- CREATE INDEX CONCURRENTLY idx_post_edition_tags ON post_edition USING GIN (tags);
- CREATE INDEX CONCURRENTLY idx_post_edition_edited ON post_edition USING BTREE (edited_at);
- CREATE INDEX CONCURRENTLY idx_expression_post_approval_change_coal ON post_approval(COALESCE(change,0));
- CREATE INDEX CONCURRENTLY idx_expression_post_approval_created_at_coal ON post_approval(COALESCE(created_at,'2022-11-20'));

**Результаты по тестам**:

params=dbname=skeakexw_indexes users=50000 posts=2500000 post_visits=500000 post_editions=500000

params=dbname=skeakexw_indexless users=50000 posts=2500000 post_visits=500000 post_editions=500000


| test_number | without indexes time | with indexes time | comment                                                   |
| ------------- | ---------------------- | ------------------- | ----------------------------------------------------------- |
| test3_1     | 13.173475ms          | 62.62767ms        |                                                           |
| test3_2     | 494.908922ms         | 236.153827ms      | Могло быть и лучше                         |
| test3_3     | 25.677738ms          | 118.605085ms      |                                                           |
| test3_4     | 182.22424ms          | 177.981498ms      | Могло быть и лучше                         |
| test3_5     | 1.850409686s         | 1.82764835s       | Совсем ужасно. Надо исправить    |
| test4_1     | 93.087218ms          | 81.036566ms       |                                                           |
| test4_2     | 83.113015ms          | 100.553016ms      |                                                           |
| test4_3     | 230.356592ms         | 228.807421ms      | Могло быть и лучше                         |
| test4_4     | 394.721926ms         | 423.217393ms      | Плоховато. Надо улучшить точно. |
| test4_5     | 14.384774ms          | 15.459392ms       |                                                           |
| test4_6     | 547.817848ms         | 581.175276ms      | Плоховато. Надо улучшить точно  |

## 7.2 итерация 2, после оптимизации

Добавили индексы в разные места

- CREATE INDEX CONCURRENTLY idx_post_author_id ON post USING BTREE (author_id); -- 3_1 // replace to HASH
- CREATE INDEX CONCURRENTLY idx_post_created ON post USING BTREE (created_at); -- 3_2, 3_3, 3_5
- CREATE INDEX CONCURRENTLY idx_post_status ON post USING BTREE (status); -- 3_3
- CREATE INDEX CONCURRENTLY idx_post_rating ON post USING BTREE (rating); -- 4_3 ORDER
- CREATE INDEX CONCURRENTLY idx_post_edition_post_id ON post_edition USING HASH (post_id); -- 3_4, 4_2
- CREATE INDEX CONCURRENTLY idx_post_edition_tags ON post_edition USING GIN (tags); -- 3_4
- CREATE INDEX CONCURRENTLY idx_post_edition_edited ON post_edition USING BTREE (edited_at); -- 3_4 ORD
- CREATE INDEX CONCURRENTLY idx_post_edition_user_id ON post_edition USING HASH (user_id); -- 4_2
- CREATE INDEX CONCURRENTLY idx_expression_post_approval_change_coal ON post_approval(COALESCE(change,0)); -- 3_5
- CREATE INDEX CONCURRENTLY idx_post_approval_post_id ON post_approval USING HASH (post_id); -- 3_5
- CREATE INDEX CONCURRENTLY idx_post_visits_day_date ON post_visits_per_day USING BTREE (day_date); -- 4_1 LIKE
- CREATE INDEX CONCURRENTLY idx_post_visits_visits ON post_visits_per_day USING BTREE (visits); -- 4_1 ORDER
- CREATE INDEX CONCURRENTLY idx_post_visits_post_id ON post_visits_per_day USING HASH (post_id); -- 4_6
- CREATE INDEX CONCURRENTLY idx_user_ratings_ratings ON user_ratings USING BTREE (rating); -- 4_5

## 5.2.2 Результаты оптимизации

params=dbname=hpemdfwd_indexes users=50000 posts=2500000 post_visits=500000 post_editions=500000

params=dbname=hpemdfwd_indexless users=50000 posts=2500000 post_visits=500000 post_editions=500000
params=dbname=klvacuzr_indexes


| test_number | without indexes time | with indexes time | comment                                                           |
| ------------- | ---------------------- | ------------------- | ------------------------------------------------------------------- |
| test3_1     | 65.486489ms          | 63.537003ms       |                                                                   |
| test3_2     | 3.391107918s         | 6.038363ms        | с индексом стало намного лучше.         |
| test3_3     | 245.633943ms         | 123.074054ms      |                                                                   |
| test3_4     | 225.958538ms         | 192.053567ms      |                                                                   |
| test3_5     | 2.87305117s          | 1.365926592s      | Очень плохой результат.                       |
| test4_1     | 45.891744ms          | 77.027456ms       |                                                                   |
| test4_2     | 45.254657ms          | 96.388864ms       |                                                                   |
| test4_3     | 224.282966ms         | 220.340384ms      |                                                                   |
| test4_4     | 374.433174ms         | 375.223436ms      | без изменений, но могло быть и лучше |
| test4_5     | 6.024233ms           | 17.840107ms       |                                                                   |
| test4_6     | 186.814461ms         | 498.238269ms      | стало хуже с индексом?                          |

## 5.3.1 Попробуем улучшить тест 3_5

EXPLAIN ANALYZE SELECT
post.id,
SUM(COALESCE(change,0)) as calculated_rating,
post.created_at as created
FROM post
LEFT JOIN post_approval pr ON pr.post_id = post.id
WHERE post.created_at BETWEEN NOW() - interval '3 year' and NOW()
GROUP BY post.id
ORDER BY created DESC
LIMIT 50


| QUERY PLAN                                                                                                                              |
| ----------------------------------------------------------------------------------------------------------------------------------------- |
| Limit  (cost=247788.18..247788.31 rows=50 width=20) (actual time=2202.961..2202.968 rows=50 loops=1)                                    |
| ->  Sort  (cost=247788.18..254038.18 rows=2500000 width=20) (actual time=2197.929..2197.934 rows=50 loops=1)                            |
| Sort Key: post.created_at DESC                                                                                                          |
| Sort Method: top-N heapsort  Memory: 31kB                                                                                               |
| ->  GroupAggregate  (cost=113.70..164739.98 rows=2500000 width=20) (actual time=0.052..1961.966 rows=2500000 loops=1)                   |
| Group Key: post.id                                                                                                                      |
| ->  Merge Left Join  (cost=113.70..120989.98 rows=2500000 width=14) (actual time=0.046..1420.538 rows=2500000 loops=1)                  |
| Merge Cond: (post.id = pr.post_id)                                                                                                      |
| ->  Index Scan using post_pkey on post  (cost=0.43..114602.26 rows=2500000 width=12) (actual time=0.035..1220.553 rows=2500000 loops=1) |
| Filter: ((created_at <= now()) AND (created_at >= (now() - '3 years'::interval)))                                                       |
| ->  Sort  (cost=113.27..117.34 rows=1630 width=6) (actual time=0.008..0.009 rows=0 loops=1)                                             |
| Sort Key: pr.post_id                                                                                                                    |
| Sort Method: quicksort  Memory: 25kB                                                                                                    |
| ->  Seq Scan on post_approval pr  (cost=0.00..26.30 rows=1630 width=6) (actual time=0.003..0.003 rows=0 loops=1)                        |
| Planning Time: 0.192 ms                                                                                                                 |
| JIT:                                                                                                                                    |
| Functions: 15                                                                                                                           |
| Options: Inlining false, Optimization false, Expressions true, Deforming true                                                           |
| Timing: Generation 1.048 ms, Inlining 0.000 ms, Optimization 0.251 ms, Emission 4.646 ms, Total 5.945 ms                                |
| Execution Time: 2204.100 ms                                                                                                             |

Судя по Explain плану, большинству цены затрачено на сортировку по ключу post.created_at DESC 🤔
CREATE INDEX CONCURRENTLY idx_post_created ON post USING BTREE (created_at);

idx_post_created CREATE INDEX idx_post_created ON public.post USING btree (created_at)

post_pkey CREATE UNIQUE INDEX post_pkey ON public.post USING btree (id)
idx_post_approval_post_id CREATE INDEX idx_post_approval_post_id ON public.post_approval USING hash (post_id) # post.id = pr.post_id Merge cond
- Try adding hash on post.id?
- try going btree onto post_approval.post_Id ?
- may be special hash for order by desc?
- may be better index for post.created_at between dates?
- try SUM(COALESCE(change,0)) onto calculated rating? :smile:



`8.` Оценить размер используемых индексов. При возможности - сократить размер созданных индексов.

hpemdfwd_indexes=# \di+

## List of relations


| Name                                     | Table               | Access method | Size       |
| ------------------------------------------ | --------------------- | --------------- | ------------ |
| comment_approval_pkey                    | comment_approval    | btree         | 8192 bytes |
| comment_pkey                             | comment             | btree         | 8192 bytes |
| idx_expression_post_approval_change_coal | post_approval       | btree         | 8192 bytes |
| idx_post_approval_post_id                | post_approval       | hash          | 80 kB      |
| idx_post_author_id                       | post                | btree         | 17 MB      |
| idx_post_created                         | post                | btree         | 54 MB      |
| idx_post_edition_edited                  | post_edition        | btree         | 11 MB      |
| idx_post_edition_post_id                 | post_edition        | hash          | 16 MB      |
| idx_post_edition_tags                    | post_edition        | gin           | 600 kB     |
| idx_post_edition_user_id                 | post_edition        | hash          | 16 MB      |
| idx_post_rating                          | post                | btree         | 17 MB      |
| idx_post_status                          | post                | btree         | 17 MB      |
| idx_post_visits_day_date                 | post_visits_per_day | btree         | 3584 kB    |
| idx_post_visits_post_id                  | post_visits_per_day | hash          | 16 MB      |
| idx_post_visits_visits                   | post_visits_per_day | btree         | 3568 kB    |
| idx_user_ratings_ratings                 | user_ratings        | btree         | 2224 kB    |
| one_approval_only_per_comment_for_user   | comment_approval    | btree         | 8192 bytes |
| one_approval_only_per_post_for_user      | post_approval       | btree         | 8192 bytes |
| only_one_visit_counter_per_post          | post_visits_per_day | btree         | 18 MB      |
| post_approval_pkey                       | post_approval       | btree         | 8192 bytes |
| post_edition_pkey                        | post_edition        | btree         | 11 MB      |
| post_pkey                                | post                | btree         | 54 MB      |
| post_visits_per_day_pkey                 | post_visits_per_day | btree         | 11 MB      |
| user__pkey                               | user_               | btree         | 1112 kB    |

Выводы:

- Btree на datetime намного более требует места чем btree на integer судя по таблицу Post
- Hash индекс может потреблять места больше чем Btree по таблицу post_edition и post_visits_per_day
