FROM cimg/go:1.17
LABEL authors="evgeniy"

ENV GOOS=linux GOARCH=amd64
USER root

COPY cmd/main.go /app/main.go
COPY go.mod /app/go.mod
WORKDIR /app

RUN go build -o main

EXPOSE 8080 8080
CMD ["/app/main"]
