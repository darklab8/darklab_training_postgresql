
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .reusable_code import query

def test_task3_2_get_n_ordered_posts(database: Database, load_task2_scheme, factories: TypeFactories):
    "2. Выбрать N опубликованных постов, отсортированных в порядке убывания даты создания;"
    N = 5
    posts = factories.post.create_batch([factories.post.template()for i in range(20)])

    with database.get_core_session() as session:
        result = session.execute(query("query3_2.sql", dict(N=N)))

        fetched_rows = result.fetchall()
        
        assert len(fetched_rows) == len(posts[:N])

        sorted_posts = sorted(posts, key=lambda x: x.created_at, reverse=True)[:N]
        fetched_posts = list([factories.post.template(*row) for row in fetched_rows])

        assert sorted_posts == fetched_posts
