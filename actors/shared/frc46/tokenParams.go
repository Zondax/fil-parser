package frc46

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
)

// FRC46TokenParams represents the payload for FRC46 token operations
type FRC46TokenParams struct {
	From         uint64          `json:"from" cborgen:"from"`
	To           uint64          `json:"to" cborgen:"to"`
	Operator     uint64          `json:"operator" cborgen:"operator"`
	Amount       abi.TokenAmount `json:"amount" cborgen:"amount"`
	OperatorData []byte          `json:"operator_data" cborgen:"operator_data"`
	TokenData    []byte          `json:"token_data" cborgen:"token_data"`
}

// FRC46TokenParams CBOR marshalling
func (ftp *FRC46TokenParams) MarshalCBOR(w io.Writer) error {
	if ftp == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, 6); err != nil {
		return err
	}

	// Write From
	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, ftp.From); err != nil {
		return fmt.Errorf("marshaling From: %w", err)
	}

	// Write To
	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, ftp.To); err != nil {
		return fmt.Errorf("marshaling To: %w", err)
	}

	// Write Operator
	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, ftp.Operator); err != nil {
		return fmt.Errorf("marshaling Operator: %w", err)
	}

	// Write Amount
	if err := ftp.Amount.MarshalCBOR(cw); err != nil {
		return fmt.Errorf("marshaling Amount: %w", err)
	}

	// Write OperatorData
	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(ftp.OperatorData))); err != nil {
		return err
	}
	if _, err := cw.Write(ftp.OperatorData); err != nil {
		return err
	}

	// Write TokenData
	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(ftp.TokenData))); err != nil {
		return err
	}
	if _, err := cw.Write(ftp.TokenData); err != nil {
		return err
	}

	return nil
}

func (ftp *FRC46TokenParams) UnmarshalCBOR(r io.Reader) error {
	*ftp = FRC46TokenParams{}

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

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// Read From
	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	ftp.From = extra

	// Read To
	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	ftp.To = extra

	// Read Operator
	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	ftp.Operator = extra

	// Read Amount
	if err := ftp.Amount.UnmarshalCBOR(cr); err != nil {
		return fmt.Errorf("unmarshaling Amount: %w", err)
	}

	// Read OperatorData
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
	ftp.OperatorData = make([]byte, extra)
	if _, err := io.ReadFull(cr, ftp.OperatorData); err != nil {
		return err
	}

	// Read TokenData
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
	ftp.TokenData = make([]byte, extra)
	if _, err := io.ReadFull(cr, ftp.TokenData); err != nil {
		return err
	}

	return nil
}
