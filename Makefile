BASE = github.com/andreiavrammsd/apixu-go
PKG = . $(BASE)/formatter $(BASE)/response $(BASE)/types
COVER_PROFILE = cover.out

.PHONY: qa

all: qa

test:
	go test $(PKG) -cover

coverage:
	go test $(PKG) -coverprofile $(COVER_PROFILE) && go tool cover -html=$(COVER_PROFILE)

qainstall:
	@set -eu; \
	go get \
		github.com/stretchr/testify/assert \
		golang.org/x/tools/cmd/goimports \
		golang.org/x/lint/golint \
		honnef.co/go/tools/cmd/megacheck \
	   	mvdan.cc/interfacer \
	   	github.com/alexkohler/prealloc \
	   	github.com/kisielk/errcheck

qa:
	go fmt $(PKG)
	go vet $(PKG)
	go test $(PKG) -cover
	golint ./...
	megacheck $(PKG)
	interfacer $(PKG)
	prealloc $(PKG)
	errcheck $(PKG)
