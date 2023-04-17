GO   = $(shell which go)
BIN  = ./bin

export GO111MODULE=on
export GCO_ENABLED=0

V = 0
Q = $(if $(filter 1,$V),,@)

## PROJECT VARIABLES
GITSHA        =$(shell git rev-parse --short HEAD)
DATE		  =$(shell date -u '+%y%m%d%I%M')
# VERSION       =$(echo $(DATE)-$(GITSHA))
VERSION       =$(shell echo $(GITSHA))

## BUILD FLAGS
GOBUILD = -a -v -trimpath='true' -buildmode='exe' -buildvcs='true' -compiler='gc' -mod='vendor'
LDFLAGS = -X github.com/blurryContour/go-webserver/version/version.cliVersion=$(VERSION)

# ==============================================================================
# MAIN TARGETS

.PHONY: bin
bin: ## Build the production binary file
	@echo "Building PROD binaries..."
	rm -rf "$(BIN)"
	mkdir -p "$(BIN)"
	@echo "VERSION: $(VERSION)"
	$(GO) generate ./...
	$(GO) build $(GOBUILD) -ldflags "$(LDFLAGS)" -o "$(BIN)/server" ./main.go

.PHONY: dev
dev: ## Build the binary file for development
	@echo "Building DEV binaries..."
	rm -rf "$(BIN)"
	mkdir -p "$(BIN)"
	@echo "VERSION: $(VERSION)"
#	$(GO) generate ./...
	go build -mod='vendor' -o "$(BIN)/server" ./main.go

.PHONY: run
run: ## Run the main.go file
	@echo "Running..."
	go run main.go

# ==============================================================================
# Install tools

install-deps: ## Install all dependencies
	@echo "installing gotest for testing https://github.com/rakyll/gotest"
	$(GO) install github.com/rakyll/gotest@latest


# ==============================================================================
# DOCS

.PHONY: docs
docs:  ## update docs
	@echo "Generating docs..."
	mkdir -p bin
	$(GO) build -o bin/docs cmd/docs/*.go
	@mkdir -p ./docs/cmd ./docs/man/man1
	@./bin/docs --target=./docs/cmd
	@./bin/docs --target=./docs/man/man1 --kind=man
	@rm -f ./bin/docs

# ==============================================================================
# HELP

.PHONY: help
help: ## Show help
	@echo Please specify a build target. The choices are:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-30s\033[0m %s\n", $$1, $$2}'