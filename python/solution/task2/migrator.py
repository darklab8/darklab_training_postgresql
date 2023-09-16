from python.shared.settings import sql_folder

class Migrations:
    task_2_1 = sql_folder / "task2" /"migrations" / "task2_1.sql"
    disable_triggers = sql_folder / "task2" / "migrations" / "task2_2_disable_triggers_for_tests.sql"
