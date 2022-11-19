-- https://www.postgresql.org/docs/current/indexes-intro.html
-- 4. Добавление нового индекса в таблицу не должно блокировать ее и ее записи.
-- https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY

-- USE word CONCURRENTLY for this
--- CREATE INDEX CONCURRENTLY
--- ON post (id);
--- P.S. disabled because it requires separate sessions


CREATE UNIQUE INDEX idx_post_id ON post USING BTREE (id);
CREATE INDEX idx_post_author_id ON post USING BTREE (author_id);
CREATE INDEX idx_post_created ON post USING BTREE (created_at);
CREATE INDEX idx_post_tags ON post USING GIN (tags);

CREATE INDEX idx_post_edition_post_id ON post_edition USING BTREE (post_id);
CREATE INDEX idx_post_edition_tags ON post_edition USING GIN (tags);
CREATE INDEX idx_post_edition_edited ON post_edition USING BTREE (edited_at);

-- DROP INDEX idx_post_id;
-- DROP INDEX idx_post_author_id;
-- DROP INDEX idx_post_created ;
-- DROP INDEX idx_post_tags;

-- DROP INDEX idx_post_edition_post_id;
-- DROP INDEX idx_post_edition_tags;
-- DROP INDEX idx_post_edition_edited;