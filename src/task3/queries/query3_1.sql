-- 1. Посчитать количество постов для пользователя с заданным ID;
SELECT count(id) FROM post
WHERE author_id = :id