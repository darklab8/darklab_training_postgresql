-- 1. Посчитать количество постов для пользователя с заданным ID;
SELECT count(id) FROM posts
WHERE author_id = {{id}}