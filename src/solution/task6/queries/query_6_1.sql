-- Копировать пост по id вместе со связанными авторами и тэгами, но без статистики, комментариев и рейтинга.
-- Скопированный пост должен быть в статусе черновик.

BEGIN;

INSERT INTO post (author_id, title, content, tags, status, created_at)
SELECT author_id, title, content, tags, 'draft', created_at FROM post WHERE id = :post_id;

INSERT INTO post_edition (post_id, user_id, edited_at, title, content, tags)
SELECT post_id, user_id, edited_at, title, content, tags FROM post_edition WHERE post_id= :post_id;

COMMIT;