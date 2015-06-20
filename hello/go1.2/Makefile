GO_FILES = $(shell find . -type f -name '*.go')
PACKAGES = $(shell ls src/)
GOPATH = $(shell pwd):$(shell pwd)/vendor:$(shell pwd)/lib/atlantis

all: build

build: $(GO_FILES)
	GOPATH=$(GOPATH) go build -o bin/hello

test: $(GO_FILES)
	GOPATH=$(GOPATH) go test $(PACKAGES)

run: build
	./bin/hello -talk

clean:
	rm -rf pkg/* bin/*

package: build
	@cp -r bin package
