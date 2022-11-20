
from utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from ..task3.reusable_code import query, measure_time


def test_task4_2(database: Database, apply_task2_migrations, factories: TypeFactories):
    "2. Найти N наиболее посещаемых постов для заданного пользователя за все время, которые создал не он, но которые он редактировал."

    amount_of_posts=100
    selected = 50

    user: UserTemplate = factories.user.create_one(factories.user.template())
    posts = list(factories.post.create_batch([factories.post.template(author_id=user.id) for i in range(amount_of_posts)]))
    factories.post_visits.create_batch((factories.post_visits.template(post_id=posts[i].id) for i in range(amount_of_posts)))

    user2: UserTemplate = factories.user.create_one(factories.user.template())
    posts2 = list(factories.post.create_batch([factories.post.template(author_id=user2.id) for i in range(amount_of_posts)]))
    factories.post_visits.create_batch((factories.post_visits.template(post_id=posts2[i].id) for i in range(amount_of_posts)))

    factories.post_edition.create_batch([factories.post_edition.template(user_id=user.id,post_id=posts2[i].id) for i in range(selected)])

    with database.get_core_session() as session:
        with measure_time(f"{amount_of_posts=}"):
            result = session.execute(query("query_4_2.sql", dict(N=selected, user_id=user.id), root=__file__))
        fetched_rows = result.fetchall()

        assert len(fetched_rows) == selected
