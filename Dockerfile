# Базовый образ для сборки Go приложения
FROM golang:1.25.1-alpine3.22 AS build

# Установка рабочей директории
WORKDIR /app

# Копируем файлы модулей для кэширования зависимостей
# Modules layer
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
# Build layer
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /myapp ./cmd/app/

FROM alpine:3.22 AS run

COPY --from=build /myapp /myapp
COPY --from=build /app/templates /templates
COPY --from=build /app/config /config
COPY --from=build /app/.env /.env
EXPOSE 8080

CMD ["/myapp"]