
from python.utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from ..task3.reusable_code import query, measure_time, Task
import secrets
import random
import datetime as dt

task = Task.task4

def test_task4_6(database: Database, apply_task2_migrations, factories: TypeFactories):
    "6. Найти N тэгов, для которых суммарное количество посещений связанных с ними постов наибольшее за неделю."

    N=100
    selected = 50

    target_tags = [secrets.token_hex(4) for i in range(50)]
    noise_tags = [secrets.token_hex(4) for i in range(50)]

    user = factories.user.create_one(factories.user.template())

    target_posts = list(factories.post.create_batch((factories.post.template(author_id=user.id) for i in range(N))))
    noise_posts = list(factories.post.create_batch((factories.post.template(author_id=user.id) for i in range(N))))

    target_post_visits = factories.post_visits.create_batch((factories.post_visits.template(post_id=target_posts[i].id, day_date=dt.datetime.now()) for i in range(N)))
    noise_post_visits = factories.post_visits.create_batch((factories.post_visits.template(post_id=noise_posts[i].id) for i in range(N)))

    factories.post_edition.create_batch((factories.post_edition.template(
        user_id=user.id,
        post_id=target_posts[i].id,
        tags=[target_tags[i % 50],random.choice(target_tags)]) for i in range(N)))

    factories.post_edition.create_batch((factories.post_edition.template(
        user_id=user.id,
        post_id=noise_posts[i].id,
        tags=[random.choice(noise_tags),random.choice(noise_tags)]) for i in range(N)))

    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query_4_6.sql", task, dict(N=selected)))
        fetched_rows = result.fetchall()

        assert len(fetched_rows) == selected