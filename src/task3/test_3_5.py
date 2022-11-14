from utils.database.sql import Database
from ..task2.factories import TypeFactories
import random
from .reusable_code import query

def test_task3_5(database: Database, apply_task2_migrations, factories: TypeFactories):
    "5. Найти N постов с наибольшим рейтингом за день/месяц/год."
    N=5
    approving_users = factories.user.create_batch([factories.user.template(id=i) for i in range(100,200)])
    posts = factories.post.create_batch([factories.post.template() for i in range(10)])
    post_approvals = factories.post_approval.create_batch([
        factories.post_approval.template(
            user_id=approving_users[i].id,
            post_id=random.choice([post.id for post in posts]),
        ) for i in range(100)])

    with database.get_core_session() as session:
        result = session.execute(query("query3_5.sql", dict(N=N)))

        fetched_rows = result.fetchall()
        fetched_posts = list([factories.post.template(*row) for row in fetched_rows])
        
        assert len(fetched_posts) == N

        for i, post in enumerate(fetched_posts[1:]):
            assert fetched_posts[i-1].rating >= post.rating
