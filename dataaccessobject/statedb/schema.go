package statedb

import (
	"bytes"
	"fmt"
	"github.com/incognitochain/incognito-chain/common"
	"sort"
	"strconv"
)

var (
	committeePrefix                    = []byte("shard-com-")
	substitutePrefix                   = []byte("shard-sub-")
	nextShardCandidatePrefix           = []byte("next-sha-cand-")
	currentShardCandidatePrefix        = []byte("cur-sha-cand-")
	nextBeaconCandidatePrefix          = []byte("next-bea-cand-")
	currentBeaconCandidatePrefix       = []byte("cur-bea-cand-")
	committeeRewardPrefix              = []byte("committee-reward-")
	rewardRequestPrefix                = []byte("reward-request-")
	blackListProducerPrefix            = []byte("black-list-")
	serialNumberPrefix                 = []byte("serial-number-")
	commitmentPrefix                   = []byte("com-value-")
	commitmentIndexPrefix              = []byte("com-index-")
	commitmentLengthPrefix             = []byte("com-length-")
	snDerivatorPrefix                  = []byte("sn-derivator-")
	outputCoinPrefix                   = []byte("output-coin-")
	tokenPrefix                        = []byte("token-")
	waitingPDEContributionPrefix       = []byte("waitingpdecontribution-")
	pdePoolPrefix                      = []byte("pdepool-")
	pdeSharePrefix                     = []byte("pdeshare-")
	pdeTradeFeePrefix                  = []byte("pdetradefee-")
	pdeContributionStatusPrefix        = []byte("pdecontributionstatus-")
	pdeTradeStatusPrefix               = []byte("pdetradestatus-")
	pdeWithdrawalStatusPrefix          = []byte("pdewithdrawalstatus-")
	pdeStatusPrefix                    = []byte("pdestatus-")
	bridgeEthTxPrefix                  = []byte("bri-eth-tx-")
	bridgeCentralizedTokenInfoPrefix   = []byte("bri-cen-token-info-")
	bridgeDecentralizedTokenInfoPrefix = []byte("bri-de-token-info-")
	bridgeStatusPrefix                 = []byte("bri-status-")
	burnPrefix                         = []byte("burn-")
)

func GetCommitteePrefixWithRole(role int, shardID int) []byte {
	switch role {
	case NextEpochShardCandidate:
		temp := []byte(string(nextShardCandidatePrefix))
		h := common.HashH(temp)
		return h[:][:prefixHashKeyLength]
	case CurrentEpochShardCandidate:
		temp := []byte(string(currentShardCandidatePrefix))
		h := common.HashH(temp)
		return h[:][:prefixHashKeyLength]
	case NextEpochBeaconCandidate:
		temp := []byte(string(nextBeaconCandidatePrefix))
		h := common.HashH(temp)
		return h[:][:prefixHashKeyLength]
	case CurrentEpochBeaconCandidate:
		temp := []byte(string(currentBeaconCandidatePrefix))
		h := common.HashH(temp)
		return h[:][:prefixHashKeyLength]
	case SubstituteValidator:
		temp := []byte(string(substitutePrefix) + strconv.Itoa(shardID))
		h := common.HashH(temp)
		return h[:][:prefixHashKeyLength]
	case CurrentValidator:
		temp := []byte(string(committeePrefix) + strconv.Itoa(shardID))
		h := common.HashH(temp)
		return h[:][:prefixHashKeyLength]
	default:
		panic("role not exist: " + strconv.Itoa(role))
	}
}
func GetCommitteeRewardPrefix() []byte {
	h := common.HashH(committeeRewardPrefix)
	return h[:][:prefixHashKeyLength]
}

func GetRewardRequestPrefix(epoch uint64) []byte {
	buf := common.Uint64ToBytes(epoch)
	temp := append(rewardRequestPrefix, buf...)
	h := common.HashH(temp)
	return h[:][:prefixHashKeyLength]
}

func GetBlackListProducerPrefix() []byte {
	h := common.HashH(blackListProducerPrefix)
	return h[:][:prefixHashKeyLength]
}

func GetSerialNumberPrefix(tokenID common.Hash, shardID byte) []byte {
	h := common.HashH(append(serialNumberPrefix, append(tokenID[:], shardID)...))
	return h[:][:prefixHashKeyLength]
}

