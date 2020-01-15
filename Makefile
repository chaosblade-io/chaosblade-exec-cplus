.PHONY: build clean

BLADE_SRC_ROOT=`pwd`
UNAME := $(shell uname)

ifeq ($(BLADE_VERSION), )
	BLADE_VERSION=0.4.0
endif

TARGET_PATH=target
BUILD_TARGET=build-target
BUILD_TARGET_DIR_NAME=chaosblade-$(BLADE_VERSION)
BUILD_TARGET_PKG_DIR=$(BUILD_TARGET)/chaosblade-$(BLADE_VERSION)
BUILD_TARGET_BIN=$(BUILD_TARGET_PKG_DIR)/bin
BUILD_TARGET_LIB=$(BUILD_TARGET_PKG_DIR)/lib/cplus
# cache downloaded file
BUILD_TARGET_CACHE=$(BUILD_TARGET)/cache
# yaml file name
CPLUS_YAML_FILE_NAME=chaosblade-cplus-spec.yaml
# agent file name
CPLUS_AGENT_FILE_NAME=chaosblade-exec-cplus.jar

build: pre_build build_cplus
	cp $(TARGET_PATH)/$(CPLUS_AGENT_FILE_NAME) $(BUILD_TARGET_LIB)
	cp -R $(TARGET_PATH)/classes/script $(BUILD_TARGET_LIB)
	cp $(TARGET_PATH)/classes/$(CPLUS_YAML_FILE_NAME) $(BUILD_TARGET_BIN)
	chmod -R 755 $(BUILD_TARGET_LIB)

pre_build:
	rm -rf $(BUILD_TARGET_PKG_DIR)
	mkdir -p $(BUILD_TARGET_BIN) $(BUILD_TARGET_LIB)

build_cplus:
	mvn clean package -Dmaven.test.skip=true -U

# test
test:
	mvn clean test -U
# clean all build result
clean:
	rm -rf $(BUILD_TARGET)

all: build test
