# syntax=docker/dockerfile:1
FROM golang:alpine

WORKDIR /app

COPY *.go flag.txt ./
RUN go build -o /app/server server.go

CMD [ "/app/server" ]
