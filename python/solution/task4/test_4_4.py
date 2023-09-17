
from python.utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from ..task3.reusable_code import query, measure_time, Task
import random
import datetime as dt

task = Task.task4

def test_task4_4(database: Database, apply_task2_migrations: None, factories: TypeFactories) -> None:
    "4. Найти N пользователей, для которых суммарный рейтинг для всех созданных ими постов максимальный среди всех пользователей младше K лет."

    N=100
    selected = 50
    K = 10

    # targets of query
    target_users = list(factories.user.create_batch((factories.user.template(id=i) for i in range(N))))
    young_users= list(factories.user.create_batch((factories.user.template(id=i+N+1, birth_date=dt.datetime.now()-dt.timedelta(days=365*10)) for i in range(N))))

    # target for query
    approved_posts = factories.post.create_batch((factories.post.template(author_id=random.choice(target_users).id, rating=random.randint(0,1000)) for i in range(int(N*2))))
    # noise data
    other_posts = factories.post.create_batch((factories.post.template(author_id=random.choice(young_users).id, rating=random.randint(0,1000)) for i in range(int(N*2))))
    
    # target for query
    post_approvals = factories.post_approval.create_batch((
        factories.post_approval.template(
            user_id=young_users[i].id,
            post_id=approved_posts[i].id,
        ) for i in range(N)))

    # for noisy data
    post_approvals = factories.post_approval.create_batch((
        factories.post_approval.template(
            user_id=young_users[i].id,
            post_id=other_posts[i].id,
        ) for i in range(N)))

    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query_4_4.sql", task, dict(N=selected, K=K)))
        fetched_rows = result.fetchall()

        assert len(fetched_rows) == selected
