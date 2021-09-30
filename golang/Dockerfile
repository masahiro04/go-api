FROM golang:1.16

COPY go.mod /app/go.mod
#COPY go.sum /code/go.sum
WORKDIR /app
RUN ls
RUN go mod download

COPY . /app

#CMD ["go run main.go"]
