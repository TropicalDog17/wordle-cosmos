package keeper

import (
	"github.com/TropicalDog17/wordle/x/wordle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetGame set a specific game in the store from its index
func (k Keeper) SetGame(ctx sdk.Context, game types.Game) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKeyPrefix))
	b := k.cdc.MustMarshal(&game)
	store.Set(types.GameKey(
		game.Index,
	), b)
}

// GetGame returns a game from its index
func (k Keeper) GetGame(
	ctx sdk.Context,
	index string,

) (val types.Game, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKeyPrefix))

	b := store.Get(types.GameKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGame removes a game from the store
func (k Keeper) RemoveGame(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKeyPrefix))
	store.Delete(types.GameKey(
		index,
	))
}

// GetAllGame returns all game
func (k Keeper) GetAllGame(ctx sdk.Context) (list []types.Game) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Game
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
