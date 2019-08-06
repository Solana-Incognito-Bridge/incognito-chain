package blockchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/incognitochain/incognito-chain/metrics"
	"github.com/incognitochain/incognito-chain/pubsub"
	"github.com/pkg/errors"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/common/base58"
	"github.com/incognitochain/incognito-chain/incognitokey"
)

/*
	// This function should receives block in consensus round
	// It verify validity of this function before sign it
	// This should be verify in the first round of consensus

	Step:
	1. Verify Pre proccessing data
	2. Retrieve beststate for new block, store in local variable
	3. Update: process local beststate with new block
	4. Verify Post processing: updated local beststate and newblock

	Return:
	- No error: valid and can be sign
	- Error: invalid new block
*/
func (blockchain *BlockChain) VerifyPreSignBeaconBlock(beaconBlock *BeaconBlock, isCommittee bool) error {
	blockchain.chainLock.Lock()
	defer blockchain.chainLock.Unlock()
	//========Verify block only
	Logger.log.Infof("BEACON | Verify block for signing process %d, with hash %+v", beaconBlock.Header.Height, *beaconBlock.Hash())
	if err := blockchain.verifyPreProcessingBeaconBlock(beaconBlock, isCommittee); err != nil {
		return err
	}
	//========Verify block with previous best state
	// Get Beststate of previous block == previous best state
	// Clone best state value into new variable
	beaconBestState := NewBeaconBestState()
	if err := beaconBestState.cloneBeaconBestState(blockchain.BestState.Beacon); err != nil {
		return err
	}
	// Verify block with previous best state
	// not verify agg signature in this function
	if err := beaconBestState.VerifyBestStateWithBeaconBlock(beaconBlock, false); err != nil {
		return err
	}
	//========Update best state with new block
	snapShotBeaconCommittee := beaconBestState.BeaconCommittee
	if err := beaconBestState.updateBeaconBestState(beaconBlock); err != nil {
		return err
	}
	//========Post verififcation: verify new beaconstate with corresponding block
	if err := beaconBestState.VerifyPostProcessingBeaconBlock(beaconBlock, snapShotBeaconCommittee); err != nil {
		return err
	}
	Logger.log.Infof("BEACON | Block %d, with hash %+v is VALID to be 🖊 signed", beaconBlock.Header.Height, *beaconBlock.Hash())
	return nil
}

func (blockchain *BlockChain) InsertBeaconBlock(block *BeaconBlock, isValidated bool) error {
	blockchain.chainLock.Lock()
	defer blockchain.chainLock.Unlock()

	blockHash := block.Header.Hash()
	Logger.log.Infof("BEACON | Check block existence for insert process %d, with hash %+v", block.Header.Height, blockHash)
	isExist, _ := blockchain.config.DataBase.HasBeaconBlock(block.Header.Hash())
	if isExist {
		return NewBlockChainError(DuplicateShardBlockError, errors.New("This block has been stored already"))
	}
	Logger.log.Infof("BEACON | Begin Insert new block %d, with hash %+v \n", block.Header.Height, blockHash)
	if !isValidated {
		Logger.log.Infof("BEACON | Verify Pre Processing Beacon Block %+v \n", blockHash)
		if err := blockchain.verifyPreProcessingBeaconBlock(block, false); err != nil {
			return err
		}
	} else {
		Logger.log.Infof("BEACON | SKIP Verify Pre Processing Block %+v \n", blockHash)
	}
	//========Verify block with previous best state
	// check with current final best state
	// block can only be insert if it match the current best state
	bestBlockHash := &blockchain.BestState.Beacon.BestBlockHash
	if !bestBlockHash.IsEqual(&block.Header.PreviousBlockHash) {
		return NewBlockChainError(BeaconError, errors.New("beacon Block does not match with any Beacon State in cache or in Database"))
	}
	if !isValidated {
		Logger.log.Infof("BEACON | Verify BestState with Beacon Block %+v \n", blockHash)
		// Verify block with previous best state
		if err := blockchain.BestState.Beacon.VerifyBestStateWithBeaconBlock(block, true); err != nil {
			return err
		}
	} else {
		Logger.log.Infof("BEACON | SKIP Verify BestState with Block %+v \n", blockHash)
	}
	// Backup beststate
	if blockchain.config.UserKeySet != nil {
		userRole, _ := blockchain.BestState.Beacon.GetPubkeyRole(blockchain.config.UserKeySet.GetPublicKeyInBase58CheckEncode(), 0)
		if userRole == common.PROPOSER_ROLE || userRole == common.VALIDATOR_ROLE {
			blockchain.config.DataBase.CleanBackup(false, 0)
			err := blockchain.BackupCurrentBeaconState(block)
			if err != nil {
				return err
			}
		}
	}
	Logger.log.Infof("BEACON | Update BestState with Beacon Block %+v \n", blockHash)
	//========Update best state with new block
	snapShotBeaconCommittee := blockchain.BestState.Beacon.BeaconCommittee
	if err := blockchain.BestState.Beacon.updateBeaconBestState(block); err != nil {
		return err
	}
	if !isValidated {
		Logger.log.Infof("BEACON | Verify Post Processing Beacon Block %+v \n", blockHash)
		//========Post verififcation: verify new beaconstate with corresponding block
		if err := blockchain.BestState.Beacon.VerifyPostProcessingBeaconBlock(block, snapShotBeaconCommittee); err != nil {
			return err
		}
	} else {
		Logger.log.Infof("BEACON | SKIP Verify Post Processing Block %+v \n", blockHash)
	}
	for shardID, shardStates := range block.Body.ShardState {
		for _, shardState := range shardStates {
			blockchain.config.DataBase.StoreAcceptedShardToBeacon(shardID, block.Header.Height, shardState.Hash)
		}
	}
	Logger.log.Infof("BEACON | Store Committee in Height %+v \n", block.Header.Height)
	if err := blockchain.config.DataBase.StoreCommitteeByHeight(block.Header.Height, blockchain.BestState.Beacon.GetShardCommittee()); err != nil {
		return NewBlockChainError(DatabaseError, err)
	}
	if err := blockchain.config.DataBase.StoreBeaconCommitteeByHeight(block.Header.Height, blockchain.BestState.Beacon.BeaconCommittee); err != nil {
		return NewBlockChainError(DatabaseError, err)
	}
	// }
	// shardCommitteeByte, err := blockchain.config.DataBase.FetchCommitteeByEpoch(block.Header.Epoch)
	// if err != nil {
	// 	fmt.Println("No committee for this epoch")
	// }
	// shardCommittee := make(map[byte][]string)
	// if err := json.Unmarshal(shardCommitteeByte, &shardCommittee); err != nil {
	// 	fmt.Println("Fail to unmarshal shard committee")
	// }
	// fmt.Println("Beacon Process/Shard Committee in Epoch ", block.Header.Epoch, shardCommittee)
	//=========Store cross shard state ==================================
	if block.Body.ShardState != nil {
		GetBeaconBestState().lock.Lock()
		lastCrossShardState := GetBeaconBestState().LastCrossShardState
		for fromShard, shardBlocks := range block.Body.ShardState {
			for _, shardBlock := range shardBlocks {
				for _, toShard := range shardBlock.CrossShard {
					if fromShard == toShard {
						continue
					}
					if lastCrossShardState[fromShard] == nil {
						lastCrossShardState[fromShard] = make(map[byte]uint64)
					}
					lastHeight := lastCrossShardState[fromShard][toShard] // get last cross shard height from shardID  to crossShardShardID
					waitHeight := shardBlock.Height
					fmt.Println("StoreCrossShardNextHeight", fromShard, toShard, lastHeight, waitHeight)
					blockchain.config.DataBase.StoreCrossShardNextHeight(fromShard, toShard, lastHeight, waitHeight)
					//beacon process shard_to_beacon in order so cross shard next height also will be saved in order
					//dont care overwrite this value
					blockchain.config.DataBase.StoreCrossShardNextHeight(fromShard, toShard, waitHeight, 0)
					if lastCrossShardState[fromShard] == nil {
						lastCrossShardState[fromShard] = make(map[byte]uint64)
					}
					lastCrossShardState[fromShard][toShard] = waitHeight //update lastHeight to waitHeight
				}
			}
			blockchain.config.CrossShardPool[fromShard].UpdatePool()
		}
		GetBeaconBestState().lock.Unlock()
	}
	// ************ Store block at last
	//========Store new Beaconblock and new Beacon bestState in cache
	Logger.log.Info("Store Beacon BestState")
	if err := blockchain.StoreBeaconBestState(); err != nil {
		return NewBlockChainError(DatabaseError, err)
	}
	Logger.log.Info("Store Beacon Block ", block.Header.Height, blockHash)
	if err := blockchain.config.DataBase.StoreBeaconBlock(block, blockHash); err != nil {
		return NewBlockChainError(DatabaseError, err)
	}
	if err := blockchain.config.DataBase.StoreBeaconBlockIndex(blockHash, block.Header.Height); err != nil {
		return NewBlockChainError(DatabaseError, err)
	}
	//=========Remove beacon block in pool
	go blockchain.config.BeaconPool.SetBeaconState(blockchain.BestState.Beacon.BeaconHeight)
	go blockchain.config.BeaconPool.RemoveBlock(blockchain.BestState.Beacon.BeaconHeight)
	//=========Remove shard to beacon block in pool
	//Logger.log.Info("Remove block from pool block with hash  ", *block.Hash(), block.Header.Height, blockchain.BestState.Beacon.BestShardHeight)
	go blockchain.config.ShardToBeaconPool.SetShardState(blockchain.BestState.Beacon.GetBestShardHeight())
	err := blockchain.updateDatabaseFromBeaconBlock(block)
	if err != nil {
		Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
		return NewBlockChainError(UnExpectedError, err)
	}
	err = blockchain.processBridgeInstructions(block)
	if err != nil {
		Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
		return NewBlockChainError(UnExpectedError, err)
	}
	go metrics.AnalyzeTimeSeriesMetricData(map[string]interface{}{
		metrics.Measurement:      metrics.NumOfBlockInsertToChain,
		metrics.MeasurementValue: float64(1),
		metrics.Tag:              metrics.ShardIDTag,
		metrics.TagValue:         metrics.Beacon,
	})
	Logger.log.Infof("Finish Insert new block %+v, with hash %+v \n", block.Header.Height, *block.Hash())
	if block.Header.Height%50 == 0 {
		BLogger.log.Debugf("Inserted beacon height: %d", block.Header.Height)
	}
	go blockchain.config.PubSubManager.PublishMessage(pubsub.NewMessage(pubsub.NewBeaconBlockTopic, block))
	go blockchain.config.PubSubManager.PublishMessage(pubsub.NewMessage(pubsub.BeaconBeststateTopic, blockchain.BestState.Beacon))
	return nil
}

