SELECT unnest(tags), SUM(post_visits_per_day.visits) as visits  FROM posts
JOIN post_visits_per_day ON post_visits_per_day.post_id = posts.id
WHERE post_visits_per_day.day_date BETWEEN NOW() - interval '1 year' and NOW()
GROUP BY unnest(tags)
ORDER BY SUM(post_visits_per_day.visits) DESC
LIMIT :N