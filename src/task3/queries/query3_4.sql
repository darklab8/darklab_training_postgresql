-- 4. Найти N недавно обновленных постов определенного тэга для K страницы (в каждой странице L постов).

SELECT post_id, max(edited_at) FROM post_editions
JOIN posts on posts.id = post_editions.post_id
WHERE :tag = ANY(posts.tags)
GROUP BY post_id
ORDER BY max(edited_at) DESC
LIMIT :shown_posts OFFSET :K * :L