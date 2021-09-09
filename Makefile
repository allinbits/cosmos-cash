PACKAGES=$(shell go list ./...)
# build paramters 
BUILD_FOLDER = build
APP_VERSION = $(git describe --tags --always)

###############################################################################
###                           Basic Golang Commands                         ###
###############################################################################

all: install

install: go.sum
	go install ./cmd/cosmos-cashd

install-debug: go.sum
	go build -gcflags="all=-N -l" ./cmd/cosmos-cashd

build: clean
	@echo build binary to $(BUILD_FOLDER)
	CGO_ENABLED=0 go build -ldflags "-w -s" -o $(BUILD_FOLDER)/ ./cmd/cosmos-cashd
	@echo computing checksum
	sha256sum $(BUILD_FOLDER)/* --tag > $(BUILD_FOLDER)/checksum.txt
	@echo copy resources
	cp -r README.md LICENSE $(BUILD_FOLDER)
	@echo done

clean:
	@echo clean build folder $(BUILD_FOLDER)
	rm -rf $(BUILD_FOLDER)
	@echo done

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES) -cover -race

lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

###############################################################################
###                           Chain Initialization                          ###
###############################################################################

start-dev: install 
	./scripts/seeds/00_start_chain.sh

seed: 
	./scripts/seeds/01_identifier_seeds.sh
	./scripts/seeds/02_verifiable_credentials_seeds.sh
	./scripts/seeds/03_issuer_seeds.sh

###############################################################################
###                                CI / CD                                  ###
###############################################################################

test-ci:
	go test -coverprofile=coverage.txt -covermode=atomic -mod=readonly $(PACKAGES)

###############################################################################
###                                RELEASE                                  ###
###############################################################################

changelog:
	git-chglog --output CHANGELOG.md

git-release-prepare:
	@echo making release $(APP_VERSION)
ifndef APP_VERSION
	$(error APP_VERSION is not set, please specifiy the version you want to tag)
endif
	git tag $(APP_VERSION)
	git-chglog --output CHANGELOG.md
	git tag $(APP_VERSION) --delete
	git add CHANGELOG.md && git commit -m "update changelog for v$(APP_VERSION)"
	@echo release complete

git-tag:
ifndef APP_VERSION
	$(error APP_VERSION is not set, please specifiy the version you want to tag)
endif
ifneq ($(shell git rev-parse --abbrev-ref HEAD),main)
	$(error you are not on the main branch. aborting)
endif
	git tag -s -a "v$(APP_VERSION)" -m "Changelog: https://github.com/allinbits/cosmos-cash/blob/main/CHANGELOG.md"

_release-patch:
	$(eval APP_VERSION = $(shell git describe --tags | awk -F '("|")' '{ print($$1)}' | awk -F. '{$$NF = $$NF + 1;} 1' | sed 's/ /./g'))
release-patch: _release-patch git-release-prepare

_release-minor:
	$(eval APP_VERSION = $(shell git describe --tags | awk -F '("|")' '{ print($$1)}' | awk -F. '{$$(NF-1) = $$(NF-1) + 1;} 1' | sed 's/ /./g' | awk -F. '{$$(NF) = 0;} 1' | sed 's/ /./g'))
release-minor: _release-minor git-release-prepare

_release-major:
	$(eval APP_VERSION = $(shell git describe --tags | awk -F '("|")' '{ print($$1)}' | awk -F. '{$$(NF-2) = $$(NF-2) + 1;} 1' | sed 's/ /./g' | awk -F. '{$$(NF-1) = 0;} 1' | sed 's/ /./g' | awk -F. '{$$(NF) = 0;} 1' | sed 's/ /./g' ))
release-major: _release-major git-release-prepare

