NO_COLOR=\x1b[0m
OK_COLOR=\x1b[32;01m
ERROR_COLOR=\x1b[31;01m
WARN_COLOR=\x1b[33;01m
OK_STRING=$(OK_COLOR)[OK]$(NO_COLOR)
ERROR_STRING=$(ERROR_COLOR)[ERRORS]$(NO_COLOR)
WARN_STRING=$(WARN_COLOR)[WARNINGS]$(NO_COLOR)

SWAGGERUI_VERSION = 3.22.3

##@ General
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

##@ Generate
gen: static/swagger-ui ## Generate API
	buf generate
	@echo "Generate API $(OK_STRING)"

##@ Build
build: gen fmt vet  ## Build service binary.
	go build -o bin/service cmd/app/main.go

run: gen fmt vet ## Run service from your laptop.
	go run ./cmd/app/main.go

##@ Test
lint: ## Run Go linter
	golangci-lint run ./...
	buf lint

test: ## Run Go tests
	go test ./...

##@ Install
install-deps: ## Install project dependencies
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go
	go install \
		github.com/bufbuild/buf/cmd/buf \
		github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking \
		github.com/bufbuild/buf/cmd/protoc-gen-buf-lint
	@echo "Install dependencies $(OK_STRING)"

static/swagger-ui:
	mkdir -p /tmp/swaggerui-temp/ || echo "Folder exists"
	curl -o /tmp/swaggerui-temp/swaggerui.zip -L \
		https://github.com/swagger-api/swagger-ui/archive/v$(SWAGGERUI_VERSION).zip
	(cd /tmp/swaggerui-temp/; unzip -q -o swaggerui.zip)
	cp -r /tmp/swaggerui-temp/swagger-ui-$(SWAGGERUI_VERSION)/dist/ ./static/swagger-ui/
	rm -rf /tmp/swaggerui-temp