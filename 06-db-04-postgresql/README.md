# Домашнее задание к занятию "6.4. PostgreSQL"

## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 13). Данные БД сохраните в volume.

```text
version: "3.1"

services:
  pgdb_13:
    container_name: netology_psql_13
    image: postgres:13
    restart: always
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=secret
      - ALLOW_EMPTY_PASSWORD=yes
    volumes:
      - "/home/dockeruser/docker/postgres_13/data:/var/lib/postgresql/data"
      - "/home/dockeruser/docker/postgres_13/backups:/var/lib/postgresql/backups"
    ports:
      - "5432:5432"
```

Подключитесь к БД PostgreSQL используя `psql`.

`docker exec -it netology_psql_13 psql -U postgres`

Воспользуйтесь командой `\?` для вывода подсказки по имеющимся в `psql` управляющим командам.

**Найдите и приведите** управляющие команды для:
- вывода списка БД

`\l[+]   [PATTERN]      list databases`

- подключения к БД

`\c[onnect] {[DBNAME|- USER|- HOST|- PORT|-] | conninfo}     connect to new database`

- вывода списка таблиц

`\dt[S+] [PATTERN]      list tables`

- вывода описания содержимого таблиц

`\d[S+]  NAME           describe table, view, sequence, or index`

- выхода из psql

`\q                     quit psql`

## Задача 2

Используя `psql` создайте БД `test_database`.

`CREATE DATABASE test_database;`

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-04-postgresql/test_data) и восстановите бэкап БД в `test_database`.

```text
docker exec -i netology_psql_13 psql -U postgres -d test_database -f /var/lib/postgresql/backups/test_dump.sql
SET
SET
SET
SET
SET
 set_config
------------

(1 row)

SET
SET
SET
SET
SET
SET
CREATE TABLE
ALTER TABLE
CREATE SEQUENCE
ALTER TABLE
ALTER SEQUENCE
ALTER TABLE
COPY 8
 setval
--------
      8
(1 row)

ALTER TABLE
```

Перейдите в управляющую консоль `psql` внутри контейнера.

`docker exec -it netology_psql_13 psql -U postgres`

Подключитесь к восстановленной БД и проведите операцию ANALYZE для сбора статистики по таблице.

```text
postgres=# \c test_database;
You are now connected to database "test_database" as user "postgres".
test_database=# \dt
         List of relations
 Schema |  Name  | Type  |  Owner
--------+--------+-------+----------
 public | orders | table | postgres
(1 row)

test_database=# ANALYZE orders;
ANALYZE
```

Используя таблицу [pg_stats](https://postgrespro.ru/docs/postgresql/12/view-pg-stats), найдите столбец таблицы `orders` 
с наибольшим средним значением размера элементов в байтах. **Приведите в ответе** команду, которую вы использовали для вычисления и полученный результат.

```text
test_database=# SELECT MAX(avg_width) max_avg_width FROM pg_stats WHERE tablename = 'orders';
 max_avg_width
---------------
            16
(1 row)
```

## Задача 3

Архитектор и администратор БД выяснили, что ваша таблица orders разрослась до невиданных размеров и
поиск по ней занимает долгое время. Вам, как успешному выпускнику курсов DevOps в нетологии предложили
провести разбиение таблицы на 2 (шардировать на orders_1 - price>499 и orders_2 - price<=499).

**Предложите SQL-транзакцию для проведения данной операции.**


Можно было бы использовать секционирование с использованием наследования:

```text
CREATE TABLE orders_1 (CHECK (price > 499)) INHERITS (orders);
CREATE TABLE orders_2 (CHECK (price <= 499)) INHERITS (orders);
```

Но в данном случае, насколько я понял, нужно дописывать триггерную функцию для распределения значений по таблицам.
Удобнее пересоздать таблицу с поддержкой секционирования, используя декларативное секционирование, и заполнить её существующими данными.

```text
BEGIN TRANSACTION;

CREATE TABLE public.orders_main (
    id integer NOT NULL,
    title character varying(80) NOT NULL,
    price integer DEFAULT 0
) PARTITION BY RANGE(price);

CREATE TABLE orders_1 PARTITION OF orders_main FOR VALUES FROM (500) TO (MAXVALUE);
CREATE TABLE orders_2 PARTITION OF orders_main FOR VALUES FROM (MINVALUE) TO (500);

INSERT INTO orders_main SELECT * FROM orders;

COMMIT;

test_database=# SELECT * FROM orders_main;
 id |        title         | price
----+----------------------+-------
  1 | War and peace        |   100
  3 | Adventure psql time  |   300
  4 | Server gravity falls |   300
  5 | Log gossips          |   123
  7 | Me and my bash-pet   |   499
  2 | My little database   |   500
  6 | WAL never lies       |   900
  8 | Dbiezdmin            |   501
(8 rows)

test_database=# SELECT * FROM orders_1;
 id |       title        | price
----+--------------------+-------
  2 | My little database |   500
  6 | WAL never lies     |   900
  8 | Dbiezdmin          |   501
(3 rows)

test_database=# SELECT * FROM orders_2;
 id |        title         | price
----+----------------------+-------
  1 | War and peace        |   100
  3 | Adventure psql time  |   300
  4 | Server gravity falls |   300
  5 | Log gossips          |   123
  7 | Me and my bash-pet   |   499
(5 rows)
```

Можно ли было изначально исключить "ручное" разбиение при проектировании таблицы orders?

`Невозможно превратить обычную таблицу в партиционированную или наоборот, поэтому при проектировании таблицы "orders" нужно было создать таблицу с поддержкой секционирования.
Это позволило бы секционировать уже существующую таблицу с данными.`

## Задача 4

Используя утилиту `pg_dump` создайте бекап БД `test_database`.

`docker exec -t netology_psql_13 pg_dump -U postgres test_database -f /var/lib/postgresql/backups/dump_test_database.sql`

Как бы вы доработали бэкап-файл, чтобы добавить уникальность значения столбца `title` для таблиц `test_database`?

```text
По данному вопросу нашёл 2 решения:

1. Ограничения уникальности (а значит и первичные ключи) в секционированных таблицах должны включать все
столбцы ключа разбиения.

CREATE UNIQUE INDEX title_unique ON orders_main (title, price);

В этом случае, уникальными должны быть пара значений - title и price.
Т.е., из нашей таблички orders_main строка "War and peace - 100" и строка "War and peace - 101" это уникальные значения.


2. Создать ограничение-исключение, охватывающее всю секционированную таблицу, нельзя; можно только поместить такое
ограничение в каждую отдельную секцию с данными. И это также является следствием того, что установить ограничения,
действующие между секциями, невозможно.

Создаём ограничение уникальности в каждой секции:

CREATE UNIQUE INDEX title_unique_1 ON orders_1 (title);
CREATE UNIQUE INDEX title_unique_2 ON orders_2 (title);

В этом же случае уникальность будет соблюдаться внутри каждой секции, но не между секциями.
Т.е., в нашей секционированной табличке orders_main уже есть строка "War and peace - 100", которая хранится в секции orders_2,
если мы попытаемся добавить строку "War and peace - 200", то получим ошибку уникального значения поля title (War and peace).
Однако, если мы добавим строку "War and peace - 777", то эта строка успешко добавится в табличку orders_main и будет размещена
в секции orders_1, где значения title="War and peace" ещё нет.

Тут уж, как говорится, каждый решает сам, что важнее в рамках общей задачи хранения данных.
```
