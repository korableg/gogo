PACKAGE=$(shell go list -m)
GO_FILES=$(shell find . -name '*.go' | grep -vE 'vendor|easyjson|mock|_gen.go|.pb.go')
VERSION?=dev
LDFLAGS="-s -w -X github.com/korableg/wow/cmd.Version=${VERSION}"

.PHONY: imports
imports:
	gci write --custom-order -s standard -s default  -s "prefix(${PACKAGE})" ${GO_FILES}

.PHONY: lint
lint:
	golangci-lint -v run

.PHONY: test
test:
	go test -race -v ./...

.PHONY: mod_tidy
mod_tidy:
	go mod tidy -go 1.23 -compat 1.23

.PHONY: precommit
precommit: mod_tidy imports lint test