DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS post_editions CASCADE;
DROP TABLE IF EXISTS post_approvals CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS comment_approvals CASCADE;
DROP TABLE IF EXISTS post_visits_per_day;

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
	tags VARCHAR(50)[],
	status TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    rating INTEGER DEFAULT 0, -- Auto calculated by TRIGGER
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE,
	CHECK (status IN ('published', 'draft', 'archived')),
    CHECK (array_length(tags, 1) < 20)
);

CREATE TABLE post_visits_per_day -- Also known as PostEditedBY in the design.jpg
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
  day_date DATE DEFAULT NOW(),
  visits BIGINT DEFAULT 0
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

CREATE INDEX
  ON posts (author_id);

CREATE INDEX
  ON posts (created_at);

-- add to users, posts, post_approvals
DO $$
  DECLARE v_users_number INT;
  DECLARE v_posts_for_each_user INT;
BEGIN
 v_users_number := 2000;
 v_posts_for_each_user := 100;

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

  UPDATE posts SET tags = array_append(tags,'abc') WHERE id % 2 = 0;
  UPDATE posts SET tags = array_append(tags,'def') WHERE id % 2 = 1;
  UPDATE posts SET tags = array_append(tags,'ghk') WHERE id % 2 = 0;
  UPDATE posts SET tags = array_append(tags,'lmn') WHERE id % 2 = 1;

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

-- add to post_editions
DO $$
  DECLARE v_users_number INT;
  DECLARE v_posts_for_each_user INT;
BEGIN
 v_users_number := 2000;
 v_posts_for_each_user := 50;

  INSERT INTO post_editions(post_id, user_id, edited_at) SELECT
    (num - 1) % (v_posts_for_each_user - 1) + 1,
    (num - 1) % (v_users_number - 1) + 1,
    DATE('2021-0' || num % 9 + 1 || '-' || num % 19 + 10)
  FROM generate_series(1, v_users_number * 2) as num;

END $$;

-- add to post_visits_per_day
DO $$
  DECLARE v_loop_amount INT;
BEGIN
 v_loop_amount := 1000;

  INSERT INTO post_visits_per_day(post_id, day_date, visits) SELECT
    (num - 1) % (50 - 1) + 1,
    DATE('2021-0' || num % 9 + 1 || '-' || num % 19 + 10),
    num % 5 +  num % 10 
  FROM generate_series(1, v_loop_amount) as num;

END $$;

SELECT * FROM posts LIMIT 10