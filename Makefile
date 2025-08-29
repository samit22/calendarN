
build:
	go mod download && go build .

test:
	mkdir -p coverage && go test ./... --cover -coverprofile coverage/coverage.out

create-tweet:
	source .env && go run ./compose