func GetCommitmentPrefix(tokenID common.Hash, shardID byte) []byte {
	h := common.HashH(append(commitmentPrefix, append(tokenID[:], shardID)...))
	return h[:][:prefixHashKeyLength]
}

func GetCommitmentIndexPrefix(tokenID common.Hash, shardID byte) []byte {
	h := common.HashH(append(commitmentIndexPrefix, append(tokenID[:], shardID)...))
	return h[:][:prefixHashKeyLength]
}

func GetCommitmentLengthPrefix() []byte {
	h := common.HashH(commitmentLengthPrefix)
	return h[:][:prefixHashKeyLength]
}

func GetSNDerivatorPrefix(tokenID common.Hash) []byte {
	h := common.HashH(append(snDerivatorPrefix, tokenID[:]...))
	return h[:][:prefixHashKeyLength]
}

func GetOutputCoinPrefix(tokenID common.Hash, shardID byte) []byte {
	h := common.HashH(append(outputCoinPrefix, append(tokenID[:], shardID)...))
	return h[:][:prefixHashKeyLength]
}

func GetTokenPrefix() []byte {
	h := common.HashH(tokenPrefix)
	return h[:][:prefixHashKeyLength]
}

func GetWaitingPDEContributionPrefixV2() []byte {
	h := common.HashH(waitingPDEContributionPrefix)
	return h[:][:prefixHashKeyLength]
}

func GetPDEPoolPairPrefixV2() []byte {
	h := common.HashH(pdePoolPrefix)
	return h[:][:prefixHashKeyLength]
}

func GetPDESharePrefixV2() []byte {
	h := common.HashH(pdeSharePrefix)
	return h[:][:prefixHashKeyLength]
}

func GetPDEStatusPrefix() []byte {
	h := common.HashH(pdeStatusPrefix)
	return h[:][:prefixHashKeyLength]
}

func GetBridgeEthTxPrefix() []byte {
	h := common.HashH(bridgeEthTxPrefix)
	return h[:][:prefixHashKeyLength]
}

func GetBridgeTokenInfoPrefix(isCentralized bool) []byte {
	if isCentralized {
		h := common.HashH(bridgeCentralizedTokenInfoPrefix)
		return h[:][:prefixHashKeyLength]
	} else {
		h := common.HashH(bridgeDecentralizedTokenInfoPrefix)
		return h[:][:prefixHashKeyLength]
	}
}

func GetBridgeStatusPrefix() []byte {
	h := common.HashH(bridgeStatusPrefix)
	return h[:][:prefixHashKeyLength]
}

func GetBurningConfirmPrefix() []byte {
	h := common.HashH(burnPrefix)
	return h[:][:prefixHashKeyLength]
}
func WaitingPDEContributionPrefix() []byte {
	return waitingPDEContributionPrefix
}
func PDEPoolPrefix() []byte {
	return pdePoolPrefix
}
func PDESharePrefix() []byte {
	return pdeSharePrefix
}
func PDETradeFeePrefix() []byte {
	return pdeTradeFeePrefix
}
func PDEContributionStatusPrefix() []byte {
	return pdeContributionStatusPrefix
}
func PDETradeStatusPrefix() []byte {
	return pdeTradeStatusPrefix
}
func PDEWithdrawalStatusPrefix() []byte {
	return pdeWithdrawalStatusPrefix
}

// GetWaitingPDEContributionKey: WaitingPDEContributionPrefix - beacon height - pairid
func GetWaitingPDEContributionKey(beaconHeight uint64, pairID string) []byte {
	prefix := append(waitingPDEContributionPrefix, []byte(fmt.Sprintf("%d-", beaconHeight))...)
	return append(prefix, []byte(pairID)...)
}

// GetPDEPoolForPairKey: PDEPoolPrefix - beacon height - token1ID - token2ID
func GetPDEPoolForPairKey(beaconHeight uint64, token1ID string, token2ID string) []byte {
	prefix := append(pdePoolPrefix, []byte(fmt.Sprintf("%d-", beaconHeight))...)
	tokenIDs := []string{token1ID, token2ID}
	sort.Strings(tokenIDs)
	return append(prefix, []byte(tokenIDs[0]+"-"+tokenIDs[1])...)
}

