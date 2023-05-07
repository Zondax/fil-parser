package actors

import (
	"bytes"
	"encoding/base64"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/builtin/v11/miner"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
)

func ParseStorageminer(txType string, msg *parser.LotusMessage, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return parseSend(msg), nil
	case parser.MethodConstructor:
		return minerConstructor(msg.Params)
	case parser.MethodControlAddresses:
		return controlAddresses(msg.Params, msgRct.Return)
	case parser.MethodChangeWorkerAddress, parser.MethodChangeWorkerAddressExported:
		return changeWorkerAddress(msg.Params)
	case parser.MethodChangePeerID, parser.MethodChangePeerIDExported:
		return changePeerID(msg.Params)
	case parser.MethodSubmitWindowedPoSt:
		return submitWindowedPoSt(msg.Params)
	case parser.MethodPreCommitSector:
		return preCommitSector(msg.Params)
	case parser.MethodProveCommitSector:
		return proveCommitSector(msg.Params)
	case parser.MethodExtendSectorExpiration:
		return extendSectorExpiration(msg.Params)
	case parser.MethodTerminateSectors:
		return terminateSectors(msg.Params, msgRct.Return)
	case parser.MethodDeclareFaults:
		return declareFaults(msg.Params)
	case parser.MethodDeclareFaultsRecovered:
		return declareFaultsRecovered(msg.Params)
	case parser.MethodOnDeferredCronEvent:
		return onDeferredCronEvent(msg.Params)
	case parser.MethodCheckSectorProven:
		return checkSectorProven(msg.Params)
	case parser.MethodApplyRewards:
		return applyRewards(msg.Params)
	case parser.MethodReportConsensusFault:
		return reportConsensusFault(msg.Params)
	case parser.MethodWithdrawBalance, parser.MethodWithdrawBalanceExported:
		return parseWithdrawBalance(msg.Params)
	case parser.MethodConfirmSectorProofsValid:
		return confirmSectorProofsValid(msg.Params)
	case parser.MethodChangeMultiaddrs, parser.MethodChangeMultiaddrsExported:
		return changeMultiaddrs(msg.Params)
	case parser.MethodCompactPartitions:
		return compactPartitions(msg.Params)
	case parser.MethodCompactSectorNumbers:
		return compactSectorNumbers(msg.Params)
	case parser.MethodConfirmChangeWorkerAddress, parser.MethodConfirmChangeWorkerAddressExported:
		return emptyParamsAndReturn()
	case parser.MethodConfirmUpdateWorkerKey: // TODO: ?
	case parser.MethodRepayDebt, parser.MethodRepayDebtExported:
		return emptyParamsAndReturn()
	case parser.MethodChangeOwnerAddress, parser.MethodChangeOwnerAddressExported: // TODO: not tested
		return changeOwnerAddress(msg.Params)
	case parser.MethodDisputeWindowedPoSt:
		return disputeWindowedPoSt(msg.Params)
	case parser.MethodPreCommitSectorBatch:
		return preCommitSectorBatch(msg.Params)
	case parser.MethodProveCommitAggregate:
		return proveCommitAggregate(msg.Params)
	case parser.MethodProveReplicaUpdates:
		return proveReplicaUpdates(msg.Params)
	case parser.MethodPreCommitSectorBatch2:
		return preCommitSectorBatch2(msg.Params)
	case parser.MethodProveReplicaUpdates2:
		return proveReplicaUpdates2(msg.Params, msgRct.Return)
	case parser.MethodChangeBeneficiary, parser.MethodChangeBeneficiaryExported:
		return changeBeneficiary(msg.Params)
	case parser.MethodGetBeneficiary:
		return getBeneficiary(msg.Params, msgRct.Return)
	case parser.MethodExtendSectorExpiration2:
		return extendSectorExpiration2(msg.Params)
	case parser.MethodGetOwner:
		return getOwner(msgRct.Return)
	case parser.MethodIsControllingAddressExported:
		return isControllingAddressExported(msg.Params, msgRct.Return)
	case parser.MethodGetSectorSize:
		return getSectorSize(msgRct.Return)
	case parser.MethodGetAvailableBalance:
		return getAvailableBalance(msgRct.Return)
	case parser.MethodGetVestingFunds:
		return getVestingFunds(msgRct.Return)
	case parser.MethodGetPeerID:
		return getPeerID(msgRct.Return)
	case parser.MethodGetMultiaddrs:
		return getMultiaddrs(msgRct.Return)
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func terminateSectors(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params miner.TerminateSectorsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var terminateReturn miner.TerminateSectorsReturn
	err = terminateReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = terminateReturn
	return metadata, nil
}

func controlAddresses(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if rawParams != nil {
		metadata[parser.ParamsKey] = base64.StdEncoding.EncodeToString(rawParams)
	}
	reader := bytes.NewReader(rawReturn)
	var controlReturn miner.GetControlAddressesReturn
	err := controlReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = parser.ControlAddress{
		Owner:        controlReturn.Owner.String(),
		Worker:       controlReturn.Worker.String(),
		ControlAddrs: getControlAddrs(controlReturn.ControlAddrs),
	}
	return metadata, nil
}

func declareFaults(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.DeclareFaultsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func declareFaultsRecovered(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.DeclareFaultsRecoveredParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func proveReplicaUpdates(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ProveReplicaUpdatesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func preCommitSectorBatch2(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.PreCommitSectorBatchParams2
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func proveReplicaUpdates2(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params miner.ProveReplicaUpdatesParams2
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r bitfield.BitField
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func proveCommitAggregate(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ProveCommitAggregateParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func preCommitSectorBatch(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.PreCommitSectorBatchParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func changeOwnerAddress(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	return metadata, nil
}

func disputeWindowedPoSt(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.DisputeWindowedPoStParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func compactSectorNumbers(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.CompactSectorNumbersParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func compactPartitions(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.CompactPartitionsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func changeMultiaddrs(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ChangeMultiaddrsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func checkSectorProven(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.CheckSectorProvenParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func extendSectorExpiration(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ExtendSectorExpirationParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func changePeerID(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ChangePeerIDParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func changeWorkerAddress(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ChangeWorkerAddressParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func reportConsensusFault(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ReportConsensusFaultParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func changeBeneficiary(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ChangeBeneficiaryParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func confirmSectorProofsValid(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ConfirmSectorProofsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func minerConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.MinerConstructorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func parseWithdrawBalance(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.WithdrawBalanceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func applyRewards(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ApplyRewardParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func preCommitSector(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.PreCommitSectorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func proveCommitSector(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ProveCommitSectorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func submitWindowedPoSt(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.SubmitWindowedPoStParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func onDeferredCronEvent(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.DeferredCronEventParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func getBeneficiary(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if rawParams != nil {
		metadata[parser.ParamsKey] = base64.StdEncoding.EncodeToString(rawParams)
	}
	reader := bytes.NewReader(rawReturn)
	var beneficiaryReturn miner.GetBeneficiaryReturn
	err := beneficiaryReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = parser.GetBeneficiryReturn{
		Active: parser.ActiveBeneficiary{
			Beneficiary: beneficiaryReturn.Active.Beneficiary.String(),
			Term: parser.BeneficiaryTerm{
				Quota:      beneficiaryReturn.Active.Term.Quota.String(),
				UsedQuota:  beneficiaryReturn.Active.Term.UsedQuota.String(),
				Expiration: int64(beneficiaryReturn.Active.Term.Expiration),
			},
		},
		Proposed: parser.Proposed{
			NewBeneficiary:        beneficiaryReturn.Proposed.NewBeneficiary.String(),
			NewQuota:              beneficiaryReturn.Proposed.NewQuota.String(),
			NewExpiration:         int64(beneficiaryReturn.Proposed.NewExpiration),
			ApprovedByBeneficiary: beneficiaryReturn.Proposed.ApprovedByBeneficiary,
			ApprovedByNominee:     beneficiaryReturn.Proposed.ApprovedByNominee,
		},
	}
	return metadata, nil
}

func isControllingAddressExported(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params miner.IsControllingAddressParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params.String()
	reader = bytes.NewReader(rawReturn)
	var terminateReturn miner.IsControllingAddressReturn
	err = terminateReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = terminateReturn
	return metadata, nil
}

func extendSectorExpiration2(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ExtendSectorExpiration2Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func getOwner(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetOwnerReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = params
	return metadata, nil
}

func getSectorSize(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	// TODO: miner.GetSectorSizeReturn does not implement UnmarshalCBOR
	// reader := bytes.NewReader(rawReturn)
	// var params abi.SectorSize
	// err := params.UnmarshalCBOR(reader)
	// if err != nil {
	// 	return metadata, err
	// }
	// metadata[ParamsKey] = params
	return metadata, nil
}

func getAvailableBalance(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetAvailableBalanceReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = params
	return metadata, nil
}

func getVestingFunds(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetVestingFundsReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = params
	return metadata, nil
}

func getPeerID(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetPeerIDReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = params
	return metadata, nil
}

func getMultiaddrs(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetMultiAddrsReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = params
	return metadata, nil
}

func getControlAddrs(addrs []address.Address) []string {
	r := make([]string, len(addrs))
	for i, addr := range addrs {
		r[i] = addr.String()
	}
	return r
}
