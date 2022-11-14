-- https://www.postgresql.org/docs/current/indexes-intro.html
-- 4. Добавление нового индекса в таблицу не должно блокировать ее и ее записи.
-- https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY

CREATE INDEX CONCURRENTLY
  ON post (id);

CREATE INDEX CONCURRENTLY
  ON post (author_id);

CREATE INDEX CONCURRENTLY
  ON post_edition (post_id);

CREATE INDEX CONCURRENTLY
  ON post (created_at);

CREATE INDEX CONCURRENTLY
  ON post USING GIN (tags);

CREATE INDEX CONCURRENTLY
  ON post USING BRIN (created_at);