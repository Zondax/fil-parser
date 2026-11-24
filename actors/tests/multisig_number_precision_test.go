package actortest

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/manifest"
	filApiTypes "github.com/filecoin-project/lotus/api/types"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/zondax/golem/pkg/logger"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	"github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	helper2 "github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/mocks"
	toolsMultisig "github.com/zondax/fil-parser/tools/multisig"
	"github.com/zondax/fil-parser/types"

	actorsV1 "github.com/zondax/fil-parser/actors/v1"
	actorsV2 "github.com/zondax/fil-parser/actors/v2"
)

// Mainnet message bafy2bzacedyg4uiezhkabfgvbbtfluie5xiukxdiivnjlfqsn7omti3r63rxy (height 3181832):
// a multisig Propose of miner ChangeBeneficiary (method 30) whose NewExpiration is MaxInt64.
// Params and receipt Return are the on-chain bytes (Filecoin.ChainGetMessage / StateSearchMsg).
const (
	changeBeneficiaryProposeCid    = "bafy2bzacedyg4uiezhkabfgvbbtfluie5xiukxdiivnjlfqsn7omti3r63rxy"
	changeBeneficiaryProposeHeight = int64(3181832)
	changeBeneficiaryProposeParams = "hEQA1NhOQQAYHlgig0UAoZyBAVEA/////////////////////xt//////////w=="
	changeBeneficiaryProposeReturn = "hAv0AEA="
	changeBeneficiaryProposeFrom   = "f3q26reguzsk5z6e3fxxkftotcozaj6bw3yttouymeazgjfqtvowgp6bgkglo7x2cu4ikvbqc6dgwc7jdc6o5q"
	changeBeneficiaryProposeTo     = "f02363802"

	// exact values decoded by Lotus StateDecodeParams for the inner params
	maxInt64Expiration = `"NewExpiration":9223372036854775807`
	maxQuota           = `"NewQuota":"340282366920938463463374607431768211455"`
	// what float64 makes of MaxInt64
	roundedExpiration = "9223372036854776000"
)

