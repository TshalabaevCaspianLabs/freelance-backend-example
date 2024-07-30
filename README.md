# MyApp

## Описание

MyApp - это пример REST API сервиса, написанного на Go, с использованием чистой архитектуры.

## Используемые технологии

- Go
- Gin
- GORM
- PostgreSQL
- Docker
- Swagger

## Запуск

1. Склонируйте репозиторий:
    ```sh
    git clone https://github.com/TshalabaevCaspianLabs/freelance-backend-example.git
    cd freelance-backend-example
    ```

2. Настройте файл конфигурации:
    Создайте файл `config.yml` в корне проекта со следующим содержимым:
    ```yaml
    db:
      host: db
      port: 5432
      user: gorm
      password: gorm
      dbname: gorm
    ```

3. Запустите Docker:
    ```sh
    docker-compose up --build
    ```

4. Откройте браузер и перейдите по адресу `http://localhost:8080/swagger/index.html` для просмотра документации API.
