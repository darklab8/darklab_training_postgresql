from ..app.utils import run_query
from ..conftest import Consts

# -- Задание 4
# -- Написать следующие запросы к БД:


def test_task_4_find_N_most_visited_posts_for_day_month_year(filled_db, engine):
    "Найти N наиболее посещаемых постов за день/месяц/год."
    N = 10

    # разница между днём, месяцем, или годом: в параметре 2021-02-12, 2021-02-%, 2021-%-%
    results = run_query(
        engine,
        rf"""

SELECT * FROM post_visits_per_day
WHERE day_date::TEXT LIKE '2021-%-%'
ORDER BY visits DESC
LIMIT :N

        """,
        {"N": N},
    )

    assert results.rowcount == N


def test_task_4_find_most_visited_posts_which_he_edited_but_did_not_create(filled_db, engine):
    "Найти N наиболее посещаемых постов для заданного пользователя за все время, которые создал не он, но которые он редактировал."
    N = 10
    user_id = 5

    results = run_query(
        engine,
        rf"""

SELECT p.id as post_id, coalesce(pv.visits,0) as visiting
FROM posts p
LEFT JOIN post_visits_per_day pv ON p.id = pv.post_id
WHERE p.author_id != :user_id AND p.author_id in (SELECT DISTINCT post_id FROM post_editions WHERE user_id = :user_id)
ORDER BY coalesce(pv.visits,0) DESC
LIMIT :N


        """,
        {"N": N, "user_id": user_id},
    )

    assert results.rowcount > 0

def test_task_4_find_N_users_with_best_summed_ratings_from_all_their_posts(filled_db, engine):
    "Найти N пользователей, для которых суммарный рейтинг для всех созданных ими постов максимальный среди всех пользователей."
    N = 10

    results = run_query(
        engine,
        rf"""

SELECT posts.author_id, SUM(post_ratings_per_day.rating) as summed_rating FROM posts
JOIN post_ratings_per_day ON posts.id = post_ratings_per_day.post_id
GROUP BY post_ratings_per_day.post_id, posts.author_id
ORDER BY SUM(post_ratings_per_day.rating) DESC
LIMIT :N

        """,
        {"N": N},
    )

    assert results.rowcount == N



# -- 4) Найти N пользователей, для которых суммарный рейтинг для всех созданных ими постов максимальный среди всех пользователей младше K лет.
# -- Возможо написать


def test_task_4_find_most_rated_users(filled_db, engine):
    "Найти N пользователей с наибольшим рейтингом."
    N = 10

    results = run_query(
        engine,
        rf"""

SELECT * FROM users
ORDER BY rating DESC
LIMIT :N

        """,
        {"N": N},
    )

    assert results.rowcount == N


# -- 6) Найти N тэгов, для которых суммарное количество посещений связанных с ними постов наибольшее за неделю.
# -- Возможо написать
