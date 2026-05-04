# Task Manager API

Go + Gin REST API. Архитектура: Handler → Service → Repository → Database.

## Стек

- **Backend:** Go + Gin, PostgreSQL + sqlx, JWT (access + refresh), Redis, Docker, Swagger
- **Frontend:** Next.js + TypeScript, Tailwind CSS, Redux Toolkit + RTK Query, React Hook Form + Zod
- **Планируется:** WebSockets

## Структура проекта

backend/
cmd/api/main.go — точка входа
docs/ — сгенерированная Swagger документация
internal/
config/config.go — конфиг из env
db/db.go — подключение к БД
db/migrate.go — запуск миграций
db/redis.go — подключение к Redis
server/handler.go — инициализация хендлеров
server/middleware.go — JWT + Rate limiting middleware
server/roter.go — роутер Gin
auth/
model.go — User, токены
repository.go — SQL запросы
service.go — бизнес-логика
service_test.go
jwt.go — генерация/валидация JWT
errors.go
task/
model.go — Task, статусы, request-типы
repository.go — SQL запросы
service.go — бизнес-логика
service_test.go
handler.go — HTTP хендлеры
errors.go
migration/
000001_create_users.up/down.sql
000002_create_tasks.up/down.sql
000003_update_tasks.up/down.sql

## Что уже реализовано

- Auth: регистрация, логин, JWT access + refresh
- JWT middleware для защищённых роутов
- Rate limiting: 60 req/min per IP через Redis
- Tasks CRUD: полная реализация (handler, service, repository) — покрыты TDD
- Swagger документация для всех эндпоинтов
- Миграции: users + tasks таблицы
- Docker + docker-compose (PostgreSQL + Redis + App)
- Healthcheck для БД в docker-compose

## В процессе

- WebSockets — real-time уведомления
- Frontend: Next.js + TypeScript + Tailwind + Redux Toolkit + RTK Query

## Соглашения по коду

- Ошибки: через `fmt.Errorf("context: %w", err)` с wrapping
- Контекст: всегда первым аргументом `ctx context.Context`
- UUID: `github.com/google/uuid` для ID
- Логирование: `slog` (не fmt.Println)
- Тесты: TDD, моки через интерфейсы

## Полезные команды

```bash
go test ./...                         # все тесты
go test ./internal/task/... -v        # тесты task пакета
go build ./cmd/api/                   # сборка
docker compose up --build             # поднять всё окружение
Как со мной работать (Антон)
Учусь Go — объясняй каждый файл по строкам
Объясняй ЗАЧЕМ это решение, не только КАК
После каждого файла делай короткий конспект
Если несколько способов — покажи лучший для Go
Учитывай best practices французского рынка: чистая архитектура, тесты, документация

```
