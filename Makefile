##@ General
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

##@ Build
swag:
	swag init -g internal/apis/rest/*/*.go --output ./swagger_gen/api

proto:
	mkdir -p ./proto_gen
	protoc --proto_path=internal/apis/grpc/v1 --go-grpc_out=./proto_gen --go_out=./proto_gen internal/apis/grpc/v1/*.proto

build: swag fmt vet ## Build service binary.
	go build -o bin/service cmd/app/main.go

run: swag fmt vet ## Run service from your laptop.
	go run ./cmd/app/main.go

##@ Test
lint:
	~/go/bin/golangci-lint run ./...

test:
	go test ./...