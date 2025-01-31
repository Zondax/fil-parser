package actors

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
)

func TestActorParser_parseSend(t *testing.T) {
	p := getActorParser()
	msg, err := deserializeMessage(manifest.MultisigKey, parser.MethodSend)
	require.NoError(t, err)
	require.NotNil(t, msg)
	got := p.parseSend(msg)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
}

func TestActorParser_parseConstructor(t *testing.T) {
	p := getActorParser()
	rawParams, err := loadFile(manifest.AccountKey, parser.MethodConstructor, parser.ParamsKey)
	require.NoError(t, err)
	require.NotNil(t, rawParams)
	got, err := p.parseConstructor(rawParams)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
}

func TestActorParser_parseUnknown(t *testing.T) {
	p := getActorParser()

	tests := []struct {
		name        string
		txType      string
		actor       string
		wantTxType  string
		wantHexOnly bool
		mockAddress string
	}{
		{
			name:        "Regular hex encoding fallback",
			txType:      parser.MethodConstructor,
			actor:       manifest.AccountKey,
			wantTxType:  "",
			wantHexOnly: true,
		},
		{
			name:        "EVM address detection",
			txType:      parser.MethodPropose,
			actor:       manifest.MultisigKey,
			wantTxType:  parser.MethodInvokeEVM,
			wantHexOnly: false,
			mockAddress: "f410f234abc",
		},
		{
			name:        "Regular address detection",
			txType:      parser.MethodProposeExported,
			actor:       manifest.MultisigKey,
			wantTxType:  parser.MethodSend,
			wantHexOnly: false,
			mockAddress: "f1abc123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rawParams, rawReturn []byte
			var err error

			if tt.mockAddress != "" {
				mockParams := struct {
					To     string `json:"To"`
					Value  string `json:"Value"`
					Method int64  `json:"Method"`
				}{
					To:     tt.mockAddress,
					Value:  "1000",
					Method: 0,
				}
				rawParams, err = json.Marshal(mockParams)
				require.NoError(t, err)
			} else {
				rawParams, err = loadFile(tt.actor, tt.txType, parser.ParamsKey)
				require.NoError(t, err)
			}
			require.NotNil(t, rawParams)

			got, err := p.unknownMetadata(rawParams, rawReturn)
			require.NoError(t, err)
			require.NotNil(t, got)

			if tt.wantHexOnly {
				require.Contains(t, got, parser.ParamsKey)
				_, isString := got[parser.ParamsKey].(string)
				require.True(t, isString)
			} else {
				params, ok := got[parser.ParamsKey].(map[string]interface{})
				require.True(t, ok)
				require.NotNil(t, params["To"])
				require.Equal(t, tt.wantTxType, got["TxTypeToExecute"])
			}
		})
	}
}
