package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
)

// GetSectorSizeReturn is implemented here because the go-state-types repo does not implement UnmarshalCBOR.
type GetSectorSizeReturn struct {
	SectorSize abi.SectorSize
}

func (t *GetSectorSizeReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = GetSectorSizeReturn{}
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

	// t.SectorSize (uint64) (uint64)
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.SectorSize = abi.SectorSize(extra)

	return nil
}
