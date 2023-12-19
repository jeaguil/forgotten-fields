BINARY_NAME := forgotten-fields

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

all: run

## tidy: formats the code
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -vm

## no-dirty: checks that there are no uncommited changes in the tracked files
.PHONY: no-dirty
no-dirty:
	git diff --exit-code

## audit: runs quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## build: build the application
.PHONY: build
build:
	go build -o ./build/tmp/bin/${BINARY_NAME}

## run: run the application
.PHONY: run
run: build
	./build/tmp/bin/${BINARY_NAME}

## clean: remove build related files
.PHONY: clean
clean:
	go clean
	rm -rf ./build
