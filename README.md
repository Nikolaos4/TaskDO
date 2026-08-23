# TaskDo

Серверное приложение для таск-трекинга. Пользователи могут создавать задачи, устанавливать для них срок выполнения, редактировать и удалять их. В приложении реализована система аутентификации на основе JWT. Приложение упаковано в Docker-контейнеры, что позволяет запускать его в любом окружении без дополнительной настройки зависимостей.

## Технологический стек

- **Язык:** Go
- **База данных:** PostgreSQL
- **Фреймворк:** Gin — REST API
- **ORM:** GORM — работа с базой данных
- **Аутентификация:** JWT (golang-jwt)
- **Хэширование паролей:** bcrypt (golang.org/x/crypto)
- **Контейнеризация:** Docker, Docker Compose

## Запуск

### Через Docker (рекомендуется)

1. Установите [Docker](https://www.docker.com/get-started) и Docker Compose
2. Склонируйте репозиторий и перейдите в папку проекта
3. Создайте файл `.env`:
   ```env
   POSTGRES_PASSWORD=ваш_пароль
   DB_URL=host=db user=postgres password=ваш_пароль dbname=taskdo port=5432 sslmode=disable
   SECRET_KEY=ваш_секретный_ключ
   ```
4. Запустите контейнеры:
   ```bash
   docker compose up -d
   ```
5. Приложение доступно по адресу: `http://localhost:8080`

### Локально

1. Установите [Go](https://golang.org/dl/) 1.21+ и [PostgreSQL](https://www.postgresql.org/download/) 15+
2. Создайте файл `.env`:
   ```env
   DB_URL=host=localhost user=postgres password=ваш_пароль dbname=taskdo port=5432 sslmode=disable
   SECRET_KEY=ваш_секретный_ключ
   ```
3. Установите зависимости и запустите:
   ```bash
   go mod download
   go run main.go
   ```
4. Приложение доступно по адресу: `http://localhost:8080`

## API

### Пользователи

| Метод | Эндпоинт | Описание | Тело запроса |
|-------|----------|----------|--------------|
| POST | `/user/create` | Регистрация | `{"login": "user", "password": "pass"}` |
| POST | `/user/login` | Вход, возвращает JWT-токен | `{"login": "user", "password": "pass"}` |
| GET | `/user/:id` | Получение информации о пользователе | — |

### Задачи

> Все эндпоинты задач требуют JWT-токен в заголовке: `Authorization: Bearer <token>`

| Метод | Эндпоинт | Описание | Тело запроса |
|-------|----------|----------|--------------|
| POST | `/task/create` | Создание задачи | `{"title": "task", "data": "2026-08-23T12:00:00Z"}` |
| PUT | `/task/:id` | Редактирование задачи | `{"title": "new_task", "data": "2026-09-23T12:00:00Z"}` |
| GET | `/task/:id` | Получение задачи по ID | — |
| GET | `/task/all_tasks` | Получение всех задач пользователя | — |
| DELETE | `/task/:id` | Удаление задачи | — |