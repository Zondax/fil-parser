package multisig_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors/multisig"
	"github.com/zondax/fil-parser/tools"
)

func TestAddVerifierValue(t *testing.T) {
	tests, err := tools.LoadTestData[map[string]any]("AddVerifierValue", expected)
	require.NoError(t, err)
	runValueTest(t, multisig.AddVerifierValue, tests)
}
