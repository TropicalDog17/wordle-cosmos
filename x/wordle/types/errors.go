package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/wordle module sentinel errors
var (
	ErrInvalidPlayer  = sdkerrors.Register(ModuleName, 1100, "Invalid player address")
	ErrInvalidCreator = sdkerrors.Register(ModuleName, 1101, "Invalid creator address")
	ErrInvalidSecret  = sdkerrors.Register(ModuleName, 1102, "Invalid secret word")
)
