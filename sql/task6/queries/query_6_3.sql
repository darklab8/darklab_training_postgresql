-- Добавить следующие столбцы к таблицам, используя alter table:
-- Статус пользователей. Статус может иметь значения активен или заблокирован. Новый столбец должен быть заполнен значением активен

CREATE TYPE user_status AS ENUM ('active', 'blocked');

ALTER TABLE user_
	ADD status user_status NOT NULL DEFAULT 'active'::user_status;
