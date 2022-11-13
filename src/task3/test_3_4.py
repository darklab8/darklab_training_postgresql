
from utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from .reusable_code import query


def test_task3_4(database: Database, load_task2_scheme, factories: TypeFactories):
    "4. Найти N недавно обновленных постов определенного тэга для K страницы (в каждой странице L постов)."
    tag_to_find = "target"
    N = 3
    K = 2
    L = 4
    post_edits_batch1 = factories.post_edition.create_batch([factories.post_edition.template(tags=["target", "other_tag"]) for i in range(20)])
    post_edits_batch2 = factories.post_edition.create_batch([factories.post_edition.template() for i in range(20)])
    posts = factories.post.create_batch([factories.post.template() for i in range(20)])

    with database.get_core_session() as session:
        result = session.execute(query("query3_4.sql", dict(N=N, L=L, K=K, tag=tag_to_find)))

        fetched_rows = result.fetchall()
        print(fetched_rows)
        
        assert len(fetched_rows) == N
