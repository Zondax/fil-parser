package common_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zondax/fil-parser/tools/common"
)

var key = "key"

func TestGetBigInt(t *testing.T) {
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
			got, err := common.GetBigInt(test.params, key, test.canBeNil)
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetInteger(t *testing.T) {
	tests := []struct {
		name     string
		params   map[string]interface{}
		canBeNil bool
		want     int
	}{
		{name: "test 1", params: map[string]interface{}{key: float64(1)}, canBeNil: false, want: 1},
		{name: "test 2", params: map[string]interface{}{key: float64(2)}, canBeNil: false, want: 2},
		{name: "test 3", params: map[string]interface{}{}, canBeNil: true, want: 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := common.GetInteger[int](test.params, key, test.canBeNil)
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetIntegerSlice(t *testing.T) {
	tests := []struct {
		name     string
		params   map[string]interface{}
		canBeNil bool
		want     []int
	}{
		{name: "test 1", params: map[string]interface{}{key: []interface{}{float64(1), float64(2)}}, canBeNil: false, want: []int{1, 2}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := common.GetIntegerSlice[int](test.params, key, test.canBeNil)
			assert.NoError(t, err)
			assert.EqualValues(t, test.want, got)
		})
	}
}
