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
	@go test -mod=readonly $(PACKAGES) -cover 

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