/*
	VerifyPreProcessingBeaconBlock
	This function DOES NOT verify new block with best state
	DO NOT USE THIS with GENESIS BLOCK
	- Producer validity
	- version
	- parent hash
	- Height = parent hash + 1
	- Epoch = blockHeight % Epoch ? Parent Epoch + 1
	- Timestamp can not excess some limit
	- Instruction hash
	- ShardStateHash
	- ShardState is sorted?
	FOR CURRENT COMMITTEES ONLY
		- Is shardState existed in pool
*/
func (blockchain *BlockChain) verifyPreProcessingBeaconBlock(beaconBlock *BeaconBlock, isCommittee bool) error {
	//verify producer sig
	hash := beaconBlock.Header.Hash()
	producerPublicKey := base58.Base58Check{}.Encode(beaconBlock.Header.ProducerAddress.Pk, common.ZeroByte)
	err := incognitokey.ValidateDataB58(producerPublicKey, beaconBlock.ProducerSig, hash.GetBytes())
	if err != nil {
		return NewBlockChainError(BeaconBlockSignatureError, fmt.Errorf("Producer Public Key %+v, Producer Signature %+v, Hash %+v", producerPublicKey, beaconBlock.ProducerSig, hash))
	}
	//verify producer via index
	producerPosition := (blockchain.BestState.Beacon.BeaconProposerIndex + beaconBlock.Header.Round) % len(blockchain.BestState.Beacon.BeaconCommittee)
	tempProducer := blockchain.BestState.Beacon.BeaconCommittee[producerPosition]
	if strings.Compare(tempProducer, producerPublicKey) != 0 {
		return NewBlockChainError(BeaconBlockProducerError, fmt.Errorf("Expect Producer Public Key to be equal but get %+v From Index, %+v From Header", tempProducer, producerPublicKey))
	}
	//verify version
	if beaconBlock.Header.Version != BEACON_BLOCK_VERSION {
		return NewBlockChainError(WrongVersionError, fmt.Errorf("Expect block version to be equal to %+v but get %+v", BEACON_BLOCK_VERSION, beaconBlock.Header.Version))
	}
	// Verify parent hash exist or not
	previousBlockHash := beaconBlock.Header.PreviousBlockHash
	parentBlockBytes, err := blockchain.config.DataBase.FetchBeaconBlock(previousBlockHash)
	if err != nil {
		return NewBlockChainError(FetchBeaconBlockError, err)
	}
	previousBeaconBlock := NewBeaconBlock()
	err = json.Unmarshal(parentBlockBytes, previousBeaconBlock)
	if err != nil {
		return NewBlockChainError(UnmashallJsonBeaconBlockError, fmt.Errorf("Failed to unmarshall parent block of block height %+v", beaconBlock.Header.Height))
	}
	// Verify block height with parent block
	if previousBeaconBlock.Header.Height+1 != beaconBlock.Header.Height {
		return NewBlockChainError(WrongBlockHeightError, fmt.Errorf("Expect receive beacon block height %+v but get %+v", previousBeaconBlock.Header.Height+1, beaconBlock.Header.Height))
	}
	// Verify epoch with parent block
	if (beaconBlock.Header.Height != 1) && (beaconBlock.Header.Height%common.EPOCH == 1) && (previousBeaconBlock.Header.Epoch != beaconBlock.Header.Epoch-1) {
		return NewBlockChainError(WrongEpochError, fmt.Errorf("Expect receive beacon block epoch %+v greater than previous block epoch %+v, 1 value", beaconBlock.Header.Epoch, previousBeaconBlock.Header.Epoch))
	}
	// Verify timestamp with parent block
	if beaconBlock.Header.Timestamp <= previousBeaconBlock.Header.Timestamp {
		return NewBlockChainError(WrongTimestampError, fmt.Errorf("Expect receive beacon block with timestamp %+v greater than previous block timestamp %+v", beaconBlock.Header.Timestamp, previousBeaconBlock.Header.Timestamp))
	}
	if !VerifyHashFromShardState(beaconBlock.Body.ShardState, beaconBlock.Header.ShardStateHash) {
		return NewBlockChainError(ShardStateHashError, fmt.Errorf("Expect shard state hash to be %+v", beaconBlock.Header.ShardStateHash))
	}
	tempInstructionArr := []string{}
	for _, strs := range beaconBlock.Body.Instructions {
		tempInstructionArr = append(tempInstructionArr, strs...)
	}
	if !VerifyHashFromStringArray(tempInstructionArr, beaconBlock.Header.InstructionHash) {
		return NewBlockChainError(InstructionHashError, fmt.Errorf("Expect instruction hash to be %+v", beaconBlock.Header.InstructionHash))
	}
	// Shard state must in right format
	// state[i].Height must less than state[i+1].Height and state[i+1].Height - state[i].Height = 1
	for _, shardStates := range beaconBlock.Body.ShardState {
		for i := 0; i < len(shardStates)-2; i++ {
			if shardStates[i+1].Height-shardStates[i].Height != 1 {
				return NewBlockChainError(ShardStateError, fmt.Errorf("Expect Shard State Height to be in the right format, %+v, %+v", shardStates[i+1].Height, shardStates[i].Height))
			}
		}
	}
	// Check if InstructionMerkleRoot is the root of merkle tree containing all instructions in this block
	flattenInsts, err := FlattenAndConvertStringInst(beaconBlock.Body.Instructions)
	if err != nil {
		return NewBlockChainError(FlattenAndConvertStringInstError, err)
	}

	root := GetKeccak256MerkleRoot(flattenInsts)
	if !bytes.Equal(root, beaconBlock.Header.InstructionMerkleRoot[:]) {
		return NewBlockChainError(FlattenAndConvertStringInstError, fmt.Errorf("Expect Instruction Merkle Root in Beacon Block Header to be %+v but get %+v", string(beaconBlock.Header.InstructionMerkleRoot[:]), string(root)))
	}

	// if pool does not have one of needed block, fail to verify
	if isCommittee {
		rewardByEpochInstruction := [][]string{}
		if beaconBlock.Header.Height%common.EPOCH == 1 {
			rewardByEpochInstruction, err = blockchain.BuildRewardInstructionByEpoch(beaconBlock.Header.Epoch - 1)
			if err != nil {
				return NewBlockChainError(BuildRewardInstructionError, err)
			}
		}
		// @UNCOMMENT TO TEST
		beaconBestState := NewBeaconBestState()
		if err := beaconBestState.cloneBeaconBestState(blockchain.BestState.Beacon); err != nil {
			return err
		}
		tempShardStates := make(map[byte][]ShardState)
		validStakers := [][]string{}
		validSwappers := make(map[byte][][]string)
		bridgeInstructions := [][]string{}
		acceptedBlockRewardInstructions := [][]string{}
		//tempMarshal, _ := json.Marshal(*blockchain.BestState.Beacon)
		//err = json.Unmarshal(tempMarshal, &beaconBestState)
		//if err != nil {
		//	return NewBlockChainError(UnExpectedError, errors.New("Fail to Unmarshal beacon beststate"))
		//}
		//beaconBestState.CandidateShardWaitingForCurrentRandom = blockchain.BestState.Beacon.CandidateShardWaitingForCurrentRandom
		//beaconBestState.CandidateShardWaitingForNextRandom = blockchain.BestState.Beacon.CandidateShardWaitingForNextRandom
		//beaconBestState.CandidateBeaconWaitingForCurrentRandom = blockchain.BestState.Beacon.CandidateBeaconWaitingForCurrentRandom
		//beaconBestState.CandidateBeaconWaitingForNextRandom = blockchain.BestState.Beacon.CandidateBeaconWaitingForNextRandom
		//if reflect.DeepEqual(beaconBestState, BeaconBestState{}) {
		//	panic(NewBlockChainError(BeaconError, errors.New("problem with beststate in producing new block")))
		//}
		allShardBlocks := blockchain.config.ShardToBeaconPool.GetValidBlock(nil)
		var keys []int
		for k := range allShardBlocks {
			keys = append(keys, int(k))
		}
		sort.Ints(keys)
		for _, value := range keys {
			shardID := byte(value)
			shardBlocks := allShardBlocks[shardID]
			if len(shardBlocks) >= len(beaconBlock.Body.ShardState[shardID]) {
				shardBlocks = shardBlocks[:len(beaconBlock.Body.ShardState[shardID])]
				shardStates := beaconBlock.Body.ShardState[shardID]
				for index, shardState := range shardStates {
					if shardBlocks[index].Header.Height != shardState.Height {
						return NewBlockChainError(ShardStateError, errors.New("shardstate fail to verify with ShardToBeacon Block in pool"))
					}
					blockHash := shardBlocks[index].Header.Hash()
					if !blockHash.IsEqual(&shardState.Hash) {
						return NewBlockChainError(ShardStateError, errors.New("shardstate fail to verify with ShardToBeacon Block in pool"))
					}
					if !reflect.DeepEqual(shardBlocks[index].Header.CrossShardBitMap, shardState.CrossShard) {
						return NewBlockChainError(ShardStateError, errors.New("shardstate fail to verify with ShardToBeacon Block in pool"))
					}
				}
				// Only accept block in one epoch
				for index, shardBlock := range shardBlocks {
					currentCommittee := blockchain.BestState.Beacon.GetAShardCommittee(shardID)
					currentPendingValidator := blockchain.BestState.Beacon.GetAShardPendingValidator(shardID)
					hash := shardBlock.Header.Hash()
					err := ValidateAggSignature(shardBlock.ValidatorsIndex, currentCommittee, shardBlock.AggregatedSig, shardBlock.R, &hash)
					if index == 0 && err != nil {
						currentCommittee, _, _, _, err = SwapValidator(currentPendingValidator, currentCommittee, blockchain.BestState.Beacon.MaxShardCommitteeSize, common.OFFSET)
						if err != nil {
							return NewBlockChainError(ShardStateError, errors.New("shardstate fail to verify with ShardToBeacon Block in pool"))
						}
						err = ValidateAggSignature(shardBlock.ValidatorsIndex, currentCommittee, shardBlock.AggregatedSig, shardBlock.R, &hash)
						if err != nil {
							return NewBlockChainError(ShardStateError, errors.New("shardstate fail to verify with ShardToBeacon Block in pool"))
						}
					}
					if index != 0 && err != nil {
						return NewBlockChainError(ShardStateError, errors.New("shardstate fail to verify with ShardToBeacon Block in pool"))
					}
				}
				for _, shardBlock := range shardBlocks {
					tempShardState, validStaker, validSwapper, bridgeInstruction, acceptedBlockRewardInstruction := blockchain.GetShardStateFromBlock(beaconBlock.Header.Height, shardBlock, shardID)
					tempShardStates[shardID] = append(tempShardStates[shardID], tempShardState[shardID])
					validStakers = append(validStakers, validStaker...)
					validSwappers[shardID] = append(validSwappers[shardID], validSwapper[shardID]...)
					bridgeInstructions = append(bridgeInstructions, bridgeInstruction...)
					acceptedBlockRewardInstructions = append(acceptedBlockRewardInstructions, acceptedBlockRewardInstruction)
				}
			} else {
				return NewBlockChainError(ShardStateError, errors.New("shardstate fail to verify with ShardToBeacon Block in pool"))
			}
		}
		beaconBestState.InitRandomClient(blockchain.config.RandomClient)
		tempInstruction := beaconBestState.GenerateInstruction(beaconBlock.Header.Height, validStakers, validSwappers, beaconBestState.CandidateShardWaitingForCurrentRandom, bridgeInstructions, acceptedBlockRewardInstructions)
		if len(rewardByEpochInstruction) != 0 {
			tempInstruction = append(tempInstruction, rewardByEpochInstruction...)
		}
		fmt.Println("BeaconProcess/tempInstruction: ", tempInstruction)
		tempInstructionArr := []string{}
		for _, strs := range tempInstruction {
			tempInstructionArr = append(tempInstructionArr, strs...)
		}
		tempInstructionHash, err := GenerateHashFromStringArray(tempInstructionArr)
		if err != nil {
			return NewBlockChainError(HashError, errors.New("Fail to generate hash for instruction"))
		}
		fmt.Println("BeaconProcess/tempInstructionHash: ", tempInstructionHash)
		fmt.Println("BeaconProcess/block.Header.InstructionHash: ", beaconBlock.Header.InstructionHash)
		if !tempInstructionHash.IsEqual(&beaconBlock.Header.InstructionHash) {
			return NewBlockChainError(InstructionHashError, errors.New("instruction hash is not correct"))
		}
	}

	return nil
}

