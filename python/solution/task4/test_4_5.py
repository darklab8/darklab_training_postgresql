
from python.utils.database.sql import Database
from ..task2.factories import TypeFactories, UserTemplate
from ..task3.reusable_code import query, measure_time
import random

def test_task4_5(database: Database, apply_task2_migrations, factories: TypeFactories):
    "5. Найти N пользователей с наибольшим рейтингом."

    N=100
    selected = 50

    list(factories.user.create_batch([factories.user.template(id=i, rating=random.randint(0,1000)) for i in range(N)]))

    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query_4_5.sql", dict(N=selected), root=__file__))
        fetched_rows = result.fetchall()

        assert len(fetched_rows) == selected