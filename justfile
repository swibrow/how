default: lint test build

build:
    go build -o how ./cmd/how

test:
    go test -race ./...

coverage:
    go test -race -coverprofile=coverage.out ./...
    go tool cover -func=coverage.out

lint:
    golangci-lint run

fmt-check:
    #!/usr/bin/env sh
    set -eu
    unformatted=$(gofmt -l .)
    if [ -n "$unformatted" ]; then
        echo "Files need formatting:"
        echo "$unformatted"
        exit 1
    fi

vet:
    go vet ./...

clean:
    rm -f how coverage.out

install:
    go install ./cmd/how
