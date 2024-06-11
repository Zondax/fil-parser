package multisig

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

const (
	VERSION       = "V1"
	MultisigStore = "multisig"

	metadataParams     = "Params"
	metadataReturn     = "Return"
	metadataTxnIDField = "TxnID"
	metadataIDField    = "ID"
	metadataMethod     = "Method"
	metadataValue      = "Value"

	txStatusOk = "ok"
)

var proposeTranslateMap = map[string]string{
	parser.MethodPropose:         parser.MethodPropose,
	parser.MethodProposeExported: parser.MethodPropose,
}

var cancelApproveTranslateMap = map[string]string{
	parser.MethodApprove:         parser.MethodApprove,
	parser.MethodApproveExported: parser.MethodApprove,
	parser.MethodCancel:          parser.MethodCancel,
	parser.MethodCancelExported:  parser.MethodCancel,
}

type EventGenerator interface {
	GenerateMultisigEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string) (*types.MultisigEvents, error)
}

type eventGenerator struct {
	helper *helper.Helper
	logger *zap.Logger
}

func NewEventGenerator(helper *helper.Helper, logger *zap.Logger) EventGenerator {
	return &eventGenerator{
		helper: helper,
		logger: logger,
	}
}

func (eg *eventGenerator) GenerateMultisigEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string) (*types.MultisigEvents, error) {
	events := &types.MultisigEvents{
		Proposals:    []*types.MultisigProposal{},
		MultisigInfo: []*types.MultisigInfo{},
	}

	for _, tx := range transactions {
		if !strings.EqualFold(tx.Status, txStatusOk) {
			eg.logger.Sugar().Debug("failed tx found, skipping it")
			continue
		}

		metadata, err := tools.ParseTxMetadata(tx.TxMetadata)
		if err != nil {
			return nil, err
		}

		isProposalType := isProposalType(tx.TxType)
		if isProposalType {
			proposal := eg.createProposal(ctx, tx, metadata, tipsetCid)
			events.Proposals = append(events.Proposals, proposal)
		} else {
			// Only consider transactions where tx.TxTo is a multisig address
			// because we only take into account those that change the state of the multisig

			addrTo, err := address.NewFromString(tx.TxTo)
			if err != nil {
				eg.logger.Sugar().Errorf("could not parse address. Err: %s", err)
				return nil, err
			}

			actorName, err := eg.helper.GetActorNameFromAddress(addrTo, int64(tx.Height), filTypes.EmptyTSK)
			if err != nil {
				eg.logger.Sugar().Errorf("could not get actor name from address. Err: %s", err)
				continue
			}
			if !strings.EqualFold(actorName, manifest.MultisigKey) {
				continue
			}

			multisigInfo, err := eg.createMultisigInfo(ctx, tx, tipsetCid)
			if err != nil {
				// TODO: Metric
				continue
			}
			events.MultisigInfo = append(events.MultisigInfo, multisigInfo)
		}
	}

	return events, nil
}

func (eg *eventGenerator) createProposal(ctx context.Context, tx *types.Transaction, metadata map[string]interface{}, tipsetCid string) *types.MultisigProposal {
	proposal := &types.MultisigProposal{
		MultisigAddress: tx.TxTo,
		Height:          tx.Height,
		TxCid:           tx.TxCid,
		Signer:          tx.TxFrom,
	}

	eg.processProposalParams(ctx, metadata, tx.TxType, proposal)
	if ret, ok := metadata[metadataReturn].(map[string]interface{}); ok {
		if txnID, ok := ret[metadataTxnIDField].(float64); ok {
			proposal.ProposalID = int64(txnID)
		}
	}

	proposal.ID = tools.BuildId(tipsetCid, proposal.MultisigAddress, fmt.Sprint(proposal.ProposalID), fmt.Sprint(tx.Height), tx.TxType)
	return proposal
}

