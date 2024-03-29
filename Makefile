#!/usr/bin/make -f

#################
###   Build   ###
#################

test:
	@echo "--> Running tests"
	go test -v ./...

test-integration:
	@echo "--> Running integration tests"
	cd integration; go test -v ./...

.PHONY: test test-integration

##################
###  Protobuf  ###
##################

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating protobuf files..."
	sh ./scripts/protocgen.sh
	@go mod tidy

proto-format:
	find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	buf lint proto/ --error-format=json

.PHONY: proto-all proto-gen proto-format proto-lint

#################
###  Linting  ###
#################

golangci_lint_cmd=golangci-lint
golangci_version=v1.51.2

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run ./... --timeout 15m

lint-fix:
	@echo "--> Running linter and fixing issues"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run ./... --fix --timeout 15m

.PHONY: lint lint-fix