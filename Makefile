NO_COLOR=\x1b[0m
OK_COLOR=\x1b[32;01m
ERROR_COLOR=\x1b[31;01m
WARN_COLOR=\x1b[33;01m
OK_STRING=$(OK_COLOR)[OK]$(NO_COLOR)
ERROR_STRING=$(ERROR_COLOR)[ERRORS]$(NO_COLOR)
WARN_STRING=$(WARN_COLOR)[WARNINGS]$(NO_COLOR)

##@ General
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

##@ Build
swag: ## Generate swagger API
	mkdir -p ./swagger_gen/api
	#swag init --parseDependency --parseInternal --generatedTime -g internal/apis/rest/*/*.go --output ./swagger_gen/api
	swag init --generatedTime --g internal/apis/rest/*/main.go --output ./swagger_gen/api

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

##@ Install
install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@echo "Install protoc $(OK_STRING)"