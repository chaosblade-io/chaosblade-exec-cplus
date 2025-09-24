.PHONY: build clean linux_amd64 linux_arm64 darwin_amd64 darwin_arm64 windows_amd64 help

BLADE_SRC_ROOT=$(shell pwd)
UNAME := $(shell uname)

# Version management - auto-detect from git tags
ifeq ($(BLADE_VERSION), )
	BLADE_VERSION=$(shell ./version/version.sh version)
endif

# Additional version information
GIT_COMMIT=$(shell ./version/version.sh commit)
BUILD_TIME=$(shell ./version/version.sh build-time)
BUILD_TYPE=$(shell ./version/version.sh build-type)
FULL_VERSION=$(shell ./version/version.sh full-version)

BUILD_TARGET=target
BUILD_TARGET_DIR_NAME=chaosblade-$(BLADE_VERSION)
BUILD_TARGET_PKG_DIR=$(BUILD_TARGET)/chaosblade-$(BLADE_VERSION)
BUILD_TARGET_YAML=$(BUILD_TARGET_PKG_DIR)/yaml
BUILD_TARGET_CPLUS_LIB=$(BUILD_TARGET_PKG_DIR)/lib/cplus
BUILD_TARGET_CPLUS_SCRIPT=$(BUILD_TARGET_CPLUS_LIB)/script

# Platform-specific directory functions
define get_platform_dir_name
chaosblade-$(BLADE_VERSION)-$(1)
endef

define get_platform_pkg_dir
$(BUILD_TARGET)/chaosblade-$(BLADE_VERSION)-$(1)
endef

define get_platform_yaml_dir
$(BUILD_TARGET)/chaosblade-$(BLADE_VERSION)-$(1)/yaml
endef

define get_platform_lib_dir
$(BUILD_TARGET)/chaosblade-$(BLADE_VERSION)-$(1)/lib/cplus
endef

define get_platform_script_dir
$(BUILD_TARGET)/chaosblade-$(BLADE_VERSION)-$(1)/lib/cplus/script
endef
# yaml file name (will be overridden by platform-specific targets)
CPLUS_YAML_FILE=$(BUILD_TARGET_YAML)/chaosblade-cplus-spec-$(BLADE_VERSION).yaml
# agent file name
CPLUS_AGENT_FILE_NAME=chaosblade-exec-cplus

GO_ENV=CGO_ENABLED=0
GO_MODULE=GO111MODULE=on
GO=env $(GO_ENV) $(GO_MODULE) go

# Cross-compilation GO command (without CGO)
GO_CROSS=env CGO_ENABLED=0 GO111MODULE=on go

# Host platform for running generators (ensure go run executes on host arch)
HOST_GOOS=$(shell go env GOHOSTOS)
HOST_GOARCH=$(shell go env GOHOSTARCH)

# Build flags for different platforms with version information
GO_FLAGS_LINUX_AMD64=-ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)"
GO_FLAGS_LINUX_ARM64=-ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)"
GO_FLAGS_DARWIN_AMD64=-ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)"
GO_FLAGS_DARWIN_ARM64=-ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)"
GO_FLAGS_WINDOWS_AMD64=-ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)"

# Common build flags
GO_FLAGS_COMMON=-ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)"

# Help target (default)
help:
	@echo "ChaosBlade C++ Executor Build System"
	@echo "===================================="
	@echo ""
	@echo "Version Information:"
	@echo "  Version:         $(BLADE_VERSION)"
	@echo "  Git Commit:      $(GIT_COMMIT)"
	@echo "  Build Time:      $(BUILD_TIME)"
	@echo "  Build Type:      $(BUILD_TYPE)"
	@echo "  Full Version:    $(FULL_VERSION)"
	@echo ""
	@echo "Available Build Targets:"
	@echo "  build          - Build current platform version"
	@echo "  build_all      - Build all platform versions"
	@echo "  linux_amd64    - Build Linux AMD64 version"
	@echo "  linux_arm64    - Build Linux ARM64 version"
	@echo "  darwin_amd64   - Build macOS AMD64 version"
	@echo "  darwin_arm64   - Build macOS ARM64 version"
	@echo "  windows_amd64  - Build Windows AMD64 version"
	@echo ""
	@echo "Other Commands:"
	@echo "  test           - Run tests"
	@echo "  format         - Format Go code using goimports and gofumpt"
	@echo "  verify         - Verify Go code formatting and import order"
	@echo "  clean          - Clean build products"
	@echo "  all            - Build and test"
	@echo "  help           - Show this help information"
	@echo "  version        - Show version information"
	@echo ""
	@echo "Environment Variables:"
	@echo "  BLADE_VERSION  - Specify build version (default: auto-detect from Git Tag)"
	@echo ""
	@echo "Usage Examples:"
	@echo "  make help                    # Show help"
	@echo "  make build                   # Build current platform version"
	@echo "  make build_all               # Build all platform versions"
	@echo "  make linux_arm64            # Build Linux ARM64 version"
	@echo "  BLADE_VERSION=1.8.0 make build  # Build with specified version"
	@echo ""

