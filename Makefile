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

.PHONY: run
## run: Runs gopwn
run: 
	@go run cmd/*.go cyclic create 25

.PHONY: test
## test: Runs go test with default values
test: 
	@go test -v -race -count=1  ./...

