.PHONY: all test qainstall coverage lint

GO111MODULE=on
COVER_PROFILE=cover.out

all: test lint

qainstall:
	@set -eu; \
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ./bin v1.16.0

test:
	go test `go list ./... | grep -v examples` -cover

coverage:
	go test `go list ./... | grep -v examples` -coverprofile $(COVER_PROFILE) && go tool cover -html=$(COVER_PROFILE)

lint:
	./bin/golangci-lint run
