package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/batch"
	"github.com/filecoin-project/go-state-types/big"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type RemoveExpiredAllocationsReturn struct {
	Considered       []uint64
	Results          batch.BatchReturn
	DataCapRecovered big.Int
}

func (t *RemoveExpiredAllocationsReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = RemoveExpiredAllocationsReturn{}

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

	// t.Considered ([]verifreg.AllocationId) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	// if extra > 8192 {
	// 	return fmt.Errorf("t.Considered: array too large (%d)", extra)
	// }

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Considered = make([]uint64, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Considered[i] = extra

			}

		}
	}
	// t.Results (batch.BatchReturn) (struct)

	{

		if err := t.Results.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.Results: %w", err)
		}

	}
	// t.DataCapRecovered (big.Int) (struct)

	{

		if err := t.DataCapRecovered.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.DataCapRecovered: %w", err)
		}

	}
	return nil
}
