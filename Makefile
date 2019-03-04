GOCMD 		= go
GOBUILD 	= $(GOCMD) build
GOCLEAN 	= $(GOCMD) clean
GOGET 		= $(GOCMD) get
GOOS		= linux
BINARY_NAME	= s3shot
PREFIX 		= /usr/local

all: clean build

build:
		$(GOBUILD) -ldflags="-s -w" -o $(BINARY_NAME)

clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)

install:
		mkdir -p $(DESTDIR)$(PREFIX)/bin
		install -m755 $(BINARY_NAME) $(DESTDIR)$(PREFIX)/bin/$(BINARY_NAME)