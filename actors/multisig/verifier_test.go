package multisig_test

import (
	"testing"

	"github.com/zondax/fil-parser/actors/multisig"
)

func TestAddVerifierValue(t *testing.T) {
	tests := []test{}
	runValueTest(t, multisig.AddVerifierValue, tests)
}
