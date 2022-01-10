##@ General
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

##@ Build
swag: ## Generate swagger API
	#swag init --parseDependency --parseInternal --generatedTime -g internal/apis/rest/*/*.go --output ./swagger_gen/api
	swag init --generatedTime -g internal/apis/rest/*/*.go --output ./swagger_gen/api

proto: ## Generate proto code
	mkdir -p ./proto_gen
	protoc --proto_path=internal/apis/grpc/v1 --go-grpc_out=./proto_gen --go_out=./proto_gen internal/apis/grpc/v1/*.proto

build: swag proto fmt vet  ## Build service binary.
	go build -o bin/service cmd/app/main.go

run: swag proto fmt vet ## Run service from your laptop.
	go run ./cmd/app/main.go

##@ Test
lint: ## Run Go linter
	~/go/bin/golangci-lint run ./...

test: ## Run Go tests
	go test ./...