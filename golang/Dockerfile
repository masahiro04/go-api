FROM golang:1.16

WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

RUN go mod download

COPY . /app
