default: test

LDFLAGS := -ldflags="-w -s"

build:
	GO111MODULE=on go build $(LDFLAGS)

install-protoc:
ifeq ($(shell command -v protoc-gen-go 2> /dev/null),)
	go get github.com/golang/protobuf/{proto,protoc-gen-go}
	go get google.golang.org/grpc
endif

protoc: install-protoc
	protoc -Iprotos --go_out=plugins=grpc:internal/service/ $(shell find protos -type f -name "*.proto")

test:
	GO111MODULE=on go test -v -race -parallel=4 ./...

.PHONY: build install-protoc protoc test
