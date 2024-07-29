package actors

import (
	"bytes"
	"encoding/base64"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	"github.com/zondax/fil-parser/parser"
)

func (p *ActorParser) ParseStorageminer(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case parser.MethodSend:
		return p.parseSend(msg), nil
	case parser.MethodConstructor:
		return p.minerConstructor(msg.Params)
	case parser.MethodControlAddresses:
		return p.controlAddresses(msg.Params, msgRct.Return)
	case parser.MethodChangeWorkerAddress, parser.MethodChangeWorkerAddressExported:
		return p.changeWorkerAddress(msg.Params)
	case parser.MethodChangePeerID, parser.MethodChangePeerIDExported:
		return p.changePeerID(msg.Params)
	case parser.MethodSubmitWindowedPoSt:
		return p.submitWindowedPoSt(msg.Params)
	case parser.MethodPreCommitSector:
		return p.preCommitSector(msg.Params)
	case parser.MethodProveCommitSector:
		return p.proveCommitSector(msg.Params)
	case parser.MethodProveCommitSectors3:
		return p.proveCommitSectors3(msg.Params, msgRct.Return)
	case parser.MethodExtendSectorExpiration:
		return p.extendSectorExpiration(msg.Params)
	case parser.MethodTerminateSectors:
		return p.terminateSectors(msg.Params, msgRct.Return)
	case parser.MethodDeclareFaults:
		return p.declareFaults(msg.Params)
	case parser.MethodDeclareFaultsRecovered:
		return p.declareFaultsRecovered(msg.Params)
	case parser.MethodOnDeferredCronEvent:
		return p.onDeferredCronEvent(msg.Params)
	case parser.MethodCheckSectorProven:
		return p.checkSectorProven(msg.Params)
	case parser.MethodApplyRewards:
		return p.applyRewards(msg.Params)
	case parser.MethodReportConsensusFault:
		return p.reportConsensusFault(msg.Params)
	case parser.MethodWithdrawBalance, parser.MethodWithdrawBalanceExported:
		return p.parseWithdrawBalance(msg.Params)
	case parser.MethodConfirmSectorProofsValid:
		return p.confirmSectorProofsValid(msg.Params)
	case parser.MethodChangeMultiaddrs, parser.MethodChangeMultiaddrsExported:
		return p.changeMultiaddrs(msg.Params)
	case parser.MethodCompactPartitions:
		return p.compactPartitions(msg.Params)
	case parser.MethodCompactSectorNumbers:
		return p.compactSectorNumbers(msg.Params)
	case parser.MethodConfirmChangeWorkerAddress, parser.MethodConfirmChangeWorkerAddressExported:
		return p.emptyParamsAndReturn()
	case parser.MethodConfirmUpdateWorkerKey: // TODO: ?
	case parser.MethodRepayDebt, parser.MethodRepayDebtExported:
		return p.emptyParamsAndReturn()
	case parser.MethodChangeOwnerAddress, parser.MethodChangeOwnerAddressExported: // TODO: not tested
		return p.changeOwnerAddress(msg.Params)
	case parser.MethodDisputeWindowedPoSt:
		return p.disputeWindowedPoSt(msg.Params)
	case parser.MethodPreCommitSectorBatch:
		return p.preCommitSectorBatch(msg.Params)
	case parser.MethodProveCommitAggregate:
		return p.proveCommitAggregate(msg.Params)
	case parser.MethodProveReplicaUpdates:
		return p.proveReplicaUpdates(msg.Params)
	case parser.MethodPreCommitSectorBatch2:
		return p.preCommitSectorBatch2(msg.Params)
	case parser.MethodProveReplicaUpdates2:
		return p.proveReplicaUpdates2(msg.Params, msgRct.Return)
	case parser.MethodChangeBeneficiary, parser.MethodChangeBeneficiaryExported:
		return p.changeBeneficiary(msg.Params)
	case parser.MethodGetBeneficiary:
		return p.getBeneficiary(msg.Params, msgRct.Return)
	case parser.MethodExtendSectorExpiration2:
		return p.extendSectorExpiration2(msg.Params)
	case parser.MethodGetOwner:
		return p.getOwner(msgRct.Return)
	case parser.MethodIsControllingAddressExported:
		return p.isControllingAddressExported(msg.Params, msgRct.Return)
	case parser.MethodGetSectorSize:
		return p.getSectorSize(msgRct.Return)
	case parser.MethodGetAvailableBalance:
		return p.getAvailableBalance(msgRct.Return)
	case parser.MethodGetVestingFunds:
		return p.getVestingFunds(msgRct.Return)
	case parser.MethodGetPeerID:
		return p.getPeerID(msgRct.Return)
	case parser.MethodGetMultiaddrs:
		return p.getMultiaddrs(msgRct.Return)
	case parser.UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, parser.ErrUnknownMethod
}

func (p *ActorParser) terminateSectors(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) controlAddresses(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) declareFaults(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) declareFaultsRecovered(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) proveReplicaUpdates(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) preCommitSectorBatch2(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) proveReplicaUpdates2(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) proveCommitAggregate(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) preCommitSectorBatch(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) changeOwnerAddress(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) disputeWindowedPoSt(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) compactSectorNumbers(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) compactPartitions(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) changeMultiaddrs(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) checkSectorProven(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) extendSectorExpiration(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) changePeerID(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) changeWorkerAddress(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) reportConsensusFault(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) changeBeneficiary(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) confirmSectorProofsValid(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) minerConstructor(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) parseWithdrawBalance(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) applyRewards(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) preCommitSector(raw []byte) (map[string]interface{}, error) {
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

// Deprecated
func (p *ActorParser) proveCommitSector(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) proveCommitSectors3(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params miner14.ProveCommitSectors3Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var returnVal miner14.ProveCommitSectors3Return
	err = returnVal.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = returnVal

	return metadata, nil
}

func (p *ActorParser) submitWindowedPoSt(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) onDeferredCronEvent(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) getBeneficiary(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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
	metadata[parser.ReturnKey] = parser.GetBeneficiaryReturn{
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

func (p *ActorParser) isControllingAddressExported(rawParams, rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) extendSectorExpiration2(raw []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) getOwner(rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) getSectorSize(rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) getAvailableBalance(rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) getVestingFunds(rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) getPeerID(rawReturn []byte) (map[string]interface{}, error) {
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

func (p *ActorParser) getMultiaddrs(rawReturn []byte) (map[string]interface{}, error) {
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
