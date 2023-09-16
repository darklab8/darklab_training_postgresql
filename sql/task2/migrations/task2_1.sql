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

CREATE TYPE status AS ENUM ('published', 'draft', 'archived');

CREATE TABLE post
(
	id SERIAL PRIMARY KEY,
	author_id INTEGER NOT NULL,
	status status NOT NULL,

	-- для удобства в потенциале было бы классно додобавить немножко денормализации
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	-- post_edition_first INTEGER, ссылка на первую редакцию поста
	-- FOREIGN KEY (post_edition_first) REFERENCES post_edition (id) ON DELETE NO ACTION,
	-- post_edition_last INTEGER, ссылка на последнюю редакцию поста
	-- FOREIGN KEY (post_edition_last) REFERENCES post_edition (id) ON DELETE NO ACTION,

    FOREIGN KEY (author_id) REFERENCES user_ (id) ON DELETE CASCADE
);

-- 3. Пользователи могут редактировать посты, созданные другими пользователями.

CREATE TABLE post_edition
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES user_ (id) ON DELETE CASCADE,

	-- New Content Of Edition
    edited_at TIMESTAMP NOT NULL DEFAULT NOW(),
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
	AFTER INSERT OR DELETE OR UPDATE
	ON post_approval
	FOR ROW
		EXECUTE PROCEDURE post_rating_function_add();

-- 9. У каждого пользователя есть рейтинг: 50% составляет средний рейтинг созданных им постов, 30% составляет средний рейтинг редактированных им постов, 20% составляет средний рейтинг его комментариев.

ALTER TABLE user_
	ADD COLUMN rating INTEGER DEFAULT 0;

CREATE OR REPLACE FUNCTION user_rating_trigger_function() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
	-- У каждого пользователя есть рейтинг:
    WITH ratings AS (
		-- 50% составляет средний рейтинг созданных им постов,
        SELECT (0.5 * avg(rating))::float as rating FROM post
        WHERE author_id = NEW.user_id
        GROUP BY author_id
		-- 30% составляет средний рейтинг редактированных им постов,
        UNION ALL
        SELECT (0.3 * avg(rating))::float as rating FROM post
        WHERE id IN (SELECT DISTINCT post_id from post_edition
                        WHERE user_id = NEW.user_id)
        UNION ALL
		-- 20% составляет средний рейтинг его комментариев.
        SELECT (0.2 * avg(change))::float as rating FROM comment_approval AS a
        JOIN comment AS c ON a.comment_id = c.id
        WHERE c.user_id = NEW.user_id
    )
    UPDATE user_
    SET rating = (SELECT sum(ratings.rating) FROM ratings)
    WHERE user_.id = NEW.user_id;

   RETURN NEW;
END $$;

CREATE TRIGGER user_rating_trigger_1
	AFTER INSERT OR DELETE OR UPDATE
	ON post_approval
	FOR ROW
		EXECUTE PROCEDURE user_rating_trigger_function();
		
CREATE TRIGGER user_rating_trigger_2
	AFTER INSERT OR DELETE OR UPDATE
	ON comment_approval
	FOR ROW
		EXECUTE PROCEDURE user_rating_trigger_function();
		
CREATE TRIGGER user_rating_trigger_3
	AFTER INSERT OR DELETE OR UPDATE
	ON post_edition
	FOR ROW
		EXECUTE PROCEDURE user_rating_trigger_function();
