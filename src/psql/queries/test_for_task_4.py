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


# -- 2) Найти N наиболее посещаемых постов для заданного пользователя за все время, которые создал не он, но которые он редактировал.
# -- Возможо написать

# -- 3) Найти N пользователей, для которых суммарный рейтинг для всех созданных ими постов максимальный среди всех пользователей.
# -- Возможо написать

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
