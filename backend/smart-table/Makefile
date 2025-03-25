# Application
APP_NAME := smart-table
MAIN_FILE := build/services/smart-table/src/main.go
BUILD_DIR := build/bin
CONFIGS := $(wildcard configs/*.env) $(wildcard configs/.env)

# Go
GO := go
OAPI_CONFIG := .oapi-codegen.yaml
OAPI_MODELS_CONFIG := .oapi-codegen.models.yaml
SQLC_CONFIG := .sqlc.yaml

# Paths
API_SPECS := $(wildcard docs/api/*.yaml)
OUTPUT_VIEWS := src/views/codegen
OUTPUT_MODELS_INTERNAL := src/codegen/intern
OUTPUT_MODELS_DEFINITIONS := src/codegen/definitions
DOCS_INTERNAL := docs/internal/*.yaml
DOCS_DEFINITIONS := docs/definitions/*.yaml

# Colors
GREEN  := \033[0;32m
BLUE   := \033[0;34m
YELLOW := \033[0;33m
RED    := \033[0;31m
NC     := \033[0m

.PHONY: all
all: build

.PHONY: build
build: gen
	@echo "${BLUE}🔹 Building application...${NC}"
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "${GREEN}✅ Application built successfully!${NC}"

.PHONY: run
run: build
	@echo "${BLUE}🔹 Starting application...${NC}"
	@$(BUILD_DIR)/$(APP_NAME)

.PHONY: test
test:
	@echo "${BLUE}🔹 Running tests...${NC}"
	@$(GO) test ./... -v -cover
	@echo "${GREEN}✅ Tests completed!${NC}"

.PHONY: lint
lint:
	@echo "${BLUE}🔹 Running linter...${NC}"
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run
	@echo "${GREEN}✅ Linting completed!${NC}"

.PHONY: clean
clean:
	@echo "${YELLOW}🔸 Cleaning up...${NC}"
	@find tests/mocks -type f -delete
	@find $(BUILD_DIR) -type f -delete
	@find . -type d -name "codegen" -exec rm -rf {} +
	@mkdir -p $(OUTPUT_VIEWS)
	@echo "${GREEN}✅ Cleanup completed!${NC}"

.PHONY: deps
deps:
	@echo "${BLUE}🔹 Updating dependencies...${NC}"
	@$(GO) mod tidy
	@echo "${GREEN}✅ Dependencies updated!${NC}"

.PHONY: gen
gen: gen-views gen-models gen-sqlc generate_env mocks
	@$(GO) mod tidy
	@echo "${GREEN}✅ All code generation completed!${NC}"

.PHONY: gen-views
gen-views:
	@echo "${BLUE}🔹 Generating views code...${NC}"
	@for file in $(API_SPECS); do \
		base_name=$$(basename $$file .yaml); \
		camel_case_name=$$(echo $$base_name | sed -E 's/(^|_)([a-z])/\U\2/g' | sed -E 's/^(.)/\L\1/') ; \
		dir_path="$(OUTPUT_VIEWS)/$$base_name"; \
		package_name="views$$(echo $$base_name | sed 's/_//g')"; \
		output_file="$$dir_path/gen.go"; \
		mkdir -p $$dir_path; \
		echo "${BLUE}🔸 Generating $$output_file with package $$package_name from $$file${NC}"; \
		go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config $(OAPI_CONFIG) -o $$output_file -package $$package_name $$file || exit 1; \
	done
	@echo "${GREEN}✅ Views generation completed!${NC}"

.PHONY: gen-models
gen-models:
	@echo "${BLUE}🔹 Generating models code...${NC}"
	@for file in $(wildcard $(DOCS_INTERNAL)); do \
		base_name=$$(basename $$file .yaml); \
		package_name="defsinternal$$(echo $$base_name | sed 's/_//g')"; \
		dir_path="$(OUTPUT_MODELS_INTERNAL)/$$base_name"; \
		output_file="$$dir_path/gen.go"; \
		mkdir -p $$dir_path; \
		echo "${BLUE}🔸 Generating $$output_file with package $$package_name from $$file${NC}"; \
		go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config $(OAPI_MODELS_CONFIG) -o $$output_file -package $$package_name $$file || exit 1; \
	done

	@for file in $(wildcard $(DOCS_DEFINITIONS)); do \
		base_name=$$(basename $$file .yaml); \
		package_name="defsdefinitions$$(echo $$base_name | sed 's/_//g')"; \
		dir_path="$(OUTPUT_MODELS_DEFINITIONS)/$$base_name"; \
		output_file="$$dir_path/gen.go"; \
		mkdir -p $$dir_path; \
		echo "${BLUE}🔸 Generating $$output_file with package $$package_name from $$file${NC}"; \
		go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config $(OAPI_MODELS_CONFIG) -o $$output_file -package $$package_name $$file || exit 1; \
	done
	@echo "${GREEN}✅ Models generation completed!${NC}"

.PHONY: gen-sqlc
gen-sqlc:
	@echo "${BLUE}🔹 Generating SQL code...${NC}"
	@echo "${BLUE}🔸 Using config: $(SQLC_CONFIG)${NC}"
	@go run github.com/sqlc-dev/sqlc/cmd/sqlc generate -f $(SQLC_CONFIG)
	@echo "${GREEN}✅ SQL code generated successfully!${NC}"

.PHONY: mocks
mocks:
	@echo "${BLUE}🔹 Generating mocks...${NC}"
	@go run github.com/vektra/mockery/v2 --config .mockery.yaml
	@$(GO) mod tidy
	@echo "${GREEN}✅ Mocks generated successfully!${NC}"

.PHONY: generate_env
generate_env:
	@echo "${BLUE}🔹 Generating .env files...${NC}"
	@chmod +x load_env.sh && ./load_env.sh
	@echo "${GREEN}✅ .env files successfully created!${NC}"

.PHONY: up
up: gen
	@echo "${BLUE}🔹 Running docker compose...${NC}"
	@echo "${BLUE}🔸 Using env files: $(CONFIGS)${NC}"
	@docker compose \
		$(foreach envfile, $(CONFIGS), --env-file=$(envfile)) \
		up --build
