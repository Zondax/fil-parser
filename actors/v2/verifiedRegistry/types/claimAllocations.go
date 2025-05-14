package types

import (
	"bytes"
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/batch"
	"github.com/filecoin-project/go-state-types/big"
	cbg "github.com/whyrusleeping/cbor-gen"
)

// ClaimAllocationsReturn is implemented in the rust builtin-actors correctly, but the message order is not consistent.
// This struct has custom parsing to allow dynamic positioning of the parameters.
type ClaimAllocationsReturn struct {
	BatchInfo    batch.BatchReturn
	ClaimedSpace []big.Int
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

	// save current position
	curr, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read all bytes: %w", err)
	}
	// check next data type
	buf := new(bytes.Buffer)
	buf.Write(curr)
	maj, _, err = cbg.CborReadHeader(cbg.NewCborReader(buf))
	if err != nil {
		return fmt.Errorf("failed to read header for ClaimAllocationsReturn element: %w", err)
	}

	// reset the reader to the original position
	cr = cbg.NewCborReader(bytes.NewReader(curr))

	// t.ClaimedSpace (big.Int) (struct) or (SectorClaimSummary) (struct)
	t.ClaimedSpace = []big.Int{}
	switch maj {
	case cbg.MajArray:
		_, extra, err := cr.ReadHeader()
		if err != nil {
			return fmt.Errorf("failed to read header for ClaimAllocationsReturn element: %w", err)
		}

		for i := 0; i < int(extra); i++ {
			var claimedSpace big.Int
			if err := claimedSpace.UnmarshalCBOR(cr); err != nil {
				return fmt.Errorf("unmarshaling t.ClaimedSpace: %w", err)
			}
			t.ClaimedSpace = append(t.ClaimedSpace, claimedSpace)
		}
	case cbg.MajByteString:
		var claimedSpace big.Int
		if err := claimedSpace.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.ClaimedSpace: %w", err)
		}
		t.ClaimedSpace = append(t.ClaimedSpace, claimedSpace)
	default:
		return fmt.Errorf("cbor input had wrong type")
	}

	return nil
}
