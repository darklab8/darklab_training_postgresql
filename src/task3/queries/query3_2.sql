-- 2. Выбрать N опубликованных постов, отсортированных в порядке убывания даты создания;

SELECT * FROM post
ORDER BY created_at DESC
LIMIT :N