package datacap_test

import (
	"testing"

	"github.com/zondax/fil-parser/actors/datacap"
)

func TestBurnExported(t *testing.T) {
	tests := []test{}

	runTest(t, datacap.BurnExported, tests)
}

func TestBurnFromExported(t *testing.T) {
	tests := []test{}

	runTest(t, datacap.BurnFromExported, tests)
}

func TestDestroyExported(t *testing.T) {
	tests := []test{}

	runTest(t, datacap.DestroyExported, tests)
}
