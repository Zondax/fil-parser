package miner

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJsonEncodedBitfieldToSectorNumbers(t *testing.T) {
	tests := []struct {
		name     string
		bitField []int
		want     []uint64
	}{
		{name: "test 1", bitField: []int{0, 4}, want: []uint64{0, 1, 2, 3}},
		{name: "test 2", bitField: []int{0, 2, 1, 3}, want: []uint64{0, 1, 3, 4, 5}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sectorNumbers, err := jsonEncodedBitfieldToSectorNumbers(test.bitField)
			fmt.Printf("want: %v, got: %v\n", test.want, sectorNumbers)
			require.NoError(t, err)
			assert.EqualValues(t, test.want, sectorNumbers)
		})
	}
}
