from pathlib import Path

class Migrations:
    task_2_1 = Path(__file__).parent / "migrations" / "task2_1.sql"
    disable_triggers = Path(__file__).parent / "migrations" / "task2_2_disable_triggers_for_tests.sql"
