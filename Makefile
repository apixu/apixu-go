.PHONY: all test coverage qainstall qa

GO111MODULE=on
COVER_PROFILE=cover.out

all: qa

test:
	go test ./... -cover

coverage:
	go test ./... -coverprofile $(COVER_PROFILE) && go tool cover -html=$(COVER_PROFILE)

qainstall:
	@set -eu; \
	GO111MODULE=off go get \
		github.com/stretchr/testify/assert \
		golang.org/x/tools/cmd/goimports \
		golang.org/x/lint/golint \
		honnef.co/go/tools/cmd/megacheck \
	   	mvdan.cc/interfacer \
	   	github.com/alexkohler/prealloc \
	   	github.com/kisielk/errcheck

qa: test
	go fmt ./...
	go vet ./...
	golint ./...
	megacheck ./...
	interfacer ./...
	prealloc ./...
	errcheck ./...
