.PHONY: all build clean build-windows build-linux build-darwin build-all build-all-arm64 run

BINARY_NAME=resizer
BUILD_DIR=build

# Определяем текущую ОС
ifeq ($(OS),Windows_NT)
    CURRENT_OS=windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        CURRENT_OS=linux
    endif
    ifeq ($(UNAME_S),Darwin)
        CURRENT_OS=darwin
    endif
endif

all: build

build:
ifeq ($(CURRENT_OS),windows)
	@if not exist "$(BUILD_DIR)" mkdir "$(BUILD_DIR)"
	go build -o $(BUILD_DIR)/$(BINARY_NAME).exe .
else
	@mkdir -p $(BUILD_DIR)
	GOOS=$(CURRENT_OS) GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) .
endif

build-windows:
	@if not exist "$(BUILD_DIR)" mkdir "$(BUILD_DIR)"
	set GOOS=windows& set GOARCH=amd64& go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .

build-linux:
	@if not exist "$(BUILD_DIR)" mkdir "$(BUILD_DIR)"
	set GOOS=linux& set GOARCH=amd64& go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .

build-darwin:
	@if not exist "$(BUILD_DIR)" mkdir "$(BUILD_DIR)"
	set GOOS=darwin& set GOARCH=amd64& go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .

build-windows-arm64:
	@if not exist "$(BUILD_DIR)" mkdir "$(BUILD_DIR)"
	set GOOS=windows& set GOARCH=arm64& go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe .

build-linux-arm64:
	@if not exist "$(BUILD_DIR)" mkdir "$(BUILD_DIR)"
	set GOOS=linux& set GOARCH=arm64& go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 .

build-darwin-arm64:
	@if not exist "$(BUILD_DIR)" mkdir "$(BUILD_DIR)"
	set GOOS=darwin& set GOARCH=arm64& go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .

build-all: build-windows build-linux build-darwin build-windows-arm64 build-linux-arm64 build-darwin-arm64

clean:
ifeq ($(CURRENT_OS),windows)
	powershell -Command "if (Test-Path '$(BUILD_DIR)') { Remove-Item -Recurse -Force '$(BUILD_DIR)' }"
else
	rm -rf $(BUILD_DIR)
endif

run: build
ifeq ($(CURRENT_OS),windows)
	$(BUILD_DIR)/$(BINARY_NAME).exe
else
	$(BUILD_DIR)/$(BINARY_NAME)
endif
