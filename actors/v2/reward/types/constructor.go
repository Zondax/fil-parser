package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
)

// ConstructorParams is the parameters for the reward actor constructor.
// https://github.com/filecoin-project/builtin-actors/blob/cd9ac2bb0afcca7a59465e57cee6569e69070d7a/actors/reward/src/lib.rs#L54
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
