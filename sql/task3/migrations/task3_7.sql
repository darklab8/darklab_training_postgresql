-- https://www.postgresql.org/docs/current/indexes-intro.html
-- 4. Добавление нового индекса в таблицу не должно блокировать ее и ее записи.
-- https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY

CREATE INDEX CONCURRENTLY idx_post_author_id ON post USING BTREE (author_id); -- 3_1 // replace to HASH
CREATE INDEX CONCURRENTLY idx_post_created ON post USING BTREE (created_at); -- 3_2, 3_3, 3_5
CREATE INDEX CONCURRENTLY idx_post_status ON post USING BTREE (status); -- 3_3
CREATE INDEX CONCURRENTLY idx_post_rating ON post USING BTREE (rating); -- 4_3 ORDER

CREATE INDEX CONCURRENTLY idx_post_edition_post_id ON post_edition USING BTREE (post_id); -- 3_4, 4_2
CREATE INDEX CONCURRENTLY idx_post_edition_tags ON post_edition USING GIN (tags); -- 3_4
CREATE INDEX CONCURRENTLY idx_post_edition_edited ON post_edition USING BTREE (edited_at); -- 3_4 ORD
CREATE INDEX CONCURRENTLY idx_post_edition_user_id ON post_edition USING BTREE (user_id); -- 4_2

CREATE INDEX CONCURRENTLY idx_expression_post_approval_change_coal ON post_approval(COALESCE(change,0)); -- 3_5
CREATE INDEX CONCURRENTLY idx_post_approval_post_id ON post_approval USING BTREE (post_id); -- 3_5

CREATE INDEX CONCURRENTLY idx_post_visits_day_date ON post_visits_per_day USING BTREE (day_date); -- 4_1 LIKE
CREATE INDEX CONCURRENTLY idx_post_visits_visits ON post_visits_per_day USING BTREE (visits); -- 4_1 ORDER
CREATE INDEX CONCURRENTLY idx_post_visits_post_id ON post_visits_per_day USING BTREE (post_id); -- 4_6

CREATE INDEX CONCURRENTLY idx_user_ratings_ratings ON user_ratings USING BTREE (rating); -- 4_5