/*
	This function will verify the validation of a block with some best state in cache or current best state
	Get beacon state of this block
	For example, new blockHeight is 91 then beacon state of this block must have height 90
	OR new block has previous has is beacon best block hash
	- Committee length and validatorIndex length
	- Producer + sig
	- Has parent hash is current best block hash in best state
	- Height
	- Epoch
	- staker
	- ShardState
*/
func (beaconBestState *BeaconBestState) VerifyBestStateWithBeaconBlock(block *BeaconBlock, isVerifySig bool) error {

	beaconBestState.lock.RLock()
	defer beaconBestState.lock.RUnlock()
	//=============Verify aggegrate signature
	if isVerifySig {
		// ValidatorIdx must > Number of Beacon Committee / 2 AND Number of Beacon Committee > 3
		if len(beaconBestState.BeaconCommittee) > 3 && len(block.ValidatorsIndex[1]) < (len(beaconBestState.BeaconCommittee)>>1) {
			return NewBlockChainError(SignatureError, errors.New("block validators and Beacon committee is not compatible "+fmt.Sprint(len(block.ValidatorsIndex))))
		}
		err := ValidateAggSignature(block.ValidatorsIndex, beaconBestState.BeaconCommittee, block.AggregatedSig, block.R, block.Hash())
		if err != nil {
			return NewBlockChainError(SignatureError, err)
		}
	}
	//=============End Verify Aggegrate signature
	if beaconBestState.BeaconHeight+1 != block.Header.Height {
		return NewBlockChainError(WrongBlockHeightError, errors.New("block height of new block should be :"+strconv.Itoa(int(block.Header.Height+1))))
	}
	if !beaconBestState.BestBlockHash.IsEqual(&block.Header.PreviousBlockHash) {
		return NewBlockChainError(BeaconBestStateNotCompatibleError, errors.New("previous us block should be :"+beaconBestState.BestBlockHash.String()))
	}
	if block.Header.Height%common.EPOCH == 1 && beaconBestState.Epoch+1 != block.Header.Epoch {
		return NewBlockChainError(EpochError, errors.New("block height and Epoch is not compatiable"))
	}
	if block.Header.Height%common.EPOCH != 1 && beaconBestState.Epoch != block.Header.Epoch {
		return NewBlockChainError(EpochError, errors.New("block height and Epoch is not compatiable"))
	}
	//=============Verify Stakers
	newBeaconCandidate, newShardCandidate := GetStakingCandidate(*block)
	if !reflect.DeepEqual(newBeaconCandidate, []string{}) {
		validBeaconCandidate := beaconBestState.GetValidStakers(newBeaconCandidate)
		if !reflect.DeepEqual(validBeaconCandidate, newBeaconCandidate) {
			return NewBlockChainError(CandidateError, errors.New("beacon candidate list is INVALID"))
		}
	}
	if !reflect.DeepEqual(newShardCandidate, []string{}) {
		validShardCandidate := beaconBestState.GetValidStakers(newShardCandidate)
		if !reflect.DeepEqual(validShardCandidate, newShardCandidate) {
			return NewBlockChainError(CandidateError, errors.New("shard candidate list is INVALID"))
		}
	}
	//=============End Verify Stakers
	// Verify shard state
	// for shardID, shardStates := range block.Body.ShardState {
	// 	// Do not check this condition with first minted block (genesis block height = 1)
	// 	if beaconBestState.BeaconHeight != 2 {
	// fmt.Printf("Beacon Process/Check ShardStates with BestState Current Shard Height %+v \n", beaconBestState.AllShardState[shardID][len(beaconBestState.AllShardState[shardID])-1].Height)
	// fmt.Printf("Beacon Process/Check ShardStates with BestState FirstShardHeight %+v \n", shardStates[0].Height)
	// if shardStates[0].Height-beaconBestState.AllShardState[shardID][len(beaconBestState.AllShardState[shardID])-1].Height != 1 {
	// 	return NewBlockChainError(ShardStateError, errors.New("Shardstates are not compatible with beacon best state"))
	// }
	// }
	// }
	return nil
}