// getMinerTargetActorParser returns a v2 actor parser whose address lookups resolve the proposal
// target to a miner actor, so the inner ChangeBeneficiary params are decoded offline.
func getMinerTargetActorParser(t *testing.T) (*actorsV2.ActorParser, *helper2.Helper) {
	t.Helper()
	actorCid := cid.MustParse("bafk2bzaceaclpbrhoqdruvsuqqgknvy2k5dywzmjoehk4uarce3uvt3w2rewu")
	lotusClient := &mocks.FullNode{}
	lotusClient.On("StateNetworkName", mock.Anything).Return(dtypes.NetworkName("mainnet"), nil)
	lotusClient.On("StateNetworkVersion", mock.Anything, mock.Anything).Return(filApiTypes.NetworkVersion(21), nil)
	lotusClient.On("StateActorCodeCIDs", mock.Anything, mock.Anything).Return(map[string]cid.Cid{
		manifest.MinerKey: actorCid,
	}, nil)

	cache := &mocks.IActorsCache{}
	cache.On("StoreAddressInfo", mock.Anything).Return(nil)
	cache.On("GetActorCode", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(actorCid.String(), nil)
	cache.On("GetActorNameFromAddress", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(manifest.MinerKey, nil)

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
	helper := helper2.NewHelper(lib, cache, lotusClient, nil, metrics.NewNoopMetricsClient())
	p := actorsV2.NewActorParser(network, helper, logger.NewDevelopmentLogger(), metrics.NewNoopMetricsClient())
	return p.(*actorsV2.ActorParser), helper
}

func TestMultisigProposeChangeBeneficiaryKeepsMaxInt64Expiration(t *testing.T) {
	p, helper := getMinerTargetActorParser(t)
	msigActor, err := p.GetActor(manifest.MultisigKey)
	require.NoError(t, err)

	rawParams, err := base64.StdEncoding.DecodeString(changeBeneficiaryProposeParams)
	require.NoError(t, err)
	rawReturn, err := base64.StdEncoding.DecodeString(changeBeneficiaryProposeReturn)
	require.NoError(t, err)
	from, err := address.NewFromString(changeBeneficiaryProposeFrom)
	require.NoError(t, err)
	to, err := address.NewFromString(changeBeneficiaryProposeTo)
	require.NoError(t, err)
	msgCid := cid.MustParse(changeBeneficiaryProposeCid)

	metadata, _, err := msigActor.Parse(t.Context(), network, changeBeneficiaryProposeHeight, parser.MethodPropose, &parser.LotusMessage{
		Cid:    msgCid,
		From:   from,
		To:     to,
		Method: abi.MethodNum(2),
		Params: rawParams,
	}, &parser.LotusMessageReceipt{ExitCode: exitcode.Ok, Return: rawReturn}, msgCid, filTypes.EmptyTSK, true)
	require.NoError(t, err)

	// tx_metadata as stored
	txMetadata, err := json.Marshal(metadata)
	require.NoError(t, err)
	require.Contains(t, string(txMetadata), `"Method":"ChangeBeneficiary"`)
	require.Contains(t, string(txMetadata), maxInt64Expiration, string(txMetadata))
	require.Contains(t, string(txMetadata), maxQuota, string(txMetadata))
	require.Contains(t, string(txMetadata), `"NewBeneficiary":"f02117153"`, string(txMetadata))
	t.Log(string(txMetadata))
	require.NotContains(t, string(txMetadata), roundedExpiration)

	// multisig_proposals value, built from the stored tx_metadata
	eg := toolsMultisig.NewEventGenerator(helper, logger.NewDevelopmentLogger(), metrics.NewNoopMetricsClient(), parser.Config{})
	events, err := eg.GenerateMultisigEvents(t.Context(), []*types.Transaction{{
		TxBasicBlockData: types.TxBasicBlockData{
			BasicBlockData: types.BasicBlockData{Height: uint64(changeBeneficiaryProposeHeight)},
		},
		TxCid:         changeBeneficiaryProposeCid,
		TxFrom:        changeBeneficiaryProposeFrom,
		TxTo:          changeBeneficiaryProposeTo,
		TxType:        parser.MethodPropose,
		TxMetadata:    string(txMetadata),
		Status:        tools.GetExitCodeOK(),
		SubcallStatus: tools.GetExitCodeOK(),
	}}, "tipset", filTypes.EmptyTSK)
	require.NoError(t, err)
	require.Len(t, events.Proposals, 1)
	proposal := events.Proposals[0]
	require.Equal(t, int64(11), proposal.ProposalID) // TxnID of the receipt
	require.Equal(t, parser.MethodChangeBeneficiary, proposal.TxTypeToExecute)
	require.Contains(t, proposal.Value, maxInt64Expiration, proposal.Value)
	require.Contains(t, proposal.Value, maxQuota, proposal.Value)
	require.False(t, strings.Contains(proposal.Value, roundedExpiration), proposal.Value)
}

// The v1 multisig metadata readers decode stored tx_metadata into maps; integers above 2^53
// must come back exactly (float64 rounds 9007199254740993 to 9007199254740992 and MaxInt64 to
// 2^63, which no longer fits an int64).
func TestActorParserV1_ParseMultisigMetadataKeepsLargeIntegers(t *testing.T) {
	tests := []struct {
		name     string
		txType   string
		metadata string
		want     string
	}{
		{
			name:     "constructor",
			txType:   parser.MethodConstructor,
			metadata: `{"Params":{"Signers":["f1h6s3zv3wb6g4ie7wiyqezcem5pdhmkpgqcboosy"],"NumApprovalsThreshold":1,"UnlockDuration":9007199254740993,"StartEpoch":9223372036854775807}}`,
			want:     `{"Signers":["f1h6s3zv3wb6g4ie7wiyqezcem5pdhmkpgqcboosy"],"NumApprovalsThreshold":1,"UnlockDuration":9007199254740993,"StartEpoch":9223372036854775807}`,
		},
		{
			name:     "constructor small values unchanged",
			txType:   parser.MethodConstructor,
			metadata: `{"Params":{"Signers":["f1h6s3zv3wb6g4ie7wiyqezcem5pdhmkpgqcboosy"],"NumApprovalsThreshold":2,"UnlockDuration":5600000000000000,"StartEpoch":-1}}`,
			want:     `{"Signers":["f1h6s3zv3wb6g4ie7wiyqezcem5pdhmkpgqcboosy"],"NumApprovalsThreshold":2,"UnlockDuration":5600000000000000,"StartEpoch":-1}`,
		},
		{
			name:     "change threshold",
			txType:   parser.MethodChangeNumApprovalsThreshold,
			metadata: `{"Params":{"NewThreshold":9007199254740993}}`,
			want:     `{"NewThreshold":9007199254740993}`,
		},
		{
			name:     "lock balance",
			txType:   parser.MethodLockBalance,
			metadata: `{"Params":{"StartEpoch":9223372036854775807,"UnlockDuration":9007199254740993,"Amount":"340282366920938463463374607431768211455"}}`,
			want:     `{"StartEpoch":9223372036854775807,"UnlockDuration":9007199254740993,"Amount":"340282366920938463463374607431768211455"}`,
		},
		{
			name:     "cancel",
			txType:   parser.MethodCancel,
			metadata: `{"Params":"{\"ID\":9007199254740993,\"ProposalHash\":null}"}`,
			want:     `{"ID":9007199254740993,"ProposalHash":""}`,
		},
		{
			name:     "approve exit code",
			txType:   parser.MethodApprove,
			metadata: `{"Params":"{\"ID\":924,\"ProposalHash\":null}","Return":{"Applied":true,"Code":16,"Ret":""}}`,
			want:     `{"ID":924,"ProposalHash":"","Return":{"Applied":true,"Code":16,"Ret":""}}`,
		},
		{
			name:     "send",
			txType:   parser.MethodSend,
			metadata: `{"Params":{"NewExpiration":9223372036854775807,"Small":1,"Fraction":0.5}}`,
			want:     `{"Params":{"Fraction":0.5,"NewExpiration":9223372036854775807,"Small":1}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := actorsV1.ParseMultisigMetadata(tt.txType, tt.metadata)
			require.NoError(t, err)
			gotJSON, err := json.Marshal(got)
			require.NoError(t, err)
			require.Equal(t, tt.want, string(gotJSON))
		})
	}
}
