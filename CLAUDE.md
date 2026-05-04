# Task Manager API

Go + Gin REST API. Architecture: Handler → Service → Repository → Database.

## Stack

- **Backend:** Go + Gin, PostgreSQL + sqlx, JWT (access + refresh), Redis, Docker, Swagger
- **Frontend:** Next.js + TypeScript, Tailwind CSS, Redux Toolkit + RTK Query, React Hook Form + Zod
- **Planned:** WebSockets

## Project Structure

backend/
cmd/api/main.go — entry point
docs/ — generated Swagger documentation
internal/
config/config.go — config from env
db/db.go — database connection
db/migrate.go — migration runner
db/redis.go — Redis connection
server/handler.go — handler initialization
server/middleware.go — JWT + Rate limiting middleware
server/router.go — Gin router
auth/
model.go — User, tokens
repository.go — SQL queries
service.go — business logic
service_test.go
jwt.go — JWT generation/validation
errors.go
task/
model.go — Task, statuses, request types
repository.go — SQL queries
service.go — business logic
service_test.go
handler.go — HTTP handlers
errors.go
migration/
000001_create_users.up/down.sql
000002_create_tasks.up/down.sql
000003_update_tasks.up/down.sql

## What's Already Implemented

- **Auth:** registration, login, JWT access + refresh tokens
- JWT middleware for protected routes
- **Rate limiting:** 60 req/min per IP via Redis
- **Tasks CRUD:** full implementation (handler, service, repository) — covered with TDD
- **Swagger documentation** for all endpoints
- **Migrations:** users + tasks tables
- **Docker + docker-compose** (PostgreSQL + Redis + App)
- Healthcheck for DB in docker-compose

## In Progress

- **WebSockets** — real-time notifications
- **Frontend:** Next.js + TypeScript + Tailwind + Redux Toolkit + RTK Query

## Code Conventions

- Errors: via `fmt.Errorf("context: %w", err)` with wrapping
- Context: always first argument `ctx context.Context`
- UUID: `github.com/google/uuid` for IDs
- Logging: `slog` (not fmt.Println)
- Tests: TDD, mocks via interfaces

## Useful Commands

```bash
go test ./...                    # run all tests
go test ./internal/task/... -v   # run task package tests
go build ./cmd/api/              # build
docker compose up --build        # start full environment
```
