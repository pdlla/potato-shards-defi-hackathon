pragma solidity ^0.5.11;
pragma experimental ABIEncoderV2;


// Simplified staking contract
// Assumes each pub key is unique and does not require proof of secret key knowledge
// Does not support slashing but does have interfaces to demonstrate what slashing might look like
contract PotatoStaking {


    struct Staker {
        uint256 bondedStake;
        uint256 unbondingTime;
        bytes[] pubKey;
    }

    mapping(address => Staker) stakers;
    address[] allStakers;
    uint numShards = 1;

    function BondStake(bytes[] memory pubKey) public payable {
        Staker storage staker = stakers[msg.sender];
        staker.bondedStake = msg.value;
        staker.pubKey = pubKey;
        allStakers.push(msg.sender);
    }

    function UnbondStake() public {
        Staker storage staker = stakers[msg.sender];
        staker.unbondingTime = block.number + 1000;
        // TODO clear from allStakers (need to switch to Set container)
    }

    function Withdraw() public {
        Staker storage staker = stakers[msg.sender];
        assert(staker.unbondingTime != 0 && block.number >= staker.unbondingTime);
        uint256 value = staker.bondedStake;
        delete stakers[msg.sender];
        msg.sender.transfer(value);
    }

    function GetValidators(uint shard) public view returns (bytes[][] memory) {
        // this is vulnerable obviously
        // instead validators can be chosen off chain using VDF
        // or selected on chain using RANDAO or just block hash (if you are willing to trust miners)
        bytes[][] memory r = new bytes[][](allStakers.length / numShards); // TODO not correct if numShards > 1
        for(uint i = 0; i * numShards < allStakers.length; i++) {
            r[i] = stakers[allStakers[i*numShards + shard]].pubKey;
        }
        return r;
    }


    // Example slashing interface
    function SlashDoubleSign(bytes[] memory evidence) public {
        // TODO check evidence of signature on two conflicting blocks within bonding time
    }

    function SlashNoVote(address signer, uint hasNotVotedSince) public payable{
        // TODO store slashing report and wait for challenge period
    }

    function ChallengeSlashNoVote(address signer, bytes[] memory evidence) public {
        // TODO check evidence for signer signature in MerkleRoot of lastCommit in block within voting time
        // if so burn deposit and clear the challenge
    }

    function FinalizeSlashNoVote(address signer) public {
        // TODO if we've passed the challenge period, slash and unbond the offender, reward reporter and return their deposit
    }

}