package parser

import "testing"

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
