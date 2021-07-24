# Ответы на домашнее задание к занятию "6.5. Elasticsearch"

## Задача 1

В этом задании вы потренируетесь в:
- установке elasticsearch
- первоначальном конфигурировании elastcisearch
- запуске elasticsearch в docker

Используя докер образ [centos:7](https://hub.docker.com/_/centos) как базовый и 
[документацию по установке и запуску Elastcisearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/targz.html):

- составьте Dockerfile-манифест для elasticsearch
- соберите docker-образ и сделайте `push` в ваш docker.io репозиторий
- запустите контейнер из получившегося образа и выполните запрос пути `/` c хост-машины

Требования к `elasticsearch.yml`:
- данные `path` должны сохраняться в `/var/lib`
- имя ноды должно быть `netology_test`

В ответе приведите:
- текст Dockerfile манифеста
- ссылку на образ в репозитории dockerhub
- ответ `elasticsearch` на запрос пути `/` в json виде

Подсказки:
- возможно вам понадобится установка пакета perl-Digest-SHA для корректной работы пакета shasum
- при сетевых проблемах внимательно изучите кластерные и сетевые настройки в elasticsearch.yml
- при некоторых проблемах вам поможет docker директива ulimit
- elasticsearch в логах обычно описывает проблему и пути ее решения

Далее мы будем работать с данным экземпляром elasticsearch.

## Ответ:

- Dockerfile:

```text
FROM centos:7

ADD https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.13.4-linux-x86_64.tar.gz /
ADD https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.13.4-linux-x86_64.tar.gz.sha512 /

RUN yum update -y && \
    yum install perl-Digest-SHA -y && \
    shasum -a 512 -c elasticsearch-7.13.4-linux-x86_64.tar.gz.sha512 && \
    tar -xzf elasticsearch-7.13.4-linux-x86_64.tar.gz && \
    cd elasticsearch-7.13.4/ && \
    useradd elasticuser && \
    chown -R elasticuser:elasticuser /elasticsearch-7.13.4/ && \
    rm -fr /elasticsearch-7.13.4-linux-x86_64.tar.gz.sha512 /elasticsearch-7.13.4-linux-x86_64.tar.gz

RUN mkdir /var/lib/{data,logs} && \
    chown -R elasticuser:elasticuser /var/lib/data && \
    chown -R elasticuser:elasticuser /var/lib/logs

WORKDIR /elasticsearch-7.13.4
RUN mkdir snapshots && \
    chown -R elasticuser:elasticuser snapshots

ADD elasticsearch.yml /elasticsearch-7.13.4/config/
RUN chown -R elasticuser:elasticuser /elasticsearch-7.13.4/config

USER elasticuser

EXPOSE 9200 9300

CMD ["./bin/elasticsearch", "-Ecluster.name=netology_cluster", "-Enode.name=netology_test"]

```

Для запуска из Dockerfile:

```text
docker build -t netology-elasticsearch:1.0 .
docker run -ti --name netology-elasticsearch-1 -p 9200:9200 netology-elasticsearch:1.0
Если возникнет проблема с ограничением, выполняем: sudo sysctl -w vm.max_map_count=262144
```

- Файл [elasticsearch.yml](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/06-db-05-elasticsearch/elasticsearch.yml)

- Образ в репозитории:

https://hub.docker.com/r/vitalymozhaev/netology-elasticsearch

Запуск из образа в репозитории:

`docker run -di --name netology-elasticsearch-1 -p 9200:9200 -p 9300:9300 vitalymozhaev/netology-elasticsearch:1`

- ответ `elasticsearch` на запрос пути `/` в json виде:

```text
curl localhost:9200/
{
  "name" : "netology_test",
  "cluster_name" : "netology_cluster",
  "cluster_uuid" : "_na_",
  "version" : {
    "number" : "7.13.4",
    "build_flavor" : "default",
    "build_type" : "tar",
    "build_hash" : "c5f60e894ca0c61cdbae4f5a686d9f08bcefc942",
    "build_date" : "2021-07-14T18:33:36.673943207Z",
    "build_snapshot" : false,
    "lucene_version" : "8.8.2",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

## Задача 2

В этом задании вы научитесь:
- создавать и удалять индексы
- изучать состояние кластера
- обосновывать причину деградации доступности данных

Ознакомьтесь с [документацией](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html) 
и добавьте в `elasticsearch` 3 индекса, в соответствии с таблицей:

| Имя | Количество реплик | Количество шард |
|-----|-------------------|-----------------|
| ind-1 | 0 | 1 |
| ind-2 | 1 | 2 |
| ind-3 | 2 | 4 |

```text
curl -X PUT "localhost:9200/ind-1?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 1,  
      "number_of_replicas": 0 
    }
  }
}
'

curl -X PUT "localhost:9200/ind-2?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 2,  
      "number_of_replicas": 1 
    }
  }
}
'

