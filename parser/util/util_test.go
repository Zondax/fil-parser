package util

import "testing"

func TestIsNil(t *testing.T) {}

func TestGetFuncName(t *testing.T) {

	tests := []struct {
		name string
		i    interface{}
		want string
	}{
		{
			name: "TestIsNil",
			i:    TestIsNil,
			want: "TestIsNil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FuncName(tt.i); got != tt.want {
				t.Errorf("GetFuncName() = %v, want %v", got, tt.want)
			}
		})
	}
}
