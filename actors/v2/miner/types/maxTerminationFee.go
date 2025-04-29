package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type MaxTerminationFeeParams struct {
	Power         big.Int
	InitialPledge abi.TokenAmount
}

type MaxTerminationFeeReturn struct {
	MaxFee abi.TokenAmount
}

func (t *MaxTerminationFeeParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MaxTerminationFeeParams{}
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

	// t.Power (big.Int) (struct)
	{

		if err := t.Power.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.Power: %w", err)
		}

	}

	// t.InitialPledge (abi.TokenAmount) (struct)
	{

		if err := t.InitialPledge.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.InitialPledge: %w", err)
		}

	}
	return nil
}

func (t *MaxTerminationFeeReturn) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MaxTerminationFeeReturn{}
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

	// t.MaxFee (abi.TokenAmount) (struct)
	{

		if err := t.MaxFee.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.MaxFee: %w", err)
		}

	}
	return nil
}
