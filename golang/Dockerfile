FROM golang:1.16

COPY . /app
WORKDIR /app

RUN go get -u github.com/cosmtrek/air
RUN go get -v github.com/rubenv/sql-migrate/...

RUN go mod download
