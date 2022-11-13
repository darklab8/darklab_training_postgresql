import pytest
from pathlib import Path
from utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from sqlalchemy import text

def query(query_filename: str, params = None) -> text:
    if params is None:
        params = dict()

    with open(str(Path(__file__).parent / "queries" / query_filename), "r") as file:
        schema_sql_code = file.read()
    
    return text(schema_sql_code).bindparams(**params)

def test_task3_1_query_count_of_posts_for_user(database: Database, load_task2_scheme, factories: TypeFactories):
    amount_of_posts = 10
    user: UserTemplate = factories.user.create_one(factories.user.template())
    factories.post.create_batch([factories.post.template(author_id=user.id) for i in range(amount_of_posts)])
    # for other users
    factories.post.create_batch([factories.post.template() for i in range(10)])


    with database.get_core_session() as session:
        result = session.execute(query("query3_1.sql", dict(id=user.id)))

        assert result.fetchone()[0] == amount_of_posts

def test_task3_2_get_n_ordered_posts(database: Database, load_task2_scheme, factories: TypeFactories):
    N = 5
    posts = factories.post.create_batch([factories.post.template()for i in range(20)])

    with database.get_core_session() as session:
        result = session.execute(query("query3_2.sql", dict(N=N)))

        fetched_rows = result.fetchall()
        
        assert len(fetched_rows) == len(posts[:N])

        sorted_posts = sorted(posts, key=lambda x: x.created_at, reverse=True)[:N]
        fetched_posts = list([factories.post.template(*row) for row in fetched_rows])

        assert sorted_posts == fetched_posts

def test_task3_3_get_drafts(database: Database, load_task2_scheme, factories: TypeFactories):
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

def test_task3_4(database: Database, load_task2_scheme, factories: TypeFactories):

    post_editions = factories.post_edition.create_batch([factories.post_edition.template() for i in range(10)])