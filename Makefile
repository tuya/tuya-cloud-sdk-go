export GO111MODULE=on

.DEFAULT_GOAL := test

.PHONY: test
test:
	go test -v -race -cover `go list ./...`

LINTER := bin/golangci-lint
$(LINTER):
	wget -q -O- https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.13

.PHONY: lint
lint: $(LINTER) ./golangci.yml  ## Run the linters
	@echo "linting..."
	$(LINTER) run --config ./golangci.yml

.PHONY: all
all: test lint
