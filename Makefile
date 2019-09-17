.PHONY: all test lint coverage integration env

GO111MODULE=on
COVER_PROFILE=cover.out

ifndef GOVERSION
GOVERSION=1.12
endif

all: test lint

test:
	go test -cover

lint:
	@[ ! -f ./bin/golangci-lint ] && curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh \
		| sh -s -- -b ./bin v1.18.0 || true
	./bin/golangci-lint run

coverage:
	go test -coverprofile $(COVER_PROFILE) && go tool cover -html=$(COVER_PROFILE)

integration:
	go test -v -tags integration

env:
	docker run --rm -ti -v $(CURDIR):/src -w /src -e APIXUKEY=$(APIXUKEY) golang:$(GOVERSION) bash
