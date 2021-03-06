# Ответы на домашнее задание к занятию "6.2. SQL"

## Введение

Перед выполнением задания вы можете ознакомиться с 
[дополнительными материалами](https://github.com/netology-code/virt-homeworks/tree/master/additional/README.md).

## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 12) c 2 volume, 
в который будут складываться данные БД и бэкапы.

Приведите получившуюся команду или docker-compose манифест.

## Ответ:

```text
version: "3.1"

services:
  pgdb_1:
    image: postgres:12
    restart: always
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=secret
      - ALLOW_EMPTY_PASSWORD=yes
    volumes:
      - "/home/dockeruser/docker/postgresql_1/data:/var/lib/postgresql/data"
      - "/home/dockeruser/docker/postgresql_1/backups:/var/lib/postgresql/backups"
    ports:
      - "5432:5432"
```

## Задача 2

В БД из задачи 1: 
- **создайте пользователя test-admin-user и БД test_db**

```text
CREATE USER "test-admin-user" WITH PASSWORD '123456';
CREATE DATABASE test_db;
\c test_db;
```

- **в БД test_db создайте таблицу orders и clients (спeцификация таблиц ниже)**

**Таблица orders:**
- id (serial primary key)
- наименование (string)
- цена (integer)

**Таблица clients:**
- id (serial primary key)
- фамилия (string)
- страна проживания (string, index)
- заказ (foreign key orders)

```text
create table orders (
    id integer PRIMARY KEY,
    name varchar(128),
    price numeric(10,2)
);

CREATE TABLE clients (
    id integer PRIMARY KEY,
    fio varchar(64),
    country varchar(64),
    order_id integer default null,
    FOREIGN KEY (order_id) REFERENCES orders (id)
);
```
По-хорошему, столбец country надо сделать integer и добавить внешний ключ на таблицу countries (id).
Таким образом, вынося повторяющиеся данные в отдельную таблицу, мы избавимся от избыточности данных.
Да и сами заказы хорошо бы вынести в отдельную таблицу, т.к. один клиент может несколько раз заказать один и тот же товар.
В этом случае в рамках этой задачи в таблице clients будет 2 записи на одно и тоже имя, что является не совсем правильным.

- **предоставьте привилегии на все операции пользователю test-admin-user на таблицы БД test_db**

```text
GRANT ALL PRIVILEGES ON orders TO "test-admin-user";
GRANT ALL PRIVILEGES ON clients TO "test-admin-user";
```

- **создайте пользователя test-simple-user и предоставьте пользователю test-simple-user права на SELECT/INSERT/UPDATE/DELETE данных таблиц БД test_db**

```text
CREATE USER "test-simple-user" WITH PASSWORD '123456';
GRANT SELECT, INSERT, UPDATE, DELETE ON orders TO "test-simple-user";
GRANT SELECT, INSERT, UPDATE, DELETE ON clients TO "test-simple-user";
```

Приведите:
- **итоговый список БД после выполнения пунктов выше**

```text
test_db-# \list
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 test_db   | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =Tc/postgres         +
           |          |          |            |            | postgres=CTc/postgres
(4 rows)

```

- **описание таблиц (describe)**

```text
test_db-# \d orders
                      Table "public.orders"
 Column |          Type          | Collation | Nullable | Default
--------+------------------------+-----------+----------+---------
 id     | integer                |           | not null |
 name   | character varying(128) |           |          |
 price  | numeric(10,2)          |           |          |
Indexes:
    "orders_pkey" PRIMARY KEY, btree (id)
Referenced by:
    TABLE "clients" CONSTRAINT "clients_order_id_fkey" FOREIGN KEY (order_id) REFERENCES orders(id)

test_db-# \d clients
                      Table "public.clients"
  Column  |         Type          | Collation | Nullable | Default
----------+-----------------------+-----------+----------+---------
 id       | integer               |           | not null |
 fio      | character varying(64) |           |          |
 country  | character varying(64) |           |          |
 order_id | integer               |           |          |
Indexes:
    "clients_pkey" PRIMARY KEY, btree (id)
Foreign-key constraints:
    "clients_order_id_fkey" FOREIGN KEY (order_id) REFERENCES orders(id)


```

- SQL-запрос для выдачи списка пользователей с правами над таблицами test_db
- список пользователей с правами над таблицами test_db

```text
 SELECT *
FROM information_schema.table_privileges
WHERE table_catalog = 'test_db'
  AND table_schema = 'public'
  AND grantee != 'postgres';

 grantor  |     grantee      | table_catalog | table_schema | table_name | privilege_type | is_grantable | with_hierarchy
----------+------------------+---------------+--------------+------------+----------------+--------------+----------------
 postgres | test-simple-user | test_db       | public       | orders     | INSERT         | NO           | NO
 postgres | test-simple-user | test_db       | public       | orders     | SELECT         | NO           | YES
 postgres | test-simple-user | test_db       | public       | orders     | UPDATE         | NO           | NO
 postgres | test-simple-user | test_db       | public       | orders     | DELETE         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | INSERT         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | SELECT         | NO           | YES
 postgres | test-admin-user  | test_db       | public       | orders     | UPDATE         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | DELETE         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | TRUNCATE       | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | REFERENCES     | NO           | NO
 postgres | test-admin-user  | test_db       | public       | orders     | TRIGGER        | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | INSERT         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | SELECT         | NO           | YES
 postgres | test-admin-user  | test_db       | public       | clients    | UPDATE         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | DELETE         | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | TRUNCATE       | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | REFERENCES     | NO           | NO
 postgres | test-admin-user  | test_db       | public       | clients    | TRIGGER        | NO           | NO
 postgres | test-simple-user | test_db       | public       | clients    | INSERT         | NO           | NO
 postgres | test-simple-user | test_db       | public       | clients    | SELECT         | NO           | YES
 postgres | test-simple-user | test_db       | public       | clients    | UPDATE         | NO           | NO
 postgres | test-simple-user | test_db       | public       | clients    | DELETE         | NO           | NO
(22 rows)

```

## Задача 3

Используя SQL синтаксис - наполните таблицы следующими тестовыми данными:

Таблица orders

|Наименование|цена|
|------------|----|
|Шоколад| 10 |
|Принтер| 3000 |
|Книга| 500 |
|Монитор| 7000|
|Гитара| 4000|

Таблица clients

|ФИО|Страна проживания|
|------------|----|
|Иванов Иван Иванович| USA |
|Петров Петр Петрович| Canada |
|Иоганн Себастьян Бах| Japan |
|Ронни Джеймс Дио| Russia|
|Ritchie Blackmore| Russia|

Используя SQL синтаксис:
- вычислите количество записей для каждой таблицы 
- приведите в ответе:
    - запросы 
    - результаты их выполнения.

## Ответ:

```text
INSERT INTO orders (id, name, price) VALUES
(1, 'Шоколад', 10),
(2, 'Принтер', 3000),
(3, 'Книга', 500),
(4, 'Монитор', 7000),
(5, 'Гитара', 4000);

INSERT INTO clients (id, fio, country) VALUES
(1, 'Иванов Иван Иванович', 'USA'),
(2, 'Петров Петр Петрович', 'Canada'),
(3, 'Иоганн Себастьян Бах', 'Japan'),
(4, 'Ронни Джеймс Дио', 'Russia'),
(5, 'Ritchie Blackmore', 'Russia');

test_db=# SELECT COUNT(*) FROM orders;
 count
-------
     5
(1 row)

test_db=# SELECT COUNT(*) FROM clients;
 count
-------
     5
(1 row)

```

## Задача 4

Часть пользователей из таблицы clients решили оформить заказы из таблицы orders.

Используя foreign keys свяжите записи из таблиц, согласно таблице:

|ФИО|Заказ|
|------------|----|
|Иванов Иван Иванович| Книга |
|Петров Петр Петрович| Монитор |
|Иоганн Себастьян Бах| Гитара |

**Приведите SQL-запросы для выполнения данных операций.**

```text
UPDATE clients SET order_id = 3 WHERE id = 1;
UPDATE clients SET order_id = 4 WHERE id = 2;
UPDATE clients SET order_id = 5 WHERE id = 3;
```

**Приведите SQL-запрос для выдачи всех пользователей, которые совершили заказ, а также вывод данного запроса.**
 
Подсказка - используйте директиву `UPDATE`.

```text
SELECT c.fio, o.name
FROM clients c
LEFT JOIN orders o ON c.order_id = o.id
WHERE o.id IS NOT NULL;

         fio          |  name
----------------------+---------
 Иванов Иван Иванович | Книга
 Петров Петр Петрович | Монитор
 Иоганн Себастьян Бах | Гитара
(3 rows)
```

## Задача 5

Получите полную информацию по выполнению запроса выдачи всех пользователей из задачи 4 
(используя директиву EXPLAIN).

**Приведите получившийся результат и объясните что значат полученные значения.**

```text
                               QUERY PLAN
-------------------------------------------------------------------------
 Hash Join  (cost=15.61..28.65 rows=239 width=420)
   Hash Cond: (c.order_id = o.id)
   ->  Seq Scan on clients c  (cost=0.00..12.40 rows=240 width=150)
   ->  Hash  (cost=12.50..12.50 rows=249 width=278)
         ->  Seq Scan on orders o  (cost=0.00..12.50 rows=249 width=278)
               Filter: (id IS NOT NULL)
(6 rows)
```

15.61. - Приблизительная стоимость запуска. Это время, которое проходит, прежде чем начнётся этап вывода данных.

.28.65 - Приблизительная общая стоимость. Она вычисляется в предположении, что узел плана выполняется до конца, то есть возвращает все доступные строки (кроме случаев с LIMIT).

rows=239 - Ожидаемое число строк, которое должен вывести этот узел плана.

width=420 - Ожидаемый средний размер строк, выводимых этим узлом плана (в байтах).

## Задача 6

Приведите список операций, который вы применяли для бэкапа данных и восстановления.

- **Создайте бэкап БД test_db и поместите его в volume, предназначенный для бэкапов (см. Задачу 1).**

```text
docker exec -t postgresql_1_pgdb_1_1 pg_dump -U postgres test_db -f /var/lib/postgresql/backups/backup_test.sql
```

- **Остановите контейнер с PostgreSQL (но не удаляйте volumes).**

```text
docker stop postgresql_1_pgdb_1_1
```

- **Поднимите новый пустой контейнер с PostgreSQL.**

```text
version: "3.1"
services:
  pgdb_2:
    image: postgres:12
    restart: always
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=secret
      - ALLOW_EMPTY_PASSWORD=yes
    volumes:
      - "/home/dockeruser/docker/postgresql_1/data2:/var/lib/postgresql/data"
      - "/home/dockeruser/docker/postgresql_1/backups:/var/lib/postgresql/backups"
    ports:
      - "5432:5432"

docker-compose up
```

- **Восстановите БД test_db в новом контейнере.**

```text
Сначала нужно создать БД и пользователя:
postgres=# \list
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
(3 rows)

postgres=# CREATE USER "test-admin-user" WITH PASSWORD '123456';
postgres=# CREATE DATABASE test_db;
postgres=# \c test_db;

Когда БД создана выполняем команду для восстановления данных
docker exec -i postgresql_1_pgdb_2_1 psql -U postgres -d test_db -f /var/lib/postgresql/backups/backup_test.sql
```


