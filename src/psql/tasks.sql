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
DROP TABLE IF EXISTS post_visits CASCADE;
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

CREATE TABLE post_editions
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
	change SMALLINT NOT NULL
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
	change SMALLINT NOT NULL
);