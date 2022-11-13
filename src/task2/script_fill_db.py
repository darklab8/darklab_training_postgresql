"""
10. Написать скрипт для заполнения таблицы тестовыми данными.
В базе данных должно быть не менее 100000 постов для разных пользователей;
Для генерации тестовых данных можно воспользоваться функцией generate_series;

P.S. При создании 10000 юзеров с 50 постами каждого, скрипт выполняется за 14 секунд
"""
from sqlalchemy.ext.automap import automap_base
import random
from utils.database.sql import Database

def random_DATE():
    return f"20{random.randint(10,22):02}-{random.randint(1,12):02}-{random.randint(1,28):02}"

def script_fill_db(database: Database):
    Base = automap_base()
    Base.prepare(database.engine, reflect=True)
    class Consts:
        users_total_amount: int = 100
        users_per_post: int = 50
        posts_total_amount: int = users_total_amount * users_per_post

    with database.get_core_session() as session:
        
        User = Base.classes.user_
        session.bulk_save_objects(
            [
                User(
                    id=i,
                    first_name=f"name_{i}",
                    second_name=f"second_name_{i}",
                    birth_date=f"{random_DATE()}",
                    email=f"email_{i}",
                    password=f"password_{i}",
                    address=f"address_{i}",
                )
                for i in range(Consts.users_total_amount)
            ]
        )

        Post = Base.classes.post
        session.bulk_save_objects(
            [
                Post(
                    id=i,
                    author_id=i % Consts.users_total_amount,
                    title=f"title_{i}",
                    content=f"content_{i}",
                    created_at=f"{random_DATE()}",
                    status=random.choice(["draft", "published", "archived"]),
                    tags=[random.choice(["abc", "def", "ghi"]), random.choice(["jkl", "mno", "pqr"])],
                )
                for i in range(Consts.posts_total_amount)
            ]
        )

        session.commit()