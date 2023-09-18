-- Найти N наиболее посещаемых постов для заданного пользователя за все время,
-- которые создал не он, но которые он редактировал.

SELECT *
FROM post p
LEFT JOIN post_visits_per_day pv ON p.id = pv.post_id
WHERE 
	p.author_id != :user_id AND
	p.id in (SELECT DISTINCT post_id FROM post_edition WHERE user_id = :user_id)
ORDER BY coalesce(pv.visits,0) DESC
LIMIT :N