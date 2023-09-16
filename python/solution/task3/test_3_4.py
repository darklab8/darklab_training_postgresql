
from python.utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from .reusable_code import query, measure_time
import random


def base_test(database: Database, factories: TypeFactories, N):
    "4. Найти N недавно обновленных постов определенного тэга для K страницы (в каждой странице L постов)."
    tag_to_find = "target"
    
    users = factories.user.create_batch([factories.user.template() for i in range(10)])
    posts = factories.post.create_batch([factories.post.template(author_id=random.choice(users).id) for i in range(N)])
    post_edits_batch1 = factories.post_edition.create_batch([
        factories.post_edition.template(tags=["target", "other_tag"],
            user_id=random.choice(users).id,post_id=random.choice(posts).id) for i in range(int(N/5))])
    post_edits_batch2 = factories.post_edition.create_batch([
        factories.post_edition.template(
            user_id=random.choice(users).id,post_id=random.choice(posts).id) for i in range(int(N/5))])

    K = int(1)
    L = int(N/10)
    N_to_query = int(N/20)
    
    Results_amount = min(N_to_query, L)
    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query3_4.sql", dict(N=N_to_query, L=L, K=K, tag=tag_to_find)))

        fetched_rows = result.fetchall()
        
        assert len(fetched_rows) == N_to_query

def test_task3_4(database: Database, apply_task2_migrations, factories: TypeFactories):
    
    base_test(database, factories, N=10)