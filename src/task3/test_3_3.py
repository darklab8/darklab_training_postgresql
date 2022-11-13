
from utils.database.sql import Database
from ..task2.factories import TypeFactories

from .reusable_code import query

def test_task3_3_get_drafts(database: Database, load_task2_scheme, factories: TypeFactories):
    "3. Выбрать N постов в статусе 'ожидает публикации', отсортированных в порядке возрастания даты создания;"
    N = 5
    draft_posts = factories.post.create_batch([factories.post.template(status='draft')for i in range(10)])
    published_posts = factories.post.create_batch([factories.post.template(status='published')for i in range(10)])
    archived_posts = factories.post.create_batch([factories.post.template(status='archived')for i in range(10)])

    with database.get_core_session() as session:
        result = session.execute(query("query3_3.sql", dict(N=N)))

        fetched_rows = result.fetchall()
        
        assert len(fetched_rows) == len(draft_posts[:N])

        sorted_posts = sorted(draft_posts, key=lambda x: x.created_at)[:N]
        fetched_posts = list([factories.post.template(*row) for row in fetched_rows])

        assert sorted_posts == fetched_posts
