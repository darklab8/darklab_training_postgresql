SELECT users.id, users.birth_date::date, SUM(posts.rating) as summed_rating FROM posts
JOIN users ON users.id = posts.author_id
WHERE (NOW()::date - users.birth_date::date) < 365 * :K
GROUP BY users.id
ORDER BY SUM(posts.rating) DESC
LIMIT :N