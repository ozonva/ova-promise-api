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

.PHONY: mock
mock:
	 mockery --all --output internal/mocks

.PHONY: gen
gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/implementation/grpc.server/protocol/promise.proto

.PHONY: migration
migration:
	migrate create -ext sql -dir migrations $(NAME)

.PHONY: build
build:
	go build -tags=jsoniter -o ./tmp/ova-promise-api ./cmd/ova-promise-api