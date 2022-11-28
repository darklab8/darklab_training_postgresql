-- Вывести рейтинг по количеству тегов под статьей пользователей
-- в порядке увеличения кол-ва тегов.

WITH sub AS (
  SELECT
    pe.post_id as post_id,
	pe.tags as tags,
	pe.created_at as created,
	row_number() OVER (
      PARTITION BY pe.post_id
      ORDER BY pe.created_at DESC
    ) AS created_recency
  FROM post_edition AS pe
)

SELECT
	RANK () OVER ( 
		PARTITION BY p.author_id
		ORDER BY array_length(sub.tags, 1)
	) as ranking,
	p.id, p.author_id, p.status, p.created_at, sub.post_id, sub.tags, sub.created
FROM sub
JOIN post p ON p.id = sub.post_id
