APP_NAME := smart-table
MAIN_FILE := build/services/smart-table/src/main.go
BUILD_DIR := build/bin
GO := go

.PHONY: all
all: build

.PHONY: build
build:
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
	rm -rf $(BUILD_DIR)

.PHONY: deps
deps:
	$(GO) mod tidy