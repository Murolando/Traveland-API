# Traveland-REST-API
# API Для Туристического приложения
## Первый опыт в создании api
---
* Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
* Работа с фреймворком gin-gonic/gin.
* Работа с БД Postgres. Запуск из Docker. Генерация файлов миграций.
## Список Возможностей:
*  Регистрация пользователя 
*  Аутентификация пользователя через jwt и refresh токены
*  CRUD по сущиностям проекта
*  Работа с DOCKER + DOCKER COMPOSE
# Для запуска приложения:
```
docker-compose up --build
```
Если приложение запускается впервые, необходимо применить миграции к базе данных:
```
migrate -path ./db/migrations -database 'postgres://postgres:<password>@localhost:5432/postgres?sslmode=disable' up
```


---
# Traveland-REST-API
# API for Tourists application
## First experience in creating an api
---
* Pure Architecture approach in building the application structure. Dependency injection technique.
* Working with the gin-tonic/gin-soda framework.
* Working with the Postgres database. Launch from docker. Generation of migration files.
## List of Features:
* User registration
* User automation via JWT and REFRESH token
* CRUD's
* Working with DOCKER + DOCKER COMPOSE
# To launch the application:
```
docker-compose -build
```
If the application is launched for the first time, you need to apply migrations to the database:
```
migrate -path ./db/migrations -database 'postgres://postgres:<password>@localhost:5432/postgres?sslmode=disable `up
'``
