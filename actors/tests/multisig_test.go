package actortest

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v11/miner"
	"github.com/filecoin-project/go-state-types/builtin/v11/verifreg"
	multisig2 "github.com/filecoin-project/go-state-types/builtin/v14/multisig"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
	"github.com/zondax/fil-parser/actors/v2/multisig"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var multisigWithParamsOrReturnTests = []struct {
	name   string
	txType string
	key    string
}{
	{
		name:   "Constructor",
		txType: parser.MethodConstructor,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Num Approvals Threshold",
		txType: parser.MethodChangeNumApprovalsThreshold,
		key:    parser.ParamsKey,
	},
	{
		name:   "Change Num Approvals Threshold",
		txType: parser.MethodChangeNumApprovalsThresholdExported,
		key:    parser.ParamsKey,
	},
	{
		name:   "Lock Balance",
		txType: parser.MethodLockBalance,
		key:    parser.ParamsKey,
	},
	{
		name:   "Lock Balance Exported",
		txType: parser.MethodLockBalanceExported,
		key:    parser.ParamsKey,
	},
	{
		name:   "Msig Universal Receiver Hook",
		txType: parser.MethodMsigUniversalReceiverHook,
		key:    parser.ParamsKey,
	},
}
var multisigWithParamsAndReturnTests = []struct {
	name   string
	txType string
}{
	{
		name:   "Propose",
		txType: parser.MethodPropose,
	},
	{
		name:   "Propose Exported",
		txType: parser.MethodProposeExported,
	},
}
var multisigParamsTests = []struct {
	name   string
	txType string
}{
	{
		name:   "Add Signer",
		txType: parser.MethodAddSigner,
	},
	{
		name:   "Add Signer Exported",
		txType: parser.MethodAddSignerExported,
	},
	{
		name:   "Remove Signer",
		txType: parser.MethodRemoveSigner,
	},
	{
		name:   "Remove Signer Exported",
		txType: parser.MethodRemoveSignerExported,
	},
	{
		name:   "Cancel",
		txType: parser.MethodCancel,
	},
	/* // TODO: will fail until https://github.com/Zondax/rosetta-filecoin-lib/pull/109 is merged
	{
		name:   "Cancel Exported",
		txType: parser.MethodCancelExported,

	},
	*/
	{
		name:   "Swap Signer",
		txType: parser.MethodSwapSigner,
	},
	{
		name:   "Swap Signer Exported",
		txType: parser.MethodSwapSignerExported,
	},
}

var multisigApproveTests = []struct {
	name   string
	method string
}{
	{
		name:   "Approve",
		method: parser.MethodApprove,
	},
	{
		name:   "Approve Exported",
		method: parser.MethodApproveExported,
	},
}

func TestActorParserV1_MultisigApprove(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range multisigApproveTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MultisigKey, tt.method)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)
			msg, err := deserializeMessage(manifest.MultisigKey, tt.method)
			require.NoError(t, err)
			require.NotNil(t, msg)
			tipSet, err := deserializeTipset(manifest.MultisigKey, tt.method)
			require.NoError(t, err)

			got, err := p.ParseMultisig(tt.method, msg, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, int64(tipSet.Height()), tipSet.Key(), true)
			require.NoError(t, err)
			require.NotNil(t, got)
		})
	}
}

func TestActorParserV1_MultisigWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range multisigWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MultisigKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := p.ParseMultisig(tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, height, filTypes.EmptyTSK, true)

			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParserV1_MultisigWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range multisigWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MultisigKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, err := p.ParseMultisig(tt.txType, msg, msgRct, height, filTypes.EmptyTSK, true)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV1_MultiSigParams(t *testing.T) {
	p := getActorParser(actorsV1.NewActorParser).(*actorsV1.ActorParser)

	for _, tt := range multisigParamsTests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := deserializeMessage(manifest.MultisigKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, msg)

			tipset, err := deserializeTipset(manifest.MultisigKey, tt.txType)
			require.NoError(t, err)

			got, err := p.ParseMultisig(tt.txType, msg, &parser.LotusMessageReceipt{
				Return: nil,
			}, int64(tipset.Height()), tipset.Key(), true)
			require.NoError(t, err)
			require.NotNil(t, got)
		})
	}
}

