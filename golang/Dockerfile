FROM golang:1.16

WORKDIR /app
COPY . /app

RUN go get -u github.com/cosmtrek/air
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0
RUN go get -v github.com/rubenv/sql-migrate/...

RUN go mod download
