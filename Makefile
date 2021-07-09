PACKAGES=$(shell go list ./...)

###############################################################################
###                           Basic Golang Commands                         ###
###############################################################################

all: install

install: go.sum
	go install ./cmd/cosmos-cashd

install-debug: go.sum
	go build -gcflags="all=-N -l" ./cmd/cosmos-cashd

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
	@echo making release
	ifndef VERSION
		$(error VERSION is not set, please specifiy the version you want to tag)
	endif
	git tag $(VERSION)
	git-chglog --output CHANGELOG.md
	git tag $(VERSION) --delete
	git add CHANGELOG.md && git commit -m "update changelog for v$(VERSION)"
	@echo release complete

git-tag:
	ifndef VERSION
		$(error VERSION is not set, please specifiy the version you want to tag)
	endif
	ifneq ($(shell git rev-parse --abbrev-ref HEAD),main)
		$(error you are not on the main branch. aborting)
	endif
	git tag -s -a "$(VERSION)" -m "Changelog: https://github.com/allinbits/cosmos-cash/blob/main/CHANGELOG.md"

_release-patch:
	$(eval VERSION = $(shell git describe --tags | awk -F '("|")' '{ print($$1)}' | awk -F. '{$$NF = $$NF + 1;} 1' | sed 's/ /./g'))
release-patch: _release-patch git-release-prepare

_release-minor:
	$(eval VERSION = $(shell git describe --tags | awk -F '("|")' '{ print($$1)}' | awk -F. '{$$(NF-1) = $$(NF-1) + 1;} 1' | sed 's/ /./g' | awk -F. '{$$(NF) = 0;} 1' | sed 's/ /./g'))
release-minor: _release-minor git-release-prepare

_release-major:
	$(eval VERSION = $(shell git describe --tags | awk -F '("|")' '{ print($$1)}' | awk -F. '{$$(NF-2) = $$(NF-2) + 1;} 1' | sed 's/ /./g' | awk -F. '{$$(NF-1) = 0;} 1' | sed 's/ /./g' | awk -F. '{$$(NF) = 0;} 1' | sed 's/ /./g' ))
release-major: _release-major git-release-prepare

