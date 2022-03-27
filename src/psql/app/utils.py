from sqlalchemy.sql import text


def run_query(engine, query, args={}, return_first=False):
    with engine.connect() as con:
        results = con.execute(text(query).bindparams(**args))

    if return_first:
        first_result, *_ = results.first()
        return first_result

    return results
