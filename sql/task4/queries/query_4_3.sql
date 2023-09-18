-- Найти N пользователей, для которых суммарный рейтинг для всех созданных ими постов максимальный среди всех пользователей.

SELECT p.author_id, SUM(p.rating) as summed_rating FROM post p
GROUP BY p.author_id
ORDER BY SUM(p.rating) DESC
LIMIT :N