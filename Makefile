#
# Makefile
#

# Version
VERSION = `date +%y.%m`

# If unable to grab the version, default to N/A
ifndef VERSION
    VERSION = "n/a"
endif

# 
# phony targets
#
.PHONY: all build install uninstall clean

#
# make targets
#

all: build

build:
	@go build -ldflags '-s -w -X main.version='${VERSION}

install:
	@go install github.com/rbisewski/golang_starter_repo

uninstall:
	@go clean -i github.com/rbisewski/golang_starter_repo

clean:
	@go clean
