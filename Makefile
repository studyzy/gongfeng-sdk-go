.PHONY: test lint fmt vet build check

build:
	go build ./...

test:
	go test -v -race -count=1 ./...

lint:
	golangci-lint run ./...

fmt:
	goimports -w .
	gofmt -w .

vet:
	go vet ./...

check: fmt vet lint test
