# Crypto Portfolio Tracker

## Проект
REST API на Go + React фронт. Трекер криптопортфеля с реальными ценами.

## Архитектура
cmd/main.go → internal/portfolio/http.go → service.go → repository.go
→ internal/prices/coingecko.go

## Стек
- Go 1.22
- PostgreSQL
- React (фронт)
- Docker + docker-compose
- GitHub Actions (CI/CD)

## Правила кода
- Каждый слой делает только своё (http, service, repository)
- Ошибки всегда обрабатываются, не игнорируются
- Никакой бизнес-логики в http.go
- Никаких SQL запросов в service.go

## Структура
crypto-portfolio/
├── cmd/main.go
├── internal/
│   ├── portfolio/
│   │   ├── http.go
│   │   ├── service.go
│   │   └── repository.go
│   └── prices/
│       └── coingecko.go
├── go.mod
└── docker-compose.yml

## Как запускать
go run cmd/main.go

## Контекст для ревью
- Я изучаю Go, пришел из JS
- Проверяй архитектуру, ошибки, Go-идиомы
- Указывай на плохие практики
- Не пиши код за меня, давай подсказки