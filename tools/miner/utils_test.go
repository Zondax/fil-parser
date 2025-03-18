package miner

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBigInt(t *testing.T) {
	key := "key"
	tests := []struct {
		name     string
		params   map[string]interface{}
		canBeNil bool
		want     *big.Int
	}{
		{name: "test 1", params: map[string]interface{}{key: "1"}, canBeNil: false, want: big.NewInt(1)},
		{name: "test 2", params: map[string]interface{}{key: "2"}, canBeNil: false, want: big.NewInt(2)},
		{name: "test 3", params: map[string]interface{}{key: ""}, canBeNil: true, want: nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := getBigInt(test.params, key, test.canBeNil)
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
