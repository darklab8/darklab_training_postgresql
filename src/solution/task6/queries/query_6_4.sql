-- Добавить следующие столбцы к таблицам, используя alter table:
-- Даты создания и обновления для постов. Новые столбцы должны быть заполнены значением текущей даты.

ALTER TABLE post
	ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN motified_at TIMESTAMP NOT NULL DEFAULT NOW();