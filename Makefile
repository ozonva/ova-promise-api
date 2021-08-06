.PHONY: lint
default:

.PHONY: lint
lint:
	golangci-lint run --fix ./...

lint-ci:
	golangci-lint run ./...

.PHONY: test
test:
	@go test --race --vet= ./... -v
