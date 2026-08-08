# News Service

REST API сервис для агрегации новостей из RSS лент.

## Стек
- Go
- PostgreSQL
- Docker

## Запуск
```bash
docker compose up -d
go run cmd/main.go
```

## Эндпоинты
| Метод | URL | Описание |
|-------|-----|----------|
| GET | /health | Проверка работы |
| GET | /articles | Список всех статей |

## Как работает
-Парсит Rss ленты
-Хранит статьи
-Отдает по категориям 
