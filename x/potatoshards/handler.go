package potatoshards

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler() sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		errMsg := fmt.Sprintf("this module does not take messages go away", msg)
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
}
