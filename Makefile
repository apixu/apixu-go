.PHONY: all test lint coverage integration

GO111MODULE=on
COVER_PROFILE=cover.out

all: test lint

test:
	go test -cover

lint:
	@[ ! -f ./bin/golangci-lint ] && curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh \
		| sh -s -- -b ./bin v1.16.0 || true
	./bin/golangci-lint run

coverage:
	go test -coverprofile $(COVER_PROFILE) && go tool cover -html=$(COVER_PROFILE)

integration:
	go test -v -tags integration
