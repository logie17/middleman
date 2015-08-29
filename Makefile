.PHONY:	clean test

APP_NAME = middleman

#embed a build number and time into the binary
TIMESTAMP := $(shell date +"%s")
BUILD_TIME := $(shell date +"%Y%m%d.%H%M%S")
VERSION = $(strip $(TIMESTAMP))
ifndef BUILD_NUMBER
	#if build number is not provided assumed to be a dev build
	BUILD_NUMBER = dev
endif
LDFLAGS = -ldflags "-X main.buildTime $(BUILD_TIME) -X main.buildNumber $(BUILD_NUMBER)"

EXEC_NAME= ./$(APP_NAME)

$(EXEC_NAME):
	go build $(LDFLAGS) -o $(EXEC_NAME) *.go

clean:
	rm -rf $(EXEC_NAME)

test:
	go test ./...

check: test
	go vet ./...
	golint ./...
