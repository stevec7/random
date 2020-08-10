VERSION=$(shell git describe --always --long --dirty)
.PHONY: all clean

all: test 

test:
	cd pkg/stack && go test -v || (echo "Tests failed"; exit 1)
	cd pkg/stackv2 && go test -v || (echo "Tests failed"; exit 1)

