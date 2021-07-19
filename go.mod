module github.com/allinbits/cosmos-cash

go 1.15

//replace github.com/cosmos/cosmos-sdk => /home/ghost/git/cosmos/cosmos-sdk

require (
	github.com/cosmos/cosmos-sdk v0.42.1
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kr/text v0.2.0 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/rs/zerolog v1.20.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/liquidity v1.2.1-rc0.0.20210311071044-1572795241e1
	github.com/tendermint/tendermint v0.34.8
	github.com/tendermint/tm-db v0.6.4
	github.com/wealdtech/go-merkletree v1.0.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
	google.golang.org/genproto v0.0.0-20210223151946-22b48be4551b
	google.golang.org/grpc v1.36.0

)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
