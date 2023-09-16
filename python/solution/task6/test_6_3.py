from python.utils.database.sql import Database
from ..task2.factories import TypeFactories
from ..task3.reusable_code import query, measure_time

def test_task6_3(database: Database, apply_task2_migrations, factories: TypeFactories):
    """
    Добавить следующие столбцы к таблицам, используя alter table:
    Статус пользователей. Статус может иметь значения активен или заблокирован. Новый столбец должен быть заполнен значением активен.
    """

    with database.get_core_session() as session:
        with measure_time(f""):
            result = session.execute(query("query_6_3.sql", dict(), root=__file__))

