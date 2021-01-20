module github.com/allinbits/cosmos-cash

go 1.15

//replace github.com/cosmos/cosmos-sdk => /home/ghost/git/cosmos/cosmos-sdk

require (
	github.com/allinbits/cosmos-cash-poc v1.0.0
	github.com/chainapsis/cosmos-sdk-interchain-account v0.1.0
	github.com/cosmos/cosmos-sdk v0.40.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/pelletier/go-toml v1.8.0
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/tendermint/tendermint v0.34.2
	github.com/tendermint/tm-db v0.6.3
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d
	google.golang.org/grpc v1.33.2
	gopkg.in/yaml.v2 v2.4.0

)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
