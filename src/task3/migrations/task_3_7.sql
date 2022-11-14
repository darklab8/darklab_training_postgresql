CREATE INDEX
  ON post (id);

CREATE INDEX
  ON post (author_id);

CREATE INDEX
  ON post_edition (post_id);

CREATE INDEX
  ON post (created_at);

CREATE INDEX
  ON post USING GIN (tags);

CREATE INDEX
  ON post USING BRIN (created_at);