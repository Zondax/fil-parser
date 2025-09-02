package multisig

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/tools/common"
	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/multisig"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
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
	metadataSigner     = "Signer"
	metadataSigners    = "Signers"
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
	GenerateMultisigEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MultisigEvents, error)
}

type eventGenerator struct {
	helper  *helper.Helper
	logger  *logger.Logger
	metrics *multisigMetricsClient
	config  parser.Config
}

func NewEventGenerator(helper *helper.Helper, logger *logger.Logger, metrics metrics.MetricsClient, config parser.Config) EventGenerator {
	return &eventGenerator{
		helper:  helper,
		logger:  logger,
		metrics: newClient(metrics, "multisigEventGenerator"),
		config:  config,
	}
}

func (eg *eventGenerator) GenerateMultisigEvents(ctx context.Context, transactions []*types.Transaction, tipsetCid string, tipsetKey filTypes.TipSetKey) (*types.MultisigEvents, error) {
	events := &types.MultisigEvents{
		Proposals:    []*types.MultisigProposal{},
		MultisigInfo: []*types.MultisigInfo{},
	}

	for _, tx := range transactions {
		if !strings.EqualFold(tx.Status, common.TxStatusOk) {
			eg.logger.Debug("failed tx found, skipping it")
			continue
		}

		metadata, err := tools.ParseTxMetadata(tx.TxMetadata)
		if err != nil {
			_ = eg.metrics.UpdateParseTxMetadataMetric(tx.TxType)
			return nil, err
		}

		isProposalType := isProposalType(tx.TxType)
		if isProposalType {
			proposal, err := eg.createProposal(ctx, tx, metadata, tipsetCid)
			if err != nil {
				return nil, fmt.Errorf("could not create proposal. Err: %s", err)
			}
			events.Proposals = append(events.Proposals, proposal)
		} else {
			// Only consider transactions where tx.TxTo is a multisig address
			// because we only take into account those that change the state of the multisig

			addrTo, err := address.NewFromString(tx.TxTo)
			if err != nil {
				eg.logger.Errorf("could not parse address. Err: %s", err)
				continue
			}

			// #nosec G115
			_, actorName, err := eg.helper.GetActorInfoFromAddress(addrTo, int64(tx.Height), tipsetKey)
			if err != nil {
				_ = eg.metrics.UpdateActorNameFromAddressMetric()
				eg.logger.Errorf("could not get actor name from address. Err: %s", err)
				continue
			}
			if !strings.EqualFold(actorName, manifest.MultisigKey) {
				continue
			}

			multisigInfo, err := eg.createMultisigInfo(ctx, tx, tipsetCid, metadata[parser.ParamsKey])
			if err != nil {
				continue
			}
			events.MultisigInfo = append(events.MultisigInfo, multisigInfo)
		}
	}

	return events, nil
}

func (eg *eventGenerator) createProposal(ctx context.Context, tx *types.Transaction, metadata map[string]interface{}, tipsetCid string) (*types.MultisigProposal, error) {
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

	proposal.ID = tools.BuildId(tipsetCid, tx.TxCid, proposal.Signer, proposal.MultisigAddress, fmt.Sprint(proposal.ProposalID), fmt.Sprint(tx.Height), tx.TxType)
	return proposal, nil
}

