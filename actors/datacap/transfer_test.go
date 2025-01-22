package datacap_test

import (
	"testing"

	"github.com/zondax/fil-parser/actors/datacap"
)

func TestTransferExported(t *testing.T) {
	tests := []test{}
	runTest(t, datacap.TransferExported, tests)
}
