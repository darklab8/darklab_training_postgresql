-- 5. Найти N постов с наибольшим рейтингом за день/месяц/год.

SELECT post.id, COALESCE(change,0), COALESCE(pr.created_at,NOW()) as calculated_rating FROM post
LEFT JOIN post_approval pr ON pr.post_id = post.id
WHERE pr.created_at BETWEEN NOW() - interval '1 year' and NOW()
ORDER BY calculated_rating DESC
LIMIT :N