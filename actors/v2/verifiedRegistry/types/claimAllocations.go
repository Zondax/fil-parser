package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/batch"
	"github.com/filecoin-project/go-state-types/big"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type ClaimAllocationsReturn struct {
	BatchInfo    batch.BatchReturn
	ClaimedSpace big.Int
}

func (t *ClaimAllocationsReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ClaimAllocationsReturn{}

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

	// t.BatchInfo (batch.BatchReturn) (struct)

	{

		if err := t.BatchInfo.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.BatchInfo: %w", err)
		}

	}

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	// t.ClaimedSpace (big.Int) (struct) or (SectorClaimSummary) (struct)
	if maj == cbg.MajArray {
		if extra != 1 {
			return fmt.Errorf("cbor input had wrong number of fields")
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
					if err := t.ClaimedSpace.UnmarshalCBOR(cr); err != nil {
						return fmt.Errorf("unmarshaling t.ClaimedSpace: %w", err)
					}
				}
			}
		}

	} else if maj == cbg.MajByteString {
		if err := t.ClaimedSpace.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.ClaimedSpace: %w", err)
		}
	} else {
		return fmt.Errorf("cbor input had wrong type")
	}

	return nil
}
