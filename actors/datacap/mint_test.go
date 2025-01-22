package datacap_test

import (
	"testing"

	"github.com/zondax/fil-parser/actors/datacap"
)

func TestMintExported(t *testing.T) {
	tests := []test{}
	runTest(t, datacap.MintExported, tests)

}