func TestActorParserV1_ParseMultisigMetadata(t *testing.T) {
	filePath := filepath.Join("..", "..", "data", "actors", "multisig", "Metadata", "metadata_test.csv")
	file, err := os.Open(filePath)
	require.NoError(t, err, "Error opening CSV file")
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	require.NoError(t, err, "Error reading CSV file")

	for _, record := range records {
		require.Len(t, record, 3, "Invalid record")

		txType := record[0]
		txMetadata := record[1]
		expectedStr := record[2]

		result, err := actorsV1.ParseMultisigMetadata(txType, txMetadata)
		require.NoError(t, err, "Error parsing metadata for txType %s", txType)

		expected, err := unmarshalExpected(txType, expectedStr)
		require.NoError(t, err, "Error unmarshaling expected for txType %s", txType)

		resultJSON, err := json.Marshal(result)
		require.NoError(t, err, "Error marshaling result to JSON")

		expectedJSON, err := json.Marshal(expected)
		require.NoError(t, err, "Error marshaling expected to JSON")

		var resultMap map[string]interface{}
		var expectedMap map[string]interface{}
		require.NoError(t, json.Unmarshal(resultJSON, &resultMap), "Error unmarshaling result JSON")
		require.NoError(t, json.Unmarshal(expectedJSON, &expectedMap), "Error unmarshaling expected JSON")

		compareRetField(t, txType, resultMap, expectedMap)

		expectedJson, err := json.Marshal(expectedMap)
		require.NoError(t, err)
		resultJson, err := json.Marshal(resultMap)
		require.NoError(t, err)
		assert.Equal(t, string(expectedJson), string(resultJson), "Mismatch for other fields in txType %s", txType)
	}
}

func TestActorParserV2_MultisigApprove(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.MultisigKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range multisigApproveTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MultisigKey, tt.method)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)
			msg, err := deserializeMessage(manifest.MultisigKey, tt.method)
			require.NoError(t, err)
			require.NotNil(t, msg)
			tipSet, err := deserializeTipset(manifest.MultisigKey, tt.method)
			require.NoError(t, err)

			got, _, err := actor.Parse(context.Background(), manifest.MultisigKey, int64(tipSet.Height()), tt.method, msg, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, msg.Cid, tipSet.Key(), true)
			require.NoError(t, err)
			require.NotNil(t, got)
		})
	}
}

func TestActorParserV2_MultisigWithParamsAndReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.MultisigKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range multisigWithParamsAndReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.MultisigKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, _, err := actor.Parse(context.Background(), manifest.MultisigKey, tools.LatestVersion(network).Height(), tt.txType, &parser.LotusMessage{
				Params: rawParams,
			}, &parser.LotusMessageReceipt{
				Return: rawReturn,
			}, cid.Undef, filTypes.EmptyTSK, true)

			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParserV2_MultisigWithParamsOrReturn(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.MultisigKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range multisigWithParamsOrReturnTests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.MultisigKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			msg := &parser.LotusMessage{}
			msgRct := &parser.LotusMessageReceipt{}

			if tt.key == parser.ReturnKey {
				msgRct.Return = rawParams
			} else {
				msg.Params = rawParams
			}

			got, _, err := actor.Parse(context.Background(), manifest.MultisigKey, tools.LatestVersion(network).Height(), tt.txType, msg, msgRct, cid.Undef, filTypes.EmptyTSK, true)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}

func TestActorParserV2_MultiSigParams(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.MultisigKey)
	require.NoError(t, err)
	require.NotNil(t, actor)

	for _, tt := range multisigParamsTests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := deserializeMessage(manifest.MultisigKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, msg)

			tipset, err := deserializeTipset(manifest.MultisigKey, tt.txType)
			require.NoError(t, err)

			got, _, err := actor.Parse(context.Background(), manifest.MultisigKey, int64(tipset.Height()), tt.txType, msg, &parser.LotusMessageReceipt{
				Return: nil,
			}, msg.Cid, tipset.Key(), true)
			require.NoError(t, err)
			require.NotNil(t, got)
		})
	}
}

