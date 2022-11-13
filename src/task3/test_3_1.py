
from utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from .reusable_code import query

def test_task3_1_query_count_of_posts_for_user(database: Database, load_task2_scheme, factories: TypeFactories):
    "1. Посчитать количество постов для пользователя с заданным ID;"
    amount_of_posts = 10
    user: UserTemplate = factories.user.create_one(factories.user.template())
    factories.post.create_batch([factories.post.template(author_id=user.id) for i in range(amount_of_posts)])
    # for other users
    factories.post.create_batch([factories.post.template() for i in range(10)])


    with database.get_core_session() as session:
        result = session.execute(query("query3_1.sql", dict(id=user.id)))

        assert result.fetchone()[0] == amount_of_posts

