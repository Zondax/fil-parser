package parser

const (
	// Backoff
	BackOffStrategyLinear      = "linear"
	BackOffStrategyExponential = "exponential"

	// Fees

	TotalFeeOp           = "Fee"
	OverEstimationBurnOp = "OverEstimationBurn"
	MinerFeeOp           = "MinerFee"
	BurnFeeOp            = "BurnFee"

	BurnAddress = "f099"
	EthPrefix   = "0x"
	FilPrefix   = "f0"

	// metadata keys
	ValueKey     = "Value"
	ParamsKey    = "Params"
	ReturnKey    = "Return"
	ParamsRawKey = "ParamsRaw"
	ReturnRawKey = "ReturnRaw"
	ErrorKey     = "Error"
	MethodNumKey = "MethodNum"
	EthHashKey   = "ethHash"
	AddressKey   = "address"
	EthLogsKey   = "ethLogs"

	UnknownStr = "unknown"

	TxTypeGenesis = "Genesis"
	GenesisHeight = 0
	TxFromGenesis = "genesis"

	// FirstExportedMethodNumber is the lowest FRC-42 method number.
	// https://github.com/filecoin-project/builtin-actors/blob/8fdbdec5e3f46b60ba0132d90533783a44c5961f/runtime/src/builtin/shared.rs#L58
	FirstExportedMethodNumber = 1 << 24

	EvmMaxReservedMethodNumber = 1023

	MultisigConstructorMethod = "Constructor"

	// Methods
	MethodSend                                = "Send"                                // Common
	MethodFee                                 = "Fee"                                 // Common
	MethodConstructor                         = "Constructor"                         // Common
	MethodCronTick                            = "CronTick"                            // Common
	MethodEpochTick                           = "EpochTick"                           // Cron
	MethodPubkeyAddress                       = "PubkeyAddress"                       // MethodsAccount
	MethodAuthenticateMessage                 = "AuthenticateMessage"                 // MethodsAccount
	MethodReceive                             = "Receive"                             // MethodsAccount // exists only in built-in actors v9
	MethodFallback                            = "Fallback"                            // MethodsAccount
	MethodExec                                = "Exec"                                // MethodsInit
	MethodExec4                               = "Exec4"                               // MethodsInit
	MethodSwapSigner                          = "SwapSigner"                          // MethodsMultisig
	MethodSwapSignerExported                  = "SwapSignerExported"                  // MethodsMultisig
	MethodAddSigner                           = "AddSigner"                           // MethodsMultisig
	MethodAddSignerExported                   = "AddSignerExported"                   // MethodsMultisig
	MethodRemoveSigner                        = "RemoveSigner"                        // MethodsMultisig
	MethodRemoveSignerExported                = "RemoveSignerExported"                // MethodsMultisig
	MethodPropose                             = "Propose"                             // MethodsMultisig
	MethodProposeExported                     = "ProposeExported"                     // MethodsMultisig
	MethodApprove                             = "Approve"                             // MethodsMultisig
	MethodApproveExported                     = "ApproveExported"                     // MethodsMultisig
	MethodCancel                              = "Cancel"                              // MethodsMultisig
	MethodCancelExported                      = "CancelExported"                      // MethodsMultisig
	MethodChangeNumApprovalsThreshold         = "ChangeNumApprovalsThreshold"         // MethodsMultisig
	MethodChangeNumApprovalsThresholdExported = "ChangeNumApprovalsThresholdExported" // MethodsMultisig
	MethodLockBalance                         = "LockBalance"                         // MethodsMultisig
	MethodLockBalanceExported                 = "LockBalanceExported"                 // MethodsMultisig
	MethodAddVerifies                         = "AddVerifies"                         // MethodsMultisig
	MethodMsigUniversalReceiverHook           = "UniversalReceiverHook"               // MethodsMultisig
	MethodAwardBlockReward                    = "AwardBlockReward"                    // MethodsReward
	MethodUpdateNetworkKPI                    = "UpdateNetworkKPI"                    // MethodsReward
	MethodThisEpochReward                     = "ThisEpochReward"                     // MethodsReward
	MethodCreateMiner                         = "CreateMiner"                         // MethodsPower
	MethodCreateMinerExported                 = "CreateMinerExported"                 // MethodsPower
	MethodUpdateClaimedPower                  = "UpdateClaimedPower"                  // MethodsPower
	MethodEnrollCronEvent                     = "EnrollCronEvent"                     // MethodsPower
	MethodSubmitPoRepForBulkVerify            = "SubmitPoRepForBulkVerify"            // MethodsPower
	MethodCurrentTotalPower                   = "CurrentTotalPower"                   // MethodsPower
	MethodUpdatePledgeTotal                   = "UpdatePledgeTotal"                   // MethodsPower
	MethodPowerDeprecated1                    = "Deprecated1"                         // MethodsPower - OnConsensusFault
	MethodNetworkRawPowerExported             = "NetworkRawPowerExported"             // MethodsPower
	MethodMinerRawPowerExported               = "MinerRawPowerExported"               // MethodsPower
	MethodMinerCountExported                  = "MinerCountExported"                  // MethodsPower
	MethodMinerConsensusCountExported         = "MinerConsensusCountExported"         // MethodsPower
	MethodOnEpochTickEnd                      = "OnEpochTickEnd"                      // MethodsPower
	MethodOnConsensusFault                    = "OnConsensusFault"                    // MethodsPower
	MethodOnDeferredCronEvent                 = "OnDeferredCronEvent"                 // MethodsMiner
	MethodPreCommitSector                     = "PreCommitSector"                     // MethodsMiner
	MethodProveCommitSector                   = "ProveCommitSector"                   // MethodsMiner - Deprecated
	MethodSubmitWindowedPoSt                  = "SubmitWindowedPoSt"                  // MethodsMiner
	MethodApplyRewards                        = "ApplyRewards"                        // MethodsMiner
	MethodWithdrawBalance                     = "WithdrawBalance"                     // MethodsMiner
	MethodWithdrawBalanceExported             = "WithdrawBalanceExported"             // MethodsMiner
	MethodChangeOwnerAddress                  = "ChangeOwnerAddress"                  // MethodsMiner
	MethodChangeOwnerAddressExported          = "ChangeOwnerAddressExported"          // MethodsMiner
	MethodChangeWorkerAddress                 = "ChangeWorkerAddress"                 // MethodsMiner
	MethodChangeWorkerAddressExported         = "ChangeWorkerAddressExported"         // MethodsMiner
	MethodConfirmUpdateWorkerKey              = "ConfirmUpdateWorkerKey"              // MethodsMiner
	MethodDeclareFaultsRecovered              = "DeclareFaultsRecovered"              // MethodsMiner
	MethodPreCommitSectorBatch                = "PreCommitSectorBatch"                // MethodsMiner
	MethodProveCommitAggregate                = "ProveCommitAggregate"                // MethodsMiner
	MethodProveReplicaUpdates                 = "ProveReplicaUpdates"                 // MethodsMiner
	MethodChangeMultiaddrs                    = "ChangeMultiaddrs"                    // MethodsMiner
	MethodChangeMultiaddrsExported            = "ChangeMultiaddrsExported"            // MethodsMiner
	MethodChangePeerID                        = "ChangePeerID"                        // MethodsMiner
	MethodChangePeerIDExported                = "ChangePeerIDExported"                // MethodsMiner
	MethodExtendSectorExpiration              = "ExtendSectorExpiration"              // MethodsMiner
	MethodControlAddresses                    = "ControlAddresses"                    // MethodsMiner
	MethodTerminateSectors                    = "TerminateSectors"                    // MethodsMiner
	MethodDeclareFaults                       = "DeclareFaults"                       // MethodsMiner
	MethodCheckSectorProven                   = "CheckSectorProven"                   // MethodsMiner
	MethodReportConsensusFault                = "ReportConsensusFault"                // MethodsMiner
	MethodConfirmSectorProofsValid            = "ConfirmSectorProofsValid"            // MethodsMiner
	MethodCompactPartitions                   = "CompactPartitions"                   // MethodsMiner
	MethodCompactSectorNumbers                = "CompactSectorNumbers"                // MethodsMiner
	MethodRepayDebt                           = "RepayDebt"                           // MethodsMiner
	MethodRepayDebtExported                   = "RepayDebtExported"                   // MethodsMiner
	MethodDisputeWindowedPoSt                 = "DisputeWindowedPoSt"                 // MethodsMiner
	MethodChangeBeneficiary                   = "ChangeBeneficiary"                   // MethodsMiner
	MethodChangeBeneficiaryExported           = "ChangeBeneficiaryExported"           // MethodsMiner
	MethodGetBeneficiary                      = "GetBeneficiary"                      // MethodsMiner
	MethodIsControllingAddressExported        = "IsControllingAddressExported"        // MethodsMiner
	MethodConfirmChangeWorkerAddress          = "ConfirmChangeWorkerAddress"          // MethodsMiner
	MethodConfirmChangeWorkerAddressExported  = "ConfirmChangeWorkerAddressExported"  // MethodsMiner
	MethodPreCommitSectorBatch2               = "PreCommitSectorBatch2"               // MethodsMiner
	MethodProveReplicaUpdates2                = "ProveReplicaUpdates2"                // MethodsMiner
	MethodExtendSectorExpiration2             = "ExtendSectorExpiration2"             // MethodsMiner
	MethodGetOwner                            = "GetOwnerExported"                    // MethodsMiner
	MethodGetSectorSize                       = "GetSectorSizeExported"               // MethodsMiner
	MethodGetAvailableBalance                 = "GetAvailableBalanceExported"         // MethodsMiner
	MethodGetVestingFunds                     = "GetVestingFundsExported"             // MethodsMiner
	MethodGetPeerID                           = "GetPeerIDExported"                   // MethodsMiner
	MethodGetMultiaddrs                       = "GetMultiaddrsExported"               // MethodsMiner
	MethodAddLockedFund                       = "AddLockedFund"                       // MethodsMiner
	MethodProveCommitSectors3                 = "ProveCommitSectors3"                 // MethodsMiner
	MethodProveCommitSectorsNI                = "ProveCommitSectorsNI"                // MethodsMiner
	MethodProveReplicaUpdates3                = "ProveReplicaUpdates3"                // MethodsMiner
	MethodInternalSectorSetupForPreseal       = "InternalSectorSetupForPreseal"       // MethodsMiner
	MethodInitialPledge                       = "InitialPledge"                       // MethodsMiner
	MethodInitialPledgeExported               = "InitialPledgeExported"               // MethodsMiner
	MethodMaxTerminationFee                   = "MaxTerminationFee"                   // MethodsMiner
	MethodMovePartitions                      = "MovePartitions"                      // MethodsMiner
	MethodMaxTerminationFeeExported           = "MaxTerminationFeeExported"           // MethodsMiner
	MethodPublishStorageDeals                 = "PublishStorageDeals"                 // MethodsMarket
	MethodPublishStorageDealsExported         = "PublishStorageDealsExported"         // MethodsMarket
	MethodAddBalance                          = "AddBalance"                          // MethodsMarket
	MethodAddBalanceExported                  = "AddBalanceExported"                  // MethodsMarket
	MethodVerifyDealsForActivation            = "VerifyDealsForActivation"            // MethodsMarket
	MethodActivateDeals                       = "ActivateDeals"                       // MethodsMarket
	MethodOnMinerSectorsTerminate             = "OnMinerSectorsTerminate"             // MethodsMarket
	MethodComputeDataCommitment               = "ComputeDataCommitment"               // MethodsMarket
	MethodGetBalance                          = "GetBalanceExported"                  // MethodsMarket
	MethodGetDealDataCommitment               = "GetDealDataCommitmentExported"       // MethodsMarket
	MethodGetDealClient                       = "GetDealClientExported"               // MethodsMarket
	MethodGetDealProvider                     = "GetDealProviderExported"             // MethodsMarket
	MethodGetDealLabel                        = "GetDealLabelExported"                // MethodsMarket
	MethodGetDealTerm                         = "GetDealTermExported"                 // MethodsMarket
	MethodGetDealTotalPrice                   = "GetDealTotalPriceExported"           // MethodsMarket
	MethodGetDealClientCollateral             = "GetDealClientCollateralExported"     // MethodsMarket
	MethodGetDealProviderCollateral           = "GetDealProviderCollateralExported"   // MethodsMarket
	MethodGetDealVerified                     = "GetDealVerifiedExported"             // MethodsMarket
	MethodGetDealActivation                   = "GetDealActivationExported"           // MethodsMarket
	MethodGetDealSectorExported               = "GetDealSectorExported"               // MethodsMarket
	MethodSettleDealPaymentsExported          = "SettleDealPaymentsExported"          // MethodsMarket
	MethodSectorContentChanged                = "SectorContentChanged"                // MethodsMarket
	MethodUpdateChannelState                  = "UpdateChannelState"                  // MethodsPaymentChannel
	MethodSettle                              = "Settle"                              // MethodsPaymentChannel
	MethodCollect                             = "Collect"                             // MethodsPaymentChannel
	MethodAddVerifiedClient                   = "AddVerifiedClient"                   // MethodsVerifiedRegistry
	MethodAddVerifiedClientExported           = "AddVerifiedClientExported"           // MethodsVerifiedRegistry
	MethodAddVerifier                         = "AddVerifier"                         // MethodsVerifiedRegistry
	MethodRemoveVerifier                      = "RemoveVerifier"                      // MethodsVerifiedRegistry
	MethodUseBytes                            = "UseBytes"                            // MethodsVerifiedRegistry
	MethodRestoreBytes                        = "RestoreBytes"                        // MethodsVerifiedRegistry
	MethodRemoveExpiredAllocations            = "RemoveExpiredAllocations"            // MethodsVerifiedRegistry
	MethodRemoveExpiredAllocationsExported    = "RemoveExpiredAllocationsExported"    // MethodsVerifiedRegistry
	MethodRemoveVerifiedClientDataCap         = "RemoveVerifiedClientDataCap"         // MethodsVerifiedRegistry
	MethodVerifiedDeprecated1                 = "Deprecated1"                         // MethodsVerifiedRegistry - UseBytes
	MethodVerifiedDeprecated2                 = "Deprecated2"                         // MethodsVerifiedRegistry - RestoreBytes
	MethodGetClaims                           = "GetClaims"                           // MethodsVerifiedRegistry
	MethodGetClaimsExported                   = "GetClaimsExported"                   // MethodsVerifiedRegistry
	MethodExtendClaimTerms                    = "ExtendClaimTerms"                    // MethodsVerifiedRegistry
	MethodExtendClaimTermsExported            = "ExtendClaimTermsExported"            // MethodsVerifiedRegistry
	MethodRemoveExpiredClaims                 = "RemoveExpiredClaims"                 // MethodsVerifiedRegistry
	MethodRemoveExpiredClaimsExported         = "RemoveExpiredClaimsExported"         // MethodsVerifiedRegistry
	MethodUniversalReceiverHook               = "UniversalReceiverHook"               // MethodsVerifiedRegistry
	MethodClaimAllocations                    = "ClaimAllocations"                    // MethodsVerifiedRegistry
	MethodInvokeContract                      = "InvokeContract"                      // MethodsEVM
	MethodGetBytecode                         = "GetBytecode"                         // MethodsEVM
	MethodGetStorageAt                        = "GetStorageAt"                        // MethodsEVM
	MethodHandleFilecoinMethod                = "HandleFilecoinMethod"                // MethodsEVM
	MethodResurrect                           = "Resurrect"                           // MethodsEVM
	MethodGetBytecodeHash                     = "GetBytecodeHash"                     // MethodsEVM
	MethodInvokeContractReadOnly              = "InvokeContractReadOnly"              // MethodsEVM
	MethodInvokeContractDelegate              = "InvokeContractDelegate"              // MethodsEVM
	MethodInvokeContractFilecoinHandler       = "InvokeContractFilecoinHandler"       // MethodsEVM
	MethodValueTransfer                       = "ValueTransfer"                       // MethodsEVM
	MethodCreate                              = "Create"                              // MethodsEam
	MethodCreate2                             = "Create2"                             // MethodsEam
	MethodCreateExternal                      = "CreateExternal"                      // MethodsEam
	MethodMint                                = "Mint"                                // MethodsDatacap: v9
	MethodMintExported                        = "MintExported"                        // MethodsDatacap
	MethodDestroy                             = "Destroy"                             // MethodsDatacap: v9
	MethodDestroyExported                     = "DestroyExported"                     // MethodsDatacap
	MethodName                                = "Name"                                // MethodsDatacap: v9
	MethodNameExported                        = "NameExported"                        // MethodsDatacap
	MethodSymbol                              = "Symbol"                              // MethodsDatacap: v9
	MethodSymbolExported                      = "SymbolExported"                      // MethodsDatacap
	MethodTotalSupply                         = "TotalSupply"                         // MethodsDatacap: v9
	MethodTotalSupplyExported                 = "TotalSupplyExported"                 // MethodsDatacap
	MethodBalanceExported                     = "BalanceExported"                     // MethodsDatacap
	MethodTransfer                            = "Transfer"                            // MethodsDatacap: v9
	MethodTransferExported                    = "TransferExported"                    // MethodsDatacap
	MethodTransferFrom                        = "TransferFrom"                        // MethodsDatacap: v9
	MethodTransferFromExported                = "TransferFromExported"                // MethodsDatacap
	MethodIncreaseAllowance                   = "IncreaseAllowance"                   // MethodsDatacap: v9
	MethodIncreaseAllowanceExported           = "IncreaseAllowanceExported"           // MethodsDatacap
	MethodDecreaseAllowance                   = "DecreaseAllowance"                   // MethodsDatacap: v9
	MethodDecreaseAllowanceExported           = "DecreaseAllowanceExported"           // MethodsDatacap
	MethodRevokeAllowance                     = "RevokeAllowance"                     // MethodsDatacap: v9
	MethodRevokeAllowanceExported             = "RevokeAllowanceExported"             // MethodsDatacap
	MethodBurn                                = "Burn"                                // MethodsDatacap: v9
	MethodBurnExported                        = "BurnExported"                        // MethodsDatacap
	MethodBurnFrom                            = "BurnFrom"                            // MethodsDatacap: v9
	MethodBurnFromExported                    = "BurnFromExported"                    // MethodsDatacap
	MethodAllowance                           = "Allowance"                           // MethodsDatacap: v9
	MethodAllowanceExported                   = "AllowanceExported"                   // MethodsDatacap
	MethodGranularityExported                 = "GranularityExported"                 // MethodsDatacap
	MethodBalanceOf                           = "BalanceOf"                           // MethodsDatacap
	MethodUnknown                             = "Unknown"                             // Common
)

