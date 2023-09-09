SELECT * FROM post_visits_per_day
WHERE day_date::TEXT LIKE '2021-%-%'
ORDER BY visits DESC
LIMIT :N