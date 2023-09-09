SELECT u.id, u.birth_date::date, SUM(p.rating) as summed_rating FROM post p
JOIN user_ u ON u.id = p.author_id
WHERE (NOW()::date - u.birth_date::date) < 365 * :K
GROUP BY u.id
ORDER BY summed_rating DESC
LIMIT :N