func TestActorParserV2_ParseMultisigMetadata(t *testing.T) {
	p := getActorParser(actorsV2.NewActorParser).(*actorsV2.ActorParser)
	actor, err := p.GetActor(manifest.MultisigKey)
	require.NoError(t, err)
	require.NotNil(t, actor)
	msigActor := actor.(*multisig.Msig)

	filePath := filepath.Join("..", "..", "data", "actors", "multisig", "Metadata", "metadata_test_v2.csv")
	file, err := os.Open(filePath)
	require.NoError(t, err, "Error opening CSV file")
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	require.NoError(t, err, "Error reading CSV file")
	height := tools.LatestVersion(network).Height()

	for _, record := range records {
		require.Len(t, record, 3, "Invalid record")

		txType := record[0]
		txMetadata := record[1]
		expectedStr := record[2]

		rawParams, rawReturn, err := cborMetadata(txType, txMetadata)
		require.NoError(t, err, "Error unmarshaling expected for txType %s", txType)

		result, _, err := msigActor.Parse(context.Background(), network, height, txType, &parser.LotusMessage{
			Params: rawParams,
		}, &parser.LotusMessageReceipt{
			Return: rawReturn,
		}, cid.Undef, filTypes.EmptyTSK, true)
		require.NoError(t, err, "Error parsing metadata for txType %s", txType)

		// expected, err := unmarshalExpected(txType, expectedStr)
		require.NoError(t, err, "Error unmarshaling expected for txType %s", txType)

		resultJSON, err := json.Marshal(result)
		require.NoError(t, err, "Error marshaling result to JSON")

		// expectedJSON, err := json.Marshal(expected)
		require.NoError(t, err, "Error marshaling expected to JSON")

		require.EqualValuesf(t, expectedStr, string(resultJSON), "Mismatch for other fields in txType %s \n expected: %s\ngot: %s", txType, expectedStr, string(resultJSON))
	}
}

func unmarshalExpected(txType, jsonStr string) (interface{}, error) {
	var v interface{}
	switch txType {
	case parser.MethodAddSigner:
		v = &multisig2.AddSignerParams{}
	case parser.MethodApprove:
		v = &actorsV1.ApproveValue{}
	case parser.MethodCancel:
		v = &actorsV1.CancelValue{}
	case parser.MethodChangeNumApprovalsThreshold:
		v = &multisig2.ChangeNumApprovalsThresholdParams{}
	case parser.MethodConstructor:
		v = &multisig2.ConstructorParams{}
	case parser.MethodLockBalance:
		v = &multisig2.LockBalanceParams{}
	case parser.MethodRemoveSigner:
		v = &multisig2.RemoveSignerParams{}
	case parser.MethodSend:
		v = &actorsV1.SendValue{}
	case parser.MethodSwapSigner:
		v = &multisig2.SwapSignerParams{}
	case parser.MethodMsigUniversalReceiverHook:
		v = &actorsV1.UniversalReceiverHookValue{}
	case parser.MethodAddVerifier:
		v = &verifreg.AddVerifierParams{}
	case parser.MethodWithdrawBalance:
		v = &miner.WithdrawBalanceParams{}
	case parser.MethodChangeOwnerAddress:
		v = &actorsV1.ChangeOwnerAddressParams{}
	case parser.MethodInvokeContract:
		v = &actorsV1.InvokeContractParams{}
	default:
		return nil, fmt.Errorf("unknown txType: %s", txType)
	}
	err := json.Unmarshal([]byte(jsonStr), v)
	return v, err
}

