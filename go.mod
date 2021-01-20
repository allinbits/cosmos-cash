module github.com/allinbits/cosmos-cash

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.40.0-rc3
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.15.2
    github.com/pelletier/go-toml v1.8.0
	github.com/regen-network/cosmos-proto v0.3.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/tendermint/tendermint v0.34.0-rc6
	github.com/tendermint/tm-db v0.6.2
	google.golang.org/genproto v0.0.0-20201014134559-03b6142f0dc9
	google.golang.org/grpc v1.33.0

)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
