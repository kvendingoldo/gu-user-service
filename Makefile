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

##@ Generate
gen_grpc: ## Generate gRPC client/server code
	mkdir -p ./gen/go
	protoc -I . --go_out ./gen/go/ --go_opt paths=source_relative --go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative ./api/v1/service.proto

gen_grpc_gw: ## Generate gRPC gateway code
	mkdir -p ./gen/go
	protoc -I . --grpc-gateway_out ./gen/go --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./api/v1/service.proto

gen: gen_grpc gen_grpc_gw
	@echo "Generate code $(OK_STRING)"

##@ Build
build: gen fmt vet  ## Build service binary.
	go build -o bin/service cmd/app/main.go

run: gen fmt vet ## Run service from your laptop.
	go run ./cmd/app/main.go

##@ Test
lint: ## Run Go linter
	~/go/bin/golangci-lint run ./...

test: ## Run Go tests
	go test ./...

##@ Install
install-deps:
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go
	go get \
		github.com/bufbuild/buf/cmd/buf \
		github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking \
		github.com/bufbuild/buf/cmd/protoc-gen-buf-lint
	@echo "Install dependencies $(OK_STRING)"


# go mod !!!