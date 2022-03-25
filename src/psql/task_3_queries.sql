-- Задание 3
-- (Обязательно к выполнению)

-- Написать следующие запросы к БД:

-- Посчитать количество постов для пользователя с заданным ID;
SELECT count(id) FROM posts
WHERE author_id = 5
-- 40 MS (I already created Index for tis one)

-- Выбрать N опубликованных постов, отсортированных в порядке убывания даты создания;
SELECT * FROM posts
ORDER BY created_at DESC
LIMIT 2000

-- Выбрать N постов в статусе "ожидает публикации", отсортированных в порядке возрастания даты создания;
SELECT * FROM posts
WHERE status = 'draft'
ORDER BY created_at ASC
LIMIT 10

-- Найти N недавно обновленных постов определенного тэга для K страницы (в каждой странице L постов).
SELECT post_id, max(edited_at) FROM post_editions
JOIN posts on posts.id = post_editions.post_id
WHERE 'def' = ANY(posts.tags)
GROUP BY post_id
ORDER BY max(edited_at) DESC
LIMIT 10 OFFSET 1*10
-- 190-130 MS before INDEXES
-- 62 MS after adding INDEX to post.id post_editions.post_id, they started to use Index Scan instead of Seq Scan


-- Найти N постов с наибольшим рейтингом за день/месяц/год.
SELECT * FROM post_visits_per_day
WHERE day_date::TEXT LIKE '2021-02-15'
ORDER BY visits DESC
LIMIT 5

SELECT * FROM post_visits_per_day
WHERE day_date::TEXT LIKE '2021-02-%'
ORDER BY visits DESC
LIMIT 5

SELECT * FROM post_visits_per_day
WHERE day_date::TEXT LIKE '2021-%-%'
ORDER BY visits DESC
LIMIT 5

-- Оценить время выполнения запросов (на достаточном количестве тестовых данных) и проанализировать план выполнения запросов.
-- Сократить время на выполнение запросов, используя подходящие индексы. Сравнить время выполнения и план запросов после создания индексов.
-- Оценить размер используемых индексов. При возможности - сократить размер созданных индексов.