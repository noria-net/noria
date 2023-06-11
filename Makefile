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
CWD=$(shell pwd)
HAS_NORIA_IMG := $(shell docker images -q noria/noriad 2> /dev/null)

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

package:
	echo "Packaging noriad $(VERSION)"
	mkdir -p release/$(VERSION)
	tar -czvf release/$(VERSION)/noria_linux_amd64.tar.gz --transform='s,^build/,,' build/noriad
	cd release/$(VERSION); \
	sha256sum noria_linux_amd64.tar.gz > sha256.txt; \
	cat sha256.txt

validate_sha256:
	cd release/$(VERSION); \
	sha256sum -c sha256.txt

clean:
	rm -rf \
		$(BUILDDIR)/ \
		artifacts/ \
		tmp-swagger-gen/

rebuild: clean build

docker-build:
	$(DOCKER) build -f ./dev/docker/Dockerfile.build -t noria_build\:$(VERSION) .; \
	$(DOCKER) run --rm -v $(CURDIR)\:/code noria_build\:$(VERSION) make build;
	$(DOCKER) run --rm -v $(CURDIR)\:/code noria_build\:$(VERSION) chown -R $(shell id -u):$(shell id -g) /code/build;

docker-image:
	$(DOCKER) build -f ./dev/docker/Dockerfile -t noria/noriad .

docker-hermes-image:
	$(DOCKER) build -f ./dev/docker/Dockerfile.hermes -t noria/hermes .

docker-test: 
	$(DOCKER) build -f ./dev/docker/Dockerfile.test -t noria_test\:$(VERSION) .
	$(DOCKER) run --rm noria_test\:$(VERSION)

multinodes_start:
	@[ ! $(shell docker images -q noria/noriad:latest 2> /dev/null) ] && make docker-image
	@echo "Initializing 3 nodes"
	@echo "Please wait..."
	@./dev/scripts/multinodes.sh init 3 1>/dev/null 
	@./dev/scripts/multinodes.sh start 3 1>/dev/null
	@echo ""
	@echo "Started 3 nodes"
	@echo ""
	@echo "The first validator exposes a few ports to your local machine for testing purposes"
	@echo "- API: 1317"
	@echo "- RPC: 26657"
	@echo "- gRPC: 9090"

multinodes_stop:
	@echo "Stopping 3 nodes and removing all data"
	@./dev/scripts/multinodes.sh clean 3

network_start:
	@[ ! $(shell docker images -q noria/noriad:latest 2> /dev/null) ] && make docker-image
	@[ ! $(shell docker images -q noria/hermes:latest 2> /dev/null) ] && make docker-hermes-image
	@echo "Initializing 2 nodes and 1 hermes relayer"
	@echo "Please wait..."
	@./dev/scripts/network.sh init 1>/dev/null 
	@./dev/scripts/network.sh start
	@echo ""
	@echo "Network started"
	@echo ""
	@echo "The first validator exposes a few ports to your local machine for testing purposes"
	@echo "- API: 1317"
	@echo "- RPC: 26657"
	@echo "- gRPC: 9090"
	@echo ""
	@echo "val1 is the validator node of the chain oasis-3"
	@echo "val2 is the validator node of the chain oasis-4"
	@echo "Both chains have a transfer channel open over channel-0"
	@echo ""
	@echo "Monitor the relayer with:"
	@echo "    docker logs -f hermes"
	@echo ""
	@echo "IBC Transfer test:"
	@echo "    docker exec val1 noriad --home /tmp/noria tx ibc-transfer transfer transfer channel-0 <val2_address in .network/val2/key.json> 123ucrd --from val1 --fees 1000000ucrd -y"
	@echo ""
	@echo "Wait a few seconds for the relayer to pick up the transaction and relay it to the other chain"
	@echo "Verify the balance on val2 with:"
	@echo "    docker exec val2 noriad --home /tmp/noria q bank balances <val2_address in .network/val2/key.json>"

network_stop:
	@echo "Stopping network and removing all data"
	@./dev/scripts/network.sh clean

run-dev: 
	killall noriad || true
	make clean
	make install
	./dev/scripts/init_local.sh
	noriad start

script-lint:
	@command -v shellcheck >/dev/null 2>&1 || { echo >&2 "shellcheck not found. Installing..."; sudo apt-get install shellcheck; }
	shellcheck scripts/estimate_block_height.sh
	shellcheck scripts/test_state_sync.sh

test:
	(which mockery || go install github.com/vektra/mockery/v2@latest)
	go generate ./...
	go test ./... -v

run-noria-upgrade:
	NORIA_FROM_VERSION=$(FROM) NORIA_TO_VERSION=$(TO) NORIA_UPGRADE_NAME=$(NAME) docker compose -f dev/upgrade/docker-compose.yml up