-- Копировать пост по id вместе со связанными авторами и тэгами, но без статистики, комментариев и рейтинга.
-- Скопированный пост должен быть в статусе черновик.

DO $$ 
DECLARE
    mypostid INTEGER;
BEGIN

INSERT INTO post(id, author_id, status, created_at)
SELECT nextval('post_id_seq'), author_id, 'draft'::status, created_at FROM post WHERE id = :post_id
RETURNING id INTO mypostid;

INSERT INTO post_edition (post_id, user_id, edited_at, title, content, tags)
SELECT mypostid, user_id, edited_at, title, content, tags FROM post_edition WHERE post_id = :post_id;

COMMIT;
END $$ LANGUAGE plpgsql
