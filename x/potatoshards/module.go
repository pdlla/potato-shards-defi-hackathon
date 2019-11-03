package potatoshards

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	abci "github.com/tendermint/tendermint/abci/types"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking/client/rest"
)


const (
	ModuleName = "potatoshards"
	StoreKey = ModuleName
	RouterKey = ModuleName
	QuerierRoute = ModuleName
)




var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the potatoshards module.
type AppModuleBasic struct{}

var _ module.AppModuleBasic = AppModuleBasic{}

// Name returns the potatoshards module's name.
func (AppModuleBasic) Name() string {
	return ModuleName
}

// RegisterCodec registers the potatoshards module's types for the given codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the potatoshards
// module.
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the potatoshards module.
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	// nothing to do :)
	return nil
}

// RegisterRESTRoutes registers the REST routes for the potatoshards module.
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

// GetTxCmd returns the root tx command for the potatoshards module.
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(StoreKey, cdc)
}

// GetQueryCmd returns no root query command for the potatoshards module.
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(StoreKey, cdc)
}

//_____________________________________
// extra helpers

// CreateValidatorMsgHelpers - used for gen-tx
func (AppModuleBasic) CreateValidatorMsgHelpers(ipDefault string) (
	fs *flag.FlagSet, nodeIDFlag, pubkeyFlag, amountFlag, defaultsDesc string) {
	return cli.CreateValidatorMsgHelpers(ipDefault)
}

// PrepareFlagsForTxCreateValidator - used for gen-tx
func (AppModuleBasic) PrepareFlagsForTxCreateValidator(config *cfg.Config, nodeID,
	chainID string, valPubKey crypto.PubKey) {
	cli.PrepareFlagsForTxCreateValidator(config, nodeID, chainID, valPubKey)
}

// BuildCreateValidatorMsg - used for gen-tx
func (AppModuleBasic) BuildCreateValidatorMsg(cliCtx context.CLIContext,
	txBldr authtypes.TxBuilder) (authtypes.TxBuilder, sdk.Msg, error) {
	return cli.BuildCreateValidatorMsg(cliCtx, txBldr)
}

//____________________________________________________________________________

// AppModule implements an application module for the potatoshards module.
type AppModule struct {
	AppModuleBasic
}

// NewAppModule creates a new AppModule object
func NewAppModule() AppModule {

	return AppModule{
		AppModuleBasic:      AppModuleBasic{},
	}
}

// Name returns the potatoshards module's name.
func (AppModule) Name() string {
	return ModuleName
}

// RegisterInvariants registers the potatoshards module invariants.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	// nothing to do
}

// Route returns the message routing key for the potatoshards module.
func (AppModule) Route() string {
	return RouterKey
}

// NewHandler returns an sdk.Handler for the potatoshards module.
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler()
}

// QuerierRoute returns the potatoshards module's querier route name.
func (AppModule) QuerierRoute() string {
	return QuerierRoute
}

// NewQuerierHandler returns the potatoshards module sdk.Querier.
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier()
}

// InitGenesis performs genesis initialization for the potatoshards module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	// TODO pass in genesis block number
	return getValidators()
}

// ExportGenesis returns the exported genesis state as raw bytes for the potatoshards
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx)
	return ModuleCdc.MustMarshalJSON(gs)
}

// BeginBlock returns the begin blocker for the potatoshards module.
func (AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the potatoshards module. It returns no validator
// updates.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return EndBlocker(ctx)
}
