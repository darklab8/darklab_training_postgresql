from sqlalchemy import create_engine

engine = create_engine("postgresql://postgres:postgres@localhost:5432/postgres", echo=True)

with engine.connect() as con:

    rs = con.execute('SELECT * FROM posts LIMIT 10')
    for r in rs:  
        print(r)