package miner

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

var lengthBufInitialPledgeReturn = []byte{130}

type InitialPledgeReturn struct {
	Amount abi.TokenAmount
}

func (t *InitialPledgeReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufInitialPledgeReturn); err != nil {
		return err
	}

	// t.Amount (big.Int) (struct)
	if err := t.Amount.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *InitialPledgeReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = InitialPledgeReturn{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Amount (big.Int) (struct)
	{

		if err := t.Amount.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Amount: %w", err)
		}

	}
	return nil
}
