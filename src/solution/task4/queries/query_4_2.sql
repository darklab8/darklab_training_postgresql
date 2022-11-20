SELECT p.id as post_id, coalesce(pv.visits,0) as visiting
FROM posts p
LEFT JOIN post_visits_per_day pv ON p.id = pv.post_id
WHERE p.author_id != :user_id AND p.author_id in (SELECT DISTINCT post_id FROM post_editions WHERE user_id = :user_id)
ORDER BY coalesce(pv.visits,0) DESC
LIMIT :N