
from utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from .reusable_code import query, measure_time

def base_test(database: Database, factories: TypeFactories, amount_of_posts):
    # user we will query
    user: UserTemplate = factories.user.create_one(factories.user.template())
    factories.post.create_batch([factories.post.template(author_id=user.id) for i in range(amount_of_posts)])
    # for other users
    user2: UserTemplate = factories.user.create_one(factories.user.template())
    factories.post.create_batch([factories.post.template(author_id=user2.id) for i in range(amount_of_posts)])

    with database.get_core_session() as session:
        with measure_time(f"3_1, {amount_of_posts=}"):
            result = session.execute(query("query3_1.sql", dict(id=user.id)))
        assert result.fetchone()[0] == amount_of_posts

def test_task3_1_query_count_of_posts_for_user(database: Database, apply_task2_migrations, factories: TypeFactories):
    "1. Посчитать количество постов для пользователя с заданным ID;"

    base_test(database, factories, amount_of_posts=20)


