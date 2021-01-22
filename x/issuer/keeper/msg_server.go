package keeper

import (
	"context"
	"fmt"
	//	"time"
	//
	//	metrics "github.com/armon/go-metrics"
	//	tmstrings "github.com/tendermint/tendermint/libs/strings"
	//
	//	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	//	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/allinbits/cosmos-cash/x/issuer/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the identity MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateIssuer(goCtx context.Context, msg *types.MsgCreateIssuer) (*types.MsgCreateIssuerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	identifer, _ := types.NewIssuer(msg.Name, msg.Token, msg.Fee, msg.State, msg.Address)
	k.Keeper.SetIssuer(ctx, []byte(msg.Address), identifer)

	id, _ := k.Keeper.GetIssuer(ctx, []byte(msg.Address))
	fmt.Println(id)

	return &types.MsgCreateIssuerResponse{}, nil
}
