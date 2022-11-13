-- 3. Выбрать N постов в статусе "ожидает публикации", отсортированных в порядке возрастания даты создания;

SELECT * FROM post
WHERE status = 'draft'
ORDER BY created_at ASC
LIMIT :N