package potatoshards

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)
// creates a querier for staking REST endpoints
func NewQuerier() sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		// TODO return validator query results
		return nil, sdk.ErrUnknownRequest("unknown staking query endpoint")
	}
}
