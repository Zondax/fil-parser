package parser

import (
	"testing"

	"github.com/filecoin-project/go-state-types/builtin/v11/datacap"
)

func TestParseParams(t *testing.T) {
	tests := []struct {
		name     string
		metadata map[string]interface{}
		want     string
	}{
		{
			name:     "nil metadata",
			metadata: nil,
			want:     "",
		},
		{
			name: "empty params",
			metadata: map[string]interface{}{
				ParamsKey: "",
			},
			want: "",
		},
		{
			name: "string in params",
			metadata: map[string]interface{}{
				ParamsKey: "f1ljefareoomkuplzvk5zkk3cjeq25fjdbs2gwzea",
			},
			want: "f1ljefareoomkuplzvk5zkk3cjeq25fjdbs2gwzea",
		},
		{
			name: "struct in params",
			metadata: map[string]interface{}{
				ParamsKey: Propose{
					To:     "f01656666",
					Value:  "0",
					Method: "",
					Params: nil,
				},
			},
			want: "{\"To\":\"f01656666\",\"Value\":\"0\",\"Method\":\"\",\"Params\":null}",
		},
		{
			name: "cbor",
			metadata: map[string]interface{}{
				ParamsKey: datacap.TransferParams{},
			},
			want: "{\"To\":\"\\u003cempty\\u003e\",\"Amount\":\"\\u003cnil\\u003e\",\"OperatorData\":null}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseParams(tt.metadata); got != tt.want {
				t.Errorf("ParseParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseReturn(t *testing.T) {
	tests := []struct {
		name     string
		metadata map[string]interface{}
		want     string
	}{
		{
			name:     "nil metadata",
			metadata: nil,
			want:     "",
		},
		{
			name: "empty params",
			metadata: map[string]interface{}{
				ReturnKey: "",
			},
			want: "",
		},
		{
			name: "string in params",
			metadata: map[string]interface{}{
				ReturnKey: "f1ljefareoomkuplzvk5zkk3cjeq25fjdbs2gwzea",
			},
			want: "f1ljefareoomkuplzvk5zkk3cjeq25fjdbs2gwzea",
		},
		{
			name: "struct in params",
			metadata: map[string]interface{}{
				ReturnKey: Propose{
					To:     "f01656666",
					Value:  "0",
					Method: "",
					Params: nil,
				},
			},
			want: "{\"To\":\"f01656666\",\"Value\":\"0\",\"Method\":\"\",\"Params\":null}",
		},
		{
			name: "cbor",
			metadata: map[string]interface{}{
				ReturnKey: datacap.TransferParams{},
			},
			want: "{\"To\":\"\\u003cempty\\u003e\",\"Amount\":\"\\u003cnil\\u003e\",\"OperatorData\":null}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseReturn(tt.metadata); got != tt.want {
				t.Errorf("ParseParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