// GetPDEShareKey: PDESharePrefix + beacon height + token1ID + token2ID + contributor address
func GetPDEShareKey(beaconHeight uint64, token1ID string, token2ID string, contributorAddress string) []byte {
	prefix := append(pdeSharePrefix, []byte(fmt.Sprintf("%d-", beaconHeight))...)
	tokenIDs := []string{token1ID, token2ID}
	sort.Strings(tokenIDs)
	return append(prefix, []byte(tokenIDs[0]+"-"+tokenIDs[1]+"-"+contributorAddress)...)
}

func GetPDEStatusKey(prefix []byte, suffix []byte) []byte {
	return append(prefix, suffix...)
}

//func GetPDETradeFeesKey(beaconHeight uint64, token1IDStr string, token2IDStr string, tokenForFeeIDStr string) []byte {
//	beaconHeightBytes := []byte(fmt.Sprintf("%d-", beaconHeight))
//	pdeTradeFeesByBCHeightPrefix := append(PDETradeFeePrefix, beaconHeightBytes...)
//	tokenIDStrs := []string{token1IDStr, token2IDStr}
//	sort.Strings(tokenIDStrs)
//	return append(pdeTradeFeesByBCHeightPrefix, []byte(tokenIDStrs[0]+"-"+tokenIDStrs[1]+"-"+tokenForFeeIDStr)...)
//}

var _ = func() (_ struct{}) {
	m := make(map[string]string)
	prefixs := [][]byte{}
	// Current validator
	for i := -1; i < 256; i++ {
		temp := GetCommitteePrefixWithRole(CurrentValidator, i)
		prefixs = append(prefixs, temp)
		if v, ok := m[string(temp)]; ok {
			panic("shard-com-" + strconv.Itoa(i) + " same prefix " + v)
		}
		m[string(temp)] = "shard-com-" + strconv.Itoa(i)
	}
	// Substitute validator
	for i := -1; i < 256; i++ {
		temp := GetCommitteePrefixWithRole(SubstituteValidator, i)
		prefixs = append(prefixs, temp)
		if v, ok := m[string(temp)]; ok {
			panic("shard-sub-" + strconv.Itoa(i) + " same prefix " + v)
		}
		m[string(temp)] = "shard-sub-" + strconv.Itoa(i)
	}
	// Current Candidate
	tempCurrentCandidate := GetCommitteePrefixWithRole(CurrentEpochShardCandidate, -2)
	prefixs = append(prefixs, tempCurrentCandidate)
	if v, ok := m[string(tempCurrentCandidate)]; ok {
		panic("cur-cand-" + " same prefix " + v)
	}
	m[string(tempCurrentCandidate)] = "cur-cand-"
	// Next candidate
	tempNextCandidate := GetCommitteePrefixWithRole(NextEpochShardCandidate, -2)
	prefixs = append(prefixs, tempNextCandidate)
	if v, ok := m[string(tempNextCandidate)]; ok {
		panic("next-cand-" + " same prefix " + v)
	}
	m[string(tempNextCandidate)] = "next-cand-"
	// serial number
	//tempSerialNumber := GetSerialNumberPrefix()
	//prefixs = append(prefixs, tempSerialNumber)
	//if v, ok := m[string(tempSerialNumber)]; ok {
	//	panic("serial-number-" + " same prefix " + v)
	//}
	//m[string(tempSerialNumber)] = "serial-number-"
	// reward receiver
	tempRewardReceiver := GetCommitteeRewardPrefix()
	prefixs = append(prefixs, tempRewardReceiver)
	if v, ok := m[string(tempRewardReceiver)]; ok {
		panic("committee-reward-" + " same prefix " + v)
	}
	m[string(tempRewardReceiver)] = "committee-reward-"
	// black list producer
	tempBlackListProducer := GetBlackListProducerPrefix()
	prefixs = append(prefixs, tempBlackListProducer)
	if v, ok := m[string(tempBlackListProducer)]; ok {
		panic("black-list-" + " same prefix " + v)
	}
	m[string(tempBlackListProducer)] = "black-list-"
	for i, v1 := range prefixs {
		for j, v2 := range prefixs {
			if i == j {
				continue
			}
			if bytes.HasPrefix(v1, v2) || bytes.HasPrefix(v2, v1) {
				panic("(prefix: " + fmt.Sprintf("%+v", v1) + ", value: " + m[string(v1)] + ")" + " is prefix or being prefix of " + " (prefix: " + fmt.Sprintf("%+v", v1) + ", value: " + m[string(v2)] + ")")
			}
		}
	}
	return
}()
