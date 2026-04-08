# Калькулятор калорий

Сервис позволяет пользователю указывать свой дневной рацион и считать потребление калорий, белков, жиров и углеводов (КБЖУ).

Используемые технологии:
- PostgreSQL (хранилище данных)
- Docker (контейнеризация)
- pgx (драйвер для работы с PostgreSQL)
- Swagger (для документации API)

Сервис написан с применением чистой архитектуры, в нём выделены основные слои: domain, service, storage, controller, DTO. В сервисе реализовано логирование с помощью пакета "log/slog" и аутентификация пользователей с помощью JWT. Реализован Graceful Shutdown для корректного завершения работы сервиса.

## Запуск

Запустить сервис можно с помощью команды `docker compose up`, но перед этим надо создать файл `.env` на основе файла `.env.example`, и файл `./config/HS256key.txt` на основе файла `./config/HS256key-example.txt`, хранящий секретный ключ для работы JWT.

## Использование сервиса

Backend приложения реализован в виде HTTP API, работающем стандартно по адресу `http://localhost:8000`. Документацию после запуска сервиса можно посмотреть по адресу `http://localhost:8080/#/`. 

Вышеперечисленные порты сервиса можно поменять в файле `compose.yaml`. Если вы будете менять порт, на котором находится документация, не забудьте поменять значение переменной окружения `SWAGGER_UI_ADDR` у сервиса с названием `backend`.

## Примеры

Далее описаны примеры некоторых запросов к HTTP API.

- [Регистрация пользователя](#регистрация-пользователя)
- [Аутентификация](#аутентификация)
- [Добавление продукта](#добавление-продукта)
- [Получение продуктов](#получение-продуктов)
- [Добавление рациона](#добавление-рациона)
- [Получение рационов](#получение-рационов)

### Регистрация пользователя

Регистрирует нового пользователя в сервисе.

```curl
curl -X 'POST' \
  'http://localhost:8000/register' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "username": "Raiden Shogun",
  "password": "qwerty"
}'
```

Пример ответа:
```json
{
  "username": "Raiden Shogun"
}
```

### Аутентификация

Позволяет аутентифицироваться в сервисе. Используется Basic Auth.
```curl
curl -X 'POST' \
  'http://localhost:8000/login' \
  -H 'accept: application/json' \
  -H 'Authorization: Basic UmFpZGVuIFNob2d1bjpxd2VydHk=' \
  -d ''
```
Возвращает Access Token, который нужно прикладывать к каждому защищённому эндпоинту (Bearer Auth с использованием JWT).

### Добавление продукта

Позволяет пользователю добавить продукт (рис, филе курицы, кетчуп и т.д.) для последующего подсчёта калорий его рациона.

```curl
curl -X 'POST' \
  'http://localhost:8000/products' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer YOUR_TOKEN' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "onion", "base_weight": 100, "base_portion": 250,
  "calories": 150.6, "fats": 4.4, "proteins": 5, "carbohydrates": 24.6
}'
```
Пример ответа:
```json
{
  "name": "onion", "base_weight": 100, "base_portion": 250,
  "calories": 150.6, "fats": 4.4, "proteins": 5, "carbohydrates": 24.6
}
```

### Получение продуктов

Выводит список всех продуктов пользователя.

```curl
curl -X 'GET' \
  'http://localhost:8000/products' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer YOUR_TOKEN'
```
Пример ответа:
```json
[
  {
    "name": "onion", "base_weight": 100, "base_portion": 250,
    "calories": 150.6, "fats": 4.4, "proteins": 5, "carbohydrates": 24.6
  }
]
```

### Добавление рациона

Позволяет пользователю указать, какие продукты он употребил в определённый день и посчитать на их основе потреблённые калории и БЖУ.

```curl
curl -X 'POST' \
  'http://localhost:8000/rations' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer YOUR_TOKEN' \
  -H 'Content-Type: application/json' \
  -d '{
  "date": "2026-03-21",
  "products": [
    {
      "name": "onion", "weight": 100, "portion": 2.5
    }
  ]
}'
```
В ответ пользователь получает посчитанные калории и БЖУ.
```json
{
  "date": "2026-03-21",
  "calories": 1091.85, "fats": 31.900000000000002,
  "proteins": 36.25, "carbohydrates": 178.35000000000002
}
```

### Получение рационов

Позволяет пользователю узнать информацию о всех добавленных им рационах.

```curl
curl -X 'GET' \
  'http://localhost:8000/rations' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer YOUR_TOKEN'
```
Пример ответа:
```json
[
  {
    "date": "2026-03-21",
    "calories": 1091.85, "fats": 31.9,
    "proteins": 36.25, "carbohydrates": 178.35
  }
]
```