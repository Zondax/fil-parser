package miner

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
