
from utils.database.sql import Database
from ..task2.factories import TypeFactories
from .reusable_code import query, measure_time

def base_test(database: Database, factories: TypeFactories, N, enable_sorting=False):
    "2. Выбрать N опубликованных постов, отсортированных в порядке убывания даты создания;"
    user = factories.user.create_one(factories.user.template())
    posts = factories.post.create_batch([factories.post.template(author_id=user.id) for i in range(N*2)])

    with database.get_core_session() as session:
        with measure_time(f"3_2, {N=}"):
            result = session.execute(query("query3_2.sql", dict(N=N)))

        fetched_rows = result.fetchall()
        
        assert len(fetched_rows) == len(posts[:N])

        if not enable_sorting:
            return
        
        sorted_posts = sorted(posts, key=lambda x: x.created_at, reverse=True)[:N]
        fetched_posts = list([factories.post.template(*row) for row in fetched_rows])

        assert sorted_posts == fetched_posts

def test_task3_2_get_n_ordered_posts(database: Database, apply_task2_migrations, factories: TypeFactories):
    base_test(database, factories, 10)
