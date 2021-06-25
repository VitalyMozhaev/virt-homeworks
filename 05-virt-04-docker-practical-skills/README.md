# Домашнее задание к занятию "5.4. Практические навыки работы с Docker"

## Задача 1 

В данном задании вы научитесь изменять существующие Dockerfile, адаптируя их под нужный инфраструктурный стек.

Измените базовый образ предложенного Dockerfile на Arch Linux c сохранением его функциональности.

```text
FROM ubuntu:latest

RUN apt-get update && \
    apt-get install -y software-properties-common && \
    add-apt-repository ppa:vincent-c/ponysay && \
    apt-get update
 
RUN apt-get install -y ponysay

ENTRYPOINT ["/usr/bin/ponysay"]
CMD ["Hey, netology”]
```

Для получения зачета, вам необходимо предоставить:
- Написанный вами Dockerfile
- Скриншот вывода командной строки после запуска контейнера из вашего базового образа
- Ссылку на образ в вашем хранилище docker-hub

## Ответ:

- Dockerfile:
```text
FROM archlinux

RUN yes | pacman -Syu && \
    yes | pacman -S ponysay

ENTRYPOINT ["/usr/bin/ponysay"]
CMD ["Hey, netology”]
```

- Скрин вывода командной строки

![Скрин вывода командной строки](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/05-virt-04-docker-practical-skills/%D0%97%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%201.png)

- Ссылка на образ: https://hub.docker.com/r/vitalymozhaev/netology-arch-ponysay


## Задача 2 

В данной задаче вы составите несколько разных Dockerfile для проекта Jenkins, опубликуем образ в `dockerhub.io` и посмотрим логи этих контейнеров.

