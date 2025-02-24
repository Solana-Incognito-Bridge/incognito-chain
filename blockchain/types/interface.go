package types

import "github.com/incognitochain/incognito-chain/common"

type BlockPoolInterface interface {
	GetPrevHash() common.Hash
	Hash() *common.Hash
	GetHeight() uint64
	GetShardID() int
	GetRound() int
}

type BlockInterface interface {
	GetVersion() int
	GetHeight() uint64
	Hash() *common.Hash
	// AddValidationField(validateData string) error
	GetProducer() string
	GetProduceTime() int64
	GetProposeTime() int64
	GetProposer() string
	GetValidationField() string
	GetRound() int
	GetRoundKey() string
	GetInstructions() [][]string
	GetConsensusType() string
	GetCurrentEpoch() uint64
	GetPrevHash() common.Hash
	Type() string
	CommitteeFromBlock() common.Hash
	BodyHash() common.Hash
	GetAggregateRootHash() common.Hash
	GetFinalityHeight() uint64
	GetShardID() int
}
