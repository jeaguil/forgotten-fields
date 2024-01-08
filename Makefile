BINARY_NAME := forgotten-fields

# Using Windows Subsystem for Linux (WSL) requires environment variable GOOS=windows when running application
WSL_FLAGS :=
UNAME_S := ${shell uname -s}
ifeq ($(UNAME_S), Linux)
	WSL_FLAGS += GOOS=windows
endif

.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

all: run

## tidy: formats the code
.PHONY: tidy
tidy:
	@echo '$@: ${BINARY_NAME}'
	go fmt ./...
	go mod tidy -v
	@echo '$@: successful'

## no-dirty: checks that there are no uncommited changes in the tracked files
.PHONY: no-dirty
no-dirty:
	git diff --exit-code

## audit: runs quality control checks
.PHONY: audit
audit:
	@echo '$@: Running quality control checks..'
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	@echo '$@: successful'

## build: build the application
.PHONY: build
build:
	@echo '$@: Building ${BINARY_NAME}...'
	$(WSL_FLAGS) go build -o ./build/tmp/bin/${BINARY_NAME}
	@echo '$@: successful'

## run: run the application
.PHONY: run
run: build
	./build/tmp/bin/${BINARY_NAME}

## clean: remove build related files
.PHONY: clean
clean:
	go clean
	rm -rf ./build