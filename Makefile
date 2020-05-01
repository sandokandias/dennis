BIN=bin
BINARY_NAME=den
CMD_PATH=./cmd/main.go
DIST=dist
DIST_MAC=$(DIST)/mac
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOCMD=go
GOTEST=$(GOCMD) test
GOTOOLCOVER=$(GOCMD) tool cover

build:
	mkdir -p $(DIST)
	#MAC
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ./$(DIST_MAC)/$(BINARY_NAME) -v $(CMD_PATH)

clean:
	rm -rf $(DIST)
	rm -rf $(BIN)