def run_raw(engine, query: str):
    with engine.connect() as con:

        results = con.execute(rf"{query}")

        return results
