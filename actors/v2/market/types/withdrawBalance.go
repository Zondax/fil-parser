package types

import (
	"bytes"
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type WithdrawBalanceReturn struct {
	AmountWithdrawn abi.TokenAmount
	Raw             []byte
}

func (t *WithdrawBalanceReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = WithdrawBalanceReturn{}

	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	cr := cbg.NewCborReader(bytes.NewReader(data))
	maj, _, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	switch maj {
	case cbg.MajByteString:
		newCr := cbg.NewCborReader(bytes.NewReader(data))
		if err := t.AmountWithdrawn.UnmarshalCBOR(newCr); err != nil {
			t.Raw = data
			return nil
		}
	default:
		return fmt.Errorf("unexpected type: %d", maj)
	}

	return nil
}
