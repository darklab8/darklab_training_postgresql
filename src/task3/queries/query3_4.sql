-- 4. Найти N недавно обновленных постов определенного тэга для K страницы (в каждой странице L постов).

SELECT
	post.id,
	max(post_edition.created_at) as latest_date
FROM post
JOIN post_edition on post.id = post_edition.post_id
WHERE :tag = ANY(post_edition.tags)
GROUP BY post.id
ORDER BY latest_date DESC
LIMIT LEAST(:N,:L) OFFSET :K * :L
