package types

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
)

// GetMultiaddrsReturn is implemented incorrectly in the go-state-types
type GetMultiaddrsReturn struct {
	Multiaddrs []byte
}

func (t *GetMultiaddrsReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = GetMultiaddrsReturn{}

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

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	switch maj {
	case cbg.MajByteString:
		t.Multiaddrs, err = io.ReadAll(cr)
		if err != nil {
			return err
		}
	case cbg.MajArray:
		if extra != 1 {
			return fmt.Errorf("cbor input had wrong number of fields")
		}
		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajByteString {
			return fmt.Errorf("cbor input should be of type byte string")
		}
		t.Multiaddrs = make([]byte, extra)
		if _, err := io.ReadFull(cr, t.Multiaddrs); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cbor input had wrong type")
	}

	return nil
}
