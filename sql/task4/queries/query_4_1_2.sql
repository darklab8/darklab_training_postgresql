SELECT * FROM post_visits_per_day
WHERE day_date >= current_date - interval '1' year and day_date < current_date
ORDER BY visits DESC
LIMIT :N
