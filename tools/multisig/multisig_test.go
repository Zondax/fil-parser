package multisig

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/manifest"
	filApiTypes "github.com/filecoin-project/lotus/api/types"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	cid "github.com/ipfs/go-cid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	filMetrics "github.com/zondax/fil-parser/metrics"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/tools/mocks"
	"github.com/zondax/fil-parser/types"
	"github.com/zondax/golem/pkg/logger"
	metrics2 "github.com/zondax/golem/pkg/metrics"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"
)

// Real mainnet multisig Constructor signers (ID form) and their account keys,
// resolved with Filecoin.StateAccountKey on a mainnet Lotus node.
const (
	signerID1     = "f01712510"
	signerID2     = "f01712508"
	signerRobust1 = "f3tcudjsbi6d3yvei67qhhbck2syrrxbeopq4ukipz2caloifsmghqpmhbop6y2ijjxezfzxsir5p7qas5ji3a"
	signerRobust2 = "f3vuvq3qourlgd27zqfsypkozibjpche53irmqxohatahv36gts3whrajbrlkhjyk72e5d5fpvciewvwacj34a"

	msigAddr       = "f02000000"
	testTipsetCid  = "bafy2bzaceaacubfrkq2wng5vcm53r7u5v5pjjs72s6kgqbm44dyfyik4y2azc"
	testTxCid      = "bafy2bzaceadjahvfnskpsz33ansyymo7uoqm7iyapx5cjpja3327pffvml7je"
	msigActorCodeS = "bafk2bzacecqtwq6hjhj2zy5gwjp3vlz3tf6f3ixbxmtzz7xgzrmmpmb7nbbhm"
)

var robustByID = map[string]string{
	signerID1: signerRobust1,
	signerID2: signerRobust2,
}

