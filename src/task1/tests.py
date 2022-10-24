def test_check(database):
    with database.get_core_session() as session:
        session.execute("SELECT 1;")
