SELECT posts.author_id, SUM(posts.rating) as summed_rating FROM posts
GROUP BY posts.author_id
ORDER BY SUM(posts.rating) DESC
LIMIT :N