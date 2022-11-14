CREATE INDEX
  ON posts (id);

CREATE INDEX
  ON posts (author_id);

CREATE INDEX
  ON post_editions (post_id);

CREATE INDEX
  ON posts (created_at);

CREATE INDEX
  ON posts USING GIN (tags);

CREATE INDEX
  ON posts USING BRIN (created_at);