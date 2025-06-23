package frc46

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
)

// TokenType represents the type identifier for FRC46 tokens
type TokenType uint64

// UniversalReceiverParams represents the parameters for the universal receiver hook
type UniversalReceiverParams struct {
	Type    TokenType `json:"type"`
	Payload []byte    `json:"payload"`
}

// UniversalReceiverParams CBOR marshalling
func (urp *UniversalReceiverParams) MarshalCBOR(w io.Writer) error {
	if urp == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, 2); err != nil {
		return err
	}

	// Write Type
	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(urp.Type)); err != nil {
		return err
	}

	// Write Payload
	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(urp.Payload))); err != nil {
		return err
	}
	if _, err := cw.Write(urp.Payload); err != nil {
		return err
	}

	return nil
}

func (urp *UniversalReceiverParams) UnmarshalCBOR(r io.Reader) error {
	*urp = UniversalReceiverParams{}

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

	// Read Type
	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	urp.Type = TokenType(extra)

	// Read Payload
	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte string")
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("byte array too large")
	}
	urp.Payload = make([]byte, extra)
	if _, err := io.ReadFull(cr, urp.Payload); err != nil {
		return err
	}

	return nil
}
