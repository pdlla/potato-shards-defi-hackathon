package potatoshards

import (
	"log"
	"math/big"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


const providerURL = "wss://rinkeby.infura.io/ws"
const contractAddress = "0x0D8A07e01Fd9b3DA4ce78109DBDFd385bE59bAE2"
const ethGenesisNumber = "12345"
const shardNumber = 0

// TODO block number as argument
func getValidators() []abci.ValidatorUpdate {
	/**
	 * Connecting to provider
	 */
	client, err := ethclient.Dial(providerURL)

	if err != nil {
		log.Fatal(err)
	}

	contract, err := NewPotato(common.HexToAddress(contractAddress), client)

	if err != nil {
		log.Fatal(err)
	}

	session := &PotatoSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	// TODO call this with a target block number once its added
	// TODO put in a real shard number
	addrs, err := session.GetValidators(big.NewInt(shardNumber))

	// TODO blocking retry logic
	if err != nil {
		log.Fatal(err)
	}

	// convert addresses from contract
	r := make([]abci.ValidatorUpdate,len(addrs))
	for i:=0; i < len(addrs); i++ {
		key := make([]byte,0)
		for j:=0; j < len(addrs[i]); j++ {
			key = append(key, addrs[i][j]...);
		}
		r[i] = abci.Ed25519ValidatorUpdate(key, 1)
	}

	return r
}
// Called every block, update validator set
func EndBlocker(ctx sdk.Context) []abci.ValidatorUpdate {
	// TODO compute expected block number based on timestamp (and compare against eth timestamps)
	return getValidators()
}
