from ..app.utils import run_query
from ..conftest import Consts

# -- Задание 3
# -- (Обязательно к выполнению)

# -- Написать следующие запросы к БД:


def test_task_3_find_users_by_id(filled_db, engine):
    "Посчитать количество постов для пользователя с заданным ID;"

    id = 5

    the_amount_of_message_user_has = run_query(
        engine,
        """SELECT count(id) FROM posts WHERE author_id = :author_id""",
        {"author_id": id},
        return_first=True,
    )

    assert the_amount_of_message_user_has == Consts.posts_per_user


def test_task_3_find_N_published_posts_sorted_by_time_of_creation(filled_db, engine):
    "Выбрать N опубликованных постов, отсортированных в порядке убывания даты создания;"

    N = 25

    results = run_query(
        engine,
        """
        SELECT * FROM posts
        ORDER BY created_at DESC
        LIMIT :N
        """,
        {"N": N},
    )

    assert results.rowcount == N


# -- 2) Выбрать N опубликованных постов, отсортированных в порядке убывания даты создания;
# SELECT * FROM posts
# ORDER BY created_at DESC
# LIMIT 2000

# -- 3) Выбрать N постов в статусе "ожидает публикации", отсортированных в порядке возрастания даты создания;
# SELECT * FROM posts
# WHERE status = 'draft'
# ORDER BY created_at ASC
# LIMIT 10

# -- 4) Найти N недавно обновленных постов определенного тэга для K страницы (в каждой странице L постов).
# SELECT post_id, max(edited_at) FROM post_editions
# JOIN posts on posts.id = post_editions.post_id
# WHERE 'def' = ANY(posts.tags)
# GROUP BY post_id
# ORDER BY max(edited_at) DESC
# LIMIT 10 OFFSET 1*10
# -- 190-130 MS before INDEXES
# -- 62 MS after adding INDEX to post.id post_editions.post_id, they started to use Index Scan instead of Seq Scan

# -- 5) Найти N постов с наибольшим рейтингом за день/месяц/год.
# -- TODO rewrite to calculate RATING per day/month/year
# -- You need to have rating_of_posts_per_day, and to have adjusted trigger

# -- Оценить время выполнения запросов (на достаточном количестве тестовых данных) и проанализировать план выполнения запросов.
# -- Сократить время на выполнение запросов, используя подходящие индексы. Сравнить время выполнения и план запросов после создания индексов.
# -- Оценить размер используемых индексов. При возможности - сократить размер созданных индексов.
