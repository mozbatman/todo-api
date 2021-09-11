# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./out/go-todo-app .

EXPOSE 8080

CMD [ "./out/go-todo-app" ]