-- Добавить следующие столбцы к таблицам, используя alter table:
-- Статус пользователей. Статус может иметь значения активен или заблокирован. Новый столбец должен быть заполнен значением активен

ALTER TABLE user_
	ADD status VARCHAR(20) NOT NULL DEFAULT 'active',
	ADD CONSTRAINT allowed_statuses CHECK (status IN ('active', 'blocked'));