-- Удалить всех пользователей, у который рейтинг меньше чем N , вместе со всеми постами и комментариями.
-- Порядок удаления сущностей: комментарии к постам пользователя, комментарии пользователя, посты пользователя, пользователь.

BEGIN;

DELETE FROM post_comment WHERE user_id IN (SELECT DISTINCT(id) FROM post WHERE author_id IN (SELECT id FROM user_ WHERE rating < N)) -- Комменты к постам
DELETE FROM post_comment WHERE user_id IN (SELECT id FROM user_ WHERE rating < N) -- Комментарии пользователя
DELETE FROM post WHERE author_id IN (SELECT id FROM user_ WHERE rating < N)
DELETE FROM user_ WHERE rating < N

COMMIT;