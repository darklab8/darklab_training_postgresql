-- Найти N пользователей с наибольшим рейтингом.

SELECT * FROM user_
ORDER BY rating DESC
LIMIT :N