/* Verify Post-processing data
- Validator root: BeaconCommittee + BeaconPendingValidator
- Beacon Candidate root: CandidateBeaconWaitingForCurrentRandom + CandidateBeaconWaitingForNextRandom
- Shard Candidate root: CandidateShardWaitingForCurrentRandom + CandidateShardWaitingForNextRandom
- Shard Validator root: ShardCommittee + ShardPendingValidator
- Random number if have in instruction
*/
func (beaconBestState *BeaconBestState) VerifyPostProcessingBeaconBlock(block *BeaconBlock, snapShotBeaconCommittee []string) error {
	beaconBestState.lock.RLock()
	defer beaconBestState.lock.RUnlock()

	var (
		strs []string
		isOk bool
	)
	//=============Verify producer signature
	producerPubkey := snapShotBeaconCommittee[beaconBestState.BeaconProposerIndex]
	blockHash := block.Header.Hash()
	if err := incognitokey.ValidateDataB58(producerPubkey, block.ProducerSig, blockHash.GetBytes()); err != nil {
		return NewBlockChainError(SignatureError, err)
	}
	//=============End Verify producer signature
	strs = append(strs, beaconBestState.BeaconCommittee...)
	strs = append(strs, beaconBestState.BeaconPendingValidator...)
	isOk = VerifyHashFromStringArray(strs, block.Header.BeaconCommitteeAndValidatorRoot)
	if !isOk {
		return NewBlockChainError(HashError, errors.New("error verify Beacon Validator root"))
	}

	strs = []string{}
	strs = append(strs, beaconBestState.CandidateBeaconWaitingForCurrentRandom...)
	strs = append(strs, beaconBestState.CandidateBeaconWaitingForNextRandom...)
	isOk = VerifyHashFromStringArray(strs, block.Header.BeaconCandidateRoot)
	if !isOk {
		return NewBlockChainError(HashError, errors.New("error verify Beacon Candidate root"))
	}

	strs = []string{}
	strs = append(strs, beaconBestState.CandidateShardWaitingForCurrentRandom...)
	strs = append(strs, beaconBestState.CandidateShardWaitingForNextRandom...)
	isOk = VerifyHashFromStringArray(strs, block.Header.ShardCandidateRoot)
	if !isOk {
		return NewBlockChainError(HashError, errors.New("error verify Shard Candidate root"))
	}

	isOk = VerifyHashFromMapByteString(beaconBestState.ShardPendingValidator, beaconBestState.ShardCommittee, block.Header.ShardCommitteeAndValidatorRoot)
	if !isOk {
		return NewBlockChainError(HashError, errors.New("error verify shard validator root"))
	}

	// COMMENT FOR TESTING
	// instructions := block.Body.Instructions
	// for _, l := range instructions {
	// 	if l[0] == "random" {
	// 		temp, err := strconv.Atoi(l[3])
	// 		if err != nil {
	// 			Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
	// 			return NewBlockChainError(UnExpectedError, err)
	// 		}
	// 		isOk, err = btc.VerifyNonceWithTimestamp(beaconBestState.CurrentRandomTimeStamp, int64(temp))
	// 		Logger.log.Infof("Verify Random number %+v", isOk)
	// 		if err != nil {
	// 			Logger.log.Error("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
	// 			return NewBlockChainError(UnExpectedError, err)
	// 		}
	// 		if !isOk {
	// 			return NewBlockChainError(RandomError, errors.New("Error verify random number"))
	// 		}
	// 	}
	// }
	return nil
}

