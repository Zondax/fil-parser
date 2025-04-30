package types

import (
	"fmt"
	"io"

	v11Market "github.com/filecoin-project/go-state-types/builtin/v11/market"
	v12Market "github.com/filecoin-project/go-state-types/builtin/v12/market"
	v13Market "github.com/filecoin-project/go-state-types/builtin/v13/market"
	v14Market "github.com/filecoin-project/go-state-types/builtin/v14/market"
	v15Market "github.com/filecoin-project/go-state-types/builtin/v15/market"
	v16Market "github.com/filecoin-project/go-state-types/builtin/v16/market"

	"github.com/filecoin-project/go-state-types/batch"
	"github.com/filecoin-project/go-state-types/big"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/tools"
)

var sectorDeals = map[string]func() cbg.CBORUnmarshaler{
	// From V20!
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.SectorDeals) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.SectorDeals) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.SectorDeals) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.SectorDeals) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.SectorDeals) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.SectorDeals) },
}

var verifiedDealInfos = map[string]func() cbg.CBORUnmarshaler{
	// From V20!
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.VerifiedDealInfo) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.VerifiedDealInfo) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.VerifiedDealInfo) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.VerifiedDealInfo) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.VerifiedDealInfo) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.VerifiedDealInfo) },
}

type BatchActivateDealsParams struct {
	version    string
	Sectors    []cbg.CBORUnmarshaler
	ComputeCID cbg.CborBool
}

type BatchActivateDealsResult struct {
	version           string
	ActivationResults batch.BatchReturn
	Activations       []SectorDealActivation
}

func NewBatchActivateDealsParams(version string) *BatchActivateDealsParams {
	return &BatchActivateDealsParams{
		version: version,
	}
}

func NewBatchActivateDealsResult(version string) *BatchActivateDealsResult {
	return &BatchActivateDealsResult{
		version: version,
	}
}

type SectorDealActivation struct {
	version              string
	NonVerifiedDealSpace big.Int
	VerifiedInfos        []cbg.CBORUnmarshaler
	UnsealedCid          cbg.CborCid
}

func (t *BatchActivateDealsParams) UnmarshalCBOR(r io.Reader) (err error) {
	version := t.version
	*t = BatchActivateDealsParams{}
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

	// t.Sectors ([]SectorDeals) (slice)
	{
		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}

		if extra > 8192 {
			return fmt.Errorf("t.Sectors: array too large (%d)", extra)
		}

		if maj != cbg.MajArray {
			return fmt.Errorf("expected cbor array")
		}

		if extra > 0 {
			t.Sectors = make([]cbg.CBORUnmarshaler, extra)
		}

		for i := 0; i < int(extra); i++ {
			t.Sectors[i] = sectorDeals[version]()
			if err := t.Sectors[i].UnmarshalCBOR(cr); err != nil {
				return fmt.Errorf("unmarshaling t.Sectors[%d]: %w", i, err)
			}
		}
	}

	// t.ComputeCID (cbg.CborBool) (bool)
	if err := t.ComputeCID.UnmarshalCBOR(cr); err != nil {
		return fmt.Errorf("unmarshaling t.ComputeCID: %w", err)
	}

	return nil
}

func (t *BatchActivateDealsResult) UnmarshalCBOR(r io.Reader) (err error) {
	version := t.version
	*t = BatchActivateDealsResult{}
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

	// t.ActivationResults (batch.BatchReturn) (struct)
	{
		if err := t.ActivationResults.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.ActivationResults: %w", err)
		}
	}

	// t.Activations ([]SectorDealActivation) (slice)
	{
		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}

		if extra > 8192 {
			return fmt.Errorf("t.Activations: array too large (%d)", extra)
		}

		if maj != cbg.MajArray {
			return fmt.Errorf("expected cbor array")
		}

		if extra > 0 {
			t.Activations = make([]SectorDealActivation, extra)
		}

		for i := 0; i < int(extra); i++ {
			t.Activations[i] = SectorDealActivation{
				version: version,
			}
			if err := t.Activations[i].UnmarshalCBOR(cr); err != nil {
				return fmt.Errorf("unmarshaling t.Activations[%d]: %w", i, err)
			}
		}
	}

	return nil
}

func (t *SectorDealActivation) UnmarshalCBOR(r io.Reader) (err error) {
	version := t.version
	*t = SectorDealActivation{}
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

	// t.NonVerifiedDealSpace (big.Int) (struct)
	{

		if err := t.NonVerifiedDealSpace.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.NonVerifiedDealSpace: %w", err)
		}

	}

	// t.VerifiedInfos ([]VerifiedDealInfo) (slice)
	{
		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}

		if extra > 8192 {
			return fmt.Errorf("t.Inputs: array too large (%d)", extra)
		}

		if maj != cbg.MajArray {
			return fmt.Errorf("expected cbor array")
		}

		if extra > 0 {
			t.VerifiedInfos = make([]cbg.CBORUnmarshaler, extra)
		}

		for i := 0; i < int(extra); i++ {
			t.VerifiedInfos[i] = verifiedDealInfos[version]()
			if err := t.VerifiedInfos[i].UnmarshalCBOR(cr); err != nil {
				return fmt.Errorf("unmarshaling t.VerifiedInfos[%d]: %w", i, err)
			}
		}
	}

	// The UnsealedCID will be a valid CID only when ActivateDealParams.ComputeCID was equal to true
	// The UnsealedCID will be= null when ActivateDealParams.ComputeCID was equal to false
	// null in CBOR: Major = MajOther and extra = 22
	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if maj == cbg.MajOther && extra == 22 {
		return nil
	}

	// t.UnsealedCid (cid.Cid) (struct)
	{
		if err := t.UnsealedCid.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.UnsealedCid: %w", err)
		}
	}

	return nil
}
