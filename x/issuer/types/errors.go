package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/issuer module sentinel errors
var (
	ErrIssuerFound                      = sdkerrors.Register(ModuleName, 1301, "issuer found")
	ErrUserFound                        = sdkerrors.Register(ModuleName, 1302, "user found")
	ErrInvalidIssuerDenom               = sdkerrors.Register(ModuleName, 1303, "invalid denom for issuer coin")
	ErrDidDocumentDoesNotExist           = sdkerrors.Register(ModuleName, 1304, "did does not exist in the store")
	ErrIncorrectControllerOfDidDocument = sdkerrors.Register(ModuleName, 1305, "sender is not a controller of the did")
	ErrIncorrectUserCredential          = sdkerrors.Register(ModuleName, 1306, "user does not have correct credential")
	ErrTokenAlreadyExists               = sdkerrors.Register(ModuleName, 1307, "token already exists and cannot be recreated")
	ErrMintingTokens                    = sdkerrors.Register(ModuleName, 1308, "error minting tokens for issuer")
	ErrBurningTokens                    = sdkerrors.Register(ModuleName, 1309, "error burning tokens for issuer")
	ErrIssuerNotFound                   = sdkerrors.Register(ModuleName, 1310, "issuer not found in data store")
)
