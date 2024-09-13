# Этап сборки
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем все остальные файлы исходного кода
COPY . .

# Сборка приложения, указываем путь к главному файлу main.go в каталоге cmd
RUN go build -o main ./cmd/main.go

# Этап выполнения
FROM alpine:3.14

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем скомпилированное приложение
COPY --from=builder /app/main .

# Копируем файл .env
COPY .env .

EXPOSE 8080

CMD ["./main"]
