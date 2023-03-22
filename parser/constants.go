package parser

const (
	// Fees

	TotalFeeOp           = "Fee"
	OverEstimationBurnOp = "OverEstimationBurn"
	MinerFeeOp           = "MinerFee"
	BurnFeeOp            = "BurnFee"

	BurnAddress = "f099"
	ethPrefix   = "0x"
	filPrefix   = "f0"

	// metadata keys
	ParamsKey  = "Params"
	ReturnKey  = "Return"
	ethHashKey = "ethHash"
	addressKey = "address"
	ethLogsKey = "ethLogs"

	UnknownStr = "unknown"

	// Methods
	MethodSend                        = "Send"                        // Common
	MethodFee                         = "Fee"                         // Common
	MethodConstructor                 = "Constructor"                 // Common
	MethodCronTick                    = "CronTick"                    // Common
	MethodEpochTick                   = "EpochTick"                   // Cron
	MethodExec                        = "Exec"                        // MethodsInit
	MethodSwapSigner                  = "SwapSigner"                  // MethodsMultisig
	MethodAddSigner                   = "AddSigner"                   // MethodsMultisig
	MethodRemoveSigner                = "RemoveSigner"                // MethodsMultisig
	MethodPropose                     = "Propose"                     // MethodsMultisig
	MethodApprove                     = "Approve"                     // MethodsMultisig
	MethodCancel                      = "Cancel"                      // MethodsMultisig
	MethodChangeNumApprovalsThreshold = "ChangeNumApprovalsThreshold" // MethodsMultisig
	MethodLockBalance                 = "LockBalance"                 // MethodsMultisig
	MethodAddVerifies                 = "AddVerifies"                 // MethodsMultisig
	MethodAwardBlockReward            = "AwardBlockReward"            // MethodsReward
	MethodUpdateNetworkKPI            = "UpdateNetworkKPI"            // MethodsReward
	MethodThisEpochReward             = "ThisEpochReward"             // MethodsReward
	MethodCreateMiner                 = "CreateMiner"                 // MethodsPower
	MethodUpdateClaimedPower          = "UpdateClaimedPower"          // MethodsPower
	MethodEnrollCronEvent             = "EnrollCronEvent"             // MethodsPower
	MethodSubmitPoRepForBulkVerify    = "SubmitPoRepForBulkVerify"    // MethodsPower
	MethodCurrentTotalPower           = "CurrentTotalPower"           // MethodsPower
	MethodUpdatePledgeTotal           = "UpdatePledgeTotal"           // MethodsPower
	MethodPowerDeprecated1            = "Deprecated1"                 // MethodsPower - OnConsensusFault
	MethodOnDeferredCronEvent         = "OnDeferredCronEvent"         // MethodsMiner
	MethodPreCommitSector             = "PreCommitSector"             // MethodsMiner
	MethodProveCommitSector           = "ProveCommitSector"           // MethodsMiner
	MethodSubmitWindowedPoSt          = "SubmitWindowedPoSt"          // MethodsMiner
	MethodApplyRewards                = "ApplyRewards"                // MethodsMiner
	MethodWithdrawBalance             = "WithdrawBalance"             // MethodsMiner
	MethodChangeOwnerAddress          = "ChangeOwnerAddress"          // MethodsMiner
	MethodChangeWorkerAddress         = "ChangeWorkerAddress"         // MethodsMiner
	MethodConfirmUpdateWorkerKey      = "ConfirmUpdateWorkerKey"      // MethodsMiner
	MethodDeclareFaultsRecovered      = "DeclareFaultsRecovered"      // MethodsMiner
	MethodPreCommitSectorBatch        = "PreCommitSectorBatch"        // MethodsMiner
	MethodProveCommitAggregate        = "ProveCommitAggregate"        // MethodsMiner
	MethodProveReplicaUpdates         = "ProveReplicaUpdates"         // MethodsMiner
	MethodChangeMultiaddrs            = "ChangeMultiaddrs"            // MethodsMiner
	MethodChangePeerID                = "ChangePeerID"                // MethodsMiner
	MethodExtendSectorExpiration      = "ExtendSectorExpiration"      // MethodsMiner
	MethodControlAddresses            = "ControlAddresses"            // MethodsMiner
	MethodTerminateSectors            = "TerminateSectors"            // MethodsMiner
	MethodDeclareFaults               = "DeclareFaults"               // MethodsMiner
	MethodCheckSectorProven           = "CheckSectorProven"           // MethodsMiner
	MethodReportConsensusFault        = "ReportConsensusFault"        // MethodsMiner
	MethodConfirmSectorProofsValid    = "ConfirmSectorProofsValid"    // MethodsMiner
	MethodCompactPartitions           = "CompactPartitions"           // MethodsMiner
	MethodCompactSectorNumbers        = "CompactSectorNumbers"        // MethodsMiner
	MethodRepayDebt                   = "RepayDebt"                   // MethodsMiner
	MethodDisputeWindowedPoSt         = "DisputeWindowedPoSt"         // MethodsMiner
	MethodChangeBeneficiary           = "ChangeBeneficiary"           // MethodsMiner
	MethodGetBeneficiary              = "GetBeneficiary"              // MethodsMiner
	MethodPublishStorageDeals         = "PublishStorageDeals"         // MethodsMarket
	MethodAddBalance                  = "AddBalance"                  // MethodsMarket
	MethodVerifyDealsForActivation    = "VerifyDealsForActivation"    // MethodsMarket
	MethodActivateDeals               = "ActivateDeals"               // MethodsMarket
	MethodOnMinerSectorsTerminate     = "OnMinerSectorsTerminate"     // MethodsMarket
	MethodComputeDataCommitment       = "ComputeDataCommitment"       // MethodsMarket
	MethodUpdateChannelState          = "UpdateChannelState"          // MethodsPaymentChannel
	MethodSettle                      = "Settle"                      // MethodsPaymentChannel
	MethodCollect                     = "Collect"                     // MethodsPaymentChannel
	MethodAddVerifiedClient           = "AddVerifiedClient"           // MethodsVerifiedRegistry
	MethodAddVerifier                 = "AddVerifier"                 // MethodsVerifiedRegistry
	MethodRemoveVerifier              = "RemoveVerifier"              // MethodsVerifiedRegistry
	MethodUseBytes                    = "UseBytes"                    // MethodsVerifiedRegistry
	MethodRestoreBytes                = "RestoreBytes"                // MethodsVerifiedRegistry
	MethodRemoveExpiredAllocations    = "RemoveExpiredAllocations"    // MethodsVerifiedRegistry
	MethodRemoveVerifiedClientDataCap = "RemoveVerifiedClientDataCap" // MethodsVerifiedRegistry
	MethodVerifiedDeprecated1         = "Deprecated1"                 // MethodsVerifiedRegistry - UseBytes
	MethodVerifiedDeprecated2         = "Deprecated2"                 // MethodsVerifiedRegistry - RestoreBytes
	MethodInvokeContract              = "InvokeContract"              // MethodsEVM
	MethodGetBytecode                 = "GetBytecode"                 // MethodsEVM
	MethodGetStorageAt                = "GetStorageAt"                // MethodsEVM
	MethodInvokeContractReadOnly      = "InvokeContractReadOnly"      // MethodsEVM
	MethodInvokeContractDelegate      = "InvokeContractDelegate"      // MethodsEVM
	MethodCreate                      = "Create"                      // MethodsEam
	MethodCreate2                     = "Create2"                     // MethodsEam
	MethodCreateExternal              = "CreateExternal"              // MethodsEam
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
