package onetime_address

import (
	"errors"
	"math/big"

	"github.com/incognitochain/incognito-chain/privacy"
	"github.com/incognitochain/incognito-chain/privacy-v2/onetime_address/address"
	"github.com/incognitochain/incognito-chain/privacy-v2/onetime_address/utxo"
)

// Create output of utxos and sum of blind values for later usage
func CreateOutputs(addressesPointer *[]address.PublicAddress, moneyPointer *[]big.Int) (*[]utxo.Utxo, *privacy.Scalar, error) {
	addr := *addressesPointer
	money := *moneyPointer

	result := make([]utxo.Utxo, len(addr))
	if len(addr) > 256 {
		return nil, nil, errors.New("Error in tx_full CreateOutputs: Cannot create too much output (maximum is 256)")
	}
	if len(addr) != len(money) {
		return nil, nil, errors.New("Error in tx_full CreateOutputs: addresses and money array length not the same")
	}
	sumBlind := new(privacy.Scalar)
	for i := 0; i < len(addr); i += 1 {
		r := privacy.RandomScalar()              // Step 1 Monero One-time address
		blind := privacy.RandomScalar()          // Also create blind for money
		sumBlind = sumBlind.Add(sumBlind, blind) // The creator of outputs should know sum of blind for signature

		addressee, txData, cachedHash := parseOTAWithCached(r, &addr[i], byte(i))
		mask, amount, commitment, err := parseMoneyToCreateOutput(blind, cachedHash, &money[i], byte(i))
		if err != nil {
			return nil, nil, errors.New("Error in tx_full CreateOutputs: money of the output is invalid")
		}
		result[i] = *utxo.NewUtxo(uint8(i), mask, amount, txData, addressee, commitment)
	}
	return &result, sumBlind, nil
}

// Check whether the utxo is from this address
func IsUtxoOfAddress(addr *address.PrivateAddress, utxo *utxo.Utxo) bool {
	rK := new(privacy.Point).ScalarMult(utxo.GetTxData(), addr.GetPrivateView())

	hashed := privacy.HashToScalar(
		append(rK.ToBytesS(), utxo.GetIndex()),
	)
	HnG := new(privacy.Point).ScalarMultBase(hashed)

	KCheck := new(privacy.Point).Sub(utxo.GetAddressee(), HnG)

	// TODO
	return *KCheck == *addr.GetPublicSpend()
}
