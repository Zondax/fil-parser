package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
	"testing"
)

func Test_parseExecActor(t *testing.T) {
	type args struct {
		actor string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "parse fil/9/multisig",
			args: args{
				actor: "fil/9/multisig",
			},
			want: "multisig",
		},
		{
			name: "parse multisig",
			args: args{
				actor: "multisig",
			},
			want: "multisig",
		},
		{
			name: "parse empty string",
			args: args{
				actor: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseExecActor(tt.args.actor); got != tt.want {
				t.Errorf("parseExecActor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActorParser_initWithParamsOrReturn(t *testing.T) {
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
			f:      p.initConstructor,
			key:    parser.ParamsKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawParams, err := loadFile(manifest.InitKey, tt.txType, tt.key)
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

func TestActorParser_exec(t *testing.T) {
	p := getActorParser()
	tests := []struct {
		name   string
		txType string
		f      func(*parser.LotusMessage, []byte) (map[string]interface{}, *types.AddressInfo, error)
	}{
		{
			name:   "Exec",
			txType: parser.MethodExec,
			f:      p.parseExec,
		},
		{
			name:   "Exec4",
			txType: parser.MethodExec4,
			f:      p.parseExec4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rawReturn, err := loadFile(manifest.InitKey, tt.txType, parser.ReturnKey)
			require.NoError(t, err)
			msg, err := deserializeMessage(manifest.InitKey, tt.txType)
			require.NoError(t, err)
			require.NotNil(t, msg)
			got, addr, err := tt.f(msg, rawReturn)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotNil(t, addr)
		})
	}
}
