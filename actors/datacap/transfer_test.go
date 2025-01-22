package datacap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/datacap"
	"github.com/zondax/fil-parser/tools"
)

func TestTransferExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "TransferExported", expected)
	require.NoError(t, err)

	runTest(t, datacap.TransferExported, tests)
}
