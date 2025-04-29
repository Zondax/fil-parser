package types

import (
	"fmt"
	"io"

	blocks "github.com/ipfs/go-block-format"
	ipldcbor "github.com/ipfs/go-ipld-cbor"
	ipldformat "github.com/ipfs/go-ipld-format"
	cbg "github.com/whyrusleeping/cbor-gen"
)

type HandleFilecoinMethodParams struct {
	Method uint64
	Args   ipldformat.Node
}

type HandleFilecoinMethodReturn struct {
	AbiBytes []byte `json:"-"`

	Result ipldformat.Node
}

func (h *HandleFilecoinMethodParams) UnmarshalCBOR(r io.Reader) error {
	*h = HandleFilecoinMethodParams{}

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

	peeker := cbg.GetPeeker(cr)
	firstMaj, _, err := cbg.CborReadHeader(peeker)
	if err != nil {
		return fmt.Errorf("failed to read header for Args element: %w", err)
	}

	// attempt parse regardless of parameter order
	var method uint64
	var node ipldformat.Node
	switch firstMaj {
	case cbg.MajByteString:
		node, err = getData(cr)
	case cbg.MajUnsignedInt:
		method, err = getMethod(cr)
	default:
		return fmt.Errorf("unexpected major type: %d", firstMaj)
	}

	if err != nil {
		return fmt.Errorf("failed to parse handleFilecoinMethod params: %w", err)
	}

	h.Method = method
	h.Args = node
	return nil
}

func getData(cr *cbg.CborReader) (ipldformat.Node, error) {
	// Read the raw bytes containing the Method number
	data, err := cbg.ReadByteArray(cr, cbg.ByteArrayMaxLen)
	if err != nil {
		return nil, fmt.Errorf("failed to read bytes for Method element: %w", err)
	}

	block := blocks.NewBlock(data)
	node, err := ipldformat.Decode(block, ipldcbor.DecodeBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to decode block: %w", err)
	}
	return node, nil
}

func getMethod(cr *cbg.CborReader) (uint64, error) {
	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return 0, fmt.Errorf("failed to read header for Args element: %w", err)
	}
	if maj != cbg.MajUnsignedInt {
		return 0, fmt.Errorf("wrong type for Args element: expected tag (Maj %d), got %d", cbg.MajUnsignedInt, maj)
	}
	return extra, nil
}

func (h *HandleFilecoinMethodReturn) UnmarshalCBOR(r io.Reader) error {
	*h = HandleFilecoinMethodReturn{}

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

	node, err := getData(cr)
	if err != nil {
		return fmt.Errorf("failed to parse handleFilecoinMethod return: %w", err)
	}

	h.Result = node
	return nil
}