curl -X PUT "localhost:9200/ind-3?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 4,  
      "number_of_replicas": 2 
    }
  }
}
'
```

Получите список индексов и их статусов, используя API и **приведите в ответе** на задание.

```text
curl -X GET "localhost:9200/_cat/indices?v"
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   ind-1 VQJbhphLRti8Kgi6GUvGbg   1   0          0            0       208b           208b
yellow open   ind-3 5IUbzftCRc6kVJCleqFuvQ   4   2          0            0       832b           832b
yellow open   ind-2 mY99_wMuR8Ka0fI8xTRJjw   2   1          0            0       416b           416b
```

Получите состояние кластера `elasticsearch`, используя API.

```text
curl -X GET "localhost:9200/_cluster/health?pretty"
{
  "cluster_name" : "netology_cluster",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 7,
  "active_shards" : 7,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 41.17647058823529
}
```

Как вы думаете, почему часть индексов и кластер находится в состоянии yellow?

```text
Состояние Yellow говорит о том, что у индексов ind-2 и ind-3 данные должны быть реплицированы на
другой или другие узелы, т.к. указаны реплики > 0, а их не существует. Это и "напрягает" Elasticsearch,
поэтому он указывает состояние этих индексов и статус кластера в "yellow"
```

Удалите все индексы.

```text
curl -X DELETE "localhost:9200/_all"
{"acknowledged":true}
```

**Важно**

При проектировании кластера elasticsearch нужно корректно рассчитывать количество реплик и шард,
иначе возможна потеря данных индексов, вплоть до полной, при деградации системы.

## Задача 3

В данном задании вы научитесь:
- создавать бэкапы данных
- восстанавливать индексы из бэкапов

Создайте директорию `{путь до корневой директории с elasticsearch в образе}/snapshots`.

`См. Dockerfile в задании 1`

Используя API [зарегистрируйте](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-register-repository.html#snapshots-register-repository) 
данную директорию как `snapshot repository` c именем `netology_backup`.

**Приведите в ответе** запрос API и результат вызова API для создания репозитория.

```text
curl -X PUT "localhost:9200/_snapshot/netology_backup?pretty" -H 'Content-Type: application/json' -d'
{
  "type": "fs",
  "settings": {
    "location": "/elasticsearch-7.13.4/snapshots"
  }
}
'
{
  "acknowledged" : true
}
```

Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.

```text
curl -X PUT "localhost:9200/test?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 1,  
      "number_of_replicas": 0 
    }
  }
}
'
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "test"
}

curl -X GET "localhost:9200/_cat/indices?v"
health status index uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test  GpQNGNHFS4eTH820S-4rKg   1   0          0            0       208b           208b
```

[Создайте `snapshot`](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-take-snapshot.html) 
состояния кластера `elasticsearch`.

```text
curl -X PUT "localhost:9200/_snapshot/netology_backup/snapshot_1?wait_for_completion=true&pretty"
{
  "snapshot" : {
    "snapshot" : "snapshot_1",
    "uuid" : "89w5NaeARHqqT4tIgDKyCg",
    "version_id" : 7130499,
    "version" : "7.13.4",
    "indices" : [
      "test"
    ],
    "data_streams" : [ ],
    "include_global_state" : true,
    "state" : "SUCCESS",
    "start_time" : "2021-07-24T21:02:14.002Z",
    "start_time_in_millis" : 1627160534002,
    "end_time" : "2021-07-24T21:02:14.002Z",
    "end_time_in_millis" : 1627160534002,
    "duration_in_millis" : 0,
    "failures" : [ ],
    "shards" : {
      "total" : 1,
      "failed" : 0,
      "successful" : 1
    },
    "feature_states" : [ ]
  }
}
```

**Приведите в ответе** список файлов в директории со `snapshot`ами.

```text
docker exec -ti netology-elasticsearch-2 ls /elasticsearch-7.13.4/snapshots
index-0  index.latest  indices  meta-89w5NaeARHqqT4tIgDKyCg.dat  my_backup_location  snap-89w5NaeARHqqT4tIgDKyCg.dat
```

Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

```text
curl -X DELETE "localhost:9200/test"
{"acknowledged":true}

curl -X PUT "localhost:9200/test-2?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 1,  
      "number_of_replicas": 0 
    }
  }
}
'
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "test-2"
}

curl -X GET "localhost:9200/_cat/indices?v"
health status index  uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test-2 Zles3WqlTCSu1BP7NJWJtw   1   0          0            0       208b           208b
```

[Восстановите](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-restore-snapshot.html) состояние
кластера `elasticsearch` из `snapshot`, созданного ранее. 

**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

```text
curl -X POST "localhost:9200/_snapshot/netology_backup/snapshot_1/_restore?pretty"
{
  "accepted" : true
}

curl -X GET "localhost:9200/_cat/indices?v"
health status index  uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   test-2 Zles3WqlTCSu1BP7NJWJtw   1   0          0            0       208b           208b
green  open   test   BYO2qki9RNeWMge8LHXg-A   1   0          0            0       208b           208b
```

Подсказки:
- возможно вам понадобится доработать `elasticsearch.yml` в части директивы `path.repo` и перезапустить `elasticsearch`

