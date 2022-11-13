-- 5. Найти N постов с наибольшим рейтингом за день/месяц/год.

SELECT post_id, sum(change) FROM posts
LEFT JOIN post_approvals pr ON pr.post_id = posts.id
WHERE pr.created_at BETWEEN NOW() - interval '1 year' and NOW()
GROUP BY post_id
ORDER BY sum(change) DESC
LIMIT :N