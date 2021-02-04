PACKAGES=$(shell go list ./...)

###############################################################################
###                           Basic Golang Commands                         ###
###############################################################################

all: install

install: go.sum
	go install -mod=readonly ./cmd/cosmos-cashd

install-debug: go.sum
	go build -mod=readonly -gcflags="all=-N -l" ./cmd/cosmos-cashd

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)

lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

###############################################################################
###                           Chain Initialization                          ###
###############################################################################

init-dev: init-chain init-validator

start-dev: 
	go run cmd/cosmos-cashd/main.go start --pruning=nothing --grpc-web.enable=false

init-chain:
	go run cmd/cosmos-cashd/main.go init --chain-id=cash cash 
	echo "y" | go run cmd/cosmos-cashd/main.go keys add validator

init-validator:
	go run cmd/cosmos-cashd/main.go add-genesis-account $(shell go run cmd/cosmos-cashd/main.go keys show validator -a) 1000000000stake
	go run cmd/cosmos-cashd/main.go gentx validator 700000000stake --chain-id cash
	go run cmd/cosmos-cashd/main.go collect-gentxs 

