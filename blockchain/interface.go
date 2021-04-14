package blockchain

import (
	"context"
	"github.com/incognitochain/incognito-chain/common/consensus"
	"github.com/incognitochain/incognito-chain/blockchain/types"
	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/incognitokey"
	"github.com/incognitochain/incognito-chain/metadata"
	"github.com/incognitochain/incognito-chain/multiview"
	portalprocessv4 "github.com/incognitochain/incognito-chain/portal/portalv4/portalprocess"
	"github.com/incognitochain/incognito-chain/pubsub"
)

type TxPool interface {
	HaveTransaction(hash *common.Hash) bool
	// RemoveTx remove tx from tx resource
	RemoveTx(txs []metadata.Transaction, isInBlock bool)
	RemoveCandidateList([]string)
	EmptyPool() bool
	MaybeAcceptTransactionForBlockProducing(metadata.Transaction, int64, *ShardBestState) (*metadata.TxDesc, error)
	MaybeAcceptBatchTransactionForBlockProducing(byte, []metadata.Transaction, int64, *ShardBestState) ([]*metadata.TxDesc, error)
}

type FeeEstimator interface {
	RegisterBlock(block *types.ShardBlock) error
}

type ConsensusEngine interface {
	ValidateProducerPosition(blk types.BlockInterface, lastProposerIdx int, committee []incognitokey.CommitteePublicKey, minCommitteeSize int) error
	ValidateProducerSig(block types.BlockInterface, consensusType string) error
	ValidateBlockCommitteSig(block types.BlockInterface, committee []incognitokey.CommitteePublicKey) error
	GetCurrentMiningPublicKey() (string, string)
	// GetCurrentValidators() []*consensus.Validator
	// GetOneValidatorForEachConsensusProcess() map[int]*consensus.Validator
	// GetMiningPublicKeyByConsensus(consensusName string) (string, error)
	GetAllMiningPublicKeys() []string
	ExtractBridgeValidationData(block types.BlockInterface) ([][]byte, []int, error)
	ExtractPortalV4ValidationData(block types.BlockInterface) ([]*portalprocessv4.PortalSig, error)
	GetAllValidatorKeyState() map[string]consensus.MiningState
	GetUserRole() (string, string, int)
	// CommitteeChange(chainName string)
}

type Server interface {
	PushBlockToAll(block types.BlockInterface, previousValidationData string, isBeacon bool) error
}

type Highway interface {
	BroadcastCommittee(uint64, []incognitokey.CommitteePublicKey, map[byte][]incognitokey.CommitteePublicKey, map[byte][]incognitokey.CommitteePublicKey)
}

type Syncker interface {
	GetCrossShardBlocksForShardProducer(toShard byte, list map[byte][]uint64) map[byte][]interface{}
	GetCrossShardBlocksForShardValidator(toShard byte, list map[byte][]uint64) (map[byte][]interface{}, error)
	SyncMissingBeaconBlock(ctx context.Context, peerID string, fromHash common.Hash)
	SyncMissingShardBlock(ctx context.Context, peerID string, sid byte, fromHash common.Hash)
}

type Pubsub interface {
	PublishMessage(message *pubsub.Message)
}

type ChainInterface interface {
	GetBestView() multiview.View
	GetFinalView() multiview.View
}
