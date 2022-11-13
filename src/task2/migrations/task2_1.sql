-- 1. Для каждого пользователя должна хранится следующая информация: имя, фамилия, дата рождения, емэил, пароль, адрес проживания

CREATE TABLE user_
(
  	id SERIAL PRIMARY KEY,
  	first_name VARCHAR(100) NOT NULL,
	second_name VARCHAR(100) NOT NULL,
  	birth_date DATE NOT NULL,
	email VARCHAR(100) NOT NULL,
	password VARCHAR(100) NOT NULL,
	address VARCHAR(200) NOT NULL
);

-- 2. Пользователь может публиковать посты и задавать для них название, содержание, тэги и статус. Пост может быть в статусе опубликован, черновик или архив.

CREATE TABLE post
(
	id SERIAL PRIMARY KEY,
	author_id INTEGER NOT NULL,
	title VARCHAR(255) NOT NULL,
	content VARCHAR(20000) NOT NULL,
	tags VARCHAR(50)[],
	status VARCHAR(20) NOT NULL,
    FOREIGN KEY (author_id) REFERENCES user_ (id) ON DELETE CASCADE,
	CONSTRAINT allowed_statuses CHECK (status IN ('published', 'draft', 'archived')),
    CONSTRAINT allowed_tags_amount CHECK (array_length(tags, 1) < 20)
);

-- P.S. we would be happy to order posts according to their creation in addition :thinking:
ALTER TABLE post
	ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT NOW();

-- 3. Пользователи могут редактировать посты, созданные другими пользователями.

CREATE TABLE post_edition
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES user_ (id) ON DELETE CASCADE,
    edited_at TIMESTAMP NOT NULL DEFAULT NOW(),

	-- New Content Of Edition
	title VARCHAR(255) NOT NULL,
	content VARCHAR(20000) NOT NULL,
	tags VARCHAR(50)[],
	CONSTRAINT allowed_tags_amount CHECK (array_length(tags, 1) < 20)
);

-- 4. Пользователь может лайкать/дизлайкать чужие посты.

CREATE TABLE post_approval
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES user_ (id) ON DELETE CASCADE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	change SMALLINT NOT NULL,
    CONSTRAINT one_approval_only_per_post_for_user UNIQUE(post_id, user_id),
    CONSTRAINT like_or_dislike_check CHECK(change = 1 OR change = -1)
);

-- 5. Пользователь может комментировать чужие посты. Комментарий содержит только текст.

CREATE TABLE comment
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES user_ (id) ON DELETE CASCADE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	content VARCHAR(2000) NOT NULL
);

-- 6. Пользователь может лайкать/дизлайкать чужие комментарии.

CREATE TABLE comment_approval
(
	id SERIAL PRIMARY KEY,
	comment_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (comment_id) REFERENCES post (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES user_ (id) ON DELETE CASCADE,
	change SMALLINT NOT NULL,
    CONSTRAINT one_approval_only_per_comment_for_user UNIQUE(comment_id, user_id),
    CONSTRAINT like_or_dislike_check CHECK(change = 1 OR change = -1)
);

-- 7. Для каждого поста хранится статистика посещений за день.

CREATE TABLE post_visits_per_day
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
	day_date DATE DEFAULT NOW(),
	visits BIGINT DEFAULT 0,
	CONSTRAINT only_one_visit_counter_per_post UNIQUE (post_id, day_date)
);

-- 8. У каждого поста есть рейтинг: количество лайков минус количество дизлайков.

ALTER TABLE post
	ADD COLUMN rating INTEGER DEFAULT 0;

CREATE FUNCTION post_rating_function_add() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
   UPDATE post
   SET rating = rating + NEW.change
   WHERE post.id = NEW.post_id;
   
   RETURN NEW;
END $$;

CREATE TRIGGER post_rating_calculating_trigger
	AFTER INSERT OR DELETE
	ON post_approval
	FOR ROW
		EXECUTE PROCEDURE post_rating_function_add();

-- 9. У каждого пользователя есть рейтинг: 50% составляет средний рейтинг созданных им постов, 30% составляет средний рейтинг редактированных им постов, 20% составляет средний рейтинг его комментариев.

-- 10. Написать скрипт для заполнения таблицы тестовыми данными. В базе данных должно быть не менее 100000 постов для разных пользователей; Для генерации тестовых данных можно воспользоваться функцией generate_series;