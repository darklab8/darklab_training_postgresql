# Task # 6
## Requirements

Задание 6

Написать следующие запросы к бд, используя транзакции:

1. Копировать пост по id вместе со связанными авторами и тэгами, но без статистики, комментариев и рейтинга. Скопированный пост должен быть в статусе черновик.

2. Удалить всех пользователей, у который рейтинг меньше чем N , вместе со всеми постами и комментариями. Порядок удаления сущностей: комментарии к постам пользователя, комментарии пользователя, посты пользователя, пользователь.

Добавить следующие столбцы к таблицам, используя alter table:

1. Статус пользователей. Статус может иметь значения активен или заблокирован. Новый столбец должен быть заполнен значением активен.
2. Даты создания и обновления для постов. Новые столбцы должны быть заполнены значением текущей даты.



## Task 6_1 Transaction: Копировать пост по id вместе со связанными авторами и тэгами, но без статистики, комментариев и рейтинга. Скопированный пост должен быть в статусе черновик

```sql
--8<-- "sql/task6/queries/query_6_1.sql"
```
```py
--8<-- "python/solution/task6/test_6_1.py"
```
## Task 6_2 2. Transaction: Удалить всех пользователей, у который рейтинг меньше чем N , вместе со всеми постами и комментариями. Порядок удаления сущностей: комментарии к постам пользователя, комментарии пользователя, посты пользователя, пользователь.

```sql
--8<-- "sql/task6/queries/query_6_2.sql"
```
```py
--8<-- "python/solution/task6/test_6_2.py"
```

## Task 6_3 Alter table to add column: Статус пользователей. Статус может иметь значения активен или заблокирован. Новый столбец должен быть заполнен значением активен.

```sql
--8<-- "sql/task6/queries/query_6_3.sql"
```
```py
--8<-- "python/solution/task6/test_6_3.py"
```

## Task 6_4 Alter table to add column: Даты создания и обновления для постов. Новые столбцы должны быть заполнены значением текущей даты.

```sql
--8<-- "sql/task6/queries/query_6_4.sql"
```
```py
--8<-- "python/solution/task6/test_6_4.py"
```
