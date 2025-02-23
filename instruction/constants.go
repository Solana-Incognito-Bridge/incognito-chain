package instruction

const (
	SWAP_SHARD_ACTION              = "swapshard"
	SWAP_ACTION                    = "swap"
	RANDOM_ACTION                  = "random"
	STAKE_ACTION                   = "stake"
	ASSIGN_ACTION                  = "assign"
	ASSIGN_SYNC_ACTION             = "assignsync"
	STOP_AUTO_STAKE_ACTION         = "stopautostake"
	SET_ACTION                     = "set"
	RETURN_ACTION                  = "return"
	UNSTAKE_ACTION                 = "unstake"
	SHARD_INST                     = "shard"
	BEACON_INST                    = "beacon"
	SPLITTER                       = ","
	TRUE                           = "true"
	FALSE                          = "false"
	FINISH_SYNC_ACTION             = "finishsync"
	ACCEPT_BLOCK_REWARD_V3_ACTION  = "acceptblockrewardv3"
	SHARD_RECEIVE_REWARD_V3_ACTION = "shardreceiverewardv3"

	SHARD_RECEIVE_REWARD_V1_ACTION = 43
	ACCEPT_BLOCK_REWARD_V1_ACTION  = 37
	SHARD_REWARD_INST              = "shardRewardInst"
)

const (
	BEACON_CHAIN_ID = -1
)

//Swap Instruction Sub Type
const (
	SWAP_BY_END_EPOCH = iota
	SWAP_BY_SLASHING
	SWAP_BY_INCREASE_COMMITTEES_SIZE
	SWAP_BY_DECREASE_COMMITTEES_SIZE
)
