.PHONY: tidy
tidy:
	go mod tidy

.PHONY: install
install:
	go get ./...

.PHONY: build
build:
	go build .

.PHONY: test
test:
	go test -v -coverprofile=coverage.out ./service
	go tool cover -func=coverage.out

.PHONY: vet
vet:
	go vet ./...

.PHONY: all
all: install vet build test