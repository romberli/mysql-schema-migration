# if enable debug mode: 0-disable, 1-enable
DEBUG ?= 0
# available platforms: linux-amd64, linux-arm64, darwin-amd64, darwin-arm64, windows-amd64
PLATFORM ?= linux-amd64

GOPROXY ?= goproxy.cn,direct

ifeq ($(DEBUG), 0)
	GC_FLAGS :=
else ifeq ($(DEBUG), 1)
	GC_FLAGS := -gcflags "all=-N -l"
else
    $(error unknown debug mode: $(DEBUG))
endif

ifeq ($(PLATFORM), linux-amd64)
	OS:=linux
	ARCH:=amd64
else ifeq ($(PLATFORM), linux-arm64)
	OS:=linux
	ARCH:=arm64
else ifeq ($(PLATFORM), darwin-amd64)
	OS:=darwin
	ARCH:=amd64
else ifeq ($(PLATFORM), darwin-arm64)
	OS:=darwin
	ARCH:=arm64
else ifeq ($(PLATFORM), windows-amd64)
	OS:=windows
	ARCH:=amd64
else
    $(error unknown platform: $(PLATFORM))
endif

COMPILE_OPTION := CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) GODEBUG=rsa1024min=0

APP_NAME := msm
APP_BIN := ./bin/$(APP_NAME)-$(PLATFORM)
GIT_TAG := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date "+%Y-%m-%d %H:%M:%S")
FULL_COMMIT := $(shell git rev-parse HEAD)
GIT_BRANCH  := $(shell git rev-parse --abbrev-ref HEAD)

LD_FLAGS := -ldflags " \
	-X 'github.com/romberli/mysql-schema-migration/config.AppName=$(APP_NAME)' \
	-X 'github.com/romberli/mysql-schema-migration/config.Version=$(GIT_TAG)' \
	-X 'github.com/romberli/mysql-schema-migration/config.BuildTime=$(BUILD_TIME)' \
	-X 'github.com/romberli/mysql-schema-migration/config.FullCommit=$(FULL_COMMIT)' \
	-X 'github.com/romberli/mysql-schema-migration/config.Branch=$(GIT_BRANCH)' \
"

.PHONY: default
default: all

.PHONY: all
all: build-linux-amd64 build-linux-arm64

.PHONY: prepare
GO_VERSION_MIN=1.24
VER_TO_INT:=awk '{split(substr($$0, match ($$0, /[0-9\.]+/)), a, "."); print a[1]*10000+a[2]*100+a[3]}'
prepare:
	@echo "preparing started..."
	@mkdir -p ./bin
	@if test $(shell go version | $(VER_TO_INT) ) -lt \
	$(shell echo "$(GO_VERSION_MIN)" | $(VER_TO_INT)); \
	then echo "go version $(GO_VERSION_MIN)+ required, found: "; go version; exit 1; \
		else echo "preparing completed.";	fi

.PHONY: build
build: prepare
	@echo "building started..."
	@export  GO111MODULE=on
	@export  GOPROXY=$(GOPROXY)
	$(COMPILE_OPTION) go build -v $(GC_FLAGS) $(LD_FLAGS) -o $(APP_BIN) main.go
	@echo "building completed."

.PHONY: build-linux-amd64
build-linux-amd64: prepare
	@echo "building linux amd64 started..."
	@rm -rf build-tool/msm-linux-amd64
	@$(MAKE) build \
    		PLATFORM=linux-amd64 \
    		APP_BIN=./bin/$(APP_NAME)-linux-amd64
	@mkdir -p build-tool/msm-linux-amd64/bin
	@cp bin/msm-linux-amd64 build-tool/msm-linux-amd64/bin/msm
	@cp -r build-tool/default/config build-tool/msm-linux-amd64/config
	@cp build-tool/default/*.sh build-tool/msm-linux-amd64/
	@cp build-tool/default/source.sql build-tool/msm-linux-amd64/
	@cp build-tool/default/target.sql build-tool/msm-linux-amd64/
	@cp build-tool/default/README.md build-tool/msm-linux-amd64/
	zip -r build-tool/msm-linux-amd64.zip build-tool/msm-linux-amd64
	@echo "building linux amd64 completed."

.PHONY: build-linux-arm64
build-linux-arm64: prepare
	@echo "building linux arm64 started..."
	@rm -rf build-tool/msm-linux-arm64
	@$(MAKE) build \
    		PLATFORM=linux-arm64 \
    		APP_BIN=./bin/$(APP_NAME)-linux-arm64
	@mkdir -p build-tool/msm-linux-arm64/bin
	@cp bin/msm-linux-arm64 build-tool/msm-linux-arm64/bin/msm
	@cp -r build-tool/default/config build-tool/msm-linux-arm64/config
	@cp build-tool/default/*.sh build-tool/msm-linux-arm64/
	@cp build-tool/default/source.sql build-tool/msm-linux-amd64/
	@cp build-tool/default/target.sql build-tool/msm-linux-amd64/
	@cp build-tool/default/README.md build-tool/msm-linux-arm64/
	zip -r build-tool/msm-linux-arm64.zip build-tool/msm-linux-arm64
	@echo "building linux arm64 completed."

.PHONY: build-darwin-arm64
build-darwin-arm64: prepare
	@echo "building darwin arm64 started..."
	@rm -rf build-tool/msm-darwin-arm64
	@$(MAKE) build \
    		PLATFORM=darwin-arm64 \
    		APP_BIN=./bin/$(APP_NAME)-darwin-arm64
	@mkdir -p build-tool/msm-darwin-arm64/bin
	@cp bin/msm-darwin-arm64 build-tool/msm-darwin-arm64/bin/msm
	@cp -r build-tool/default/config build-tool/msm-darwin-arm64/config
	@cp build-tool/default/*.sh build-tool/msm-darwin-arm64/
	@cp build-tool/default/source.sql build-tool/msm-darwin-arm64/
	@cp build-tool/default/target.sql build-tool/msm-darwin-arm64/
	@cp build-tool/default/README.md build-tool/msm-darwin-arm64/
	zip -r build-tool/msm-darwin-arm64.zip build-tool/msm-darwin-arm64
	@echo "building darwin arm64 completed."

# Cleans our projects: deletes binaries
.PHONY: clean
clean:
	@echo "cleaning files created during building started..."
	rm -f ./bin/*
	rm -rf ./build-tool/msm-linux-amd64/
	rm -rf ./build-tool/msm-linux-arm64/
	rm -rf ./build-tool/msm-darwin-arm64/
	rm -rf ./build-tool/msm-linux-amd64.zip
	rm -rf ./build-tool/msm-linux-arm64.zip
	rm -rf ./build-tool/msm-darwin-arm64.zip
	@echo "cleaning files created during building completed."