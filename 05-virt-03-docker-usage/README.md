# Ответы на домашнее задание к занятию "5.3. Контейнеризация на примере Docker"

## Задача 1

Посмотрите на сценарии ниже и ответьте на вопрос: "Подходит ли в этом сценарии использование докера? Или лучше подойдет виртуальная машина, физическая машина? Или возможны разные варианты?"
Детально опишите и обоснуйте свой выбор.

Сценарии:

- **Высоконагруженное монолитное java веб-приложение;**

Для размещения монолитного приложение в контейнерах, его необходимо переписать, разделив на микросервисы. Целиком размещать монолит в контейнере не разумно, особенно учитывая высокую нагрузку. А для обеспечения высокой нагрузки и отказоустойчивости, лучше размещать монолитное приложение на физическом сервере. Это позволит экономить ресурсы сервера и исключить дополнительную точку отказа в виде гипервизора.

- **Go-микросервис для генерации отчетов;**

Микросервисы - это отличное применение для docker контейнеров. Именно такие маленькие сервисы удобно крутить в контейнерах, их удобно обновлять и перезапускать, при этом остальные сервисы работают в штатном режиме.

- **Nodejs веб-приложение;**

В принципе, можно запускать и в контейнере, можно и на отдельной виртуалке. 

- **Мобильное приложение c версиями для Android и iOS;**

Docker не имеет графической оболочки, поэтому для подобных приложений удобно использовать виртуалку.

- **База данных postgresql используемая, как кэш;**

Сами контейнеры не подразумевают хранение информации (состояние stateless), поэтому базы данных лучше разворачивать на сервере или виртуалке.

- **Шина данных на базе Apache Kafka;**

Примерно понимая специфику работы Apache Kafka, полагаю самым удачным решением будет размещение на виртуальной машине, так можно будет масштабировать и реплицировать данные.

- **Очередь для Logstash на базе Redis;**

Учитывая, что Redis работает в оперативной памяти, контейнер лучше виртуалки, т.к. все процессы работают напрямую с физической памятью.

- **Elastic stack для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana;**

Средства оркестрации, управляющие docker контейнерами смогут реализовать эту задачу, образы в докерхаб есть.

- **Мониторинг-стек на базе prometheus и grafana;**

Удобнее и быстрее разворачивать в docker контейнере, по сравнению с настройкой виртуальной машины. Настроенный образ легко переносить и масштабировать.

- **Mongodb, как основное хранилище данных для java-приложения;**

Базы данных лучше ставить на отдельный физический сервер, однако, отдельный сервер под БД не всегда рентабельно закупать, поэтому если java-приложение не будет иметь высокой нагрузки удобнее и дешевле БД установить на виртуальной машине. Кроме того, надо помнить и об облачных технологиях, позволяющих в несколько кликов разворачивать виртуалки.

- **Jenkins-сервер.**

Очень удобно и быстро настроить в контейнере и, пробросив настройки и нужные данные через volumes, запускать и управлять сервером.


## Задача 2 

Сценарий выполения задачи:

- создайте свой репозиторий на докерхаб; 
- выберете любой образ, который содержит апачи веб-сервер;
- создайте свой форк образа;
- реализуйте функциональность: 
запуск веб-сервера в фоне с индекс-страницей, содержащей HTML-код ниже: 
```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m kinda DevOps now</h1>
</body>
</html>
```
Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на докерхаб-репо.

## Ответ:

https://hub.docker.com/r/vitalymozhaev/netology-docker-usage

Результат работы контейнера:

![Результат работы контейнера](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/05-virt-03-docker-usage/%D0%97%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%202.png)

## Задача 3 

- Запустите первый контейнер из образа centos c любым тэгом в фоновом режиме, подключив папку info из текущей рабочей директории на хостовой машине в /share/info контейнера;
- Запустите второй контейнер из образа debian:latest в фоновом режиме, подключив папку info из текущей рабочей директории на хостовой машине в /info контейнера;
- Подключитесь к первому контейнеру с помощью exec и создайте текстовый файл любого содержания в /share/info ;
- Добавьте еще один файл в папку info на хостовой машине;
- Подключитесь во второй контейнер и отобразите листинг и содержание файлов в /info контейнера.

## Ответ:

Работающие контейнеры:

![Работающие контейнеры](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/05-virt-03-docker-usage/%D0%97%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%203_1.png)

Результат выполнения:

![Результат работы контейнера](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/05-virt-03-docker-usage/%D0%97%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%203_2.png)

