# Potato Shards

#### Potato Shards is a alternative shard-compatible staking module for Cosmos

Built for 2019 DeFi hackathon. This is a proof of concept implementation. See Future Work section.

### Concept

This module allows you to launch a chain that where validators are selected from an Ethereum smart contract.
Would be validators stake in on an Ethereum contract. The contract sorts and selects validators for each shard.
This module reads the contract and reports the new set of validators to Tendermint.
Shards can communicate with each other using IBC.

### DeFi

Decentralizing existing financial products on blockchain is an important effort to remove the crucial, exploitable and unreliable element of trust for finance.
However, if decentralization is as much a paradigm shift to technology as evangelists claim there should entirely new finance innovations out there that are incommensurable with anything in existing finance.
In this regard, the existing blockchain economy has already introduce myriad new and unprecedented financial products.

One such product is staking which is still VERY NEW. No PoS blockchain currently has the market cap of leading PoW chains and their behavior is highly dependent on initial stake distribution.
The economics and behavior behind staking are largely unexplored. 
Thus the claim is that staking should be considered a first class DeFi product and one that has huge amounts of unexplored potential.

This project aims to create another mechanism for staking to demonstrate how simple it is to come up with new powerful configurations.

In practice, sharded chains that select validators based on staking conditions in Ethereum have many use cases.
It leverages security guarantees of Ethereum and speed of classical consensus (see hybrid consensus by Shi et. al).
It is extremely simple for the average user to understand and builds on the most widely used smart contract platform.
It allows for small and secure validator sets to run low-value chains securely.
In particular, the inspiration for this idea was taken from free 2 play single player games that run on centralized servers (such as farmville).
In this scenario, each user's game can thus be its own shard and anyone can stake in on the contract to run validator nodes and perhaps earn rewards in game in return.
this is very promising use case for `potatoshards` and could largely replace centralized servers for this genre of games. 

As future work, the author of this project is also excited to use blockchain games that double as closed economies to explore and play with new radical and unprecedented markets.    

### Usage

You must first deploy the `potatostaking.sol` contract on Ethereum.

Update potatoreader.go to contain the contract address, genesis block number, shard number, and ethereum RPC provider URL. 
Next, simply incorporate this module into your project instead of `staking` module and your chain will use validators pulled from the Ethereum smart contract.
Of course make sure you stake in validators on the eth smart contract first :).

The genesis validator set is pulled directly from the smart contract.
Note that potato-shards is currently not compatible `slashing` and `distributor` modules. 

There is no support for dynamic shard switching as a validator AFAIK so a single validator must run one node for each shard.
If the staking contract does not reorganize validators frequently this can be done manually.

### Future Work

This implementation is a proof of concept and has many missing features. It is intended as a template for a full implementation.

It is intended that each shard connect to each other with IBC. This is left to the developer to implement it as they see fit but a standard base token transfer IBC between shards would be a good feature to have for this module.

As noted in comments of potatostaking.sol, it's possible to implement slashing on the Ethereum smart contract.
Perhaps a better way to do it is to relay on a bidirectional peg chain and IBC. Thus all slashing operations happen in the peg chain and users need not concern themselves with Ethereum transactions.

It's reasonable to reward validators on the cosmos chain in the base token so interface for `potatoshards` should be updated to be compatible with `distributor`

potatostaking.sol is a basic implementation and has many missing (important) features.
It should be updated to an interface so that different validator selection criteria can be used based on need.

### Notes

potatostaking.go is generated from potatostaking.abi using the go-ethereum abigen tool  




