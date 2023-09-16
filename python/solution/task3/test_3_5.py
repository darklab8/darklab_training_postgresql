from python.utils.database.sql import Database
from ..task2.factories import TypeFactories
import random
from .reusable_code import query, measure_time, Task

task = Task.task3

def base_test(database: Database, factories: TypeFactories, N: int):
    "5. Найти N постов с наибольшим рейтингом за день/месяц/год."
    approving_users = list(factories.user.create_batch((factories.user.template(id=i) for i in range(N))))
    posts = factories.post.create_batch((factories.post.template(author_id=approving_users[i].id) for i in range(N)))
    post_approvals = factories.post_approval.create_batch((
        factories.post_approval.template(
            user_id=approving_users[i].id,
            post_id=random.choice([post.id for post in posts]),
        ) for i in range(N)))

    N_to_query = int(N/10)

    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query3_5.sql", task, dict(N=N_to_query)))

        fetched_rows = result.fetchall()
        fetched_posts = list([factories.post.template(*row) for row in fetched_rows])
        
        assert len(fetched_posts) > 0 and len(fetched_posts) < N_to_query

        for i, post in enumerate(fetched_posts[1:]):
            assert fetched_posts[i-1].rating >= post.rating

def test_task3_5(database: Database, apply_task2_migrations, factories: TypeFactories):
    base_test(database=database, factories=factories, N=500)
    

