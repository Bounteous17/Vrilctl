include .env

GOCMD=go
GOBUILD=$(GOCMD) build  -o $(BINARY_NAME) -v
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_PATH=bin/
BINARY_NAME=$(BINARY_PATH)start

all: test build

build:
	@echo "Building $(GOFILES) to ./bin"
	$(GOBUILD)

run:
	$(GOBUILD)
	./$(BINARY_NAME)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_PATH)*