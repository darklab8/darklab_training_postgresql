-- Найти N пользователей с наибольшим рейтингом.

SELECT * FROM user_ratings
ORDER BY rating DESC
LIMIT :N
