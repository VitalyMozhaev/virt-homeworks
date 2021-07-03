# Ответы на домашнее задание к занятию "6.1. Типы и структура СУБД"

## Введение

Перед выполнением задания вы можете ознакомиться с 
[дополнительными материалами](https://github.com/netology-code/virt-homeworks/tree/master/additional/README.md).

## Задача 1

Архитектор ПО решил проконсультироваться у вас, какой тип БД 
лучше выбрать для хранения определенных данных.

Он вам предоставил следующие типы сущностей, которые нужно будет хранить в БД. Выберите подходящие типы СУБД для каждой сущности и объясните свой выбор.

- **Электронные чеки в json виде**

MongoDB прекрасный вариант для регистрации событий электронной коммерции и хранения данных в формате json.

- **Склады и автомобильные дороги для логистической компании**

Любая реляционная СУБД позволит организовать взаимосвязанные структуры для решения логистических задач, например, MySQL или PostgreSQL отлично подойдут.

- **Генеалогические деревья**

По моделе данных подходит иерархический тип СУБД, однако этот тип данных несколько эксклюзивен и не имеет широкого распространения, что скажется на разработке и поддержке, поэтому в принципе подойдёт любая реляционная СУБД, например, MySQL или PostgreSQL.

- **Кэш идентификаторов клиентов с ограниченным временем жизни для движка аутенфикации**

Подойдёт любая СУБД на основе "ключ - значение", например, Memcahed или Redis, обе СУБД имеют Time-To-Live, размещают данные в оперативной памяти, это значительно ускоряет работу с данными, что и нужно для кэш. Но именно для этой задачи больше подойдёт Redis, т.к. он имеет возможность синхронизировать данные на диск и при сбое или перезапуске большая часть данных будет сохранена и пользователи продолжат работу без лишней авторизации.

- **Отношения клиент-покупка для интернет-магазина**

Подойдёт любая СУБД на основе "ключ - значение", в данном случае отличным решением будет Redis, т.к. он хранит данные в оперативной памяти и имеет возможность синхронизировать данные на диск.

## Задача 2

Вы создали распределенное высоконагруженное приложение и хотите классифицировать его согласно 
CAP-теореме. Какой классификации по CAP-теореме соответствует ваша система?
А согласно PACELC-теореме, как бы вы классифицировали данные реализации?
(каждый пункт - это отдельная реализация вашей системы и для каждого пункта надо привести классификацию):

- **Данные записываются на все узлы с задержкой до часа (асинхронная запись)**

CAP: AP (достигают "конечной согласованности" за счет репликации и проверки)  
PACELC: PC/EC

- **При сетевых сбоях, система может разделиться на 2 раздельных кластера**

CAP: CA (имеют проблемы с разделами, и работают с репликацией)  
PACELC: PA/EL

- **Система может не прислать корректный ответ или сбросить соединение**

CAP: CP (имеют проблемы с доступностью при сохранении согласованности данных между разделенными узлами)  
PACELC: PC/EC

## Задача 3

Могут ли в одной системе сочетаться принципы BASE и ACID? Почему?

## Ответ:

По своему подходу ACID и BASE кажутся противоположными, но на самом деле это разные варианты трех характеристик CAP-теоремы. При проектировании распределенных систем её компоненты имеют разные требования к согласованности, поэтому ACID и BASE в зависимости от задач могут и будут использоваться в той или иной комбинации.

## Задача 4

Вам дали задачу написать системное решение, основой которого бы послужили:

- фиксация некоторых значений с временем жизни
- реакция на истечение таймаута

Вы слышали о key-value хранилище, которое имеет механизм [Pub/Sub](https://habr.com/ru/post/278237/). 
Что это за система? Какие минусы выбора данной системы?

## Ответ:

Key-value хранилище с механизмом Pub/Sub - это Redis.

Минусы Redis:

- ограничение по оперативной памяти
- шардинг приводит к задержкам
- при сбое в работе сервера или перезагрузке часть несинхронизированных на диск данных будет потеряна
- это NoSQL, т.е. никакого языка SQL
- нет сегментации на пользователей или группы пользователей, доступ по общему паролю.
