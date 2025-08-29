
CONSUMER_KEY?=
CONSUMER_SECRET?=
ACCESS_TOKEN?=
ACCESS_TOKEN_SECRET?=

build:
	go mod download && go build .

test:
	mkdir -p coverage && go test ./... --cover -coverprofile coverage/coverage.out

create-tweet:
	CONSUMER_KEY=${CONSUMER_KEY} CONSUMER_SECRET=${CONSUMER_SECRET} ACCESS_TOKEN=${ACCESS_TOKEN} ACCESS_TOKEN_SECRET=${ACCESS_TOKEN_SECRET}  go run ./compose

.PHONY: build test create-tweet
