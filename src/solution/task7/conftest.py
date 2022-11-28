from utils.database.sql import Database
from ..task2.factories import TypeFactories
import random
import pytest

@pytest.fixture
def task_7_setup(database: Database, apply_task2_migrations, factories: TypeFactories):
    N=50
    tag_generator = lambda: random.choices(
        list([f"a{i}" for i in range(100)]),k=random.randint(1,16)
    )

    users = list(factories.user.create_batch([factories.user.template() for i in range(N)]))
    posts = list(factories.post.create_batch([factories.post.template(
            author_id=random.choice(users).id,
         ) for i in range(N)]))

    factories.post_edition.create_batch((factories.post_edition.template(
        post_id=posts[i].id,
        user_id=random.choice(users).id,
        tags=tag_generator(),
    ) for i in range(N)))

    factories.post_edition.create_batch((factories.post_edition.template(
        post_id=random.choice(posts).id,
        user_id=random.choice(users).id,
        tags=tag_generator(),
    ) for i in range(50)))

    return (N,)