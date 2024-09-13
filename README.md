# Booking-Service

## Описание

**Booking Service** - это сервис для управления бронированиями. Он представляет API для управления пользователями, бронированиями и другими функциями. Для обеспечения четкого разделения обязанностей и поддерживаемости, в разработке использованы принципы *чистой архитектуры*. Это позволяет легко расширять и поддерживать сервис, а также минимизировать взаимозависимости между различными частями системы.

## Содержание

- [Особенности](#особенности)
- [Технологии](#технологии)
- [Установка и запуск](#установка-и-запуск)
- [Конфигурация](#конфигурация)
- [Использование](#использование)
- [Примеры использование](#примеры-использование)
- [Тестирование](#тестирование)

## Особенности

- Управление пользователями и бронированиями.
- API для взаимодействия с приложением.
- Использовние PostgreSQL для хранения данных.
- Полная поддержка Docker для изоляции среды.

## Технологии

- **Go**: Язык программирования, используемый для разработки основного приложения.
- **Docker**: Платформа для контейнеризации приложений.
- **PostgreSQL**: Система управления базами данных.
- **Gorilla**: Набор инструментов для разработки веб-приложений на Go.
- **Swagger**: Инструмент для документирования API.

## Установка и запуск

1. Убедитесь, что у вас установлен Docker и Docker Compose.
2. Клонируйте репозиторий и перейдите в директорию с репозиторием:
```bash
git clone https://github.com/AramLab/booking-service.git
cd booking-service
```
3. При необходимости поменяйте файл `.env` с переменными окружения.
4. Запустите Docker Compose:
```bash
docker-compose up --build
```
Это создаст и запустит два контейнера: один для приложения и один для базы данных PostgreSQL.

## Конфигурация

Параметры конфигурации определяются в файле `.env`. Вот описание используемых переменных:

- `DB_USER`: Имя пользователя для подключения к базе данных PostgreSQL.
- `DB_PASS`: Пароль пользователя для подключения к базе данных PostgreSQL.
- `DB_NAME`: Имя базы данных.
- `DB_PORT`: Порт для подключения к базе данных PostgreSQL.
- `DB_SSLMODE`: Режим SSL для подключения к базе данных.
- `DB_SERVER_PORT`: Порт, на котором будет работать приложение.

## Использование

Вы можете взаимодействовать с API через следующие эндпоинты:

- `POST /user`: Создать нового пользователя.
- `DELETE /user/{id}`: Удалить пользователя по его ID.

- `POST /booking`: Создать новое бронирование.
- `DELETE /booking/{id}`: Удалить бронирование по его ID.
- `GET /bookings`: Получить список всех бронирований.

Для получения дополнительной информации о запросах и ответах **после запуска проекта**, пожалуйста, смотрите [документацию API](http://localhost:8080/swagger/index.html) (документация будет доступна после запуска проекта).

## Примеры использования

Ниже представлены примеры использования API с помощью Postman:

**Запрос**: `POST /user`

**Тело запроса**:
```bash
{
  "username": "vasyliy25",
  "password": "securepassword"
}
```

**Запрос**: `POST /booking`

**Тело запроса**:
```bash
{
    "user_id": 1,
    "start_time": "2024-09-15T15:00:00Z",
    "end_time": "2024-09-15T16:00:00Z"
}
```

## Тестирование

Тесты находятся в папке `./server/booking` и `./server/user`. 

Запустить все тесты можно с помощью команды: 
```bash
go test ./...
```

Или отдельно тесты в каждой папке:
```bash
go test ./server/booking
go test ./server/user
```
