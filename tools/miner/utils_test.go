package miner

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
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
			got, err := getBigInt(test.params, key, test.canBeNil)
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
			got, err := getInteger[int](test.params, key, test.canBeNil)
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
			got, err := getIntegerSlice[int](test.params, key, test.canBeNil)
			assert.NoError(t, err)
			assert.EqualValues(t, test.want, got)
		})
	}
}

func TestSectorProofToBigInt(t *testing.T) {
	tests := []struct {
		name  string
		proof int64
		want  *big.Int
	}{
		{name: "test 1", proof: 0, want: big.NewInt(2 << 10)},
		{name: "test 2", proof: 1, want: big.NewInt(8 << 20)},
		{name: "test 3", proof: 2, want: big.NewInt(512 << 20)},
		{name: "test 4", proof: 334343434, want: big.NewInt(0)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sectorProofToBigInt(test.proof)
			assert.Equal(t, test.want, got)
		})
	}
}
