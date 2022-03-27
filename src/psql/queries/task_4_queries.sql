-- Задание 4
-- Написать следующие запросы к БД:


-- 1) Найти N наиболее посещаемых постов за день/месяц/год.
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

-- 2) Найти N наиболее посещаемых постов для заданного пользователя за все время, которые создал не он, но которые он редактировал.
-- Возможо написать

-- 3) Найти N пользователей, для которых суммарный рейтинг для всех созданных ими постов максимальный среди всех пользователей.
-- Возможо написать

-- 4) Найти N пользователей, для которых суммарный рейтинг для всех созданных ими постов максимальный среди всех пользователей младше K лет.
-- Возможо написать

-- 5) Найти N пользователей с наибольшим рейтингом.
SELECT * FROM users
ORDER BY rating DESC
LIMIT 10

-- 6) Найти N тэгов, для которых суммарное количество посещений связанных с ними постов наибольшее за неделю.
-- Возможо написать