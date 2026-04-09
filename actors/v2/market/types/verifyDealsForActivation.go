package types

import (
	"fmt"
	"io"

	v11Market "github.com/filecoin-project/go-state-types/builtin/v11/market"
	cbg "github.com/whyrusleeping/cbor-gen"
	"github.com/zondax/fil-parser/tools"
)

type VerifyDealsForActivationParams struct {
	version string
	// SectorDeals
	Sectors []cbg.CBORUnmarshaler
}

func NewVerifyDealsForActivationParams(version string) *VerifyDealsForActivationParams {
	return &VerifyDealsForActivationParams{
		version: version,
	}
}

func (t *VerifyDealsForActivationParams) UnmarshalCBOR(r io.Reader) (err error) {
	version := t.version
	*t = VerifyDealsForActivationParams{}
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
		return fmt.Errorf("cbor input should be an array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input should have 1 element")
	}

	// t.Sectors ([]SectorDeals) (slice)
	{
		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}

		if extra > 8192 {
			return fmt.Errorf("t.Sectors: array too large (%d)", extra)
		}

		if maj != cbg.MajArray {
			return fmt.Errorf("expected cbor array")
		}

		if extra > 0 {
			t.Sectors = make([]cbg.CBORUnmarshaler, extra)
		}

		for i := 0; i < int(extra); i++ {
			t.Sectors[i] = customSectorDeals[version]()
			if err := t.Sectors[i].UnmarshalCBOR(cr); err != nil {
				return fmt.Errorf("unmarshaling t.Sectors[%d]: %w", i, err)
			}

			// required because of breaking change in go-state-types: https://github.com/filecoin-project/go-state-types/issues/435
			if version == tools.V21.String() {
				tmp, ok := t.Sectors[i].(*SectorDeals)
				if !ok {
					return fmt.Errorf("error handliing VerifyDealsForActivationParams.SectorDeals V21(v12-actors) edge-case")
				}

				// use compatible struct to avoid adding the extra field from SectorDeals.
				t.Sectors[i] = &v11Market.SectorDeals{
					SectorType:   tmp.SectorType,
					SectorExpiry: tmp.SectorExpiry,
					DealIDs:      tmp.DealIDs,
				}
			}
		}
	}

	return nil
}
