# i-moscow-backend

Backend-проект нашей команды 2 трека с хакатона [Moscow City Hack](https://moscityhack.innoagency.ru/#main)

Серверная часть представляет из себя REST API для взаимодействия с приложением. 

## Стек

* Go 1.16
* JWT
* Firebase
* MongoDB

![](https://github.com/semyon-dev/i.moscow/blob/master/stack.png)

## Пример .env файла

```
ACCESS_SECRET=secret
MONGO_URL="mongodb://127.0.0.1:27017/?compressors=zlib&readPreference=primary&ssl=false"
PORT=8080
FIREBASE_FILE_NAME="i-moscow-firebase-adminsdk.json"
```

Так же для работы с уведомлениями в корне проекта должен лежать firebase admin sdk

## Запуск проекта

`go run app/main.go`

## Команда origin dev

* Семен Новиков [semyon_dev](https://github.com/semyon-dev) (Project Manager & Backend)
* Валерий Алюшин [alyush1n](https://github.com/alyush1n)  (Design)
* Андрей Небогатиков [Dronicho](https://github.com/Dronicho)  (Mobile)
