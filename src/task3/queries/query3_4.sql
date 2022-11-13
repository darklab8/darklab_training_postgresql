-- 4. Найти N недавно обновленных постов определенного тэга для K страницы (в каждой странице L постов).

SELECT
	post.id,
	max(GREATEST(post.created_at, post_edition.edited_at)) as latest_date
FROM post
LEFT JOIN post_edition on post.id = post_edition.post_id
WHERE :tag = ANY(post.tags) OR :tag = ANY(post_edition.tags)
GROUP BY post.id
ORDER BY latest_date DESC
LIMIT :N OFFSET :K * :L
