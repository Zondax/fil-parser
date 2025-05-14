package types

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
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
		}
	}

	return nil
}