/*
	Update Beststate with new Block
*/
func (beaconBestState *BeaconBestState) updateBeaconBestState(newBlock *BeaconBlock) error {
	//beaconBestState.lock.Lock()
	//defer beaconBestState.lock.Unlock()
	newBeaconCandidate := []string{}
	newShardCandidate := []string{}
	// Logger.log.Infof("Start processing new block at height %d, with hash %+v", newBlock.Header.Height, *newBlock.Hash())
	if newBlock == nil {
		return errors.New("null pointer")
	}
	// signal of random parameter from beacon block
	randomFlag := false
	// update BestShardHash, BestBlock, BestBlockHash
	beaconBestState.PreviousBestBlockHash = beaconBestState.BestBlockHash
	beaconBestState.BestBlockHash = *newBlock.Hash()
	beaconBestState.BestBlock = *newBlock
	beaconBestState.Epoch = newBlock.Header.Epoch
	beaconBestState.BeaconHeight = newBlock.Header.Height
	if newBlock.Header.Height == 1 {
		beaconBestState.BeaconProposerIndex = 0
	} else {
		beaconBestState.BeaconProposerIndex = common.IndexOfStr(base58.Base58Check{}.Encode(newBlock.Header.ProducerAddress.Pk, common.ZeroByte), beaconBestState.BeaconCommittee)
	}

	allShardState := newBlock.Body.ShardState
	// if beaconBestState.AllShardState == nil {
	// 	beaconBestState.AllShardState = make(map[byte][]ShardState)
	// 	for index := 0; index < common.MAX_SHARD_NUMBER; index++ {
	// 		beaconBestState.AllShardState[byte(index)] = []ShardState{
	// 			ShardState{
	// 				Height: 1,
	// 			},
	// 		}
	// 	}
	// }
	if beaconBestState.BestShardHash == nil {
		beaconBestState.BestShardHash = make(map[byte]common.Hash)
	}
	if beaconBestState.BestShardHeight == nil {
		beaconBestState.BestShardHeight = make(map[byte]uint64)
	}
	// Update new best new block hash
	for shardID, shardStates := range allShardState {
		beaconBestState.BestShardHash[shardID] = shardStates[len(shardStates)-1].Hash
		beaconBestState.BestShardHeight[shardID] = shardStates[len(shardStates)-1].Height
		//if _, ok := beaconBestState.AllShardState[shardID]; !ok {
		//	beaconBestState.AllShardState[shardID] = []ShardState{}
		//}
		//beaconBestState.AllShardState[shardID] = append(beaconBestState.AllShardState[shardID], shardStates...)
	}

	//cross shard state

	// update param
	instructions := newBlock.Body.Instructions
	for _, l := range instructions {
		if len(l) < 1 {
			continue
		}

		if l[0] == SwapAction {
			fmt.Println("SWAP", l)
			// format
			// ["swap" "inPubkey1,inPubkey2,..." "outPupkey1, outPubkey2,..." "shard" "shardID"]
			// ["swap" "inPubkey1,inPubkey2,..." "outPupkey1, outPubkey2,..." "beacon"]
			inPubkeys := strings.Split(l[1], ",")
			outPubkeys := strings.Split(l[2], ",")
			fmt.Println("SWAP l1", l[1])
			fmt.Println("SWAP l2", l[2])
			fmt.Println("SWAP inPubkeys", inPubkeys)
			fmt.Println("SWAP outPubkeys", outPubkeys)
			if l[3] == "shard" {
				temp, err := strconv.Atoi(l[4])
				if err != nil {
					Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
					return NewBlockChainError(UnExpectedError, err)
				}
				shardID := byte(temp)
				// delete in public key out of sharding pending validator list
				if len(l[1]) > 0 {
					fmt.Println("Beacon Process/Update Before, ShardPendingValidator", beaconBestState.ShardPendingValidator[shardID])
					tempShardPendingValidator, err := RemoveValidator(beaconBestState.ShardPendingValidator[shardID], inPubkeys)
					fmt.Println("Beacon Process/Update After, ShardPendingValidator", beaconBestState.ShardPendingValidator[shardID])
					if err != nil {
						Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
						return NewBlockChainError(UnExpectedError, err)
					}
					beaconBestState.ShardPendingValidator[shardID] = tempShardPendingValidator
					// append in public key to committees
					beaconBestState.ShardCommittee[shardID] = append(beaconBestState.ShardCommittee[shardID], inPubkeys...)
					fmt.Println("Beacon Process/Update Add New, ShardCommitees", beaconBestState.ShardCommittee[shardID])
				}
				// delete out public key out of current committees
				if len(l[2]) > 0 {
					tempShardCommittees, err := RemoveValidator(beaconBestState.ShardCommittee[shardID], outPubkeys)
					fmt.Println("Beacon Process/Update Remove Old, ShardCommitees", beaconBestState.ShardCommittee[shardID])
					if err != nil {
						Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
						return NewBlockChainError(UnExpectedError, err)
					}
					beaconBestState.ShardCommittee[shardID] = tempShardCommittees
				}
			} else if l[3] == "beacon" {
				if len(l[1]) > 0 {
					tempBeaconPendingValidator, err := RemoveValidator(beaconBestState.BeaconPendingValidator, inPubkeys)
					if err != nil {
						Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
						return NewBlockChainError(UnExpectedError, err)
					}
					beaconBestState.BeaconPendingValidator = tempBeaconPendingValidator
					beaconBestState.BeaconCommittee = append(beaconBestState.BeaconCommittee, inPubkeys...)
				}
				if len(l[2]) > 0 {
					tempBeaconCommittes, err := RemoveValidator(beaconBestState.BeaconCommittee, outPubkeys)
					if err != nil {
						Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
						return NewBlockChainError(UnExpectedError, err)
					}
					beaconBestState.BeaconCommittee = tempBeaconCommittes
				}
			}
		}
		// ["random" "{nonce}" "{blockheight}" "{timestamp}" "{bitcoinTimestamp}"]
		if l[0] == RandomAction {
			temp, err := strconv.Atoi(l[1])
			if err != nil {
				Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
				return NewBlockChainError(UnExpectedError, err)
			}
			beaconBestState.CurrentRandomNumber = int64(temp)
			Logger.log.Infof("Random number found %+v", beaconBestState.CurrentRandomNumber)
			randomFlag = true
		}
		// Update candidate
		// get staking candidate list and store
		// store new staking candidate
		if l[0] == StakeAction && l[2] == "beacon" {
			beacon := strings.Split(l[1], ",")
			newBeaconCandidate = append(newBeaconCandidate, beacon...)
		}

		if l[0] == StakeAction && l[2] == "shard" {
			shard := strings.Split(l[1], ",")
			newShardCandidate = append(newShardCandidate, shard...)
		}

	}

	if beaconBestState.BeaconHeight == 1 {
		// Assign committee with genesis block
		Logger.log.Infof("Proccessing Genesis Block")
		//Test with 1 member
		beaconBestState.BeaconCommittee = make([]string, beaconBestState.MaxBeaconCommitteeSize)
		copy(beaconBestState.BeaconCommittee, newBeaconCandidate[:beaconBestState.MaxBeaconCommitteeSize])
		for shardID := 0; shardID < beaconBestState.ActiveShards; shardID++ {
			beaconBestState.ShardCommittee[byte(shardID)] = append(beaconBestState.ShardCommittee[byte(shardID)], newShardCandidate[shardID*beaconBestState.MinShardCommitteeSize:(shardID+1)*beaconBestState.MinShardCommitteeSize]...)
			fmt.Println(beaconBestState.ShardCommittee[byte(shardID)])
		}
		beaconBestState.Epoch = 1
	} else {
		beaconBestState.CandidateBeaconWaitingForNextRandom = append(beaconBestState.CandidateBeaconWaitingForNextRandom, newBeaconCandidate...)
		beaconBestState.CandidateShardWaitingForNextRandom = append(beaconBestState.CandidateShardWaitingForNextRandom, newShardCandidate...)
		// fmt.Println("Beacon Process/Before: CandidateShardWaitingForNextRandom: ", beaconBestState.CandidateShardWaitingForNextRandom)
	}

	if beaconBestState.BeaconHeight%common.EPOCH == 1 && beaconBestState.BeaconHeight != 1 {
		beaconBestState.IsGetRandomNumber = false
		// Begin of each epoch
	} else if beaconBestState.BeaconHeight%common.EPOCH < common.RANDOM_TIME {
		// Before get random from bitcoin
	} else if beaconBestState.BeaconHeight%common.EPOCH >= common.RANDOM_TIME {
		// After get random from bitcoin
		if beaconBestState.BeaconHeight%common.EPOCH == common.RANDOM_TIME {
			// snapshot candidate list
			beaconBestState.CandidateShardWaitingForCurrentRandom = beaconBestState.CandidateShardWaitingForNextRandom
			beaconBestState.CandidateBeaconWaitingForCurrentRandom = beaconBestState.CandidateBeaconWaitingForNextRandom
			Logger.log.Critical("==================Beacon Process: Snapshot candidate====================")
			Logger.log.Critical("Beacon Process: CandidateShardWaitingForCurrentRandom: ", beaconBestState.CandidateShardWaitingForCurrentRandom)
			Logger.log.Critical("Beacon Process: CandidateBeaconWaitingForCurrentRandom: ", beaconBestState.CandidateBeaconWaitingForCurrentRandom)
			// reset candidate list
			beaconBestState.CandidateShardWaitingForNextRandom = []string{}
			beaconBestState.CandidateBeaconWaitingForNextRandom = []string{}
			Logger.log.Critical("Beacon Process/After: CandidateShardWaitingForNextRandom: ", beaconBestState.CandidateShardWaitingForNextRandom)
			Logger.log.Critical("Beacon Process/After: CandidateBeaconWaitingForCurrentRandom: ", beaconBestState.CandidateBeaconWaitingForCurrentRandom)
			// assign random timestamp
			beaconBestState.CurrentRandomTimeStamp = newBlock.Header.Timestamp
		}
		// if get new random number
		// Assign candidate to shard
		// assign CandidateShardWaitingForCurrentRandom to ShardPendingValidator with CurrentRandom
		if randomFlag {
			beaconBestState.IsGetRandomNumber = true
			//fmt.Println("Beacon Process/Update/RandomFlag: Shard Candidate Waiting for Current Random Number", beaconBestState.CandidateShardWaitingForCurrentRandom)
			//Logger.log.Critical("beaconBestState.ShardPendingValidator", beaconBestState.ShardPendingValidator)
			//Logger.log.Critical("beaconBestState.CandidateShardWaitingForCurrentRandom", beaconBestState.CandidateShardWaitingForCurrentRandom)
			//Logger.log.Critical("beaconBestState.CurrentRandomNumber", beaconBestState.CurrentRandomNumber)
			//Logger.log.Critical("beaconBestState.ActiveShards", beaconBestState.ActiveShards)
			err := AssignValidatorShard(beaconBestState.ShardPendingValidator, beaconBestState.CandidateShardWaitingForCurrentRandom, beaconBestState.CurrentRandomNumber, beaconBestState.ActiveShards)
			if err != nil {
				Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
				return NewBlockChainError(UnExpectedError, err)
			}
			// delete CandidateShardWaitingForCurrentRandom list
			beaconBestState.CandidateShardWaitingForCurrentRandom = []string{}
			//fmt.Println("Beacon Process/Update/RandomFalg: Shard Pending Validator", beaconBestState.ShardPendingValidator)
			// Shuffle candidate
			// shuffle CandidateBeaconWaitingForCurrentRandom with current random number
			//fmt.Println("Beacon Process/Update/RandomFlag: Beacon Candidate Waiting for Current Random Number", beaconBestState.CandidateBeaconWaitingForCurrentRandom)
			newBeaconPendingValidator, err := ShuffleCandidate(beaconBestState.CandidateBeaconWaitingForCurrentRandom, beaconBestState.CurrentRandomNumber)
			//fmt.Println("Beacon Process/Update/RandomFalg: NewBeaconPendingValidator", newBeaconPendingValidator)
			if err != nil {
				Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
				return NewBlockChainError(UnExpectedError, err)
			}
			beaconBestState.CandidateBeaconWaitingForCurrentRandom = []string{}
			beaconBestState.BeaconPendingValidator = append(beaconBestState.BeaconPendingValidator, newBeaconPendingValidator...)
			//fmt.Println("Beacon Process/Update/RandomFalg: Beacon Pending Validator", beaconBestState.BeaconPendingValidator)
			if err != nil {
				return err
			}
		}
	} else if beaconBestState.BeaconHeight%common.EPOCH == 0 {
		// At the end of each epoch, eg: block 200, 400, 600 with epoch is 200 blocks
		// Swap pending validator in committees, pop some of public key in committees out
		// ONLY SWAP FOR BEACON
		// SHARD WILL SWAP ITblockchain
		var (
			beaconSwapedCommittees []string
			beaconNewCommittees    []string
			err                    error
		)
		beaconBestState.BeaconPendingValidator, beaconBestState.BeaconCommittee, beaconSwapedCommittees, beaconNewCommittees, err = SwapValidator(beaconBestState.BeaconPendingValidator, beaconBestState.BeaconCommittee, beaconBestState.MaxBeaconCommitteeSize, common.OFFSET)
		if err != nil {
			Logger.log.Errorf("Blockchain Error %+v", NewBlockChainError(UnExpectedError, err))
			return NewBlockChainError(UnExpectedError, err)
		}
		Logger.log.Info("Swap: Out committee %+v", beaconSwapedCommittees)
		Logger.log.Info("Swap: In committee %+v", beaconNewCommittees)
	}
	return nil
}

