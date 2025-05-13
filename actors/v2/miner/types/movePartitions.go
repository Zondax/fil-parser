package types

import (
	"fmt"
	"io"

	"github.com/filecoin-project/go-bitfield"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

// MovePartitionsParams is implemented in a fork https://github.com/ipfs-force-community/builtin-actors/blob/99642572098400e6bbdff27c5126714781350fce/actors/miner/src/lib.rs#L131
// and does not exist in go-state-types,spec-actors or the main builtin-actors repo.
// Exists in multiple txs: [bafy2bzacedynmef6vlfoepcxemq4skigwck7dgoqrkge6v3cu7232qkfc4thq,bafy2bzacecvvi7tkvdqs54sffkcy25iqzmivgoegxuyad7mbu5cme3zi7qbts,...] on calibration.
type MovePartitionsParams struct {
	OrigDeadline uint64
	DestDeadline uint64
	Partitions   bitfield.BitField
}

func (t *MovePartitionsParams) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MovePartitionsParams{}

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

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.OrigDeadline (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.OrigDeadline = uint64(extra)

	}
	// t.DestDeadline (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.DestDeadline = uint64(extra)

	}
	// t.Partitions (bitfield.BitField) (struct)

	{

		if err := t.Partitions.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Partitions: %w", err)
		}

	}
	return nil
}
