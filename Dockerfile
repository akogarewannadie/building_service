# Используем официальное изображение Golang
FROM golang:1.18-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все файлы в рабочую директорию контейнера
COPY . .

# Устанавливаем зависимости Go
RUN go mod tidy

# Собираем приложение
RUN go build -o building_service .

# Устанавливаем переменные окружения для подключения к базе данных
ENV DB_HOST=postgres_container
ENV DB_PORT=5432
ENV DB_USER=user
ENV DB_PASSWORD=password
ENV DB_NAME=building_service

# Экспонируем порт 8080 для доступа к приложению
EXPOSE 8080

# Команда для запуска приложения
CMD ["./building_service"]
