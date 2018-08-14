GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = github.com/spf13/cobra golang.org/x/crypto/ssh/terminal

# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=fullmoon
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
install: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) -v -u $(GOPACKAGES)