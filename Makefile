.PHONY: tidy
tidy:
	go mod tidy

.PHONY: install
install:
	go get ./...

.PHONY: build
build:
	go build .

.PHONY: docker-build
docker-build:
	docker build -t gonedrive .

.PHONY: test
test:
	go test -v -race -covermode atomic -coverprofile=profile.cov ./service/azure ./service/secret ./crypto
	go tool cover -func=profile.cov

.PHONY: vet
vet:
	go vet ./...

.PHONY: goreport
goreport:
	goreportcard-cli -v

.PHONY: all
all: install vet goreport build test
