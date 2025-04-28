package types

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"

	blocks "github.com/ipfs/go-block-format"
	ipldcbor "github.com/ipfs/go-ipld-cbor"
	ipldformat "github.com/ipfs/go-ipld-format"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

type HandleFilecoinMethodParams struct {
	AbiBytes []byte `json:"-"`

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

	// t.Method (uint64) (uint64)
	// 2. Parse t.Method (uint64) - Encoded within the First Element (Byte String)
	{
		// // Read the header for the first element, expecting a byte string
		// maj, methodLen, err := cr.ReadHeader()
		// if err != nil {
		// 	return fmt.Errorf("failed to read header for Method element: %w", err)
		// }

		// // Check if it's a byte string as confirmed by testing
		// if maj != cbg.MajByteString {
		// 	return fmt.Errorf("wrong type for Method element: expected byte string (Maj %d), got %d", cbg.MajByteString, maj)
		// }

		// Read the raw bytes containing the Method number
		methodBytes, err := cbg.ReadByteArray(cr, cbg.ByteArrayMaxLen)
		if err != nil {
			return fmt.Errorf("failed to read bytes for Method element: %w", err)
		}
		// cr2 := cbg.NewCborReader(bytes.NewReader(methodBytes))

		// tmp, err := cbg.ReadByteArray(cr2, cbg.ByteArrayMaxLen)
		// if err != nil {
		// 	return fmt.Errorf("failed to read bytes for Method element: %w", err)
		// }

		h.AbiBytes = methodBytes
		// tmp := v10Market.DealProposal{}
		// err = tmp.UnmarshalCBOR(bytes.NewReader(methodBytes))
		// if err != nil {
		// 	return fmt.Errorf("failed to unmarshal Method DEAL PROPOSAL element: %w", err)
		// }
		// x, _ := json.Marshal(tmp)
		// fmt.Println("TMP", string(x))
		// h.Method = tmp.Proposal.PieceRef.Piece
		block := blocks.NewBlock(methodBytes)
		node, err := ipldformat.Decode(block, ipldcbor.DecodeBlock)
		if err != nil {
			return fmt.Errorf("failed to decode block: %w", err)
		} else {
			fmt.Println("node: ", node.Loggable())
		}

		h.Args = node

		fmt.Printf("Parsed Method from bytes: %d\n", h.Method) // Optional: Debug print
	}

	// t.Args (*blocks.Block) (struct)
	{
		maj, _, err := cr.ReadHeader()
		if err != nil {
			return fmt.Errorf("failed to read header for Args element: %w", err)
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for Args element: expected tag (Maj %d), got %d", cbg.MajUnsignedInt, maj)
		}

		fmt.Println("method: ", extra)
		h.Method = extra
		// h.Args = node.
	}

	return nil

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

	// t.Result (*blocks.Block) (struct)
	{
		if err := ipldcbor.DecodeReader(cr, h.Result); err != nil {
			return xerrors.Errorf("unmarshaling t.Result: %w", err)
		}
		fmt.Println("RESULT: ", h.Result.String())
	}

	return nil

}
func DecodeFilecoinMethodInput(data []byte) (selector []byte, method []byte, codec []byte, paramsOffset []byte, paramsLen []byte, params []byte) {
	// Read selector (first 4 bytes)
	selector = data[:4]

	// Read method (last 8 bytes of the 32-byte word after selector)
	method = data[4+24 : 4+32]

	// // Read codec (last 8 bytes of the next 32-byte word)
	codec = data[4+32+24 : 4+64]

	// // Read params offset (last 8 bytes of the next 32-byte word)
	// paramsOffset = data[4+64+24 : 4+96]

	// // Read params length (last 8 bytes of the next 32-byte word)
	// paramsLen = data[4+96+24 : 4+128]

	// // Extract params
	// paramsLenValue := int(data[4+96+31]) // Use the last byte for simplicity
	// params = data[4+128 : 4+128+paramsLenValue]

	return
}

func DecodeHex(hexData string) {
	data, _ := hex.DecodeString(hexData)

	// data, _ := hex.DecodeString(hexData)

	// 1. Function selector (4 bytes)
	selector := data[:4]
	fmt.Printf("Function selector: 0x%x\n", selector)

	// 2. Method (32 bytes)
	methodWord := data[4:36]
	methodValue := binary.BigEndian.Uint64(methodWord[24:32])
	fmt.Printf("Method: %d\n", methodValue)

	// 3. Codec (32 bytes) - interpreted as a CID
	codecWord := data[36:68]
	fmt.Printf("Codec (raw): 0x%x\n", codecWord)

	// 4. Params offset (32 bytes)
	offsetWord := data[68:100]
	offsetValue := binary.BigEndian.Uint64(offsetWord[24:32])
	fmt.Printf("Params offset: %d\n", offsetValue)

	// 5. Params length (32 bytes)
	lengthWord := data[100:132]
	lengthValue := binary.BigEndian.Uint64(lengthWord[24:32])
	fmt.Printf("Params length: %d\n", lengthValue)

	// 6. Parameter data
	if len(data) > 132 {
		paramStart := 4 + int(offsetValue)
		paramEnd := paramStart + int(lengthValue)

		if paramEnd > len(data) {
			paramEnd = len(data)
		}

		if paramStart < len(data) {
			paramData := data[paramStart:paramEnd]
			fmt.Printf("Param data: 0x%x\n", paramData)
		}
	}
}

func HexDump(hexData string) {
	data, _ := hex.DecodeString(hexData)
	// Print function selector
	fmt.Printf("Function Selector: 0x%x\n\n", data[:4])

	// Print each 32-byte word with detailed view
	for i := 0; i < len(data[4:])/32; i++ {
		start := 4 + i*32
		end := start + 32
		if end > len(data) {
			end = len(data)
		}

		word := data[start:end]
		fmt.Printf("Word %d (bytes %d-%d):\n", i+1, start, end-1)
		fmt.Printf("  Hex: %x\n", word)

		// Try looking at the last 8 bytes as a uint64 in both endiannesses
		if len(word) >= 8 {
			be := uint64(0)
			le := uint64(0)

			// Big endian
			for j := 0; j < 8; j++ {
				be = (be << 8) | uint64(word[len(word)-8+j])
			}

			// Little endian
			for j := 0; j < 8; j++ {
				le = le | (uint64(word[len(word)-8+j]) << (j * 8))
			}

			fmt.Printf("  Last 8 bytes as BE uint64: %d\n", be)
			fmt.Printf("  Last 8 bytes as LE uint64: %d\n", le)
		}

		fmt.Println()
	}
}
