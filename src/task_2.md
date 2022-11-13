# Task # 2
## Requirements
Задание 2

Спроектировать схему БД для следующего приложения:

1. Для каждого пользователя должна хранится следующая информация: имя, фамилия, дата рождения, емэил, пароль, адрес проживания

2. Пользователь может публиковать посты и задавать для них название, содержание, тэги и статус. Пост может быть в статусе опубликован, черновик или архив.
3. Пользователи могут редактировать посты, созданные другими пользователями.
4. Пользователь может лайкать/дизлайкать чужие посты.
5. Пользователь может комментировать чужие посты. Комментарий содержит только текст.
6. Пользователь может лайкать/дизлайкать чужие комментарии.
7. Для каждого поста хранится статистика посещений за день.
8. У каждого поста есть рейтинг: количество лайков минус количество дизлайков.
9. У каждого пользователя есть рейтинг: 50% составляет средний рейтинг созданных им постов, 30% составляет средний рейтинг редактированных им постов, 20% составляет средний рейтинг его комментариев.
10. Написать скрипт для заполнения таблицы тестовыми данными. В базе данных должно быть не менее 100000 постов для разных пользователей; Для генерации тестовых данных можно воспользоваться функцией generate_series;

Пример генерации тестовых данных
Для таблиц со следующей структурой:

```sql
CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE posts
(
    id SERIAL PRIMARY KEY,
    author_id INTEGER NOT NULL,
    title VARCHAR(255)   NOT NULL,
    text TEXT           NOT NULL,
    status INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (author_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);
```


Можно воспользоваться процедурой (создаёт 10000 пользователей, по 50 постов у каждого, каждый десятый пост в статусе "ожидает модерации" (status=2), остальные посты "опубликованы" (status = 3)):

```sql
DO $$
    DECLARE v_users_number INT;
    DECLARE v_posts_for_each_user INT;
    BEGIN
    v_users_number := 10000;
    v_posts_for_each_user := 50;

    INSERT INTO users 
    SELECT num, concat('name', num), concat('email', num, '@example.com')
    FROM generate_series(1, v_users_number) as num;

    INSERT INTO posts
    SELECT num, (num - 1) % v_users_number + 1, LEFT(MD5(num::varchar), 10), MD5(num::varchar), 3, NOW()
    FROM generate_series(1, v_posts_for_each_user * v_users_number) as num;

    UPDATE posts SET status = 2 WHERE id % 10 = 0;
END $$;
```

## Solution

### Entity Relationship Diagram

<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
</head>
<body>
  <div class="mermaid">
erDiagram
    CUSTOMER ||--o{ POST : contains
    POST ||--o{ POST_VISITS_PER_DAY : contains
    POST ||--o{ POST_EDITION : contains
    CUSTOMER ||--o{ POST_EDITION : contains
    POST ||--o{ POST_APPROVAL : contains
    CUSTOMER ||--o{ POST_APPROVAL : contains
    POST ||--o{ COMMENT : contains
    CUSTOMER ||--o{ COMMENT : contains
    COMMENT ||--o{ COMMENT_APPROVAL : contains
    CUSTOMER ||--o{ COMMENT_APPROVAL : contains
    CUSTOMER {
        int id PK
        string name
        string second_name
        date birthday
        string email
        string password
        string address
        int rating
    }
    POST_VISITS_PER_DAY {
        int id PK
        int post_id FK
        int user_id FK
        date edited_at
    }
    POST {
        int id PK
        int author_id FK
        text title
        text content
        array tags
        text status
        int rating
        date created_at
    }
    POST_EDITION {
        int id PK
        int post_id FK
        int user_id FK
        date edited_at
    }
    POST_APPROVAL {
        int id PK
        int post_id FK
        int user_id FK
        date created_at
        bool change
    }
    COMMENT {
        int id PK
        int post_id FK
        int user_id FK
        date created_at
        text content
    }
    COMMENT_APPROVAL {
        int id PK
        int comment_id FK
        int user_id FK
        bool change
    }


  </div>
 <script src="../shared/mermaid.min.js"></script>
 <script>mermaid.initialize({startOnLoad:true});
</script>
</body>
</html>

### SQL Code

```sql
--8<-- "src/task2/solution/code.sql"
```
