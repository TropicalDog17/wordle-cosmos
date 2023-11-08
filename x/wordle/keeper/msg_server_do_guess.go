package keeper

import (
	"context"

	"github.com/TropicalDog17/wordle/x/wordle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DoGuess(goCtx context.Context, msg *types.MsgDoGuess) (*types.MsgDoGuessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDoGuessResponse{}, nil
}
