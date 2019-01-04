NAME := esattory

GO ?= go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get

BUILD_DIR=./build
BINARY ?= $(BUILD_DIR)/$(NAME)
BINARY_UNIX ?= $(BUILD_DIR)/$(NAME)_unix

DOCKER ?= docker
DOCKER_REPOSITORY ?= nagaa052/$(NAME)

.PHONY: all
all: test build

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY) -v

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY)

.PHONY: deps
deps:
	dep ensure

.PHONY: run
run: build
run:
	$(GOBUILD) -o $(BINARY) -v
	$(BINARY)

.PHONY: docker-build
docker-build: DOCKER_TAG ?= latest
docker-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
	$(DOCKER) build -t $(DOCKER_REPOSITORY):$(DOCKER_TAG) .
	rm -f $(BINARY_UNIX)
