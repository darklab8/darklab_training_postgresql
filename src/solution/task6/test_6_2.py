from utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from ..task3.reusable_code import query, measure_time
import random

def test_task6_2(database: Database, apply_task2_migrations, factories: TypeFactories):
    """
    Удалить всех пользователей, у который рейтинг меньше чем N , вместе со всеми постами и комментариями.
    Порядок удаления сущностей: комментарии к постам пользователя, комментарии пользователя, посты пользователя, пользователь.
    """

    N=50
    post_id=6

    target_users = list(factories.user.create_batch([factories.user.template(rating=random.randint(0,9)) for i in range(N)]))
    target_posts = list(factories.post.create_batch([factories.post.template(author_id=target_users[i].id) for i in range(N)]))

    # noise
    noise_users = list(factories.user.create_batch([factories.user.template(rating=random.randint(10,99)) for i in range(N)]))
    noise_posts = list(factories.post.create_batch([factories.post.template(author_id=noise_users[i].id) for i in range(N)]))

    # target
    target_editions = factories.post_edition.create_batch((factories.post_edition.template(post_id=random.choice(target_posts).id, user_id=random.choice(target_users).id) for i in range(50)))
    noise_editions = factories.post_edition.create_batch((factories.post_edition.template(post_id=random.choice(noise_posts).id, user_id=random.choice(noise_users).id) for i in range(50)))

    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query_6_2.sql", dict(), root=__file__))

            assert (session.execute("SELECT COUNT(id) FROM post_edition;").fetchone())[0] == 60
            assert (session.execute("SELECT COUNT(id) FROM post;").fetchone())[0] == N+2
