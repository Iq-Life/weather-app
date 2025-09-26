# Запуск локально

```bash
go run cmd/app/main.go cmd/app/weather.go
```

### Запуск как бинарный файл

Билдим приложение в бинарный исполняемый файл и запускаем его:
```bash
  go build -o myapp cmd/app/main.go cmd/app/weather.go && ./myapp
```

### Запуск в Docker контейнере

Делаем то же самое, но в Docker контейнере:
```bash
  docker build -t myapp:latest . && docker run -p 8080:8080 myapp:latest
```

### Запуск в Docker контейнере через Docker Compose

Делаем то же самое, но через Docker Compose:
```bash
  docker compose up --build
```
