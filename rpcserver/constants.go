package rpcserver

// rpc cmd method
const (
	// test rpc server
	testHttpServer = "testrpcserver"
	startProfiling = "startprofiling"
	stopProfiling  = "stopprofiling"
	exportMetrics  = "exportmetrics"

	getNetworkInfo       = "getnetworkinfo"
	getConnectionCount   = "getconnectioncount"
	getAllConnectedPeers = "getallconnectedpeers"
	getAllPeers          = "getallpeers"
	getNodeRole          = "getnoderole"
	getInOutMessages     = "getinoutmessages"
	getInOutMessageCount = "getinoutmessagecount"

	estimateFee              = "estimatefee"
	estimateFeeV2            = "estimatefeev2"
	estimateFeeWithEstimator = "estimatefeewithestimator"

	getActiveShards    = "getactiveshards"
	getMaxShardsNumber = "getmaxshardsnumber"

	getMiningInfo                 = "getmininginfo"
	getRawMempool                 = "getrawmempool"
	getNumberOfTxsInMempool       = "getnumberoftxsinmempool"
	getMempoolEntry               = "getmempoolentry"
	removeTxInMempool             = "removetxinmempool"
	getBeaconPoolState            = "getbeaconpoolstate"
	getShardPoolState             = "getshardpoolstate"
	getShardPoolLatestValidHeight = "getshardpoollatestvalidheight"
	//getShardToBeaconPoolState     = "getshardtobeaconpoolstate"
	//getCrossShardPoolState        = "getcrossshardpoolstate"
	getNextCrossShard           = "getnextcrossshard"
	getShardToBeaconPoolStateV2 = "getshardtobeaconpoolstatev2"
	getCrossShardPoolStateV2    = "getcrossshardpoolstatev2"
	getShardPoolStateV2         = "getshardpoolstatev2"
	getBeaconPoolStateV2        = "getbeaconpoolstatev2"
	//getFeeEstimator             = "getfeeestimator"
	setBackup                   = "setbackup"
	getLatestBackup             = "getlatestbackup"
	getBestBlock                = "getbestblock"
	getBestBlockHash            = "getbestblockhash"
	getBlocks                   = "getblocks"
	retrieveBlock               = "retrieveblock"
	retrieveBlockByHeight       = "retrieveblockbyheight"
	retrieveBeaconBlock         = "retrievebeaconblock"
	retrieveBeaconBlockByHeight = "retrievebeaconblockbyheight"
	getBlockChainInfo           = "getblockchaininfo"
	getBlockCount               = "getblockcount"
	getBlockHash                = "getblockhash"

	listOutputCoins                              = "listoutputcoins"
	createRawTransaction                         = "createtransaction"
	sendRawTransaction                           = "sendtransaction"
	createAndSendTransaction                     = "createandsendtransaction"
	createAndSendTransactionV2                   = "createandsendtransactionv2"
	createAndSendCustomTokenTransaction          = "createandsendcustomtokentransaction"
	sendRawCustomTokenTransaction                = "sendrawcustomtokentransaction"
	createRawCustomTokenTransaction              = "createrawcustomtokentransaction"
	createRawPrivacyCustomTokenTransaction       = "createrawprivacycustomtokentransaction"
	sendRawPrivacyCustomTokenTransaction         = "sendrawprivacycustomtokentransaction"
	createAndSendPrivacyCustomTokenTransaction   = "createandsendprivacycustomtokentransaction"
	createAndSendPrivacyCustomTokenTransactionV2 = "createandsendprivacycustomtokentransactionv2"
	getMempoolInfo                               = "getmempoolinfo"
	getPendingTxsInBlockgen                      = "getpendingtxsinblockgen"
	getCandidateList                             = "getcandidatelist"
	getCommitteeList                             = "getcommitteelist"
	canPubkeyStake                               = "canpubkeystake"
	getTotalTransaction                          = "gettotaltransaction"
	listUnspentCustomToken                       = "listunspentcustomtoken"
	getBalanceCustomToken                        = "getbalancecustomtoken"
	getTransactionByHash                         = "gettransactionbyhash"
	gettransactionhashbyreceiver                 = "gettransactionhashbyreceiver"
	gettransactionhashbyreceiverv2               = "gettransactionhashbyreceiverv2"
	gettransactionbyreceiver                     = "gettransactionbyreceiver"
	gettransactionbyreceiverv2                   = "gettransactionbyreceiverv2"
	listCustomToken                              = "listcustomtoken"
	listPrivacyCustomToken                       = "listprivacycustomtoken"
	getPrivacyCustomToken                        = "getprivacycustomtoken"
	listPrivacyCustomTokenByShard                = "listprivacycustomtokenbyshard"
	getBalancePrivacyCustomToken                 = "getbalanceprivacycustomtoken"
	customTokenTxs                               = "customtoken"
	listCustomTokenHolders                       = "customtokenholder"
	privacyCustomTokenTxs                        = "privacycustomtoken"
	checkHashValue                               = "checkhashvalue"
	getListCustomTokenBalance                    = "getlistcustomtokenbalance"
	getListPrivacyCustomTokenBalance             = "getlistprivacycustomtokenbalance"
	getBlockHeader                               = "getheader"
	getCrossShardBlock                           = "getcrossshardblock"
	randomCommitments                            = "randomcommitments"
	hasSerialNumbers                             = "hasserialnumbers"
	hasSnDerivators                              = "hassnderivators"
	listSnDerivators                             = "listsnderivators"
	listSerialNumbers                            = "listserialnumbers"
	listCommitments                              = "listcommitments"
	listCommitmentIndices                        = "listcommitmentindices"
	createAndSendStakingTransaction              = "createandsendstakingtransaction"
	createAndSendStakingTransactionV2            = "createandsendstakingtransactionv2"
	createAndSendStopAutoStakingTransaction      = "createandsendstopautostakingtransaction"
	createAndSendStopAutoStakingTransactionV2    = "createandsendstopautostakingtransactionv2"
	decryptoutputcoinbykeyoftransaction          = "decryptoutputcoinbykeyoftransaction"

	//===========For Testing and Benchmark==============
	getAndSendTxsFromFile   = "getandsendtxsfromfile"
	getAndSendTxsFromFileV2 = "getandsendtxsfromfilev2"
	unlockMempool           = "unlockmempool"
	getAutoStakingByHeight  = "getautostakingbyheight"
	getCommitteeState       = "getcommitteestate"
	getRewardAmountByEpoch  = "getrewardamountbyepoch"
	//==================================================

	getShardBestState        = "getshardbeststate"
	getShardBestStateDetail  = "getshardbeststatedetail"
	getBeaconBestState       = "getbeaconbeststate"
	getBeaconBestStateDetail = "getbeaconbeststatedetail"

	// Wallet rpc cmd
	listAccounts               = "listaccounts"
	getAccount                 = "getaccount"
	getAddressesByAccount      = "getaddressesbyaccount"
	getAccountAddress          = "getaccountaddress"
	dumpPrivkey                = "dumpprivkey"
	importAccount              = "importaccount"
	removeAccount              = "removeaccount"
	listUnspentOutputCoins     = "listunspentoutputcoins"
	getBalance                 = "getbalance"
	getBalanceByPrivatekey     = "getbalancebyprivatekey"
	getBalanceByPaymentAddress = "getbalancebypaymentaddress"
	getReceivedByAccount       = "getreceivedbyaccount"
	setTxFee                   = "settxfee"

	// walletsta
	getPublicKeyFromPaymentAddress = "getpublickeyfrompaymentaddress"
	defragmentAccount              = "defragmentaccount"
	defragmentAccountV2            = "defragmentaccountv2"
	defragmentAccountToken         = "defragmentaccounttoken"
	defragmentAccountTokenV2       = "defragmentaccounttokenv2"

	getStackingAmount = "getstackingamount"

	// utils
	hashToIdenticon = "hashtoidenticon"
	generateTokenID = "generatetokenid"

	createIssuingRequest               = "createissuingrequest"
	sendIssuingRequest                 = "sendissuingrequest"
	createAndSendIssuingRequest        = "createandsendissuingrequest"
	createAndSendIssuingRequestV2      = "createandsendissuingrequestv2"
	createAndSendContractingRequest    = "createandsendcontractingrequest"
	createAndSendContractingRequestV2  = "createandsendcontractingrequestv2"
	createAndSendBurningRequest        = "createandsendburningrequest"
	createAndSendBurningRequestV2      = "createandsendburningrequestv2"
	createAndSendTxWithIssuingETHReq   = "createandsendtxwithissuingethreq"
	createAndSendTxWithIssuingETHReqV2 = "createandsendtxwithissuingethreqv2"
	checkETHHashIssued                 = "checkethhashissued"
	getAllBridgeTokens                 = "getallbridgetokens"
	getETHHeaderByHash                 = "getethheaderbyhash"
	getBridgeReqWithStatus             = "getbridgereqwithstatus"

	// Incognito -> Ethereum bridge
	getBeaconSwapProof       = "getbeaconswapproof"
	getLatestBeaconSwapProof = "getlatestbeaconswapproof"
	getBridgeSwapProof       = "getbridgeswapproof"
	getLatestBridgeSwapProof = "getlatestbridgeswapproof"
	getBurnProof             = "getburnproof"

	// reward
	CreateRawWithDrawTransaction = "withdrawreward"
	getRewardAmount              = "getrewardamount"
	getRewardAmountByPublicKey   = "getrewardamountbypublickey"
	listRewardAmount             = "listrewardamount"

	revertbeaconchain = "revertbeaconchain"
	revertshardchain  = "revertshardchain"

	enableMining                = "enablemining"
	getChainMiningStatus        = "getchainminingstatus"
	getPublickeyMining          = "getpublickeymining"
	getPublicKeyRole            = "getpublickeyrole"
	getRoleByValidatorKey       = "getrolebyvalidatorkey"
	getIncognitoPublicKeyRole   = "getincognitopublickeyrole"
	getMinerRewardFromMiningKey = "getminerrewardfromminingkey"

	// slash
	getProducersBlackList       = "getproducersblacklist"
	getProducersBlackListDetail = "getproducersblacklistdetail"

	// pde
	getPDEState                                = "getpdestate"
	createAndSendTxWithWithdrawalReq           = "createandsendtxwithwithdrawalreq"
	createAndSendTxWithWithdrawalReqV2         = "createandsendtxwithwithdrawalreqv2"
	createAndSendTxWithPDEFeeWithdrawalReq     = "createandsendtxwithpdefeewithdrawalreq"
	createAndSendTxWithPTokenTradeReq          = "createandsendtxwithptokentradereq"
	createAndSendTxWithPTokenCrossPoolTradeReq = "createandsendtxwithptokencrosspooltradereq"
	createAndSendTxWithPRVTradeReq             = "createandsendtxwithprvtradereq"
	createAndSendTxWithPRVCrossPoolTradeReq    = "createandsendtxwithprvcrosspooltradereq"
	createAndSendTxWithPTokenContribution      = "createandsendtxwithptokencontribution"
	createAndSendTxWithPRVContribution         = "createandsendtxwithprvcontribution"
	createAndSendTxWithPTokenContributionV2    = "createandsendtxwithptokencontributionv2"
	createAndSendTxWithPRVContributionV2       = "createandsendtxwithprvcontributionv2"
	convertNativeTokenToPrivacyToken           = "convertnativetokentoprivacytoken"
	convertPrivacyTokenToNativeToken           = "convertprivacytokentonativetoken"
	getPDEContributionStatus                   = "getpdecontributionstatus"
	getPDEContributionStatusV2                 = "getpdecontributionstatusv2"
	getPDETradeStatus                          = "getpdetradestatus"
	getPDEWithdrawalStatus                     = "getpdewithdrawalstatus"
	getPDEFeeWithdrawalStatus                  = "getpdefeewithdrawalstatus"
	convertPDEPrices                           = "convertpdeprices"
	extractPDEInstsFromBeaconBlock             = "extractpdeinstsfrombeaconblock"

	// get burning address
	getBurningAddress = "getburningaddress"

	// portal
	createAndSendTxWithCustodianDeposit           = "createandsendtxwithcustodiandeposit"
	createAndSendTxWithReqPToken                  = "createandsendtxwithreqptoken"
	getPortalState                                = "getportalstate"
	getPortalCustodianDepositStatus               = "getportalcustodiandepositstatus"
	createAndSendRegisterPortingPublicTokens      = "createandsendregisterportingpublictokens"
	createAndSendPortalExchangeRates              = "createandsendportalexchangerates"
	getPortalFinalExchangeRates                   = "getportalfinalexchangerates"
	getPortalPortingRequestByKey                  = "getportalportingrequestbykey"
	getPortalPortingRequestByPortingId            = "getportalportingrequestbyportingid"
	convertExchangeRates                          = "convertexchangerates"
	getPortalReqPTokenStatus                      = "getportalreqptokenstatus"
	getPortingRequestFees                         = "getportingrequestfees"
	createAndSendTxWithRedeemReq                  = "createandsendtxwithredeemreq"
	createAndSendTxWithReqUnlockCollateral        = "createandsendtxwithrequnlockcollateral"
	getPortalReqUnlockCollateralStatus            = "getportalrequnlockcollateralstatus"
	getPortalReqRedeemStatus                      = "getportalreqredeemstatus"
	createAndSendCustodianWithdrawRequest         = "createandsendcustodianwithdrawrequest"
	getCustodianWithdrawByTxId                    = "getcustodianwithdrawbytxid"
	getCustodianLiquidationStatus                 = "getcustodianliquidationstatus"
	createAndSendTxWithReqWithdrawRewardPortal    = "createandsendtxwithreqwithdrawrewardportal"
	createAndSendRedeemLiquidationExchangeRates   = "createandsendredeemliquidationexchangerates"
	createAndSendLiquidationCustodianDeposit      = "createandsendliquidationcustodiandeposit"
	createAndSendTopUpWaitingPorting              = "createandsendtopupwaitingporting"
	getAmountNeededForCustodianDepositLiquidation = "getamountneededforcustodiandepositliquidation"
	getLiquidationExchangeRatesPool               = "getliquidationtpexchangeratespool"
	getPortalReward                               = "getportalreward"
	getRequestWithdrawPortalRewardStatus          = "getrequestwithdrawportalrewardstatus"
	createAndSendTxWithReqMatchingRedeem          = "createandsendtxwithreqmatchingredeem"
	getReqMatchingRedeemStatus                    = "getreqmatchingredeemstatus"
	getPortalCustodianTopupStatus                 = "getcustodiantopupstatus"
	getPortalCustodianTopupWaitingPortingStatus   = "getcustodiantopupwaitingportingstatus"
	getAmountTopUpWaitingPorting                  = "getamounttopupwaitingporting"
	getPortalReqRedeemByTxIDStatus                = "getreqredeemstatusbytxid"
	getReqRedeemFromLiquidationPoolByTxIDStatus   = "getreqredeemfromliquidationpoolbytxidstatus"

	// relaying
	createAndSendTxWithRelayingBNBHeader = "createandsendtxwithrelayingbnbheader"
	createAndSendTxWithRelayingBTCHeader = "createandsendtxwithrelayingbtcheader"
	getRelayingBNBHeaderState            = "getrelayingbnbheaderstate"
	getRelayingBNBHeaderByBlockHeight    = "getrelayingbnbheaderbyblockheight"
	getBTCRelayingBestState              = "getbtcrelayingbeststate"
	getBTCBlockByHash                    = "getbtcblockbyhash"
	getLatestBNBHeaderBlockHeight        = "getlatestbnbheaderblockheight"

	// incognito mode for sc
	getBurnProofForDepositToSC                  = "getburnprooffordeposittosc"
	createAndSendBurningForDepositToSCRequest   = "createandsendburningfordeposittoscrequest"
	createAndSendBurningForDepositToSCRequestV2 = "createandsendburningfordeposittoscrequestv2"

	getBeaconPoolInfo     = "getbeaconpoolinfo"
	getShardPoolInfo      = "getshardpoolinfo"
	getCrossShardPoolInfo = "getcrossshardpoolinfo"
	getAllView            = "getallview"
	getAllViewDetail      = "getallviewdetail"

	// feature rewards
	getRewardFeature = "getrewardfeature"

	getTotalStaker = "gettotalstaker"

	//validator state
	getValKeyState = "getvalkeystate"

	// init privacy custom token
	createAndSendPTokenInitRequest = "createandsendPTokenInitrequest"
)

