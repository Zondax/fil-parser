package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/batch"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

type AllocationsResponse struct {

	// On Calibnet some transactions have 4 paramaters e.g txCid='bafy2bzacecdk5ebxm4husinvyhobs2fzk4jpfh5trlnheg6gba34wcqzff2fu'
	// This field is left for downstream indexers with the extra data if needed.
	ExtraData []byte

	AllocationResults batch.BatchReturn
	ExtensionResults  batch.BatchReturn
	NewAllocations    []uint64
}

func (t *AllocationsResponse) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AllocationsResponse{}

	cr := cbg.NewCborReader(r)

	maj, _, err := cr.ReadHeader()
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

	// allow more parameters
	// if extra != 3 {
	// 	return fmt.Errorf("cbor input had wrong number of fields")
	// }
	_, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}

	tmp := make([]byte, extra)
	cbg.Read(r, tmp)
	t.ExtraData = tmp

	// t.AllocationResults (batch.BatchReturn) (struct)

	if err := t.AllocationResults.UnmarshalCBOR(cr); err != nil {
		return xerrors.Errorf("unmarshaling t.AllocationResults: %w", err)
	}

	// t.ExtensionResults (batch.BatchReturn) (struct)

	if err := t.ExtensionResults.UnmarshalCBOR(cr); err != nil {
		return xerrors.Errorf("unmarshaling t.ExtensionResults: %w", err)
	}

	// t.NewAllocations ([]verifreg.AllocationId) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	// allow more allocations
	// if extra > 8192 {
	// 	return fmt.Errorf("t.NewAllocations: array too large (%d)", extra)
	// }

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.NewAllocations = make([]uint64, extra)
	}

	for i := 0; i < int(extra); i++ {
		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.NewAllocations[i] = extra
	}

	extraData, _ := io.ReadAll(cr)
	t.ExtraData = extraData

	return nil
}
