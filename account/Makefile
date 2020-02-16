GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=account-service
DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

build-image: build-linux
	$(DOCKERBUILD) --no-cache . -t cyub/go-kit-account

run:
	./$(BINARY_NAME) --consul.host 192.168.33.10 --consul.port 8500
