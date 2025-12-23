# i-moscow-backend

The backend project of our team (Origin Dev), 2nd track of the hackathon Moscow City Hack 2021.

The server side is implemented as a REST API for interacting with the application. Go was chosen as the programming language because it allows building modern, easily scalable backend services. For the same reasons, MongoDB was selected as the database. The hh_skills.csv file was imported into MongoDB for more convenient data handling. To build project competency profiles, we used full-text search (fullTextSearch) and also added indexes to ensure fast query performance.

## Tech

* Go 1.16+
* JWT
* Firebase
* MongoDB

![](https://github.com/semyon-dev/i.moscow/blob/master/stack.png)

## .env example

```
ACCESS_SECRET=secret
MONGO_URL="mongodb://127.0.0.1:27017/?compressors=zlib&readPreference=primary&ssl=false"
PORT=8080
FIREBASE_FILE_NAME="i-moscow-firebase-adminsdk.json"
```

Additionally, to work with notifications, the Firebase Admin SDK must be located in the root of the project.

## Run project

`go run app/main.go`

## Team origin dev

* Semyon Novikov [semyon_dev](https://github.com/semyon-dev) (Project Manager & Backend)
* Valeriy Alyushin [alyush1n](https://github.com/alyush1n)  (Design)
* Andrey Nebogatikov [Dronicho](https://github.com/Dronicho)  (Mobile)
