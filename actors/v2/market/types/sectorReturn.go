package types

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
)

type SectorReturn []bool

func (s *SectorReturn) UnmarshalCBOR(r io.Reader) error {
	*s = SectorReturn{}
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

	// first byte tells us the length of the array
	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	*s = make(SectorReturn, extra)
	for i := uint64(0); i < extra; i++ {
		b, err := cr.ReadByte()
		if err != nil {
			return err
		}

		switch b {
		case 0xf5: // cbor true
			(*s)[i] = true
		case 0xf4: // cbor false
			(*s)[i] = false
		default:
			return fmt.Errorf("invalid value: %d", b)
		}
	}

	return nil
}
