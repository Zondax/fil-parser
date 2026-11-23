package v2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zondax/fil-parser/actors/v2/miner"
	"github.com/zondax/fil-parser/actors/v2/reward"
	"github.com/zondax/fil-parser/tools"
)

// nv29Reward lists the nine FRC-42 methods builtin-actors v19 added to the reward actor, and
// nv29Miner the single plain-numbered method it added to the miner actor (number 37).
var (
	nv29Reward = []string{
		"SetWeightRecordsExported", "StepWeightRecordsExported", "RegisterStreamExported",
		"RemoveStreamExported", "SetDistributionExported", "SetSharesExported",
		"ReplaceAddressExported", "CancelPendingExported", "ClaimExported",
	}
	nv29Miner = []string{"UpgradeSectorQuality"}
)

// TestNV29MethodsAppearOnlyAtV29 pins the activation boundary. A new method must be reachable
// at V29 and must NOT be visible at V28 — registering it in a shared table (e.g. miner's
// customMethods(), which is merged into every version from V0) would make it appear to have
// existed for the whole chain history.
func TestNV29MethodsAppearOnlyAtV29(t *testing.T) {
	ctx := context.Background()
	const net = tools.CalibrationNetwork

	// Calibration heights: V28 (FireHorse) activated at 3694534; V29 (Solstice) activates at
	// 4109133 (2026-09-28T12:59:30Z). Methods() takes a tipset height: the tipset AT the upgrade
	// height is still executed by V28 (the migration runs before the next tipset's messages), so
	// the methods switch at upgrade+1.
	const (
		upgrade   = int64(4109133)
		v29Height = upgrade + 1
	)
	require.Equal(t, "V29", tools.VersionFromHeight(net, upgrade).String())
	require.Equal(t, "V28", tools.ExecutedVersionFromHeight(net, upgrade).String())
	require.Equal(t, "V29", tools.ExecutedVersionFromHeight(net, v29Height).String())

	for _, tc := range []struct {
		actor   string
		methods []string
		at      func(int64) (map[string]bool, error)
	}{
		{"reward", nv29Reward, func(h int64) (map[string]bool, error) {
			mm, err := reward.New(nil).Methods(ctx, net, h)
			if err != nil {
				return nil, err
			}
			out := map[string]bool{}
			for _, v := range mm {
				out[v.Name] = true
			}
			return out, nil
		}},
		{"miner", nv29Miner, func(h int64) (map[string]bool, error) {
			mm, err := miner.New(nil).Methods(ctx, net, h)
			if err != nil {
				return nil, err
			}
			out := map[string]bool{}
			for _, v := range mm {
				out[v.Name] = true
			}
			return out, nil
		}},
	} {
		t.Run(tc.actor+"/present at V29", func(t *testing.T) {
			got, err := tc.at(v29Height)
			require.NoError(t, err)
			for _, m := range tc.methods {
				require.True(t, got[m], "%s should expose %s at V29", tc.actor, m)
			}
		})
		for _, h := range []int64{3694534 + 10, upgrade} {
			t.Run(fmt.Sprintf("%s/absent at V28 tipset %d", tc.actor, h), func(t *testing.T) {
				got, err := tc.at(h)
				require.NoError(t, err)
				for _, m := range tc.methods {
					require.False(t, got[m], "%s must NOT expose %s at V28 (NV29 method leaked into history)", tc.actor, m)
				}
			})
		}
	}
}
