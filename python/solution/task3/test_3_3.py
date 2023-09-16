
from python.utils.database.sql import Database
from ..task2.factories import TypeFactories
from .reusable_code import query, measure_time

def base_test(database: Database, factories: TypeFactories, N: int):
    user = factories.user.create_one(factories.user.template())
    draft_posts = factories.post.create_batch([factories.post.template(status='draft', author_id=user.id)for i in range(N)])
    published_posts = factories.post.create_batch([factories.post.template(status='published', author_id=user.id)for i in range(N)])
    archived_posts = factories.post.create_batch([factories.post.template(status='archived', author_id=user.id)for i in range(N)])

    N_to_query = int(N/2)
    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query3_3.sql", dict(N=N_to_query)))

        fetched_rows = result.fetchall()
        
        assert len(fetched_rows) == len(draft_posts[:N_to_query])

def test_task3_3_get_drafts(database: Database, apply_task2_migrations, factories: TypeFactories):
    "3. Выбрать N постов в статусе 'ожидает публикации', отсортированных в порядке возрастания даты создания;"

    base_test(database, factories, N=5)

