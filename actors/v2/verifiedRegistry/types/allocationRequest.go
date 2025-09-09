package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

// We need custom parsing for AllocationRequest in NV17 and NV18 because the Provider field in go-state-types is a uint64 but in builtin-actors it is an address.Address
type AllocationRequests[T cbg.CBORUnmarshaler] struct {
	Allocations []AllocationRequest `json:"Allocations"`
	Extensions  []T                 `json:"Extensions"`
}

type AllocationRequest struct {
	Provider   address.Address
	Data       cid.Cid
	Size       abi.PaddedPieceSize
	TermMin    abi.ChainEpoch
	TermMax    abi.ChainEpoch
	Expiration abi.ChainEpoch
}

func (t *AllocationRequests[T]) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AllocationRequests[T]{}

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

	// t.Allocations ([]verifreg.AllocationRequest) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > 8192 {
		return fmt.Errorf("t.Allocations: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Allocations = make([]AllocationRequest, extra)
	}

	for i := 0; i < int(extra); i++ {
		if err := t.Allocations[i].UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Allocations[i]: %w", err)
		}
	}
	// t.Extensions ([]verifreg.ClaimExtensionRequest) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > 8192 {
		return fmt.Errorf("t.Extensions: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Extensions = make([]T, extra)
	}

	for i := 0; i < int(extra); i++ {
		if err := t.Extensions[i].UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Extensions[i]: %w", err)
		}
	}
	return nil
}

func (t *AllocationRequest) UnmarshalCBOR(r io.Reader) (err error) {
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

	// t.Provider (addr.Address) (uint64)

	if err := t.Provider.UnmarshalCBOR(cr); err != nil {
		return xerrors.Errorf("unmarshaling t.Provider: %w", err)

	}

	// t.Data (cid.Cid) (struct)

	c, err := cbg.ReadCid(cr)
	if err != nil {
		return xerrors.Errorf("failed to read cid field t.Data: %w", err)
	}

	t.Data = c

	// t.Size (abi.PaddedPieceSize) (uint64)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Size = abi.PaddedPieceSize(extra)

	// t.TermMin (abi.ChainEpoch) (int64)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	var extraI int64
	switch maj {
	case cbg.MajUnsignedInt:
		//nolint:gosec,G115 // Allowing integer overflow conversion
		extraI = int64(extra)
		if extraI < 0 {
			return fmt.Errorf("int64 positive overflow")
		}
	case cbg.MajNegativeInt:
		//nolint:gosec,G115 // Allowing integer overflow conversion
		extraI = int64(extra)
		if extraI < 0 {
			return fmt.Errorf("int64 negative overflow")
		}
		extraI = -1 - extraI
	default:
		return fmt.Errorf("wrong type for int64 field: %d", maj)
	}

	t.TermMin = abi.ChainEpoch(extraI)

	// t.TermMax (abi.ChainEpoch) (int64)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	switch maj {
	case cbg.MajUnsignedInt:
		//nolint:gosec,G115 // Allowing integer overflow conversion
		extraI = int64(extra)
		if extraI < 0 {
			return fmt.Errorf("int64 positive overflow")
		}
	case cbg.MajNegativeInt:
		//nolint:gosec,G115 // Allowing integer overflow conversion
		extraI = int64(extra)
		if extraI < 0 {
			return fmt.Errorf("int64 negative overflow")
		}
		extraI = -1 - extraI
	default:
		return fmt.Errorf("wrong type for int64 field: %d", maj)
	}

	t.TermMax = abi.ChainEpoch(extraI)

	// t.Expiration (abi.ChainEpoch) (int64)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	switch maj {
	case cbg.MajUnsignedInt:
		//nolint:gosec,G115 // Allowing integer overflow conversion
		extraI = int64(extra)
		if extraI < 0 {
			return fmt.Errorf("int64 positive overflow")
		}
	case cbg.MajNegativeInt:
		//nolint:gosec,G115 // Allowing integer overflow conversion
		extraI = int64(extra)
		if extraI < 0 {
			return fmt.Errorf("int64 negative overflow")
		}
		extraI = -1 - extraI
	default:
		return fmt.Errorf("wrong type for int64 field: %d", maj)
	}

	t.Expiration = abi.ChainEpoch(extraI)

	return nil
}
