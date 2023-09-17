
from python.utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from ..task3.reusable_code import query, measure_time, Task
import random

task = Task.task4

def test_task4_3(database: Database, apply_task2_migrations: None, factories: TypeFactories) -> None:
    "3. Найти N пользователей, для которых суммарный рейтинг для всех созданных ими постов максимальный среди всех пользователей."

    N=100
    selected = 50

    users = list(factories.user.create_batch((factories.user.template(id=i) for i in range(N))))
    factories.post.create_batch((factories.post.template(author_id=random.choice(users).id, rating=random.randint(0,1000)) for i in range(int(N*2))))

    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query_4_3.sql", task, dict(N=selected)))
        fetched_rows = result.fetchall()

        assert len(fetched_rows) == selected
