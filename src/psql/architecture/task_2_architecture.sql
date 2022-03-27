DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS post_editions CASCADE;
DROP TABLE IF EXISTS post_approvals CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS comment_approvals CASCADE;
DROP TABLE IF EXISTS post_visits_per_day CASCADE;
DROP TABLE IF EXISTS post_ratings_per_day CASCADE;

CREATE TABLE users
(
  	id SERIAL PRIMARY KEY,
  	first_name VARCHAR(100) NOT NULL,
	second_name VARCHAR(100) NOT NULL,
  	birth_date DATE NOT NULL,
	email VARCHAR(100) NOT NULL,
	password VARCHAR(100) NOT NULL,
	address VARCHAR(200) NOT NULL,
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
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE,
	CHECK (status IN ('published', 'draft', 'archived')),
    CHECK (array_length(tags, 1) < 20)
);

CREATE TABLE post_ratings_per_day 
(
	id SERIAL PRIMARY KEY,
	post_id INTEGER NOT NULL,
	FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
	day_date DATE DEFAULT NOW(),
	rating BIGINT DEFAULT 0
);

CREATE TABLE post_visits_per_day
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

	
CREATE OR REPLACE FUNCTION user_rating_trigger_function() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
BEGIN

	-- У каждого пользователя есть рейтинг:
	-- 50% составляет средний рейтинг созданных им постов,
	-- 30% составляет средний рейтинг редактированных им постов,
	-- 20% составляет средний рейтинг его комментариев.
    
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

-- CREATE TRIGGER user_rating_trigger_1
-- 	AFTER INSERT OR DELETE
-- 	ON post_approvals
-- 	FOR ROW
-- 		EXECUTE PROCEDURE user_rating_trigger_function();
		
-- CREATE TRIGGER user_rating_trigger_2
-- 	AFTER INSERT OR DELETE
-- 	ON comment_approvals
-- 	FOR ROW
-- 		EXECUTE PROCEDURE user_rating_trigger_function();
		
-- CREATE TRIGGER user_rating_trigger_3
-- 	AFTER INSERT OR DELETE
-- 	ON post_editions
-- 	FOR ROW
-- 		EXECUTE PROCEDURE user_rating_trigger_function();

CREATE INDEX
  ON posts (id);

CREATE INDEX
  ON posts (author_id);

CREATE INDEX
  ON post_editions (post_id);

CREATE INDEX
  ON posts (created_at);

CREATE INDEX
  ON posts USING GIN (tags);

CREATE INDEX
  ON posts USING BRIN (created_at);