// SupportedOperations operations that will be parsed
var SupportedOperations = map[string]bool{
	MethodSend:                   true,
	MethodFee:                    true,
	MethodExec:                   true,
	MethodSwapSigner:             true,
	MethodAddSigner:              true,
	MethodRemoveSigner:           true,
	MethodPropose:                true,
	MethodApprove:                true,
	MethodCancel:                 true,
	MethodAwardBlockReward:       true,
	MethodOnDeferredCronEvent:    true,
	MethodPreCommitSector:        true,
	MethodProveCommitSector:      true,
	MethodProveCommitSectors3:    true,
	MethodSubmitWindowedPoSt:     true,
	MethodApplyRewards:           true,
	MethodWithdrawBalance:        true,
	MethodChangeOwnerAddress:     true,
	MethodChangeWorkerAddress:    true,
	MethodConfirmUpdateWorkerKey: true,
	MethodDeclareFaultsRecovered: true,
	MethodPreCommitSectorBatch:   true,
	MethodProveCommitAggregate:   true,
	MethodProveReplicaUpdates:    true,
	MethodCreateMiner:            true,
	MethodChangeMultiaddrs:       true,
	MethodChangePeerID:           true,
	MethodExtendSectorExpiration: true,
	MethodPublishStorageDeals:    true,
	MethodAddBalance:             true,
	MethodAddVerifiedClient:      true,
	MethodAddVerifier:            true,
	MethodRemoveVerifier:         true,
	MethodInvokeContract:         true,
	MethodGetBytecode:            true,
	MethodGetStorageAt:           true,
	MethodInvokeContractReadOnly: true,
	MethodInvokeContractDelegate: true,
}

