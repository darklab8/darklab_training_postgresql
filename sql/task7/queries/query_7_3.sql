-- К заданию 2 добавить столбцы со значениями:
-- сумма количества тегов в первый день, предыдущий день, следующий день и последний день периода.

-- Вывести суммарное накопленное количество тегов на текущую дату по каждому пользователю.

WITH data_per_day AS (
	WITH post_edition_latest AS (
	  SELECT
		pe.post_id as post_id,
		pe.tags as tags,
		pe.edited_at as edited,
		row_number() OVER (
		  PARTITION BY pe.post_id
		  ORDER BY pe.edited_at DESC
		) AS created_order
	  FROM post_edition AS pe
	)


	SELECT
		SUM(array_length(pel.tags,1)) OVER (
			PARTITION BY p.author_id, pel.edited
		) as tag_sum_per_day,
		post_id, tags,edited, author_id, status, rating
	FROM post_edition_latest pel
	JOIN post p ON p.id = pel.post_id
	WHERE pel.created_order = 1
)

SELECT u.id, COALESCE(first_year.tag_sum_per_day,0) as first_year_tags, COALESCE(second_year.tag_sum_per_day,0) as second_year_tags
FROM user_ u
LEFT JOIN data_per_day first_year ON u.id = first_year.author_id AND first_year.edited BETWEEN NOW() - interval '1 year' and NOW()
LEFT JOIN data_per_day second_year ON u.id = second_year.author_id AND second_year.edited BETWEEN NOW() - interval '2 year' and NOW() - interval '1 year'
ORDER BY first_year_tags DESC
