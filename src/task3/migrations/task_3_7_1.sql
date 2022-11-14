-- https://www.postgresql.org/docs/current/indexes-intro.html
-- 4. Добавление нового индекса в таблицу не должно блокировать ее и ее записи.
-- https://www.postgresql.org/docs/current/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY

CREATE INDEX CONCURRENTLY
  ON post (id);
  