# Application
APP_NAME := smart-table
MAIN_FILE := build/services/smart-table/src/main.go
BUILD_DIR := build/bin

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

# SQLC Paths
SQL_QUERIES := src/domains/orders/infra/pg/sql
SQL_SCHEMA := postgresql/smart-table
SQL_OUTPUT := src/domains/orders/infra/pg/codegen

.PHONY: all
all: build

.PHONY: build
build: gen
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

.PHONY: run
run: build
	$(BUILD_DIR)/$(APP_NAME)

.PHONY: test
test:
	$(GO) test ./... -v -cover

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	find $(BUILD_DIR) -type f -delete
	find $(OUTPUT_VIEWS) -type f -name "gen.go" -delete
	find $(OUTPUT_MODELS_INTERNAL) -type f -name "gen.go" -delete
	find $(OUTPUT_MODELS_DEFINITIONS) -type f -name "gen.go" -delete
	find $(SQL_OUTPUT) -type f -name "*.go" -delete
	find $(OUTPUT_VIEWS) -type d -mindepth 1 -empty -delete
	find $(OUTPUT_MODELS_INTERNAL) -type d -empty -delete
	find $(OUTPUT_MODELS_DEFINITIONS) -type d -empty -delete
	find $(SQL_OUTPUT) -type d -empty -delete
	mkdir -p $(OUTPUT_VIEWS)

.PHONY: deps
deps:
	$(GO) mod tidy

.PHONY: gen
gen: gen-views gen-models gen-sqlc
	$(GO) mod tidy

.PHONY: gen-views
gen-views:
	@for file in $(API_SPECS); do \
		base_name=$$(basename $$file .yaml); \
		camel_case_name=$$(echo $$base_name | sed -E 's/(^|_)([a-z])/\U\2/g' | sed -E 's/^(.)/\L\1/') ; \
		dir_path="$(OUTPUT_VIEWS)/$$base_name"; \
		package_name="views$$camel_case_name"; \
		output_file="$$dir_path/gen.go"; \
		mkdir -p $$dir_path; \
		echo "Generating $$output_file with package $$package_name from $$file"; \
		oapi-codegen --config $(OAPI_CONFIG) -o $$output_file -package $$package_name $$file; \
	done

.PHONY: gen-models
gen-models:
	@for file in $(wildcard $(DOCS_INTERNAL)); do \
		base_name=$$(basename $$file .yaml); \
		camel_case_name=$$(echo $$base_name | sed -E 's/(^|_)([a-z])/\U\2/g' | sed -E 's/^(.)/\L\1/') ; \
		dir_path="$(OUTPUT_MODELS_INTERNAL)/$$base_name"; \
		package_name="defsInternal$$camel_case_name"; \
		output_file="$$dir_path/gen.go"; \
		mkdir -p $$dir_path; \
		echo "Generating $$output_file with package $$package_name from $$file"; \
		oapi-codegen --config $(OAPI_MODELS_CONFIG) -o $$output_file -package $$package_name $$file; \
	done

	@for file in $(wildcard $(DOCS_DEFINITIONS)); do \
		base_name=$$(basename $$file .yaml); \
		camel_case_name=$$(echo $$base_name | sed -E 's/(^|_)([a-z])/\U\2/g' | sed -E 's/^(.)/\L\1/') ; \
		dir_path="$(OUTPUT_MODELS_DEFINITIONS)/$$base_name"; \
		package_name="defsDefinitions$$camel_case_name"; \
		output_file="$$dir_path/gen.go"; \
		mkdir -p $$dir_path; \
		echo "Generating $$output_file with package $$package_name from $$file"; \
		oapi-codegen --config $(OAPI_MODELS_CONFIG) -o $$output_file -package $$package_name $$file; \
	done

.PHONY: gen-sqlc
gen-sqlc:
	mkdir -p $(SQL_OUTPUT)
	echo "Generating SQL code from $(SQLC_CONFIG)..."
	sqlc generate -f $(SQLC_CONFIG)

.PHONY: mocks
mocks:
	mockery --config .mockery.yaml
	@go mod tidy
