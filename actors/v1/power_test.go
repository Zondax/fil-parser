package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"testing"
)

func TestActorParser_powertWithParamsOrReturn(t *testing.T) {
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
			f:      p.powerConstructor,
			key:    parser.ParamsKey,
		},
		{
			name:   "Current Total Power",
			txType: parser.MethodCurrentTotalPower,
			f:      p.currentTotalPower,
			key:    parser.ReturnKey,
		},
		{
			name:   "Enroll Cron Event",
			txType: parser.MethodEnrollCronEvent,
			f:      p.enrollCronEvent,
			key:    parser.ParamsKey,
		},
		{
			name:   "Submit PoRep For Bulk Verify",
			txType: parser.MethodSubmitPoRepForBulkVerify,
			f:      p.submitPoRepForBulkVerify,
			key:    parser.ParamsKey,
		},
		{
			name:   "Update Claimed Power",
			txType: parser.MethodUpdateClaimedPower,
			f:      p.updateClaimedPower,
			key:    parser.ParamsKey,
		},
		{
			name:   "Update Pledge Total",
			txType: parser.MethodUpdatePledgeTotal,
			f:      p.updatePledgeTotal,
			key:    parser.ParamsKey,
		},
		{
			name:   "Network Raw Power Exported",
			txType: parser.MethodNetworkRawPowerExported,
			f:      p.networkRawPower,
			key:    parser.ReturnKey,
		},
		{
			name:   "Miner Count Exported",
			txType: parser.MethodMinerCountExported,
			f:      p.minerCount,
			key:    parser.ReturnKey,
		},
		{
			name:   "Miner Consensus Count Exported",
			txType: parser.MethodMinerConsensusCountExported,
			f:      p.minerConsensusCount,
			key:    parser.ReturnKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.PowerKey, tt.txType, tt.key)
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

func TestActorParser_powerWithParamsAndReturn(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func([]byte, []byte) (map[string]interface{}, error)
	}{
		{
			name:   "Miner Raw Power Exported",
			txType: parser.MethodMinerRawPowerExported,
			f:      p.minerRawPower,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, rawReturn, err := getParamsAndReturn(manifest.PowerKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, rawParams)
			require.NotNil(t, rawReturn)

			got, err := tt.f(rawParams, rawReturn)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Contains(t, got, parser.ParamsKey, "Params could no be found in metadata")
			require.NotNil(t, got[parser.ParamsKey])
			require.Contains(t, got, parser.ReturnKey, "Return could no be found in metadata")
			require.NotNil(t, got[parser.ReturnKey])
		})
	}
}

func TestActorParser_parseCreateMiner(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		method string
	}{
		{
			name:   "Create Miner",
			method: parser.MethodCreateMiner,
		},
		{
			name:   "Create Miner Exported",
			method: parser.MethodCreateMinerExported,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawReturn, err := loadFile(manifest.PowerKey, tt.method, parser.ReturnKey)
			require.NoError(t, err)
			require.NotNil(t, rawReturn)

			msg, err := deserializeMessage(manifest.PowerKey, tt.method)
			require.NoError(t, err)
			require.NotNil(t, msg)

			got, addr, err := p.parseCreateMiner(msg, rawReturn)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotNil(t, addr)
			require.Contains(t, got, parser.ReturnKey)
		})
	}
}
