GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO_MOD=$(GOCMD) mod
GO_ENV=$(GOCMD) env
BINARY_NAME=irita-sync
BINARY_UNIX=$(BINARY_NAME)-unix
export GO111MODULE = on
build_tags = netgo
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))
BUILD_FLAGS := -tags "$(build_tags)"
all: get_deps build

get_deps:
	@rm -rf vendor/
	@echo "--> Downloading dependencies"
	$(GO_MOD) download
	$(GO_MOD) vendor

build:
	$(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_NAME) .

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