func (eg *eventGenerator) processProposalParams(ctx context.Context, metadata map[string]interface{}, txType string, proposal *types.MultisigProposal) {
	if isCancelOrApprove(txType) {
		proposal.ActionType = cancelApproveTranslateMap[txType]
		proposal.TxTypeToExecute = ""

		metadata[metadataParams] = eg.parseParamsString(ctx, metadata)

		if params, ok := metadata[metadataParams].(map[string]interface{}); ok {
			if metadataID, ok := params[metadataIDField].(float64); ok {
				proposal.ProposalID = int64(metadataID)
			}
			eg.processNestedParams(ctx, params, proposal)
		}
	} else {
		proposal.ActionType = proposeTranslateMap[txType]
		proposal.TxTypeToExecute = parser.MethodUnknown

		params, ok := metadata[metadataParams].(map[string]interface{})
		if !ok {
			return
		}
		eg.processNestedParams(ctx, params, proposal)

		method, _ := params[metadataMethod].(string)
		if method != "" {
			proposal.TxTypeToExecute = method
			return
		}

		metadataJSON, _ := json.Marshal(metadata)
		eg.logger.Sugar().Debug(ctx, fmt.Sprintf("unknown method with metadata %v", string(metadataJSON)))
		proposal.Value = string(metadataJSON)
	}
}

func (eg *eventGenerator) processNestedParams(ctx context.Context, params map[string]interface{}, proposal *types.MultisigProposal) {
	if nestedParams, ok := params[metadataParams].(map[string]interface{}); ok {
		jsonParams, err := json.Marshal(nestedParams)
		if err != nil {
			eg.logger.Sugar().Error(ctx, fmt.Sprintf("Error marshaling nested params: %v", err))
			return
		}
		proposal.Value = string(jsonParams)
		return
	}

	if valueStr, ok := params[metadataValue].(string); ok {
		params = map[string]interface{}{"Value": valueStr}
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		eg.logger.Sugar().Error(ctx, fmt.Sprintf("Error marshaling params: %v", err))
		return
	}

	eg.logger.Sugar().Debug(ctx, fmt.Sprintf("zero value with params: %v", string(jsonParams)))
	proposal.Value = string(jsonParams)

}

func (eg *eventGenerator) createMultisigInfo(ctx context.Context, tx *types.Transaction, tipsetCid string) (*types.MultisigInfo, error) {
	value, err := actors.ParseMultisigMetadata(tx.TxType, tx.TxMetadata)
	if err != nil {
		eg.logger.Sugar().Error(ctx, fmt.Sprintf("Multisig error parsing metadata: %s", err.Error()))
		value = tx.TxMetadata // if there is an error then we need to store the raw metadata
	}

	b, err := json.Marshal(value)
	if err != nil {
		eg.logger.Sugar().Error(ctx, fmt.Sprintf("Multisig error marshaling value: %s", err.Error()))
		return nil, err
	}

	return &types.MultisigInfo{
		ID:              tools.BuildId(tipsetCid, tx.TxTo, fmt.Sprint(tx.Height), tx.TxCid, tx.TxType),
		MultisigAddress: tx.TxTo,
		Height:          tx.Height,
		TxCid:           tx.TxCid,
		Signer:          tx.TxFrom,
		ActionType:      tx.TxType,
		Value:           string(b),
	}, nil
}

func (eg *eventGenerator) parseParamsString(ctx context.Context, metadata map[string]interface{}) map[string]interface{} {
	var params map[string]interface{}
	if paramsStr, ok := metadata[metadataParams].(string); ok {
		if err := json.Unmarshal([]byte(paramsStr), &params); err != nil {
			eg.logger.Sugar().Error(fmt.Sprintf("Error deserializing params string: %v", err))
			return nil
		}
	}
	return params
}

func isProposalType(txType string) bool {
	return proposeTranslateMap[txType] != "" || cancelApproveTranslateMap[txType] != ""
}

func isCancelOrApprove(txType string) bool {
	return cancelApproveTranslateMap[txType] != ""
}
