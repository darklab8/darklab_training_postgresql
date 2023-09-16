from python.utils.database.sql import Database
from ..task2.factories import TypeFactories
from ..task3.reusable_code import query, measure_time, Task

task = Task.task7

def test_task7_2(database: Database, apply_task2_migrations, factories: TypeFactories, task_7_setup):
    """
    Вывести суммарное накопленное количество тегов на текущую дату по каждому пользователю.
    """
    N, *_ = task_7_setup

    with database.get_core_session() as session:
        with measure_time(f"{N=}"):
            result = session.execute(query("query_7_2.sql", task))
            rows = result.fetchall()
            assert len(rows) > 0