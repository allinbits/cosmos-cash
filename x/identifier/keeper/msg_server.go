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
	"github.com/allinbits/cosmos-cash/x/identifier/types"
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

func (k msgServer) CreateIdentifier(goCtx context.Context, msg *types.MsgCreateIdentifier) (*types.MsgCreateIdentifierResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	acc, _ := sdk.AccAddressFromBech32(msg.Id)

	identifer, _ := types.NewIdentifier(acc)
	k.Keeper.SetIdentifier(ctx, []byte(msg.Id), identifer)

	id, _ := k.Keeper.GetIdentifier(ctx, []byte(msg.Id))
	fmt.Println(id)

	return &types.MsgCreateIdentifierResponse{}, nil
}
