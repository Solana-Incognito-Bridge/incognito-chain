package pdexv3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/dataaccessobject/statedb"
	metadataCommon "github.com/incognitochain/incognito-chain/metadata/common"
	"github.com/incognitochain/incognito-chain/privacy"
)

type UserMintNftRequest struct {
	metadataCommon.MetadataBase
	otaReceive string
	amount     uint64
}

func NewUserMintNftRequest() *UserMintNftRequest {
	return &UserMintNftRequest{}
}

func NewUserMintNftRequestWithValue(otaReceive string, amount uint64) *UserMintNftRequest {
	metadataBase := metadataCommon.MetadataBase{
		Type: metadataCommon.Pdexv3UserMintNftRequestMeta,
	}
	return &UserMintNftRequest{
		otaReceive:   otaReceive,
		amount:       amount,
		MetadataBase: metadataBase,
	}
}

func (request *UserMintNftRequest) ValidateTxWithBlockChain(
	tx metadataCommon.Transaction,
	chainRetriever metadataCommon.ChainRetriever,
	shardViewRetriever metadataCommon.ShardViewRetriever,
	beaconViewRetriever metadataCommon.BeaconViewRetriever,
	shardID byte,
	transactionStateDB *statedb.StateDB,
) (bool, error) {
	if err := beaconViewRetriever.IsValidMintNftRequireAmount(request.amount); err != nil {
		return false, err
	}
	return true, nil
}

func (request *UserMintNftRequest) ValidateSanityData(
	chainRetriever metadataCommon.ChainRetriever,
	shardViewRetriever metadataCommon.ShardViewRetriever,
	beaconViewRetriever metadataCommon.BeaconViewRetriever,
	beaconHeight uint64,
	tx metadataCommon.Transaction,
) (bool, bool, error) {
	otaReceive := privacy.OTAReceiver{}
	err := otaReceive.FromString(request.otaReceive)
	if err != nil {
		return false, false, metadataCommon.NewMetadataTxError(metadataCommon.PDEInvalidMetadataValueError, err)
	}
	if !otaReceive.IsValid() {
		return false, false, metadataCommon.NewMetadataTxError(metadataCommon.PDEInvalidMetadataValueError, errors.New("ReceiveAddress is not valid"))
	}
	if request.amount == 0 {
		return false, false, metadataCommon.NewMetadataTxError(metadataCommon.PDEInvalidMetadataValueError, errors.New("request.amount is 0"))
	}

	isBurned, burnCoin, burnedTokenID, err := tx.GetTxBurnData()
	if err != nil || !isBurned {
		return false, false, metadataCommon.NewMetadataTxError(metadataCommon.PDENotBurningTxError, err)
	}
	if !bytes.Equal(burnedTokenID[:], common.PRVCoinID[:]) {
		return false, false, metadataCommon.NewMetadataTxError(metadataCommon.PDEInvalidMetadataValueError, errors.New("Wrong request info's token id, it should be equal to tx's token id"))
	}
	if burnCoin.GetValue() != request.amount {
		err := fmt.Errorf("Burnt amount is not valid expect %v but get %v", request.amount, burnCoin.GetValue())
		return false, false, metadataCommon.NewMetadataTxError(metadataCommon.PDEInvalidMetadataValueError, err)
	}
	if tx.GetType() != common.TxNormalType {
		return false, false, metadataCommon.NewMetadataTxError(metadataCommon.PDEInvalidMetadataValueError, errors.New("Tx type must be normal privacy type"))
	}
	return true, true, nil
}

func (request *UserMintNftRequest) ValidateMetadataByItself() bool {
	return request.Type == metadataCommon.Pdexv3UserMintNftRequestMeta
}

func (request *UserMintNftRequest) Hash() *common.Hash {
	record := request.MetadataBase.Hash().String()
	record += request.otaReceive
	record += strconv.FormatUint(uint64(request.amount), 10)
	// final hash
	hash := common.HashH([]byte(record))
	return &hash
}

func (request *UserMintNftRequest) CalculateSize() uint64 {
	return metadataCommon.CalculateSize(request)
}

func (request *UserMintNftRequest) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(struct {
		OtaReceive string `json:"OtaReceive"`
		Amount     uint64 `json:"Amount"`
		metadataCommon.MetadataBase
	}{
		Amount:       request.amount,
		OtaReceive:   request.otaReceive,
		MetadataBase: request.MetadataBase,
	})
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (request *UserMintNftRequest) UnmarshalJSON(data []byte) error {
	temp := struct {
		OtaReceive string `json:"OtaReceive"`
		Amount     uint64 `json:"Amount"`
		metadataCommon.MetadataBase
	}{}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	request.amount = temp.Amount
	request.otaReceive = temp.OtaReceive
	request.MetadataBase = temp.MetadataBase
	return nil
}

func (request *UserMintNftRequest) OtaReceive() string {
	return request.otaReceive
}

func (request *UserMintNftRequest) Amount() uint64 {
	return request.amount
}

func (request *UserMintNftRequest) GetOTADeclarations() []metadataCommon.OTADeclaration {
	var result []metadataCommon.OTADeclaration
	currentTokenID := common.ConfidentialAssetID
	otaReceive := privacy.OTAReceiver{}
	otaReceive.FromString(request.otaReceive)
	result = append(result, metadataCommon.OTADeclaration{
		PublicKey: otaReceive.PublicKey.ToBytes(), TokenID: currentTokenID,
	})
	return result
}
