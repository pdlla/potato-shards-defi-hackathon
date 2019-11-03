package potatoshards
import(
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all staking state that must be provided at genesis
type GenesisState struct {
}

// get raw genesis raw message for testing
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func ExportGenesis(ctx sdk.Context) GenesisState {
	return GenesisState{}
}