//===================================Util for Beacon=============================
func GetStakingCandidate(beaconBlock BeaconBlock) ([]string, []string) {
	beacon := []string{}
	shard := []string{}
	beaconBlockBody := beaconBlock.Body
	for _, v := range beaconBlockBody.Instructions {
		if len(v) < 1 {
			continue
		}
		if v[0] == StakeAction && v[2] == "beacon" {
			beacon = strings.Split(v[1], ",")
		}
		if v[0] == StakeAction && v[2] == "shard" {
			shard = strings.Split(v[1], ",")
		}
	}

	return beacon, shard
}

// Assumption:
// validator and candidate public key encode as base58 string
// assume that candidates are already been checked
// Check validation of candidate in transaction
func AssignValidator(candidates []string, rand int64, activeShards int) (map[byte][]string, error) {
	pendingValidators := make(map[byte][]string)
	for _, candidate := range candidates {
		shardID := calculateCandidateShardID(candidate, rand, activeShards)
		pendingValidators[shardID] = append(pendingValidators[shardID], candidate)
	}
	return pendingValidators, nil
}

// AssignValidatorShard, param for better convenice than AssignValidator
func AssignValidatorShard(currentCandidates map[byte][]string, shardCandidates []string, rand int64, activeShards int) error {
	for _, candidate := range shardCandidates {
		shardID := calculateCandidateShardID(candidate, rand, activeShards)
		currentCandidates[shardID] = append(currentCandidates[shardID], candidate)
	}
	return nil
}

