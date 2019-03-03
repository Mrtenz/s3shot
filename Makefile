GOCMD 		= go
GOBUILD 	= $(GOCMD) build
GOCLEAN 	= $(GOCMD) clean
GOGET 		= $(GOCMD) get -u
GOOS		= linux
BINARY_NAME	= s3shot
PREFIX 		= /usr/local

all: dependencies clean build

build:
		$(GOBUILD) -ldflags="-s -w" -o $(BINARY_NAME)

clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)

dependencies:
		$(GOGET) gopkg.in/urfave/cli.v1
		$(GOGET) github.com/aws/aws-sdk-go

install:
		mkdir -p $(DESTDIR)$(PREFIX)/bin
		install -m755 $(BINARY_NAME) $(DESTDIR)$(PREFIX)/bin/$(BINARY_NAME)