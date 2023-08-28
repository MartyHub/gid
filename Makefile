default: watch

all: tidy vet lint test

lint:
	golangci-lint run

test:
	go test -race -timeout 10s ./...

tidy:
	go mod tidy

vet:
	go vet ./...

watch:
	modd --file=.modd.conf

.PHONY: all lint test tidy vet watch
