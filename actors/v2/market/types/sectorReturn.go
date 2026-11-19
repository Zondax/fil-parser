package types

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
)

// SectorContentChangedReturn is the return of the market actor's SectorContentChanged method.
//
// On the wire every level is serde-transparent, so it is an array with one entry per notified
// sector, in the order of the params, each holding one `accepted` flag per piece of that sector:
//
//	[[bool, ...], [bool, ...], ...]
//
// See builtin-actors actors/miner/src/types.rs (SectorContentChangedReturn, SectorReturn,
// PieceReturn), unchanged from v13 to v19.
//
// The previous decoder (SectorReturn []bool) read only the first sector's array, so every
// sector after the first had no return data.
type SectorContentChangedReturn [][]bool

func (s *SectorContentChangedReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*s = SectorContentChangedReturn{}
	cr := cbg.NewCborReader(r)

	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	maj, sectors, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}
	if sectors > cbg.MaxLength {
		return fmt.Errorf("sector return array too large (%d)", sectors)
	}

	result := make(SectorContentChangedReturn, sectors)
	for i := uint64(0); i < sectors; i++ {
		maj, pieces, err := cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajArray {
			return fmt.Errorf("sector return %d should be of type array", i)
		}
		if pieces > cbg.MaxLength {
			return fmt.Errorf("piece return array too large (%d)", pieces)
		}

		accepted := make([]bool, pieces)
		for j := uint64(0); j < pieces; j++ {
			maj, extra, err := cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("piece return %d/%d should be a bool", i, j)
			}
			switch extra {
			case 21: // cbor true
				accepted[j] = true
			case 20: // cbor false
				accepted[j] = false
			default:
				return fmt.Errorf("invalid bool value in piece return %d/%d: %d", i, j, extra)
			}
		}
		result[i] = accepted
	}

	*s = result
	return nil
}
