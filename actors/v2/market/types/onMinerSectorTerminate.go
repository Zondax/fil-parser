package types

import (
	"fmt"
	"io"
	"math"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type OnMinerSectorsTerminateParams struct {
	Epoch   abi.ChainEpoch
	Sectors bitfield.BitField
}

func (t *OnMinerSectorsTerminateParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = OnMinerSectorsTerminateParams{}

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

	// t.Epoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		if err != nil {
			return err
		}
		var extraI int64
		switch maj {
		case cbg.MajUnsignedInt:
			if extra > uint64(math.MaxInt64) {
				return fmt.Errorf("int64 positive overflow")
			}
			extraI = int64(extra)
		case cbg.MajNegativeInt:
			if extra > uint64(math.MaxInt64) {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - int64(extra)
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Epoch = abi.ChainEpoch(extraI)
	}

	// t.Sectors (bitfield.BitField) (struct)
	{
		if err := t.Sectors.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.Sectors: %w", err)
		}

	}

	return nil
}
