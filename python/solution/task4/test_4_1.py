
from python.utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from ..task3.reusable_code import query, measure_time, Task

task = Task.task4

def test_task4_1(database: Database, apply_task2_migrations, factories: TypeFactories):
    "1. Найти N наиболее посещаемых постов за день/месяц/год.;"

    amount_of_posts=50
    selected = 10

    user: UserTemplate = factories.user.create_one(factories.user.template())
    posts = list(factories.post.create_batch((factories.post.template(author_id=user.id) for i in range(amount_of_posts))))
    factories.post_visits.create_batch((factories.post_visits.template(post_id=posts[i].id) for i in range(amount_of_posts)))

    with database.get_core_session() as session:
        with measure_time(f"{amount_of_posts=}"):
            result = session.execute(query("query_4_1.sql", task, dict(N=selected)))
        fetched_rows = result.fetchall()
        assert len(fetched_rows) == selected