func VerifyValidator(candidate string, rand int64, shardID byte, activeShards int) (bool, error) {
	res := calculateCandidateShardID(candidate, rand, activeShards)
	if shardID == res {
		return true, nil
	} else {
		return false, nil
	}
}

// Formula ShardID: LSB[hash(candidatePubKey+randomNumber)]
// Last byte of hash(candidatePubKey+randomNumber)
func calculateCandidateShardID(candidate string, rand int64, activeShards int) (shardID byte) {

	seed := candidate + strconv.Itoa(int(rand))
	hash := common.HashB([]byte(seed))
	// fmt.Println("Candidate public key", candidate)
	// fmt.Println("Hash of candidate serialized pubkey and random number", hash)
	// fmt.Printf("\"%d\",\n", hash[len(hash)-1])
	// fmt.Println("Shard to be assign", hash[len(hash)-1])
	shardID = byte(int(hash[len(hash)-1]) % activeShards)
	Logger.log.Critical("calculateCandidateShardID/shardID", shardID)
	return shardID
}

// consider these list as queue structure
// unqueue a number of validator out of currentValidators list
// enqueue a number of validator into currentValidators list <=> unqueue a number of validator out of pendingValidators list
// return value: #1 remaining pendingValidators, #2 new currentValidators #3 swapped out validator, #4 incoming validator #5 error
func SwapValidator(pendingValidators []string, currentValidators []string, maxCommittee int, offset int) ([]string, []string, []string, []string, error) {
	if maxCommittee < 0 || offset < 0 {
		panic("committee can't be zero")
	}
	if offset == 0 {
		return []string{}, pendingValidators, currentValidators, []string{}, errors.New("can't not swap 0 validator")
	}
	// if number of pending validator is less or equal than offset, set offset equal to number of pending validator
	if offset > len(pendingValidators) {
		offset = len(pendingValidators)
	}
	// if swap offset = 0 then do nothing
	if offset == 0 {
		return pendingValidators, currentValidators, []string{}, []string{}, errors.New("no pending validator for swapping")
	}
	if offset > maxCommittee {
		return pendingValidators, currentValidators, []string{}, []string{}, errors.New("trying to swap too many validator")
	}
	tempValidators := []string{}
	swapValidator := []string{}
	// if len(currentValidator) < maxCommittee then push validator until it is full
	if len(currentValidators) < maxCommittee {
		diff := maxCommittee - len(currentValidators)
		if diff >= offset {
			tempValidators = append(tempValidators, pendingValidators[:offset]...)
			currentValidators = append(currentValidators, tempValidators...)
			pendingValidators = pendingValidators[offset:]
			return pendingValidators, currentValidators, swapValidator, tempValidators, nil
		} else {
			offset -= diff
			tempValidators := append(tempValidators, pendingValidators[:diff]...)
			pendingValidators = pendingValidators[diff:]
			currentValidators = append(currentValidators, tempValidators...)
		}
	}
	fmt.Println("Swap Validator/Before: pendingValidators", pendingValidators)
	fmt.Println("Swap Validator/Before: currentValidators", currentValidators)
	fmt.Println("Swap Validator: offset", offset)
	// out pubkey: swapped out validator
	swapValidator = append(swapValidator, currentValidators[:offset]...)
	// unqueue validator with index from 0 to offset-1 from currentValidators list
	currentValidators = currentValidators[offset:]
	// in pubkey: unqueue validator with index from 0 to offset-1 from pendingValidators list
	tempValidators = append(tempValidators, pendingValidators[:offset]...)
	// enqueue new validator to the remaning of current validators list
	currentValidators = append(currentValidators, pendingValidators[:offset]...)
	// save new pending validators list
	pendingValidators = pendingValidators[offset:]
	fmt.Println("Swap Validator: pendingValidators", pendingValidators)
	fmt.Println("Swap Validator: currentValidators", currentValidators)
	fmt.Println("Swap Validator: swapValidator", swapValidator)
	fmt.Println("Swap Validator: tempValidators", tempValidators)
	if len(currentValidators) > maxCommittee {
		panic("Length of current validator greater than max committee in Swap validator ")
	}
	return pendingValidators, currentValidators, swapValidator, tempValidators, nil
}