- Составьте 2 Dockerfile:

    - Общие моменты:
        - Образ должен запускать [Jenkins server](https://www.jenkins.io/download/)
        
    - Спецификация первого образа:
        - Базовый образ - [amazoncorreto](https://hub.docker.com/_/amazoncorretto)
        - Присвоить образу тэг `ver1` 
    
    - Спецификация второго образа:
        - Базовый образ - [ubuntu:latest](https://hub.docker.com/_/ubuntu)
        - Присвоить образу тэг `ver2` 

- Соберите 2 образа по полученным Dockerfile
- Запустите и проверьте их работоспособность
- Опубликуйте образы в своём dockerhub.io хранилище

Для получения зачета, вам необходимо предоставить:
- Наполнения 2х Dockerfile из задания
- Скриншоты логов запущенных вами контейнеров (из командной строки)
- Скриншоты веб-интерфейса Jenkins запущенных вами контейнеров (достаточно 1 скриншота на контейнер)
- Ссылки на образы в вашем хранилище docker-hub

## Ответ:

## 2.1 Amazoncorretto

- Dockerfile:

```text
FROM amazoncorretto

ADD https://pkg.jenkins.io/redhat-stable/jenkins.repo /etc/yum.repos.d/

RUN rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io.key && \
    yum install -y jenkins
EXPOSE 8080
CMD ["java", "-jar", "/usr/lib/jenkins/jenkins.war"]
```

- Скрин логов:

[Задание 2.1 - amazoncorretto - logs.png](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/05-virt-04-docker-practical-skills/%D0%97%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%202%20-%20amazoncorretto%20-%20logs.png)

- Скрин Web-интерфейса Jenkins:

[Задание 2.1 - amazoncorretto - web.png](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/05-virt-04-docker-practical-skills/%D0%97%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%202%20-%20amazoncorretto%20-%20web.png)

- Образ:

https://hub.docker.com/r/vitalymozhaev/amazoncorretto-jenkins

## 2.2 Ubuntu

- Dockerfile:

```text
FROM ubuntu

ADD https://pkg.jenkins.io/debian/jenkins.io.key /tmp/

RUN apt-get update && \
    apt-get install -y gnupg ca-certificates && \
    apt-key add /tmp/jenkins.io.key && \
    sh -c 'echo deb https://pkg.jenkins.io/debian-stable binary/ >> /etc/apt/sources.list' && \
    apt-get update && \
    apt-get install -y openjdk-11-jdk openjdk-11-jre jenkins
EXPOSE 8080
CMD ["java", "-jar", "/usr/share/jenkins/jenkins.war"]
```

- Скрин логов:

[Задание 2.2 - ubuntu - logs.png](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/05-virt-04-docker-practical-skills/%D0%97%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%202%20-%20ubuntu%20-%20logs.png)

- Скрин Web-интерфейса Jenkins:

[Задание 2.2 - ubuntu - web.png](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/05-virt-04-docker-practical-skills/%D0%97%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%202%20-%20ubuntu%20-%20web.png)

- Образ:

https://hub.docker.com/r/vitalymozhaev/ubuntu-jenkins

## Задача 3 

В данном задании вы научитесь:
- объединять контейнеры в единую сеть
- исполнять команды "изнутри" контейнера

Для выполнения задания вам нужно:
- Написать Dockerfile: 
    - Использовать образ https://hub.docker.com/_/node как базовый
    - Установить необходимые зависимые библиотеки для запуска npm приложения https://github.com/simplicitesoftware/nodejs-demo
    - Выставить у приложения (и контейнера) порт 3000 для прослушки входящих запросов  
    - Соберите образ и запустите контейнер в фоновом режиме с публикацией порта

- Запустить второй контейнер из образа ubuntu:latest
- Создайте `docker network` и добавьте в нее оба запущенных контейнера
- Используя `docker exec` запустить командную строку контейнера `ubuntu` в интерактивном режиме
- Используя утилиту `curl` вызвать путь `/` контейнера с npm приложением  

Для получения зачета, вам необходимо предоставить:
- Наполнение Dockerfile с npm приложением
- Скриншот вывода вызова команды списка docker сетей (docker network cli)
- Скриншот вызова утилиты curl с успешным ответом

## Ответ:

- Dockerfile с npm приложением

```text
FROM node

ADD https://github.com/simplicitesoftware/nodejs-demo/archive/refs/heads/master.zip /

RUN apt-get update && \
    unzip master.zip
WORKDIR "/nodejs-demo-master"
RUN npm install
EXPOSE 3000
CMD ["npm", "start", "0.0.0.0"]
```

- docker network ls

```text
docker network create -d bridge node-ubuntu
docker network connect node-ubuntu ubuntu_l
docker network connect node-ubuntu netology-node
docker network ls
NETWORK ID     NAME          DRIVER    SCOPE
87977a6f0994   bridge        bridge    local
dd7ffcefb8dc   host          host      local
a674133452b7   node-ubuntu   bridge    local
f9c18a85e48b   none          null      local
```


Тут видим настройки сети и подключенный контейнеры:

- docker network inspect node-ubuntu

```text
[
    {
        "Name": "node-ubuntu",
        "Id": "a674133452b7fcb04ee96bf4fa7c6d2b4ae0fc2f67bc45bbdc4159a9b178e138",
        "Created": "2021-06-25T15:54:46.412440344Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "172.18.0.0/16",
                    "Gateway": "172.18.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "5fa7abc0c7a92d7144b8eb89658bee7d7d5238d4a4554f8d4454cd96373c9d13": {
                "Name": "ubuntu_l",
                "EndpointID": "e7825889502526d5de7218532c2c1a34c5f6632228e28e547b64ce165d50a736",
                "MacAddress": "02:42:ac:12:00:02",
                "IPv4Address": "172.18.0.2/16",
                "IPv6Address": ""
            },
            "6e4e2382c64e86756b684a5c5c04fa1f86edfecf4882778a82689e34fe32addd": {
                "Name": "netology-node",
                "EndpointID": "2f921e500d6530728a63aef28736802f415416a6a26e1aafd4497528ec2a2065",
                "MacAddress": "02:42:ac:12:00:03",
                "IPv4Address": "172.18.0.3/16",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {}
    }
]
```

- [Вывод `curl 172.18.0.3:3000` из контейнера с ubuntu](https://github.com/VitalyMozhaev/virt-homeworks/blob/main/05-virt-04-docker-practical-skills/%D0%97%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%203%20-%20%D0%B2%D1%8B%D0%B2%D0%BE%D0%B4%20curl.png)

