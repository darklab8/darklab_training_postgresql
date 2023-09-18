-- Найти N наиболее посещаемых постов за день/месяц/год.

SELECT * FROM post_visits_per_day
WHERE day_date::TEXT LIKE '2023-%-%'
ORDER BY visits DESC
LIMIT :N
