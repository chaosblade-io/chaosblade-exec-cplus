.PHONY: build clean

BLADE_SRC_ROOT=$(shell pwd)
UNAME := $(shell uname)

ifeq ($(BLADE_VERSION), )
	BLADE_VERSION=1.7.3
endif

BUILD_TARGET=build-target
BUILD_TARGET_DIR_NAME=chaosblade-$(BLADE_VERSION)
BUILD_TARGET_PKG_DIR=$(BUILD_TARGET)/chaosblade-$(BLADE_VERSION)
BUILD_TARGET_YAML=$(BUILD_TARGET_PKG_DIR)/yaml
BUILD_TARGET_CPLUS_LIB=$(BUILD_TARGET_PKG_DIR)/lib/cplus
BUILD_TARGET_CPLUS_SCRIPT=$(BUILD_TARGET_CPLUS_LIB)/script
# yaml file name
CPLUS_YAML_FILE=$(BUILD_TARGET_YAML)/chaosblade-cplus-spec.yaml
# agent file name
CPLUS_AGENT_FILE_NAME=chaosblade-exec-cplus

GO_ENV=CGO_ENABLED=1
GO_MODULE=GO111MODULE=on
GO=env $(GO_ENV) $(GO_MODULE) go
ifeq ($(GOOS), linux)
	GO_FLAGS=-ldflags="-linkmode external -extldflags -static"
endif

build: pre_build build_cplus build_yaml
	cp -R script/* $(BUILD_TARGET_CPLUS_SCRIPT)
	chmod -R 755 $(BUILD_TARGET_CPLUS_LIB)

pre_build:
	rm -rf $(BUILD_TARGET_PKG_DIR)
	mkdir -p $(BUILD_TARGET_YAML) $(BUILD_TARGET_CPLUS_SCRIPT)

build_yaml: build/spec.go
	$(GO) run $< $(CPLUS_YAML_FILE)

build_cplus: main.go
	$(GO) build $(GO_FLAGS) -o $(BUILD_TARGET_CPLUS_LIB)/chaosblade-exec-cplus $<

# build chaosblade linux version by docker image
build_linux:
	docker build -f build/image/musl/Dockerfile -t chaosblade-cplus-build-musl:latest build/image/musl
	docker run --rm \
		-v $(shell echo -n ${GOPATH}):/go \
		-v $(BLADE_SRC_ROOT):/chaosblade-exec-cplus \
		-w /chaosblade-exec-cplus \
		chaosblade-cplus-build-musl:latest

# test
test:
	mvn clean test -U
# clean all build result
clean:
	$(GO) clean ./...
	rm -rf $(BUILD_TARGET)

all: build test