# Default target
.DEFAULT_GOAL := help

# Version info target
version:
	@echo "Version Information:"
	@echo "  Version:         $(BLADE_VERSION)"
	@echo "  Git Commit:      $(GIT_COMMIT)"
	@echo "  Build Time:      $(BUILD_TIME)"
	@echo "  Build Type:      $(BUILD_TYPE)"
	@echo "  Full Version:    $(FULL_VERSION)"
	@echo "  Is Tagged:       $(shell ./version/version.sh is-tagged)"

build: pre_build build_cplus build_yaml
	cp -R script/* $(BUILD_TARGET_CPLUS_SCRIPT)
	chmod -R 755 $(BUILD_TARGET_CPLUS_LIB)

pre_build:
	rm -rf $(BUILD_TARGET_PKG_DIR)
	mkdir -p $(BUILD_TARGET_YAML) $(BUILD_TARGET_CPLUS_SCRIPT)

build_yaml: build/spec.go
	GOOS=$(HOST_GOOS) GOARCH=$(HOST_GOARCH) $(GO) run -ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)" $< $(CPLUS_YAML_FILE)

build_cplus: main.go
	$(GO) build $(GO_FLAGS_COMMON) -o $(BUILD_TARGET_CPLUS_LIB)/$(CPLUS_AGENT_FILE_NAME) $<

# Multi-platform build targets
linux_amd64:
	$(eval PLATFORM := linux_amd64)
	$(eval PLATFORM_PKG_DIR := $(call get_platform_pkg_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_DIR := $(call get_platform_yaml_dir,$(PLATFORM)))
	$(eval PLATFORM_LIB_DIR := $(call get_platform_lib_dir,$(PLATFORM)))
	$(eval PLATFORM_SCRIPT_DIR := $(call get_platform_script_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_FILE := $(PLATFORM_YAML_DIR)/chaosblade-cplus-spec-$(BLADE_VERSION).yaml)
	rm -rf $(PLATFORM_PKG_DIR)
	mkdir -p $(PLATFORM_YAML_DIR) $(PLATFORM_SCRIPT_DIR)
	GOOS=linux GOARCH=amd64 $(GO_CROSS) build $(GO_FLAGS_LINUX_AMD64) -o $(PLATFORM_LIB_DIR)/$(CPLUS_AGENT_FILE_NAME) main.go
	GOOS=$(HOST_GOOS) GOARCH=$(HOST_GOARCH) $(GO) run -ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)" build/spec.go $(PLATFORM_YAML_FILE)
	cp -R script/* $(PLATFORM_SCRIPT_DIR)
	chmod -R 755 $(PLATFORM_LIB_DIR)

linux_arm64:
	$(eval PLATFORM := linux_arm64)
	$(eval PLATFORM_PKG_DIR := $(call get_platform_pkg_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_DIR := $(call get_platform_yaml_dir,$(PLATFORM)))
	$(eval PLATFORM_LIB_DIR := $(call get_platform_lib_dir,$(PLATFORM)))
	$(eval PLATFORM_SCRIPT_DIR := $(call get_platform_script_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_FILE := $(PLATFORM_YAML_DIR)/chaosblade-cplus-spec-$(BLADE_VERSION).yaml)
	rm -rf $(PLATFORM_PKG_DIR)
	mkdir -p $(PLATFORM_YAML_DIR) $(PLATFORM_SCRIPT_DIR)
	GOOS=linux GOARCH=arm64 $(GO_CROSS) build $(GO_FLAGS_LINUX_ARM64) -o $(PLATFORM_LIB_DIR)/$(CPLUS_AGENT_FILE_NAME) main.go
	GOOS=$(HOST_GOOS) GOARCH=$(HOST_GOARCH) $(GO) run -ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)" build/spec.go $(PLATFORM_YAML_FILE)
	cp -R script/* $(PLATFORM_SCRIPT_DIR)
	chmod -R 755 $(PLATFORM_LIB_DIR)

darwin_amd64:
	$(eval PLATFORM := darwin_amd64)
	$(eval PLATFORM_PKG_DIR := $(call get_platform_pkg_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_DIR := $(call get_platform_yaml_dir,$(PLATFORM)))
	$(eval PLATFORM_LIB_DIR := $(call get_platform_lib_dir,$(PLATFORM)))
	$(eval PLATFORM_SCRIPT_DIR := $(call get_platform_script_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_FILE := $(PLATFORM_YAML_DIR)/chaosblade-cplus-spec-$(BLADE_VERSION).yaml)
	rm -rf $(PLATFORM_PKG_DIR)
	mkdir -p $(PLATFORM_YAML_DIR) $(PLATFORM_SCRIPT_DIR)
	GOOS=darwin GOARCH=amd64 $(GO) build $(GO_FLAGS_DARWIN_AMD64) -o $(PLATFORM_LIB_DIR)/$(CPLUS_AGENT_FILE_NAME) main.go
	GOOS=$(HOST_GOOS) GOARCH=$(HOST_GOARCH) $(GO) run -ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)" build/spec.go $(PLATFORM_YAML_FILE)
	cp -R script/* $(PLATFORM_SCRIPT_DIR)
	chmod -R 755 $(PLATFORM_LIB_DIR)

darwin_arm64:
	$(eval PLATFORM := darwin_arm64)
	$(eval PLATFORM_PKG_DIR := $(call get_platform_pkg_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_DIR := $(call get_platform_yaml_dir,$(PLATFORM)))
	$(eval PLATFORM_LIB_DIR := $(call get_platform_lib_dir,$(PLATFORM)))
	$(eval PLATFORM_SCRIPT_DIR := $(call get_platform_script_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_FILE := $(PLATFORM_YAML_DIR)/chaosblade-cplus-spec-$(BLADE_VERSION).yaml)
	rm -rf $(PLATFORM_PKG_DIR)
	mkdir -p $(PLATFORM_YAML_DIR) $(PLATFORM_SCRIPT_DIR)
	GOOS=darwin GOARCH=arm64 $(GO) build $(GO_FLAGS_DARWIN_ARM64) -o $(PLATFORM_LIB_DIR)/$(CPLUS_AGENT_FILE_NAME) main.go
	GOOS=$(HOST_GOOS) GOARCH=$(HOST_GOARCH) $(GO) run -ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)" build/spec.go $(PLATFORM_YAML_FILE)
	cp -R script/* $(PLATFORM_SCRIPT_DIR)
	chmod -R 755 $(PLATFORM_LIB_DIR)

windows_amd64:
	$(eval PLATFORM := windows_amd64)
	$(eval PLATFORM_PKG_DIR := $(call get_platform_pkg_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_DIR := $(call get_platform_yaml_dir,$(PLATFORM)))
	$(eval PLATFORM_LIB_DIR := $(call get_platform_lib_dir,$(PLATFORM)))
	$(eval PLATFORM_SCRIPT_DIR := $(call get_platform_script_dir,$(PLATFORM)))
	$(eval PLATFORM_YAML_FILE := $(PLATFORM_YAML_DIR)/chaosblade-cplus-spec-$(BLADE_VERSION).yaml)
	rm -rf $(PLATFORM_PKG_DIR)
	mkdir -p $(PLATFORM_YAML_DIR) $(PLATFORM_SCRIPT_DIR)
	GOOS=windows GOARCH=amd64 $(GO_CROSS) build $(GO_FLAGS_WINDOWS_AMD64) -o $(PLATFORM_LIB_DIR)/$(CPLUS_AGENT_FILE_NAME).exe main.go
	GOOS=$(HOST_GOOS) GOARCH=$(HOST_GOARCH) $(GO) run -ldflags="-X main.Version=$(BLADE_VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.BuildType=$(BUILD_TYPE)" build/spec.go $(PLATFORM_YAML_FILE)
	cp -R script/* $(PLATFORM_SCRIPT_DIR)
	chmod -R 755 $(PLATFORM_LIB_DIR)

# test
test:
	mvn clean test -U

# clean all build result
clean:
	$(GO) clean ./...
	rm -rf $(BUILD_TARGET)

# Build all platforms
build_all: linux_amd64 linux_arm64 darwin_amd64 darwin_arm64 windows_amd64
	@echo "=========================================="
	@echo "All platform builds completed successfully!"
	@echo "Generated directories:"
	@echo "  - chaosblade-$(BLADE_VERSION)-linux_amd64"
	@echo "  - chaosblade-$(BLADE_VERSION)-linux_arm64"
	@echo "  - chaosblade-$(BLADE_VERSION)-darwin_amd64"
	@echo "  - chaosblade-$(BLADE_VERSION)-darwin_arm64"
	@echo "  - chaosblade-$(BLADE_VERSION)-windows_amd64"
	@echo "=========================================="

all: build test

.PHONY: format
format:
	@echo "Running goimports and gofumpt to format Go code..."
	@./hack/update-imports.sh
	@./hack/update-gofmt.sh

.PHONY: verify
verify:
	@echo "Verifying Go code formatting and import order..."
	@./hack/verify-gofmt.sh
	@./hack/verify-imports.sh