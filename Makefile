VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BINARY_NAME = calendarN
LDFLAGS = -s -w -X main.version=$(VERSION)

CONSUMER_KEY?=
CONSUMER_SECRET?=
ACCESS_TOKEN?=
ACCESS_TOKEN_SECRET?=

.PHONY: build test create-tweet install clean release-darwin release-linux release-all

build:
	go mod download && go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) .

test:
	mkdir -p coverage && go test ./... --cover -coverprofile coverage/coverage.out

install: build
	mv $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
	rm -rf dist/

# Cross-compilation targets
release-darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-arm64 .

release-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-linux-arm64 .

release-all: clean release-darwin release-linux
	cd dist && shasum -a 256 * > checksums.txt

create-tweet:
	CONSUMER_KEY=${CONSUMER_KEY} CONSUMER_SECRET=${CONSUMER_SECRET} ACCESS_TOKEN=${ACCESS_TOKEN} ACCESS_TOKEN_SECRET=${ACCESS_TOKEN_SECRET} go run ./compose

# Update homebrew formula with new version SHA
update-formula:
	@echo "To update the formula:"
	@echo "1. Create a new release tag: git tag v1.x.x && git push origin v1.x.x"
	@echo "2. Download the tarball and compute SHA256:"
	@echo "   curl -sL https://github.com/samit22/calendarN/archive/refs/tags/v1.x.x.tar.gz | shasum -a 256"
	@echo "3. Update Formula/calendarn.rb with new version and SHA256"
