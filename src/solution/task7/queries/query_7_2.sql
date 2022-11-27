-- Вывести суммарное накопленное количество тегов на текущую дату по каждому пользователю.

SELECT
    p.id as post_id,
    MAX(COALESCE(edited_at, created_at)) as latest,
    COALESCE(pe.tags,p.tags) as some_tags,
    ARRAY_LENGTH(COALESCE(pe.tags,p.tags),1) as tag_count,
    row_number() OVER (ORDER BY tag_count ASC) as rating
FROM post p
LEFT JOIN post_edition pe ON pe.post_id = p.id
GROUP BY p.id, some_tags
ORDER BY tag_count ASC