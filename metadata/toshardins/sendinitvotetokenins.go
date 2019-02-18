package toshardins

import (
	"encoding/json"
	"github.com/ninjadotorg/constant/common"
	"github.com/ninjadotorg/constant/database"
	"github.com/ninjadotorg/constant/metadata"
	"github.com/ninjadotorg/constant/privacy"
	"github.com/ninjadotorg/constant/transaction"
	"strconv"
)

type TxSendInitDCBVoteTokenMetadataIns struct {
	Amount                 uint32
	ReceiverPaymentAddress privacy.PaymentAddress
}

func (txSendInitDCBVoteTokenMetadataIns *TxSendInitDCBVoteTokenMetadataIns) GetStringFormat() ([]string, error) {
	content, err := json.Marshal(txSendInitDCBVoteTokenMetadataIns)
	if err != nil {
		return nil, err
	}
	shardID := GetShardIDFromPaymentAddressBytes(txSendInitDCBVoteTokenMetadataIns.ReceiverPaymentAddress)
	return []string{
		strconv.Itoa(metadata.SendInitDCBVoteTokenMeta),
		strconv.Itoa(int(shardID)),
		string(content),
	}, nil
}

func NewTxSendInitDCBVoteTokenMetadataIns(amount uint32, receiverPaymentAddress privacy.PaymentAddress) *TxSendInitDCBVoteTokenMetadataIns {
	return &TxSendInitDCBVoteTokenMetadataIns{Amount: amount, ReceiverPaymentAddress: receiverPaymentAddress}
}

func (txSendInitDCBVoteTokenMetadataIns *TxSendInitDCBVoteTokenMetadataIns) BuildTransaction(
	minerPrivateKey *privacy.SpendingKey,
	db database.DatabaseInterface,
) metadata.Transaction {
	meta := metadata.NewSendInitDCBVoteTokenMetadata(
		txSendInitDCBVoteTokenMetadataIns.Amount,
		txSendInitDCBVoteTokenMetadataIns.ReceiverPaymentAddress,
	)
	sendVoteTokenTransaction := transaction.NewEmptyTx(
		minerPrivateKey,
		db,
		meta,
	)
	return sendVoteTokenTransaction
}

type TxSendInitGOVVoteTokenMetadataIns struct {
	Amount                 uint32
	ReceiverPaymentAddress privacy.PaymentAddress
}

func (txSendInitGOVVoteTokenMetadataIns *TxSendInitGOVVoteTokenMetadataIns) GetStringFormat() ([]string, error) {
	content, err := json.Marshal(txSendInitGOVVoteTokenMetadataIns)
	if err != nil {
		return nil, err
	}
	shardID := GetShardIDFromPaymentAddressBytes(txSendInitGOVVoteTokenMetadataIns.ReceiverPaymentAddress)
	return []string{
		strconv.Itoa(metadata.SendInitGOVVoteTokenMeta),
		strconv.Itoa(int(shardID)),
		string(content),
	}, nil
}

func NewTxSendInitGOVVoteTokenMetadataIns(amount uint32, receiverPaymentAddress privacy.PaymentAddress) *TxSendInitGOVVoteTokenMetadataIns {
	return &TxSendInitGOVVoteTokenMetadataIns{Amount: amount, ReceiverPaymentAddress: receiverPaymentAddress}
}

func (txSendInitGOVVoteTokenMetadataIns *TxSendInitGOVVoteTokenMetadataIns) BuildTransaction(
	minerPrivateKey *privacy.SpendingKey,
	db database.DatabaseInterface,
) metadata.Transaction {
	meta := metadata.NewSendInitGOVVoteTokenMetadata(
		txSendInitGOVVoteTokenMetadataIns.Amount,
		txSendInitGOVVoteTokenMetadataIns.ReceiverPaymentAddress,
	)
	sendVoteTokenTransaction := transaction.NewEmptyTx(
		minerPrivateKey,
		db,
		meta,
	)
	return sendVoteTokenTransaction
}

func NewTxSendInitVoteTokenMetadataIns(
	boardType byte,
	amount uint32,
	receiverPaymentAddress privacy.PaymentAddress,
) Instruction {
	if boardType == common.DCBBoard {
		return NewTxSendInitDCBVoteTokenMetadataIns(amount, receiverPaymentAddress)
	} else {
		return NewTxSendInitGOVVoteTokenMetadataIns(amount, receiverPaymentAddress)
	}
}
