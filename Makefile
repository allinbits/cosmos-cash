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
	# TODO: find race condition
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
	./scripts/seeds/04_user_seeds.sh

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

git-release:
	@echo making release
	git tag $(GIT_DESCR)
	git-chglog --output CHANGELOG.md
	git tag $(GIT_DESCR) --delete
	git add CHANGELOG.md && git commit -m "update changelog for v$(GIT_DESCR)"
	git tag -s -a "$(GIT_DESCR)" -m "Changelog: https://github.com/allinbits/cosmos-cash/blob/main/CHANGELOG.md"
	@echo release complete


_release-patch:
	$(eval GIT_DESCR = $(shell git describe --tags | awk -F '("|")' '{ print($$1)}' | awk -F. '{$$NF = $$NF + 1;} 1' | sed 's/ /./g'))
release-patch: _release-patch git-release

_release-minor:
	$(eval GIT_DESCR = $(shell git describe --tags | awk -F '("|")' '{ print($$1)}' | awk -F. '{$$(NF-1) = $$(NF-1) + 1;} 1' | sed 's/ /./g' | awk -F. '{$$(NF) = 0;} 1' | sed 's/ /./g'))
release-minor: _release-minor git-release

_release-major:
	$(eval GIT_DESCR = $(shell git describe --tags | awk -F '("|")' '{ print($$1)}' | awk -F. '{$$(NF-2) = $$(NF-2) + 1;} 1' | sed 's/ /./g' | awk -F. '{$$(NF-1) = 0;} 1' | sed 's/ /./g' | awk -F. '{$$(NF) = 0;} 1' | sed 's/ /./g' ))
release-major: _release-major git-release 

