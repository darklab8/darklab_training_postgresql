from python.utils.database.sql import Database
from ..task2.factories import TypeFactories
from ..task3.reusable_code import query, measure_time, Task
from sqlalchemy import text

task = Task.task6

def test_task6_4(database: Database, apply_task2_migrations: None, factories: TypeFactories) -> None:
    """
    Добавить следующие столбцы к таблицам, используя alter table:
    Статус пользователей. Статус может иметь значения активен или заблокирован. Новый столбец должен быть заполнен значением активен.
    """

    with database.get_core_session() as session:
        with measure_time(f""):
            result = session.execute(text("""
ALTER TABLE post
	DROP COLUMN created_at;
            """))
        session.commit()

    with database.get_core_session() as session:
        with measure_time(f""):
            result = session.execute(query("query_6_4.sql", task))
