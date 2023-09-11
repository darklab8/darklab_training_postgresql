-- 5. Найти N постов с наибольшим рейтингом за день/месяц/год.

SELECT
	post.id,
	SUM(COALESCE(change,0)) as calculated_rating,
	post.created_at as created
FROM post
LEFT JOIN post_approval pr ON pr.post_id = post.id
WHERE post.created_at BETWEEN NOW() - interval '1 year' and NOW()
GROUP BY post.id
ORDER BY created DESC
LIMIT @N
