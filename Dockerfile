# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
RUN go build -o ./docker-go-example

EXPOSE 3000

CMD ["./docker-go-example"]