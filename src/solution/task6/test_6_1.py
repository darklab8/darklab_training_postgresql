

from utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from ..task3.reusable_code import query, measure_time
import random

def test_task6_1(database: Database, apply_task2_migrations, factories: TypeFactories):
    """
    Копировать пост по id вместе со связанными авторами и тэгами, но без статистики, комментариев и рейтинга.
    Скопированный пост должен быть в статусе черновик.
    """

    N=50
    post_id=6

    target_user: UserTemplate = factories.user.create_one(factories.user.template())
    target_post: UserTemplate = factories.post.create_one(factories.post.template(author_id=target_user.id,id=post_id))

    # noise
    noise_users = list(factories.user.create_batch([factories.user.template() for i in range(N)]))
    noise_posts = list(factories.post.create_batch([factories.post.template(author_id=noise_users[i].id) for i in range(N)]))

    # target
    target_editions = factories.post_edition.create_batch((factories.post_edition.template(post_id=target_post.id, user_id=random.choice(noise_users).id) for i in range(10)))
    noise_editions = factories.post_edition.create_batch((factories.post_edition.template(post_id=random.choice(noise_posts).id, user_id=random.choice(noise_users).id) for i in range(40)))

    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query_6_1.sql", dict(post_id=target_post.id), root=__file__))

            assert (session.execute("SELECT COUNT(id) FROM post_edition;").fetchone())[0] == 60
            assert (session.execute("SELECT COUNT(id) FROM post;").fetchone())[0] == N+2

