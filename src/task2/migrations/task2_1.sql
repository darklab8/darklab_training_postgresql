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
	content TEXT NOT NULL,
	tags VARCHAR(50)[],
	status TEXT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES user_ (id) ON DELETE CASCADE,
	CHECK (status IN ('published', 'draft', 'archived')),
    CHECK (array_length(tags, 1) < 20)
);
