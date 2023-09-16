-- Найти N тэгов, для которых суммарное количество посещений связанных с ними постов наибольшее за неделю.#

WITH sub AS (
  SELECT
    pe.post_id as post_id,
	pe.tags as tags,
	pe.edited_at as created,
	row_number() OVER (
      PARTITION BY pe.post_id
      ORDER BY pe.edited_at DESC
    ) AS created_recency
  FROM post_edition AS pe
)

SELECT
	distinct(unnest(sub.tags)) as unnested_tag,
	SUM(COALESCE(pv.visits,0)) as summed_visits
FROM sub
LEFT JOIN post_visits_per_day pv ON pv.post_id = sub.post_id
GROUP BY unnested_tag
ORDER BY summed_visits DESC
LIMIT :N