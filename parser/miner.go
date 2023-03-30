package parser

import (
	"bytes"
	"encoding/base64"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/builtin/v11/miner"
	filTypes "github.com/filecoin-project/lotus/chain/types"
)

func (p *Parser) parseStorageminer(txType string, msg *filTypes.Message, msgRct *filTypes.MessageReceipt) (map[string]interface{}, error) {
	switch txType {
	case MethodSend:
		return p.parseSend(msg), nil
	case MethodConstructor:
		return p.minerConstructor(msg.Params)
	case MethodControlAddresses:
		return p.controlAddresses(msg.Params, msgRct.Return)
	case MethodChangeWorkerAddress, MethodChangeWorkerAddressExported:
		return p.changeWorkerAddress(msg.Params)
	case MethodChangePeerID, MethodChangePeerIDExported:
		return p.changePeerID(msg.Params)
	case MethodSubmitWindowedPoSt:
		return p.submitWindowedPoSt(msg.Params)
	case MethodPreCommitSector:
		return p.preCommitSector(msg.Params)
	case MethodProveCommitSector:
		return p.proveCommitSector(msg.Params)
	case MethodExtendSectorExpiration:
		return p.extendSectorExpiration(msg.Params)
	case MethodTerminateSectors:
		return p.terminateSectors(msg.Params, msgRct.Return)
	case MethodDeclareFaults:
		return p.declareFaults(msg.Params)
	case MethodDeclareFaultsRecovered:
		return p.declareFaultsRecovered(msg.Params)
	case MethodOnDeferredCronEvent:
		return p.onDeferredCronEvent(msg.Params)
	case MethodCheckSectorProven:
		return p.checkSectorProven(msg.Params)
	case MethodApplyRewards:
		return p.applyRewards(msg.Params)
	case MethodReportConsensusFault:
		return p.reportConsensusFault(msg.Params)
	case MethodWithdrawBalance, MethodWithdrawBalanceExported:
		return p.parseWithdrawBalance(msg.Params)
	case MethodConfirmSectorProofsValid:
		return p.confirmSectorProofsValid(msg.Params)
	case MethodChangeMultiaddrs, MethodChangeMultiaddrsExported:
		return p.changeMultiaddrs(msg.Params)
	case MethodCompactPartitions:
		return p.compactPartitions(msg.Params)
	case MethodCompactSectorNumbers:
		return p.compactSectorNumbers(msg.Params)
	case MethodConfirmChangeWorkerAddress, MethodConfirmChangeWorkerAddressExported:
		return p.emptyParamsAndReturn()
	case MethodConfirmUpdateWorkerKey: // TODO: ?
	case MethodRepayDebt, MethodRepayDebtExported:
		return p.emptyParamsAndReturn()
	case MethodChangeOwnerAddress, MethodChangeOwnerAddressExported: // TODO: not tested
		return p.changeOwnerAddress(msg.Params)
	case MethodDisputeWindowedPoSt:
		return p.disputeWindowedPoSt(msg.Params)
	case MethodPreCommitSectorBatch:
		return p.preCommitSectorBatch(msg.Params)
	case MethodProveCommitAggregate:
		return p.proveCommitAggregate(msg.Params)
	case MethodProveReplicaUpdates:
		return p.proveReplicaUpdates(msg.Params)
	case MethodPreCommitSectorBatch2:
		return p.preCommitSectorBatch2(msg.Params)
	case MethodProveReplicaUpdates2:
		return p.proveReplicaUpdates2(msg.Params, msgRct.Return)
	case MethodChangeBeneficiary, MethodChangeBeneficiaryExported:
		return p.changeBeneficiary(msg.Params)
	case MethodGetBeneficiary:
		return p.getBeneficiary(msg.Params, msgRct.Return)
	case MethodExtendSectorExpiration2:
		return p.extendSectorExpiration2(msg.Params)
	case MethodGetOwner:
		return p.getOwner(msgRct.Return)
	case MethodIsControllingAddressExported:
		return p.isControllingAddressExported(msg.Params, msgRct.Return)
	case MethodGetSectorSize:
		return p.getSectorSize(msgRct.Return)
	case MethodGetAvailableBalance:
		return p.getAvailableBalance(msgRct.Return)
	case MethodGetVestingFunds:
		return p.getVestingFunds(msgRct.Return)
	case MethodGetPeerID:
		return p.getPeerID(msgRct.Return)
	case MethodGetMultiaddrs:
		return p.getMultiaddrs(msgRct.Return)
	case UnknownStr:
		return p.unknownMetadata(msg.Params, msgRct.Return)
	}
	return map[string]interface{}{}, errUnknownMethod
}

func (p *Parser) terminateSectors(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params miner.TerminateSectorsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var terminateReturn miner.TerminateSectorsReturn
	err = terminateReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = terminateReturn
	return metadata, nil
}

