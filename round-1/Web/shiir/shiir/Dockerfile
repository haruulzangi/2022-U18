# syntax=docker/dockerfile:1
FROM golang:alpine

WORKDIR /app

COPY . .
RUN go build -o /app/server server.go

CMD [ "/app/server" ]
