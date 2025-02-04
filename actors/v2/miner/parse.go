package miner

import (
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func (m *Miner) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, _ cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		resp, err := m.Constructor(network, height, msg.Params)
		return resp, nil, err
	case parser.MethodControlAddresses:
		// return m.ControlAddresses(network, height, msg.Params, msgRct.Return)
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
		// return m.emptyParamsAndReturn()
	case parser.MethodConfirmUpdateWorkerKey: // TODO: ?
	case parser.MethodRepayDebt, parser.MethodRepayDebtExported:
		// return p.emptyParamsAndReturn()
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
		// return m.GetBeneficiaryExported(network, height, msg.Params, msgRct.Return)
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
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, nil, parser.ErrUnknownMethod
}

func (m *Miner) TransactionTypes() []string {
	return []string{
		parser.MethodSend,
		parser.MethodConstructor,
		parser.MethodControlAddresses,
		parser.MethodChangeWorkerAddress,
		parser.MethodChangeWorkerAddressExported,
		parser.MethodChangePeerID,
		parser.MethodChangePeerIDExported,
		parser.MethodSubmitWindowedPoSt,
		parser.MethodPreCommitSector,
		parser.MethodProveCommitSector,
		parser.MethodProveCommitSectors3,
		parser.MethodExtendSectorExpiration,
		parser.MethodTerminateSectors,
		parser.MethodDeclareFaults,
		parser.MethodDeclareFaultsRecovered,
		parser.MethodOnDeferredCronEvent,
		parser.MethodCheckSectorProven,
		parser.MethodApplyRewards,
		parser.MethodReportConsensusFault,
		parser.MethodWithdrawBalance,
		parser.MethodWithdrawBalanceExported,
		parser.MethodConfirmSectorProofsValid,
		parser.MethodChangeMultiaddrs,
		parser.MethodChangeMultiaddrsExported,
		parser.MethodCompactPartitions,
		parser.MethodCompactSectorNumbers,
		parser.MethodConfirmChangeWorkerAddress,
		parser.MethodConfirmChangeWorkerAddressExported,
		parser.MethodConfirmUpdateWorkerKey,
		parser.MethodRepayDebt,
		parser.MethodRepayDebtExported,
		parser.MethodChangeOwnerAddress,
		parser.MethodChangeOwnerAddressExported,
		parser.MethodDisputeWindowedPoSt,
		parser.MethodPreCommitSectorBatch,
		parser.MethodProveCommitAggregate,
		parser.MethodProveReplicaUpdates,
		parser.MethodPreCommitSectorBatch2,
		parser.MethodProveReplicaUpdates2,
		parser.MethodChangeBeneficiary,
		parser.MethodChangeBeneficiaryExported,
		parser.MethodGetBeneficiary,
		parser.MethodExtendSectorExpiration2,
		parser.MethodGetOwner,
		parser.MethodIsControllingAddressExported,
		parser.MethodGetSectorSize,
		parser.MethodGetAvailableBalance,
		parser.MethodGetVestingFunds,
		parser.MethodGetPeerID,
		parser.MethodGetMultiaddrs,
	}
}
