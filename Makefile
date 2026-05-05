.PHONY: build
build:
	go build ./...

clean:
	go clean

install_lint:
	# golangci-lint v1.64.8 was the last v1.x release and was built with Go 1.24,
	# so it can't analyze Go 1.25 code. Bumped to v2.x.
	# TODO: full v2 migration — fix pre-existing issues exposed by v2's broader default
	# linter set (errcheck, govet, staticcheck/QF1008, goconst hits in legacy actor parsers,
	# gosec G115 in tools/). Run `make lint` without --default=none to see the backlog (~70).
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v2.12.1

check-modtidy:
	go mod tidy
	git diff --exit-code -- go.mod go.sum

lint:
	golangci-lint --version
	# Linter set + exclusions live in .golangci.yml (v2 config format).
	# v2 reclassified gofmt as a formatter (no longer -E'able) and added stricter
	# defaults; .golangci.yml restores the v1.64.8 behavior with TODO suppressions
	# documented inline.
	golangci-lint run --timeout 5m

test:
	go test -timeout 120m -race ./...
