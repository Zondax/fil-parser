package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type ConstructorParams struct {
	Power abi.StoragePower
}

func (t *ConstructorParams) UnmarshalCBOR(r io.Reader) (err error) {
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

	if maj != cbg.MajArray || extra != 1 {
		return fmt.Errorf("wrong number of fields")
	}

	if err := t.Power.UnmarshalCBOR(cr); err != nil {
		return fmt.Errorf("unmarshaling t.Power: %w", err)
	}

	return nil
}
