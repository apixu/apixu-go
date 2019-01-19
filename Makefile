.PHONY: all test qainstall coverage lint

GO111MODULE=on
COVER_PROFILE=cover.out

all: test lint

qainstall:
	@set -eu; \
	GO111MODULE=off go get github.com/stretchr/testify/assert; \
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b /usr/local/bin v1.12.5

test:
	go test ./... -cover

coverage:
	go test ./... -coverprofile $(COVER_PROFILE) && go tool cover -html=$(COVER_PROFILE)

lint:
	golangci-lint run
