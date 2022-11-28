-- Вывести суммарное накопленное количество тегов на текущую дату по каждому пользователю.

WITH post_edition_latest AS (
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

SELECT *, COALESCE(tfcd.tag_amount,0) as coalesced_tag_today
FROM user_
LEFT JOIN (
	SELECT
		SUM(array_length(pel.tags,1)) as tag_amount,
		p.author_id as author_id
	FROM post_edition_latest pel
	JOIN post p ON p.id = pel.post_id
	WHERE pel.created :: DATE BETWEEN '2020-05-10' AND '2020-07-10' -- selected greater length just to have some data
	GROUP BY p.author_id
) tfcd ON tfcd.author_id = user_.id