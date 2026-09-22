package reward

import (
	"bytes"
	"context"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	rewardv19 "github.com/filecoin-project/go-state-types/builtin/v19/reward"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

// nv29Height is any height that tools.VersionFromHeight resolves to V29 on calibration.
const nv29Height = int64(999999999999999)

func mustMarshal(t *testing.T, m cbg.CBORMarshaler) []byte {
	t.Helper()
	var buf bytes.Buffer
	require.NoError(t, m.MarshalCBOR(&buf))
	return buf.Bytes()
}

// TestNV29RewardMethodsRoundTrip proves each new NV29 reward method is wired to the CORRECT
// param type, not merely registered. Registration tests (TestVersionCoverage,
// TestABIMethodNumberToMethodName) pass empty payloads, so they would not notice a method
// pointing at the wrong struct. Here a real payload is encoded with the type the method is
// supposed to accept and pushes it through Parse(). This catches a method wired to a
// structurally different type. It cannot catch a swap between two identical types --
// SetWeightRecordsParams and StepWeightRecordsParams are the same shape upstream -- so it is a
// decode guard, not a full type guard.
func TestNV29RewardMethodsRoundTrip(t *testing.T) {
	require.Equal(t, "V29", tools.VersionFromHeight(tools.CalibrationNetwork, nv29Height).String())

	addr, err := address.NewFromString("f01234")
	require.NoError(t, err)

	tests := []struct {
		txType string
		params cbg.CBORMarshaler
		ret    cbg.CBORMarshaler
	}{
		{parser.MethodSetWeightRecordsExported, &rewardv19.SetWeightRecordsParams{}, nil},
		{parser.MethodStepWeightRecordsExported, &rewardv19.StepWeightRecordsParams{}, nil},
		{parser.MethodRegisterStreamExported, &rewardv19.RegisterStreamParams{}, nil},
		{parser.MethodRemoveStreamExported, &rewardv19.RemoveStreamParams{}, nil},
		{parser.MethodSetDistributionExported, &rewardv19.SetDistributionParams{Writer: addr}, nil},
		{parser.MethodSetSharesExported, &rewardv19.SetSharesParams{}, nil},
		{parser.MethodReplaceAddressExported, &rewardv19.ReplaceAddressParams{OldAddress: addr, NewAddress: addr}, nil},
		{parser.MethodCancelPendingExported, &rewardv19.CancelPendingParams{}, nil},
		{parser.MethodClaimExported, &rewardv19.ClaimParams{Wallets: []address.Address{addr}},
			&rewardv19.ClaimReturn{Amounts: []abi.TokenAmount{big.Zero()}}},
	}

	r := New(nil)
	for _, tt := range tests {
		t.Run(tt.txType, func(t *testing.T) {
			msg := &parser.LotusMessage{Params: mustMarshal(t, tt.params)}
			rct := &parser.LotusMessageReceipt{}
			if tt.ret != nil {
				rct.Return = mustMarshal(t, tt.ret)
			}
			got, _, err := r.Parse(context.Background(), tools.CalibrationNetwork, nv29Height,
				tt.txType, msg, rct, cid.Undef, filTypes.EmptyTSK, true)
			require.NoError(t, err, "%s must decode its own params", tt.txType)
			require.Contains(t, got, parser.ParamsKey)
			if tt.ret != nil {
				require.Contains(t, got, parser.ReturnKey, "%s must decode its return", tt.txType)
			}
		})
	}
}
