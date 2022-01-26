.PHONY: all clean test docker latest init db swag

export CGO_ENABLED=0
export GOOS=darwin
export GOARCH=amd64

BINARY=go-lgtm-api
VERSION=$(shell git describe --abbrev=0 --tags 2> /dev/null || echo "0.1.0")
BUILD=$(shell git rev-parse HEAD 2> /dev/null || echo "undefined")
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"
MIGRATE_OPTIONS=-config=db/dbconfig.yml -env development
GOPATH=$(shell go env GOPATH)

run:
	air -c .air.toml

all:
	go build -o $(BINARY) $(LDFLAGS)

init:
	go mod download && \
	go install github.com/go-delve/delve/cmd/dlv@latest && \
	go install github.com/rubenv/sql-migrate/...@latest && \
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(GOPATH)/bin && \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin

migrate:
	sql-migrate status $(MIGRATE_OPTIONS)
	sql-migrate up $(MIGRATE_OPTIONS)
	sql-migrate status $(MIGRATE_OPTIONS)

test:
	go test ./...

dao_test:
	go test ./adapters/dao.blogRW

uc_test:
	go test ./usecases

cover:
	go test -cover ./... -coverprofile=cover.out.tmp | grep coverage:
	# 自動生成コードをカバレッジ対象から外し、カバレッジファイルを作成
	cat cover.out.tmp | grep -v "**_mock.go" | grep -v "wire_gen.go" > cover.out
	rm cover.out.tmp
	go tool cover -html=cover.out -o cover.html
	open cover.html

lint:
	golangci-lint run --fix

clean:
	if [ -f $(BINARY) ] ; then rm $(BINARY) ; fi

mock:
	mockgen -source=./usecases/INTERACTOR.go -destination=./adapters/uc.mock/interactor.go -package=mock && \
	  mockgen -source=./usecases/HANDLER.go -destination=./adapters/uc.mock/handler.go -package=mock

swag:
	swag init
