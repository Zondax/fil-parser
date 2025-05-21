package types

import (
	"bytes"
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type WithdrawBalanceReturn struct {
	AmountWithdrawn abi.TokenAmount
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
		cr = cbg.NewCborReader(bytes.NewReader(data))
		if err := t.AmountWithdrawn.UnmarshalCBOR(cr); err != nil {
			rawBytes, err := io.ReadAll(cr)
			if err != nil {
				return err
			}
			tmp := big.PositiveFromUnsignedBytes(rawBytes)
			t.AmountWithdrawn = abi.TokenAmount(tmp)
			return nil
		}
	default:
		return fmt.Errorf("unexpected type: %d", maj)
	}

	return nil
}
