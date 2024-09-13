
FROM golang:1.22.0 AS builder


WORKDIR /app


ENV GOOS=linux


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go build -o main .


FROM debian:latest


RUN apt-get update && apt-get install -y ca-certificates postgresql-client && rm -rf /var/lib/apt/lists/*


COPY --from=builder /app/main /app/main


COPY init.sql /docker-entrypoint-initdb.d/init.sql


ENV DB_HOST=db
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=12345678
ENV DB_NAME=building_service


CMD ["/app/main"]
