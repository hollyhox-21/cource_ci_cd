# базовый образ
FROM cimg/go:1.17
# лэйбл образа
LABEL authors="evgeniy"

# переменные окружения для сборки приложения
ENV GOOS=linux GOARCH=amd64
# переключение на супер пользователя. Для сборки приложения
USER root

# копирование файлов в образ
COPY cmd/main.go /app/main.go
COPY go.mod /app/go.mod

# установка рабочей директории
WORKDIR /app

# команда сборки приложения
RUN go build -o main

# отрытие порта
EXPOSE 8080 8080

# команда запуска приложения для удержания контейнера в рабочем состоянии
CMD ["/app/main"]