func GetSupportedOps() []string {
	var result []string
	for k, v := range SupportedOperations {
		if v {
			result = append(result, k)
		}
	}
	return result
}

// These are actors that were not deployed in the genesis block, but created on following upgrades
// https://github.com/filecoin-project/community/discussions/74#discussioncomment-4313888
// EAM deploy tipset mainnet bafy2bzaceazaznpb47y6ljwmqrtkwcogpkuaslm2345w6khxdjq7r3cemurdw (height: 2683348)
// EAM deploy tipset calibration bafy2bzacebnt7vyhpmgl45egwgmm7svpxmzapmtn4kb6x5sdenxd4lufkkazu (height: 322354)

// https://github.com/filecoin-project/community/discussions/74#discussioncomment-3825422
// DataCap deploy tipset mainnet bafy2bzaceanlssigz6edrspkrzlkovhgr2ywqfbhp4p3cxxelf4oixhu44q32 (height: 2383680)
// DataCap deploy tipset calibration bafy2bzaceaezolcmjxugz7whrso62dwg3xyrybpsqkt6mrovtdnvjo622ed2g (height: 16800)

// actors deployed after genesis, for reference
var MainnetPostGenesisActors = [][]string{
	{"f010", "bafy2bzaceazaznpb47y6ljwmqrtkwcogpkuaslm2345w6khxdjq7r3cemurdw"}, // EAM mainnet
	{"f07", "bafy2bzaceanlssigz6edrspkrzlkovhgr2ywqfbhp4p3cxxelf4oixhu44q32"},  // DataCap mainnet
}

var CalibrationPostGenesisActors = [][]string{
	{"f010", "bafy2bzacebnt7vyhpmgl45egwgmm7svpxmzapmtn4kb6x5sdenxd4lufkkazu"}, // EAM calibration
	{"f07", "bafy2bzaceaezolcmjxugz7whrso62dwg3xyrybpsqkt6mrovtdnvjo622ed2g"},  // DataCap calibration
}