func newTestGenerator(t *testing.T, cfg parser.Config) *eventGenerator {
	t.Helper()
	log := logger.NewDevelopmentLogger()
	metrics := filMetrics.NewMetricsClient(metrics2.NewNoopMetrics())

	msigCode, err := cid.Parse(msigActorCodeS)
	require.NoError(t, err)

	node := &mocks.FullNode{}
	node.On("StateNetworkName", mock.Anything).Return(dtypes.NetworkName("mainnet"), nil)
	node.On("StateNetworkVersion", mock.Anything, mock.Anything).Return(filApiTypes.NetworkVersion(16), nil)
	node.On("StateActorCodeCIDs", mock.Anything, mock.Anything).Return(map[string]cid.Cid{
		manifest.MultisigKey: msigCode,
	}, nil)

	cache := &mocks.IActorsCache{}
	cache.On("StoreAddressInfo", mock.Anything).Return(nil)
	cache.On("GetActorCode", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(msigActorCodeS, nil)
	// ID -> robust resolution for the known signers.
	cache.On("GetRobustAddress", mock.Anything, mock.Anything, mock.Anything).Return(
		func(_ context.Context, addr address.Address, _ bool) string { return robustByID[addr.String()] },
		func(_ context.Context, addr address.Address, _ bool) error {
			if _, ok := robustByID[addr.String()]; ok {
				return nil
			}
			return errors.New("not found")
		},
	)
	// Robust signers: no short address known, so they are kept as they are.
	cache.On("GetShortAddress", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("not found"))

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(node)
	h := helper.NewHelper(lib, cache, node, log, metrics)
	require.NotNil(t, h)

	gen, ok := NewEventGenerator(h, log, metrics, cfg).(*eventGenerator)
	require.True(t, ok)
	return gen
}

func constructorTx(signers ...string) *types.Transaction {
	signersJSON, _ := json.Marshal(signers)
	return &types.Transaction{
		TxBasicBlockData: types.TxBasicBlockData{BasicBlockData: types.BasicBlockData{Height: 1000}},
		TxCid:            testTxCid,
		TxFrom:           "f01",
		TxTo:             msigAddr,
		TxType:           parser.MethodConstructor,
		Status:           "Ok",
		SubcallStatus:    "Ok",
		TxMetadata:       `{"Params":{"NumApprovalsThreshold":2,"Signers":` + string(signersJSON) + `,"StartEpoch":0,"UnlockDuration":0}}`,
	}
}

func signersFromValue(t *testing.T, value string) []string {
	t.Helper()
	var out struct {
		NumApprovalsThreshold int
		Signers               []string
		StartEpoch            int
		UnlockDuration        int
	}
	require.NoError(t, json.Unmarshal([]byte(value), &out))
	require.Equal(t, 2, out.NumApprovalsThreshold)
	return out.Signers
}

func TestCreateMultisigInfoConstructorSigners(t *testing.T) {
	tests := []struct {
		name        string
		consolidate bool
		bestEffort  bool
		signers     []string
		want        []string
	}{
		{
			name:        "ID signers are consolidated to robust addresses",
			consolidate: true,
			signers:     []string{signerID1, signerID2},
			want:        []string{signerRobust1, signerRobust2},
		},
		{
			name:        "robust signers stay unchanged",
			consolidate: true,
			signers:     []string{signerRobust1, signerRobust2},
			want:        []string{signerRobust1, signerRobust2},
		},
		{
			name:        "mixed ID and robust signers",
			consolidate: true,
			signers:     []string{signerRobust1, signerID2},
			want:        []string{signerRobust1, signerRobust2},
		},
		{
			name:        "ConsolidateRobustAddress disabled leaves signers unchanged",
			consolidate: false,
			signers:     []string{signerID1, signerID2},
			want:        []string{signerID1, signerID2},
		},
		{
			name:        "unresolvable signer leaves all signers unchanged",
			consolidate: true,
			signers:     []string{signerID1, "f09999999"},
			want:        []string{signerID1, "f09999999"},
		},
		{
			name:        "best effort keeps an unresolvable signer and consolidates the rest",
			consolidate: true,
			bestEffort:  true,
			signers:     []string{signerID1, "f09999999"},
			want:        []string{signerRobust1, "f09999999"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eg := newTestGenerator(t, parser.Config{ConsolidateRobustAddress: tt.consolidate, RobustAddressBestEffort: tt.bestEffort})
			tx := constructorTx(tt.signers...)
			metadata, err := tools.ParseTxMetadata(tx.TxMetadata)
			require.NoError(t, err)

			info, err := eg.createMultisigInfo(context.Background(), tx, testTipsetCid, metadata[parser.ParamsKey])
			require.NoError(t, err)
			require.Equal(t, tt.want, signersFromValue(t, info.Value))
		})
	}
}

func TestCreateMultisigInfoConstructorNonStringSigner(t *testing.T) {
	eg := newTestGenerator(t, parser.Config{ConsolidateRobustAddress: true})
	tx := constructorTx()
	tx.TxMetadata = `{"Params":{"NumApprovalsThreshold":2,"Signers":["f01712510",42],"StartEpoch":0,"UnlockDuration":0}}`
	metadata, err := tools.ParseTxMetadata(tx.TxMetadata)
	require.NoError(t, err)

	info, err := eg.createMultisigInfo(context.Background(), tx, testTipsetCid, metadata[parser.ParamsKey])
	require.NoError(t, err)
	require.JSONEq(t, `{"NumApprovalsThreshold":2,"Signers":["f01712510",42],"StartEpoch":0,"UnlockDuration":0}`, info.Value)
}

func TestGenerateMultisigEventsConstructorSigners(t *testing.T) {
	eg := newTestGenerator(t, parser.Config{ConsolidateRobustAddress: true})
	tx := constructorTx(signerID1, signerID2)

	events, err := eg.GenerateMultisigEvents(context.Background(), []*types.Transaction{tx}, testTipsetCid, filTypes.EmptyTSK)
	require.NoError(t, err)
	require.Len(t, events.MultisigInfo, 1)
	require.JSONEq(t,
		`{"NumApprovalsThreshold":2,"Signers":["`+signerRobust1+`","`+signerRobust2+`"],"StartEpoch":0,"UnlockDuration":0}`,
		events.MultisigInfo[0].Value,
	)
}