const (
	testSubcrice                                = "testsubcribe"
	subcribeNewShardBlock                       = "subcribenewshardblock"
	subcribeNewBeaconBlock                      = "subcribenewbeaconblock"
	subcribePendingTransaction                  = "subcribependingtransaction"
	subcribeShardCandidateByPublickey           = "subcribeshardcandidatebypublickey"
	subcribeShardPendingValidatorByPublickey    = "subcribeshardpendingvalidatorbypublickey"
	subcribeShardCommitteeByPublickey           = "subcribeshardcommitteebypublickey"
	subcribeBeaconCandidateByPublickey          = "subcribebeaconcandidatebypublickey"
	subcribeBeaconPendingValidatorByPublickey   = "subcribebeaconpendingvalidatorbypublickey"
	subcribeBeaconCommitteeByPublickey          = "subcribebeaconcommitteebypublickey"
	subcribeCrossOutputCoinByPrivateKey         = "subcribecrossoutputcoinbyprivatekey"
	subcribeCrossCustomTokenByPrivateKey        = "subcribecrosscustomtokenbyprivatekey"
	subcribeCrossCustomTokenPrivacyByPrivateKey = "subcribecrosscustomtokenprivacybyprivatekey"
	subcribeMempoolInfo                         = "subcribemempoolinfo"
	subcribeShardBestState                      = "subcribeshardbeststate"
	subcribeBeaconBestState                     = "subcribebeaconbeststate"
	subcribeBeaconPoolBeststate                 = "subcribebeaconpoolbeststate"
	subcribeShardPoolBeststate                  = "subcribeshardpoolbeststate"
)
