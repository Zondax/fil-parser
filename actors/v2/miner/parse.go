package miner

import (
	"github.com/zondax/fil-parser/parser"
)

func (m *Miner) Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		// return p.parseSend(msg), nil
	case parser.MethodConstructor:
		return m.Constructor(network, height, msg.Params)
	case parser.MethodControlAddresses:
		// return m.ControlAddresses(network, height, msg.Params, msgRct.Return)
	case parser.MethodChangeWorkerAddress, parser.MethodChangeWorkerAddressExported:
		return m.ChangeWorkerAddressExported(network, height, msg.Params)
	case parser.MethodChangePeerID, parser.MethodChangePeerIDExported:
		return m.ChangePeerIDExported(network, height, msg.Params)
	case parser.MethodSubmitWindowedPoSt:
		return m.SubmitWindowedPoSt(network, height, msg.Params)
	case parser.MethodPreCommitSector:
		return m.PreCommitSector(network, height, msg.Params)
	case parser.MethodProveCommitSector:
		return m.ProveCommitSector(network, height, msg.Params)
	case parser.MethodProveCommitSectors3:
		return m.ProveCommitSectors3(network, height, msg.Params, msgRct.Return)
	case parser.MethodExtendSectorExpiration:
		return m.ExtendSectorExpiration(network, height, msg.Params)
	case parser.MethodTerminateSectors:
		return m.TerminateSectors(network, height, msg.Params, msgRct.Return)
	case parser.MethodDeclareFaults:
		return m.DeclareFaults(network, height, msg.Params)
	case parser.MethodDeclareFaultsRecovered:
		return m.DeclareFaultsRecovered(network, height, msg.Params)
	case parser.MethodOnDeferredCronEvent:
		return m.OnDeferredCronEvent(network, height, msg.Params)
	case parser.MethodCheckSectorProven:
		return m.CheckSectorProven(network, height, msg.Params)
	case parser.MethodApplyRewards:
		return m.ApplyRewards(network, height, msg.Params)
	case parser.MethodReportConsensusFault:
		return m.ReportConsensusFault(network, height, msg.Params)
	case parser.MethodWithdrawBalance, parser.MethodWithdrawBalanceExported:
		return m.WithdrawBalanceExported(network, height, msg.Params)
	case parser.MethodConfirmSectorProofsValid:
		return m.ConfirmSectorProofsValid(network, height, msg.Params)
	case parser.MethodChangeMultiaddrs, parser.MethodChangeMultiaddrsExported:
		return m.ChangeMultiaddrsExported(network, height, msg.Params)
	case parser.MethodCompactPartitions:
		return m.CompactPartitions(network, height, msg.Params)
	case parser.MethodCompactSectorNumbers:
		return m.CompactSectorNumbers(network, height, msg.Params)
	case parser.MethodConfirmChangeWorkerAddress, parser.MethodConfirmChangeWorkerAddressExported:
		// return m.emptyParamsAndReturn()
	case parser.MethodConfirmUpdateWorkerKey: // TODO: ?
	case parser.MethodRepayDebt, parser.MethodRepayDebtExported:
		// return p.emptyParamsAndReturn()
	case parser.MethodChangeOwnerAddress, parser.MethodChangeOwnerAddressExported: // TODO: not tested
		return m.ChangeOwnerAddressExported(network, height, msg.Params)
	case parser.MethodDisputeWindowedPoSt:
		return m.DisputeWindowedPoSt(network, height, msg.Params)
	case parser.MethodPreCommitSectorBatch:
		return m.PreCommitSectorBatch(network, height, msg.Params)
	case parser.MethodProveCommitAggregate:
		return m.ProveCommitAggregate(network, height, msg.Params)
	case parser.MethodProveReplicaUpdates:
		return m.ProveReplicaUpdates(network, height, msg.Params)
	case parser.MethodPreCommitSectorBatch2:
		return m.PreCommitSectorBatch2(network, height, msg.Params)
	case parser.MethodProveReplicaUpdates2:
		return m.ProveReplicaUpdates2(network, height, msg.Params, msgRct.Return)
	case parser.MethodChangeBeneficiary, parser.MethodChangeBeneficiaryExported:
		return m.ChangeBeneficiaryExported(network, height, msg.Params)
	case parser.MethodGetBeneficiary:
		// return m.GetBeneficiaryExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodExtendSectorExpiration2:
		return m.ExtendSectorExpiration2(network, height, msg.Params)
	case parser.MethodGetOwner:
		return m.GetOwnerExported(network, height, msgRct.Return)
	case parser.MethodIsControllingAddressExported:
		return m.IsControllingAddressExported(network, height, msg.Params, msgRct.Return)
	case parser.MethodGetSectorSize:
		return m.GetSectorSize(network, height, msgRct.Return)
	case parser.MethodGetAvailableBalance:
		return m.GetAvailableBalanceExported(network, height, msgRct.Return)
	case parser.MethodGetVestingFunds:
		return m.GetVestingFundsExported(network, height, msgRct.Return)
	case parser.MethodGetPeerID:
		return m.GetPeerIDExported(network, height, msgRct.Return)
	case parser.MethodGetMultiaddrs:
		return m.GetMultiaddrsExported(network, height, msgRct.Return)
	case parser.UnknownStr:
		// return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}
