SELECT
	distinct(unnest(tags)) as unnested_tag,
	SUM(COALESCE(pv.visits,0)) as summed_visits
FROM post p
LEFT JOIN post_visits_per_day pv ON pv.post_id = p.id
GROUP BY unnested_tag
ORDER BY summed_visits DESC
LIMIT :N