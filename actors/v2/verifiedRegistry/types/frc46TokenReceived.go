package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type FRC46TokenReceived struct {
	network string
	height  int64

	// From is an Actor ID address (address.NewIDAddress(x))
	From    uint64
	FromStr string

	// To is an Actor ID address (address.NewIDAddress(x))
	To    uint64
	ToStr string

	// Operator is an Actor ID address (address.NewIDAddress(x))
	Operator    uint64
	OperatorStr string

	Amount abi.TokenAmount

	// OperatorData is encoded verifreg.AllocationRequests
	OperatorData []byte

	TokenData []byte
}

func NewFRC46TokenReceived(network string, height int64) *FRC46TokenReceived {
	return &FRC46TokenReceived{
		network: network,
		height:  height,
	}
}

func (f *FRC46TokenReceived) UnmarshalCBOR(r io.Reader) error {
	cr := cbg.NewCborReader(r)

	maj, extra, err := cbg.CborReadHeader(cr)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor reader expected array")
	}
	if extra != 6 {
		return fmt.Errorf("cbor reader expected 6 elements")
	}

	// From (uint64)
	from, err := decodeUint64(cr)
	if err != nil {
		return err
	}
	f.From = from
	f.FromStr, err = decodeActorID(from)
	if err != nil {
		return err
	}

	// To (uint64)
	to, err := decodeUint64(cr)
	if err != nil {
		return err
	}
	f.To = to
	f.ToStr, err = decodeActorID(to)
	if err != nil {
		return err
	}

	// Operator (uint64)
	operator, err := decodeUint64(cr)
	if err != nil {
		return err
	}
	f.Operator = operator
	f.OperatorStr, err = decodeActorID(operator)
	if err != nil {
		return err
	}

	// Amount (uint64)
	if err := f.Amount.UnmarshalCBOR(cr); err != nil {
		return err
	}

	// OperatorData (RawBytes)
	operatorData, err := cbg.ReadByteArray(cr, cbg.ByteArrayMaxLen)
	if err != nil {
		return err
	}

	f.OperatorData = operatorData

	// TokenData (RawBytes)
	tokenData, err := cbg.ReadByteArray(cr, cbg.ByteArrayMaxLen)
	if err != nil {
		return err
	}
	f.TokenData = tokenData

	return nil
}

func decodeActorID(id uint64) (string, error) {
	idAddr, err := address.NewIDAddress(id)
	if err != nil {
		return "", err
	}
	return idAddr.String(), nil
}

func decodeUint64(cr *cbg.CborReader) (uint64, error) {
	maj, extra, err := cbg.CborReadHeader(cr)
	if err != nil {
		return 0, err
	}
	if maj != cbg.MajUnsignedInt {
		return 0, fmt.Errorf("cbor reader expected unsigned int")
	}
	return extra, nil
}
