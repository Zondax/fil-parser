package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
)

// GetNominalSectorExpirationParams wraps the bare-uint64 SectorNumber parameter.
// The v18 builtin-actors method takes *abi.SectorNumber directly (primitive CBOR uint),
// so we wrap it here to satisfy parseGeneric's struct-with-UnmarshalCBOR interface.
type GetNominalSectorExpirationParams struct {
	SectorNumber abi.SectorNumber
}

func (t *GetNominalSectorExpirationParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = GetNominalSectorExpirationParams{}
	cr := cbg.NewCborReader(r)
	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for SectorNumber: %d", maj)
	}
	t.SectorNumber = abi.SectorNumber(extra)
	return nil
}

// GetNominalSectorExpirationReturn wraps the bare-int ChainEpoch return value.
type GetNominalSectorExpirationReturn struct {
	Epoch abi.ChainEpoch
}

func (t *GetNominalSectorExpirationReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = GetNominalSectorExpirationReturn{}
	cr := cbg.NewCborReader(r)
	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	switch maj {
	case cbg.MajUnsignedInt:
		t.Epoch = abi.ChainEpoch(extra)
	case cbg.MajNegativeInt:
		t.Epoch = abi.ChainEpoch(-int64(extra) - 1)
	default:
		return fmt.Errorf("wrong type for ChainEpoch: %d", maj)
	}
	return nil
}