func (eg *eventGenerator) processProposalParams(ctx context.Context, metadata map[string]interface{}, txType string, proposal *types.MultisigProposal) {
	if isCancelOrApprove(txType) {
		proposal.ActionType = cancelApproveTranslateMap[txType]
		proposal.TxTypeToExecute = ""

		metadata[metadataParams] = eg.parseParams(ctx, metadata)

		if params, ok := metadata[metadataParams].(map[string]interface{}); ok {
			if metadataID, ok := params[metadataIDField].(float64); ok {
				proposal.ProposalID = int64(metadataID)
			}
			jsonParams, err := json.Marshal(params)
			if err != nil {
				eg.logger.Errorf("Error marshaling nested params: %v", err)
				return
			}
			proposal.Value = string(jsonParams)
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
		eg.logger.Debugf("unknown method with metadata %v", string(metadataJSON))
		proposal.Value = string(metadataJSON)
	}
}

func (eg *eventGenerator) processNestedParams(ctx context.Context, params map[string]interface{}, proposal *types.MultisigProposal) {
	if nestedParams, ok := params[metadataParams].(map[string]interface{}); ok {
		jsonParams, err := json.Marshal(nestedParams)
		if err != nil {
			eg.logger.Errorf("Error marshaling nested params: %v", err)
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
		eg.logger.Errorf("Error marshaling params: %v", err)
		return
	}

	eg.logger.Debugf("zero value with params: %v", string(jsonParams))
	proposal.Value = string(jsonParams)

}

func (eg *eventGenerator) createMultisigInfo(ctx context.Context, tx *types.Transaction, tipsetCid string, parsedMetadata any) (*types.MultisigInfo, error) {
	if eg.config.ConsolidateRobustAddress {
		if isAddOrRemoveSigner(tx.TxType) {
			signerData, ok := parsedMetadata.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("parsedMetadata is not a map[string]interface{}")
			}
			signerAddrStr, ok := signerData[metadataSigner].(string)
			if ok {
				signerAddr, err := address.NewFromString(signerAddrStr)
				if err != nil {
					return nil, fmt.Errorf("address.NewFromString(%s): %s", signerAddrStr, err)
				}
				signerAddrStr, err = actors.ConsolidateToRobustAddress(signerAddr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
				if err != nil {
					return nil, fmt.Errorf("actors.ConsolidateToRobustAddress(%s): %s", signerAddrStr, err)
				}
				signerData[metadataSigner] = signerAddrStr
				parsedMetadata = signerData
			}
		}

		if isConstructor(tx.TxType) {
			constructorData, ok := parsedMetadata.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("parsedMetadata is not a map[string]interface{}")
			}
			signers, ok := constructorData[metadataSigners].([]string)
			if ok {
				consolidatedSigners := make([]string, 0, len(signers))
				for _, signer := range signers {
					addr, err := address.NewFromString(signer)
					if err != nil {
						eg.logger.Errorf("address.NewFromString(%s): %s", signer, err)
						break
					}
					robustAddr, err := actors.ConsolidateToRobustAddress(addr, eg.helper, eg.logger, eg.config.RobustAddressBestEffort)
					if err != nil {
						eg.logger.Errorf("actors.ConsolidateToRobustAddress(%s): %s", addr, err)
						break
					}
					consolidatedSigners = append(consolidatedSigners, robustAddr)
				}
				// only add the consolidated signers if all the signers were consolidated successfully
				if len(consolidatedSigners) == len(signers) {
					constructorData[metadataSigners] = consolidatedSigners
					parsedMetadata = constructorData
				}
			}
		}
	}

	b, err := json.Marshal(parsedMetadata)
	if err != nil {
		_ = eg.metrics.UpdateMarshalMultisigMetadataMetric(tx.TxType)
		eg.logger.Errorf("Multisig error marshaling value: %s", err.Error())
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

func (eg *eventGenerator) parseParams(ctx context.Context, metadata map[string]interface{}) map[string]interface{} {
	var params map[string]interface{}
	switch data := metadata[metadataParams].(type) {
	case string:
		if err := json.Unmarshal([]byte(data), &params); err != nil {
			eg.logger.Errorf("Error deserializing params string: %v", err)
			return nil
		}
	case map[string]any:
		params = data
	default:
		eg.logger.Errorf("Unknown params type: %T", data)
	}

	return params
}

func isProposalType(txType string) bool {
	return proposeTranslateMap[txType] != "" || cancelApproveTranslateMap[txType] != ""
}

func isCancelOrApprove(txType string) bool {
	return cancelApproveTranslateMap[txType] != ""
}

func isAddOrRemoveSigner(txType string) bool {
	switch txType {
	case parser.MethodAddSigner, parser.MethodAddSignerExported, parser.MethodRemoveSigner, parser.MethodRemoveSignerExported:
		return true
	}
	return false
}

func isConstructor(txType string) bool {
	return txType == parser.MethodConstructor
}

/*
GenerateGenesisMultisigData generates the multisig data for an  address in the genesis.
Ref: https://github.com/filecoin-project/lotus/blob/2714a84248095f877f52ce20e737d9c8843a352a/cli/multisig.go#L188

	{
		"Signers":["f1tnzpesy6ddygdfymv3iktnd4cshbpbjlm7qgxhq","f12bpw2u2syy7coh67cidpoydcm5ysqjzxuxdog7y"],
		"NumApprovalsThreshold": 1,
		"UnlockDuration": 3153600,
		"StartEpoch":           147120
	}
*/
func GenerateGenesisMultisigData(ctx context.Context, api api.FullNode, addr address.Address, genesisTipset *types.ExtendedTipSet) (map[string]any, error) {
	store := adt.WrapStore(ctx, cbor.NewCborStore(blockstore.NewAPIBlockstore(api)))

	act, err := api.StateGetActor(ctx, addr, genesisTipset.Key())
	if err != nil {
		return nil, fmt.Errorf("api.StateGetActor(): %s", err)
	}

	mstate, err := multisig.Load(store, act)
	if err != nil {
		return nil, fmt.Errorf("multisig.Load(): %s", err)
	}

	signers, err := mstate.Signers()
	if err != nil {
		return nil, err
	}

	var signerActors []string
	for _, s := range signers {
		// the address is already a robust address
		signerActor, err := api.StateAccountKey(ctx, s, filTypes.EmptyTSK)
		if err != nil {
			return nil, fmt.Errorf("api.StateAccountKey(): %s", err)
		}
		signerActors = append(signerActors, signerActor.String())
	}

	threshold, err := mstate.Threshold()
	if err != nil {
		return nil, fmt.Errorf("mstate.Threshold(): %s", err)
	}

	startEpoch, err := mstate.StartEpoch()
	if err != nil {
		return nil, fmt.Errorf("mstate.StartEpoch(): %s", err)
	}

	unlockDuration, err := mstate.UnlockDuration()
	if err != nil {
		return nil, fmt.Errorf("mstate.UnlockDuration(): %s", err)
	}

	metadata := map[string]interface{}{
		"Signers":               signerActors,
		"NumApprovalsThreshold": threshold,
		"UnlockDuration":        unlockDuration,
		"StartEpoch":            startEpoch,
	}
	return metadata, nil
}
