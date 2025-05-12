package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-bitfield"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

type MovePartitionsParams struct {
	OrigDeadline uint64
	DestDeadline uint64
	Partitions   bitfield.BitField
}

func (t *MovePartitionsParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MovePartitionsParams{}

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

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.OrigDeadline (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.OrigDeadline = uint64(extra)

	}
	// t.DestDeadline (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.DestDeadline = uint64(extra)

	}
	// t.Partitions (bitfield.BitField) (struct)

	{

		if err := t.Partitions.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Partitions: %w", err)
		}

	}
	return nil
}
