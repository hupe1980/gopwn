PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: help
## help: Prints this help message
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: setup
## setup: Setup installes dependencies
setup:
	@go mod tidy

.PHONY: run
## run: Runs gopwn
run: 
	@go run $$(ls -1 cmd/*.go | grep -v _test.go) checksec testdata/cat

.PHONY: build
## build: Builds a beta version of gopwn
build:
	go build -o dist/gopwn cmd/*.go

.PHONY: ci
## ci: Run all the tests and code checks
ci: build test

.PHONY: test
## test: Runs go test with default values
test: 
	@go test -v -race -count=1  ./...