func (p *Parser) controlAddresses(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if rawParams != nil {
		metadata[ParamsKey] = base64.StdEncoding.EncodeToString(rawParams)
	}
	reader := bytes.NewReader(rawReturn)
	var controlReturn miner.GetControlAddressesReturn
	err := controlReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = controlAddress{
		Owner:        controlReturn.Owner.String(),
		Worker:       controlReturn.Worker.String(),
		ControlAddrs: getControlAddrs(controlReturn.ControlAddrs),
	}
	return metadata, nil
}

func (p *Parser) declareFaults(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.DeclareFaultsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) declareFaultsRecovered(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.DeclareFaultsRecoveredParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) proveReplicaUpdates(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ProveReplicaUpdatesParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) preCommitSectorBatch2(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.PreCommitSectorBatchParams2
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) proveReplicaUpdates2(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params miner.ProveReplicaUpdatesParams2
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params

	reader = bytes.NewReader(rawReturn)
	var r bitfield.BitField
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = r
	return metadata, nil
}

func (p *Parser) proveCommitAggregate(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ProveCommitAggregateParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) preCommitSectorBatch(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.PreCommitSectorBatchParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) changeOwnerAddress(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params address.Address
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params.String()
	return metadata, nil
}

func (p *Parser) disputeWindowedPoSt(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.DisputeWindowedPoStParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) compactSectorNumbers(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.CompactSectorNumbersParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) compactPartitions(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.CompactPartitionsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) changeMultiaddrs(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ChangeMultiaddrsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) checkSectorProven(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.CheckSectorProvenParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) extendSectorExpiration(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ExtendSectorExpirationParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) changePeerID(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ChangePeerIDParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) changeWorkerAddress(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ChangeWorkerAddressParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) reportConsensusFault(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ReportConsensusFaultParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) changeBeneficiary(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ChangeBeneficiaryParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) confirmSectorProofsValid(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ConfirmSectorProofsParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) minerConstructor(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.MinerConstructorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) parseWithdrawBalance(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.WithdrawBalanceParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) applyRewards(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ApplyRewardParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) preCommitSector(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.PreCommitSectorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) proveCommitSector(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ProveCommitSectorParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) submitWindowedPoSt(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.SubmitWindowedPoStParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) onDeferredCronEvent(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.DeferredCronEventParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) getBeneficiary(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if rawParams != nil {
		metadata[ParamsKey] = base64.StdEncoding.EncodeToString(rawParams)
	}
	reader := bytes.NewReader(rawReturn)
	var beneficiaryReturn miner.GetBeneficiaryReturn
	err := beneficiaryReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = getBeneficiryReturn{
		Active: activeBeneficiary{
			Beneficiary: beneficiaryReturn.Active.Beneficiary.String(),
			Term: beneficiaryTerm{
				Quota:      beneficiaryReturn.Active.Term.Quota.String(),
				UsedQuota:  beneficiaryReturn.Active.Term.UsedQuota.String(),
				Expiration: int64(beneficiaryReturn.Active.Term.Expiration),
			},
		},
		Proposed: proposed{
			NewBeneficiary:        beneficiaryReturn.Proposed.NewBeneficiary.String(),
			NewQuota:              beneficiaryReturn.Proposed.NewQuota.String(),
			NewExpiration:         int64(beneficiaryReturn.Proposed.NewExpiration),
			ApprovedByBeneficiary: beneficiaryReturn.Proposed.ApprovedByBeneficiary,
			ApprovedByNominee:     beneficiaryReturn.Proposed.ApprovedByNominee,
		},
	}
	return metadata, nil
}

func (p *Parser) isControllingAddressExported(rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params miner.IsControllingAddressParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params.String()
	reader = bytes.NewReader(rawReturn)
	var terminateReturn miner.IsControllingAddressReturn
	err = terminateReturn.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = terminateReturn
	return metadata, nil
}

func (p *Parser) extendSectorExpiration2(raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params miner.ExtendSectorExpiration2Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) getOwner(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetOwnerReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = params
	return metadata, nil
}

func (p *Parser) getSectorSize(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	// TODO: miner.GetSectorSizeReturn does not implement UnmarshalCBOR
	//reader := bytes.NewReader(rawReturn)
	//var params abi.SectorSize
	//err := params.UnmarshalCBOR(reader)
	//if err != nil {
	//	return metadata, err
	//}
	//metadata[ParamsKey] = params
	return metadata, nil
}

func (p *Parser) getAvailableBalance(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetAvailableBalanceReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = params
	return metadata, nil
}

func (p *Parser) getVestingFunds(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetVestingFundsReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = params
	return metadata, nil
}

func (p *Parser) getPeerID(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetPeerIDReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = params
	return metadata, nil
}

func (p *Parser) getMultiaddrs(rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	var params miner.GetMultiAddrsReturn
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[ReturnKey] = params
	return metadata, nil
}

func getControlAddrs(addrs []address.Address) []string {
	r := make([]string, len(addrs))
	for i, addr := range addrs {
		r[i] = addr.String()
	}
	return r
}