func getInnerParamAndReturnJson(jsonStr string) ([]byte, []byte, error) {
	var data map[string]map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return nil, nil, err
	}

	var paramsBytes []byte
	var retBytes []byte
	var err error

	params, ok := data["Params"]
	if ok {
		paramsBytes, err = json.Marshal(params)
		if err != nil {
			return nil, nil, err
		}
	}
	ret, ok := data["Return"]
	if ok {
		retBytes, err = json.Marshal(ret)
		if err != nil {
			return nil, nil, err
		}
	}

	return paramsBytes, retBytes, nil
}
func cborMetadata(txType, jsonStr string) ([]byte, []byte, error) {
	var bufParams bytes.Buffer
	var bufReturn bytes.Buffer
	switch txType {
	case parser.MethodAddSigner:
		tmp := &multisig2.AddSignerParams{}
		if err := json.Unmarshal([]byte(jsonStr), tmp); err != nil {
			return nil, nil, err
		}
		err := tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodApprove:
		params, ret, err := getInnerParamAndReturnJson(jsonStr)
		if err != nil {
			return nil, nil, err
		}
		p := &multisig2.TxnIDParams{
			ProposalHash: []byte{},
		}
		if err := json.Unmarshal(params, p); err != nil {
			return nil, nil, err
		}
		err = p.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
		r := &multisig2.ApproveReturn{}
		if err := json.Unmarshal(ret, r); err != nil {
			return nil, nil, err
		}
		err = r.MarshalCBOR(&bufReturn)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodCancel:
		params, _, err := getInnerParamAndReturnJson(jsonStr)
		if err != nil {
			return nil, nil, err
		}
		tmp := &multisig2.TxnIDParams{
			ProposalHash: []byte{},
		}
		if err := json.Unmarshal(params, tmp); err != nil {
			return nil, nil, err
		}
		err = tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodChangeNumApprovalsThreshold:
		params, _, err := getInnerParamAndReturnJson(jsonStr)
		if err != nil {
			return nil, nil, err
		}
		tmp := &multisig2.ChangeNumApprovalsThresholdParams{}
		if err := json.Unmarshal(params, tmp); err != nil {
			return nil, nil, err
		}
		err = tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodConstructor:
		params, _, err := getInnerParamAndReturnJson(jsonStr)
		if err != nil {
			return nil, nil, err
		}
		tmp := &multisig2.ConstructorParams{}
		if err := json.Unmarshal(params, tmp); err != nil {
			return nil, nil, err
		}
		err = tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodLockBalance:
		params, _, err := getInnerParamAndReturnJson(jsonStr)
		if err != nil {
			return nil, nil, err
		}
		tmp := &multisig2.LockBalanceParams{}
		if err := json.Unmarshal(params, tmp); err != nil {
			return nil, nil, err
		}
		err = tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodRemoveSigner:
		params, _, err := getInnerParamAndReturnJson(jsonStr)
		if err != nil {
			return nil, nil, err
		}
		tmp := &multisig2.RemoveSignerParams{}
		if err := json.Unmarshal(params, tmp); err != nil {
			return nil, nil, err
		}
		err = tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodSend:
		tmp := abi.CborBytes([]byte("txData"))
		err := tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodSwapSigner:
		tmp := &multisig2.SwapSignerParams{}
		if err := json.Unmarshal([]byte(jsonStr), tmp); err != nil {
			return nil, nil, err
		}
		err := tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodMsigUniversalReceiverHook:
		tmp := abi.CborBytes([]byte("txData"))
		err := tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodAddVerifier:
		params, _, err := getInnerParamAndReturnJson(jsonStr)
		if err != nil {
			return nil, nil, err
		}
		tmp := &verifreg.AddVerifierParams{}
		if err := json.Unmarshal(params, tmp); err != nil {
			return nil, nil, err
		}
		err = tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodWithdrawBalance:
		params, _, err := getInnerParamAndReturnJson(jsonStr)
		if err != nil {
			return nil, nil, err
		}
		tmp := &miner.WithdrawBalanceParams{}
		if err := json.Unmarshal(params, tmp); err != nil {
			return nil, nil, err
		}
		err = tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}
	case parser.MethodChangeOwnerAddress:
		params, _, err := getInnerParamAndReturnJson(jsonStr)
		if err != nil {
			return nil, nil, err
		}
		tmp := &address.Address{}
		if err := json.Unmarshal(params, tmp); err != nil {
			return nil, nil, err
		}
		err = tmp.MarshalCBOR(&bufParams)
		if err != nil {
			return nil, nil, err
		}

	default:
		return nil, nil, fmt.Errorf("unknown txType: %s", txType)
	}
	return bufParams.Bytes(), bufReturn.Bytes(), nil
}

func compareRetField(t *testing.T, txType string, resultMap, expectedMap map[string]interface{}) {
	if strings.EqualFold(txType, parser.MethodApprove) {
		resultReturn := resultMap["Return"].(map[string]interface{})
		expectedReturn := expectedMap["Return"].(map[string]interface{})

		resultRet, resultRetExists := resultReturn["Ret"].(string)
		expectedRet, expectedRetExists := expectedReturn["Ret"].(string)

		if resultRetExists && expectedRetExists {
			resultRetBytes, err := base64.StdEncoding.DecodeString(resultRet)
			require.NoError(t, err, "Error decoding result Ret from Base64")

			assert.Equal(t, string(resultRetBytes), expectedRet, "Mismatch for Ret field in txType %s", txType)
		}

		delete(resultReturn, "Ret")
		delete(expectedReturn, "Ret")
	}
}
