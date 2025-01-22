package datacap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/datacap"
	"github.com/zondax/fil-parser/tools"
)

func TestBurnExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "BurnExported", expected)
	require.NoError(t, err)

	runTest(t, datacap.BurnExported, tests)
}

func TestBurnFromExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "BurnFromExported", expected)
	require.NoError(t, err)

	runTest(t, datacap.BurnFromExported, tests)
}

func TestDestroyExported(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any](network, "DestroyExported", expected)
	require.NoError(t, err)

	runTest(t, datacap.DestroyExported, tests)
}
