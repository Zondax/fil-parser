package types

import (
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

// InitialPledgeReturn is implemented in the rust builtin-actors but not the golang version
type InitialPledgeReturn struct {
	Amount abi.TokenAmount
}

func (t *InitialPledgeReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = InitialPledgeReturn{}

	cr := cbg.NewCborReader(r)

	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	// t.Amount (big.Int) (struct)
	{

		if err := t.Amount.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Amount: %w", err)
		}

	}
	return nil
}
