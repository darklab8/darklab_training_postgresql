/* TASK DESCRIPTION:
Для каждого пользователя должна хранится следующая информация: имя, фамилия, дата рождения, емэил, пароль, адрес проживания
Пользователь может публиковать посты и задавать для них название, содержание, тэги и статус. Пост может быть в статусе опубликован, черновик или архив.
Пользователи могут редактировать посты, созданные другими пользователями.
Пользователь может лайкать/дизлайкать чужие посты.
Пользователь может комментировать чужие посты. Комментарий содержит только текст.
Пользователь может лайкать/дизлайкать чужие комментарии.
Для каждого поста хранится статистика посещений за день. 
У каждого поста есть рейтинг: количество лайков минус количество дизлайков.
У каждого пользователя есть рейтинг: 50% составляет средний рейтинг созданных им постов, 30% составляет средний рейтинг редактированных им постов, 20% составляет средний рейтинг его комментариев.
*/
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS post_editions CASCADE;
DROP TABLE IF EXISTS post_approvals CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS comment_approvals CASCADE;

CREATE TABLE users
(
  	id SERIAL PRIMARY KEY,
  	first_name CHAR(100) NOT NULL,
	second_name CHAR(100) NOT NULL,
  	birth_date DATE NOT NULL,
	email CHAR(100) NOT NULL,
	password CHAR(100) NOT NULL,
	address CHAR(200) NOT NULL,
    rating INTEGER DEFAULT 0 -- Auto calculated by TRIGGER
);

CREATE TABLE posts
(
	id SERIAL PRIMARY KEY,
	author_id INTEGER NOT NULL,
	title VARCHAR(255)   NOT NULL,
	content VARCHAR(20000)           NOT NULL,
	tags CHAR(50)[],
	status TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    rating INTEGER DEFAULT 0, -- Auto calculated by TRIGGER
    visits_today BIGINT DEFAULT 0,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE,
	CHECK (status IN ('published', 'draft', 'archived')),
    CHECK (array_length(tags, 1) < 20)
);

CREATE TABLE post_editions -- Also known as PostEditedBY in the design.jpg
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    edited_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE post_approvals
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
	change SMALLINT NOT NULL,
    UNIQUE(post_id, user_id),
    CHECK(change = 1 OR change = -1)
);

CREATE TABLE comments
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	content VARCHAR(2000)           NOT NULL
);

CREATE TABLE comment_approvals
(
	id SERIAL PRIMARY KEY,
	comment_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (comment_id) REFERENCES posts (id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
	change SMALLINT NOT NULL,
    UNIQUE(comment_id, user_id),
    CHECK(change = 1 OR change = -1)
);

CREATE OR REPLACE FUNCTION post_rating_trigger_function() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
   -- trigger logic
   UPDATE posts
   SET rating = rating + NEW.change
   WHERE posts.id = NEW.post_id;
   
   RETURN NEW;
END $$;

CREATE TRIGGER post_rating_trigger
	AFTER INSERT OR DELETE
	ON post_approvals
	FOR ROW
		EXECUTE PROCEDURE post_rating_trigger_function();
		
CREATE OR REPLACE FUNCTION user_rating_trigger_function() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN
    RAISE NOTICE 'test1'; 
    
    WITH ratings AS (
        SELECT (0.5 * avg(rating))::float as rating FROM posts
        WHERE author_id = NEW.user_id
        GROUP BY author_id
        UNION ALL
        SELECT (0.2 * avg(change))::float as rating FROM comment_approvals AS a
        JOIN comments AS c ON a.comment_id = c.id
        WHERE c.user_id = NEW.user_id
        UNION ALL
        SELECT (0.3 * avg(rating))::float as rating FROM POSTS
        WHERE id IN (SELECT DISTINCT post_id from post_editions
                        WHERE user_id = NEW.user_id)
    )
    UPDATE users
    SET rating = (SELECT sum(ratings.rating) FROM ratings)
    WHERE users.id = NEW.user_id;

    RAISE NOTICE 'var is %', NEW.user_id; 

    RAISE NOTICE 'test12'; 

   RETURN NEW;
END $$;

CREATE TRIGGER user_rating_trigger_1
	AFTER INSERT OR DELETE
	ON post_approvals
	FOR ROW
		EXECUTE PROCEDURE user_rating_trigger_function();
		
CREATE TRIGGER user_rating_trigger_2
	AFTER INSERT OR DELETE
	ON comment_approvals
	FOR ROW
		EXECUTE PROCEDURE user_rating_trigger_function();
		
CREATE TRIGGER user_rating_trigger_3
	AFTER INSERT OR DELETE
	ON post_editions
	FOR ROW
		EXECUTE PROCEDURE user_rating_trigger_function();


DO $$
  DECLARE v_users_number INT;
  DECLARE v_posts_for_each_user INT;
BEGIN
 v_users_number := 1000;
 v_posts_for_each_user := 50;

  INSERT INTO users(first_name, second_name, birth_date, email, password, address) SELECT
    concat('first_name', num),
    concat('second_name', num),
    DATE('1990-04-01'),
    concat('email', num, '@example.com'),
    LEFT(MD5(num::varchar), 100),
    concat('address', num)
  FROM generate_series(1, v_users_number) as num;
  
   INSERT INTO posts(author_id, title, content, status) SELECT
    (num - 1) % v_users_number + 1,
    LEFT(MD5(num::varchar), 5),
    MD5(num::varchar),
    'draft'
  FROM generate_series(1, v_posts_for_each_user * v_users_number) as num;

  UPDATE posts SET status = 'published' WHERE id % 10 = 0;

  INSERT INTO post_approvals(post_id, user_id, change) SELECT
    (num - 1) % (v_posts_for_each_user - 1) + 1,
    num,
    1
  FROM generate_series(1, v_users_number) as num;

END $$;

INSERT INTO post_approvals(post_id, user_id, change)
VALUES (5, 1000, -1);

INSERT INTO post_approvals(post_id, user_id, change)
VALUES (5, 999, -1);

INSERT INTO comments(user_id, post_id, content)
VALUES (5, 1, '1'),
        (5, 1, '2'),
        (5, 1, '3');

INSERT INTO comment_approvals(user_id, comment_id, change)
VALUES (5, 1, 1),
        (5, 2, -1),
        (5, 3, 1);

INSERT INTO post_editions(post_id, user_id)
VALUES (21, 5),
        (22, 5),
        (23, 6),
        (22, 5),
        (22, 5);

SELECT * FROM users
WHERE id = 5;
