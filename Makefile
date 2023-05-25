#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
MOCKSDIR ?= $(CURDIR)/mocks
SIMAPP = ./app
HTTPS_GIT := https://github.com/noria-net/noria.git
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf

ifneq ($(OS),Windows_NT)
	UNAME_S = $(shell uname -s)
endif

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
	ifeq ($(OS),Windows_NT)
		GCCEXE = $(shell where gcc.exe 2> NUL)
		ifeq ($(GCCEXE),)
			$(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
		else
			build_tags += ledger
		endif
	else
		ifeq ($(UNAME_S),OpenBSD)
			$(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
		else
			GCC = $(shell command -v gcc 2> /dev/null)
			ifeq ($(GCC),)
				$(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
			else
				build_tags += ledger
			endif
		endif
	endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
	build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=noria \
			-X github.com/cosmos/cosmos-sdk/version.AppName=noriad \
			-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
			-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
			-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
	ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
	ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
	CGO_ENABLED=1
	BUILD_TAGS += rocksdb
	ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
	BUILD_TAGS += boltdb
	ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif

ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
	ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
	BUILD_FLAGS += -trimpath
endif

install: go.sum 
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/noriad

build:
	(which mockery || go install github.com/vektra/mockery/v2@latest)
	go generate ./...
	echo $(BUILD_FLAGS)
	go build $(BUILD_FLAGS) -o ./build/noriad ./cmd/noriad

clean:
	rm -rf \
		$(BUILDDIR)/ \
		artifacts/ \
		tmp-swagger-gen/

rebuild: clean build

docker-build:
	$(DOCKER) build -f ./docker/Dockerfile.build -t noria\:$(VERSION) .; \
	$(DOCKER) run --rm -v $(CURDIR)\:/code noria\:$(VERSION) make build;

docker-test:
	$(DOCKER) build -f ./docker/Dockerfile.test -t noria_test\:$(VERSION) .; \
	$(DOCKER) run noria_test\:$(VERSION);

script-lint:
	@command -v shellcheck >/dev/null 2>&1 || { echo >&2 "shellcheck not found. Installing..."; sudo apt-get install shellcheck; }
	shellcheck scripts/estimate_block_height.sh
	shellcheck scripts/test_state_sync.sh

test:
	(which mockery || go install github.com/vektra/mockery/v2@latest)
	go generate ./...
	go test ./... -v
