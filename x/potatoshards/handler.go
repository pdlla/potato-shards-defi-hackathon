package potatoshards

import (
	"fmt"
	"log"
	"math/big"

	// Vendor imports
	"github.com/tendermint/tendermint/libs/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const ModuleName = "potatoshards"

func NewHandler() sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		errMsg := fmt.Sprintf("this module does not take messages go away", msg)
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
}
