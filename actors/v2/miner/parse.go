package miner

import (
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (m *Miner) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid, _ filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		resp := actors.ParseSend(msg)
		return resp, nil, nil
	case parser.MethodConstructor:
		resp, err := m.Constructor(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodControlAddresses:
		resp, err := m.ControlAddresses(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodChangeWorkerAddress, parser.MethodChangeWorkerAddressExported:
		resp, err := m.ChangeWorkerAddressExported(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodChangePeerID, parser.MethodChangePeerIDExported:
		resp, err := m.ChangePeerIDExported(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodSubmitWindowedPoSt:
		resp, err := m.SubmitWindowedPoSt(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodPreCommitSector:
		resp, err := m.PreCommitSector(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodProveCommitSector:
		resp, err := m.ProveCommitSector(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodProveCommitSectors3:
		resp, err := m.ProveCommitSectors3(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodExtendSectorExpiration:
		resp, err := m.ExtendSectorExpiration(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodTerminateSectors:
		resp, err := m.TerminateSectors(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodDeclareFaults:
		resp, err := m.DeclareFaults(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodDeclareFaultsRecovered:
		resp, err := m.DeclareFaultsRecovered(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodOnDeferredCronEvent:
		resp, err := m.OnDeferredCronEvent(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodCheckSectorProven:
		resp, err := m.CheckSectorProven(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodApplyRewards:
		resp, err := m.ApplyRewards(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodReportConsensusFault:
		resp, err := m.ReportConsensusFault(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodWithdrawBalance, parser.MethodWithdrawBalanceExported:
		resp, err := m.WithdrawBalanceExported(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodConfirmSectorProofsValid:
		resp, err := m.ConfirmSectorProofsValid(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodChangeMultiaddrs, parser.MethodChangeMultiaddrsExported:
		resp, err := m.ChangeMultiaddrsExported(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodCompactPartitions:
		resp, err := m.CompactPartitions(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodCompactSectorNumbers:
		resp, err := m.CompactSectorNumbers(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodConfirmChangeWorkerAddress, parser.MethodConfirmChangeWorkerAddressExported:
		resp, err := actors.ParseEmptyParamsAndReturn()
		return resp, nil, err
	case parser.MethodRepayDebt, parser.MethodRepayDebtExported:
		resp, err := actors.ParseEmptyParamsAndReturn()
		return resp, nil, err
	case parser.MethodChangeOwnerAddress, parser.MethodChangeOwnerAddressExported: // TODO: not tested
		resp, err := m.ChangeOwnerAddressExported(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodDisputeWindowedPoSt:
		resp, err := m.DisputeWindowedPoSt(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodPreCommitSectorBatch:
		resp, err := m.PreCommitSectorBatch(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodProveCommitAggregate:
		resp, err := m.ProveCommitAggregate(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodProveReplicaUpdates:
		resp, err := m.ProveReplicaUpdates(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodPreCommitSectorBatch2:
		resp, err := m.PreCommitSectorBatch2(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodProveReplicaUpdates2:
		resp, err := m.ProveReplicaUpdates2(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodChangeBeneficiary, parser.MethodChangeBeneficiaryExported:
		resp, err := m.ChangeBeneficiaryExported(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodGetBeneficiary:
		resp, err := m.GetBeneficiary(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodExtendSectorExpiration2:
		resp, err := m.ExtendSectorExpiration2(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodGetOwner:
		resp, err := m.GetOwnerExported(network, height, msgRct.Return)
		return resp, nil, err
	case parser.MethodIsControllingAddressExported:
		resp, err := m.IsControllingAddressExported(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetSectorSize:
		resp, err := m.GetSectorSize(network, height, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetAvailableBalance:
		resp, err := m.GetAvailableBalanceExported(network, height, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetVestingFunds:
		resp, err := m.GetVestingFundsExported(network, height, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetPeerID:
		resp, err := m.GetPeerIDExported(network, height, msgRct.Return)
		return resp, nil, err
	case parser.MethodGetMultiaddrs:
		resp, err := m.GetMultiaddrsExported(network, height, msgRct.Return)
		return resp, nil, err
	case parser.MethodAddLockedFund:
		resp, err := m.AddLockedFund(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodProveReplicaUpdates3:
		resp, err := m.ProveReplicaUpdates3(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodInternalSectorSetupForPreseal:
		resp, err := m.InternalSectorSetupForPreseal(network, height, msg.Params, msgRct.Return)
		return resp, nil, err
	case parser.MethodConfirmUpdateWorkerKey:
		resp, err := m.ConfirmUpdateWorkerKey(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodProveCommitSectorsNI:
		resp, err := m.ProveCommitSectorsNI(network, height, msg.Params)
		return resp, nil, err
	case parser.UnknownStr:
		resp, err := actors.ParseUnknownMetadata(msg.Params, msgRct.Return)
		return resp, nil, err
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (m *Miner) TransactionTypes() map[string]any {
	return map[string]any{
		parser.MethodSend:                               actors.ParseSend,
		parser.MethodConstructor:                        m.Constructor,
		parser.MethodControlAddresses:                   m.ControlAddresses,
		parser.MethodChangeWorkerAddress:                m.ChangeWorkerAddressExported,
		parser.MethodChangeWorkerAddressExported:        m.ChangeWorkerAddressExported,
		parser.MethodChangePeerID:                       m.ChangePeerIDExported,
		parser.MethodChangePeerIDExported:               m.ChangePeerIDExported,
		parser.MethodSubmitWindowedPoSt:                 m.SubmitWindowedPoSt,
		parser.MethodPreCommitSector:                    m.PreCommitSector,
		parser.MethodProveCommitSector:                  m.ProveCommitSector,
		parser.MethodProveCommitSectors3:                m.ProveCommitSectors3,
		parser.MethodExtendSectorExpiration:             m.ExtendSectorExpiration,
		parser.MethodTerminateSectors:                   m.TerminateSectors,
		parser.MethodDeclareFaults:                      m.DeclareFaults,
		parser.MethodDeclareFaultsRecovered:             m.DeclareFaultsRecovered,
		parser.MethodOnDeferredCronEvent:                m.OnDeferredCronEvent,
		parser.MethodCheckSectorProven:                  m.CheckSectorProven,
		parser.MethodApplyRewards:                       m.ApplyRewards,
		parser.MethodReportConsensusFault:               m.ReportConsensusFault,
		parser.MethodWithdrawBalance:                    m.WithdrawBalanceExported,
		parser.MethodWithdrawBalanceExported:            m.WithdrawBalanceExported,
		parser.MethodConfirmSectorProofsValid:           m.ConfirmSectorProofsValid,
		parser.MethodChangeMultiaddrs:                   m.ChangeMultiaddrsExported,
		parser.MethodChangeMultiaddrsExported:           m.ChangeMultiaddrsExported,
		parser.MethodCompactPartitions:                  m.CompactPartitions,
		parser.MethodCompactSectorNumbers:               m.CompactSectorNumbers,
		parser.MethodConfirmChangeWorkerAddress:         actors.ParseEmptyParamsAndReturn,
		parser.MethodConfirmChangeWorkerAddressExported: actors.ParseEmptyParamsAndReturn,
		parser.MethodConfirmUpdateWorkerKey:             m.ConfirmUpdateWorkerKey,
		parser.MethodRepayDebt:                          actors.ParseEmptyParamsAndReturn,
		parser.MethodRepayDebtExported:                  actors.ParseEmptyParamsAndReturn,
		parser.MethodChangeOwnerAddress:                 m.ChangeOwnerAddressExported,
		parser.MethodChangeOwnerAddressExported:         m.ChangeOwnerAddressExported,
		parser.MethodDisputeWindowedPoSt:                m.DisputeWindowedPoSt,
		parser.MethodPreCommitSectorBatch:               m.PreCommitSectorBatch,
		parser.MethodProveCommitAggregate:               m.ProveCommitAggregate,
		parser.MethodProveReplicaUpdates:                m.ProveReplicaUpdates,
		parser.MethodPreCommitSectorBatch2:              m.PreCommitSectorBatch2,
		parser.MethodProveReplicaUpdates2:               m.ProveReplicaUpdates2,
		parser.MethodChangeBeneficiary:                  m.ChangeBeneficiaryExported,
		parser.MethodChangeBeneficiaryExported:          m.ChangeBeneficiaryExported,
		parser.MethodGetBeneficiary:                     m.GetBeneficiary,
		parser.MethodExtendSectorExpiration2:            m.ExtendSectorExpiration2,
		parser.MethodGetOwner:                           m.GetOwnerExported,
		parser.MethodIsControllingAddressExported:       m.IsControllingAddressExported,
		parser.MethodGetSectorSize:                      m.GetSectorSize,
		parser.MethodGetAvailableBalance:                m.GetAvailableBalanceExported,
		parser.MethodGetVestingFunds:                    m.GetVestingFundsExported,
		parser.MethodGetPeerID:                          m.GetPeerIDExported,
		parser.MethodGetMultiaddrs:                      m.GetMultiaddrsExported,
		parser.MethodAddLockedFund:                      m.AddLockedFund,
		parser.MethodInternalSectorSetupForPreseal:      m.InternalSectorSetupForPreseal,
		parser.MethodProveCommitSectorsNI:               m.ProveCommitSectorsNI,
		parser.MethodProveReplicaUpdates3:               m.ProveReplicaUpdates3,
	}
}
