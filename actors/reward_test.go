package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"testing"
)

func TestActorParser_rewardWithParamsOrReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte) (map[string]interface{}, error)
		key    string
	}{
		{
			name:   "Constructor",
			txType: parser.MethodConstructor,
			f:      p.rewardConstructor,
			key:    parser.ParamsKey,
		},
		{
			name:   "Award Block Reward",
			txType: parser.MethodAwardBlockReward,
			f:      p.awardBlockReward,
			key:    parser.ParamsKey,
		},
		{
			name:   "This Epoch Reward",
			txType: parser.MethodThisEpochReward,
			f:      p.thisEpochReward,
			key:    parser.ReturnKey,
		},
		{
			name:   "Update Network KPI",
			txType: parser.MethodUpdateNetworkKPI,
			f:      p.updateNetworkKpi,
			key:    parser.ParamsKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.RewardKey, tt.txType, tt.key)
			require.NoError(t, err)
			require.NotNil(t, rawParams)

			got, err := tt.f(rawParams)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, tt.key, fmt.Sprintf("%s could no be found in metadata", tt.key))
			require.NotNil(t, got[tt.key])
		})
	}
}
