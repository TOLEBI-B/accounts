# Scheta Kubernetes Service

## 🚀 Описание
Простой сервис на Go + PostgreSQL, разворачиваемый в Kubernetes (k3s).  
Автодеплой через GitHub Actions -> DockerHub -> k3s.

## 📂 Структура
- `.github/workflows/deploy.yaml` — CI/CD
- `k3s/` — манифесты Kubernetes (app, postgres, ingress)
- `cmd/server/` — main.go
- `internal/` — db и handlers
- `migrations/` — SQL миграции

## 🔧 Запуск локально
```bash
docker-compose up -d postgres
go run ./cmd/server
```

## ☸️ Деплой в Kubernetes
```bash
kubectl apply -f k3s/postgres.yaml
kubectl apply -f k3s/deployment.yaml
kubectl apply -f k3s/service.yaml
kubectl apply -f k3s/ingress.yaml
```

## 🌐 API
- `GET /accounts` → список счетов
