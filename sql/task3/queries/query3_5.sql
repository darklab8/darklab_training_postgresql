-- 5. Найти N постов с наибольшим рейтингом за день/месяц/год.

SELECT
	id,
	rating,
	created_at 
FROM post
WHERE created_at BETWEEN NOW() - interval '1 year' and NOW()
-- WHERE created_at BETWEEN NOW() - interval '1 month' and NOW()
-- WHERE created_at BETWEEN NOW() - interval '1 day' and NOW()
ORDER BY created_at DESC
LIMIT :N
