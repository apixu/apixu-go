.PHONY: all test qainstall coverage lint

GO111MODULE=on
COVER_PROFILE=cover.out

all: test lint

qainstall:
	@set -eu; \
	GO111MODULE=off go get github.com/stretchr/testify/assert; \
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ./bin v1.14.0

test:
	go test ./... -cover

coverage:
	go test ./... -coverprofile $(COVER_PROFILE) && go tool cover -html=$(COVER_PROFILE)

lint:
	./bin/golangci-lint run