// return: #param1: validator list after remove
// in parameter: #param1: list of full validator
// in parameter: #param2: list of removed validator
// removed validators list must be a subset of full validator list and it must be first in the list
func RemoveValidator(validators []string, removedValidators []string) ([]string, error) {
	// if number of pending validator is less or equal than offset, set offset equal to number of pending validator
	if len(removedValidators) > len(validators) {
		return validators, errors.New("trying to remove too many validators")
	}
	for index, validator := range removedValidators {
		if strings.Compare(validators[index], validator) == 0 {
			validators = validators[1:]
		} else {
			// not found wanted validator
			return validators, errors.New("remove Validator with Wrong Format")
		}
	}
	return validators, nil
}

/*
	Shuffle Candidate:
		Candidate Value Concatenate with Random Number
		Then Hash and Obtain Hash Value
		Sort Hash Value Then Re-arrange Candidate corresponding to Hash Value
*/
func ShuffleCandidate(candidates []string, rand int64) ([]string, error) {
	fmt.Println("Beacon Process/Shuffle Candidate: Candidate Before Sort ", candidates)
	hashes := []string{}
	m := make(map[string]string)
	sortedCandidate := []string{}
	for _, candidate := range candidates {
		seed := candidate + strconv.Itoa(int(rand))
		hash := common.HashB([]byte(seed))
		hashes = append(hashes, string(hash[:32]))
		m[string(hash[:32])] = candidate
	}
	sort.Strings(hashes)
	for _, candidate := range m {
		sortedCandidate = append(sortedCandidate, candidate)
	}
	fmt.Println("Beacon Process/Shuffle Candidate: Candidate After Sort ", sortedCandidate)
	return sortedCandidate, nil
}

/*
	Kick a list of candidate out of current validators list
	Candidates will be eliminated as the list order (from 0 index to last index)
	A candidate will be click out of list if it match those condition:
		- candidate pubkey found in current validators list
		- size of current validator list is greater or equal to min committess size
	Return params:
	#1 kickedValidator, #2 remain candidates (not kick yet), #3 new current validator list
*/
func kickValidatorByPubkeyList(candidates []string, currentValidators []string, minCommitteeSize int) ([]string, []string, []string) {
	removedCandidates := []string{}
	remainedCandidates := []string{}
	remainedIndex := 0
	for index, candidate := range candidates {
		remainedIndex = index
		if len(currentValidators) == minCommitteeSize {
			break
		}
		if index := common.IndexOfStr(candidate, currentValidators); index < 0 {
			remainedCandidates = append(remainedCandidates, candidate)
			continue
		} else {
			removedCandidates = append(removedCandidates, candidate)
			currentValidators = append(currentValidators[:index], currentValidators[index+1:]...)
		}
	}
	if remainedIndex < len(candidates)-1 {
		remainedCandidates = append(remainedCandidates, candidates[remainedIndex:]...)
	}
	return removedCandidates, remainedCandidates, currentValidators
}
func kickValidatorByPubkey(candidate string, currentValidators []string, minCommitteeSize int) (bool, []string) {
	if index := common.IndexOfStr(candidate, currentValidators); index < 0 {
		return false, currentValidators
	} else {
		currentValidators = append(currentValidators[:index], currentValidators[index+1:]...)
		return true, currentValidators
	}
}
