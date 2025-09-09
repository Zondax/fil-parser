package types

import (
	"bytes"
	"fmt"
	"io"
	"math"

	"github.com/filecoin-project/go-state-types/abi"
	v11Market "github.com/filecoin-project/go-state-types/builtin/v11/market"
	v12Market "github.com/filecoin-project/go-state-types/builtin/v12/market"
	v13Market "github.com/filecoin-project/go-state-types/builtin/v13/market"
	v14Market "github.com/filecoin-project/go-state-types/builtin/v14/market"
	v15Market "github.com/filecoin-project/go-state-types/builtin/v15/market"
	v16Market "github.com/filecoin-project/go-state-types/builtin/v16/market"
	v17Market "github.com/filecoin-project/go-state-types/builtin/v17/market"

	"github.com/filecoin-project/go-state-types/batch"
	"github.com/filecoin-project/go-state-types/big"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/zondax/fil-parser/tools"
)

// some sector deals DO NOT deals include the sector number
var customSectorDeals = map[string]func() cbg.CBORUnmarshaler{
	// From V20!
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.SectorDeals) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.SectorDeals) },
	//
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(SectorDeals) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(SectorDeals) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(SectorDeals) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(SectorDeals) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(SectorDeals) },
}

// some sector deals INCLUDE the sector number
var canonicalSectorDeals = map[string]func() cbg.CBORUnmarshaler{
	// From V20!
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.SectorDeals) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.SectorDeals) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.SectorDeals) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.SectorDeals) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.SectorDeals) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.SectorDeals) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(v17Market.SectorDeals) },
}

var verifiedDealInfos = map[string]func() cbg.CBORUnmarshaler{
	// From V20!
	tools.V20.String(): func() cbg.CBORUnmarshaler { return new(v11Market.VerifiedDealInfo) },
	tools.V21.String(): func() cbg.CBORUnmarshaler { return new(v12Market.VerifiedDealInfo) },
	tools.V22.String(): func() cbg.CBORUnmarshaler { return new(v13Market.VerifiedDealInfo) },
	tools.V23.String(): func() cbg.CBORUnmarshaler { return new(v14Market.VerifiedDealInfo) },
	tools.V24.String(): func() cbg.CBORUnmarshaler { return new(v15Market.VerifiedDealInfo) },
	tools.V25.String(): func() cbg.CBORUnmarshaler { return new(v16Market.VerifiedDealInfo) },
	tools.V26.String(): func() cbg.CBORUnmarshaler { return new(v17Market.VerifiedDealInfo) },
}

type SectorDeals struct {
	SectorNumber uint64
	SectorType   abi.RegisteredSealProof
	SectorExpiry abi.ChainEpoch
	DealIDs      []abi.DealID
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

	nextSectorDealsBytes []byte
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

	data, err := io.ReadAll(cr)
	if err != nil {
		return fmt.Errorf("reading t.Sectors: %w", err)
	}

	err = t.customSectorDeal(version, bytes.NewReader(data))
	if err != nil {
		err = t.canonicalSectorDeal(version, bytes.NewReader(data))
	}
	if err != nil {
		return fmt.Errorf("t.Sectors error: %w", err)
	}

	return nil
}

func (t *BatchActivateDealsParams) customSectorDeal(version string, data io.Reader) error {
	cr := cbg.NewCborReader(data)
	maj, extra, err := cr.ReadHeader()
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
		t.Sectors[i] = customSectorDeals[version]()
		err := t.Sectors[i].UnmarshalCBOR(cr)
		if err != nil {
			return fmt.Errorf("unmarshaling version: %s, t.Sectors[%d]: %w", version, i, err)
		}
	}
	if err := t.ComputeCID.UnmarshalCBOR(cr); err != nil {
		return fmt.Errorf("unmarshaling t.ComputeCID: %w", err)
	}
	return nil
}

func (t *BatchActivateDealsParams) canonicalSectorDeal(version string, data io.Reader) error {
	cr := cbg.NewCborReader(data)
	maj, extra, err := cr.ReadHeader()
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
		t.Sectors[i] = canonicalSectorDeals[version]()

		err := t.Sectors[i].UnmarshalCBOR(cr)

		if err != nil {
			return fmt.Errorf("unmarshaling version: %s, t.Sectors[%d]: %w", version, i, err)
		}
	}

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
				return fmt.Errorf("unmarshaling version: %s, t.Activations[%d]: %w", version, i, err)
			}
			// reset reader to next bytes
			if t.Activations[i].nextSectorDealsBytes != nil {
				cr = cbg.NewCborReader(bytes.NewReader(t.Activations[i].nextSectorDealsBytes))
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

	if extra != 3 && extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// read all bytes
	// NOTE: Once we do this, and this structure is part of an array we need to return the remaining bytes to the caller to continue processing.( t.nextSectorDealsBytes )
	data, err := io.ReadAll(cr)
	if err != nil {
		return fmt.Errorf("reading t.NonVerifiedDealSpace: %w", err)
	}
	// read next data type
	cr = cbg.NewCborReader(bytes.NewReader(data))
	maj, _, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	// reset cbor reader to read actual data
	cr = cbg.NewCborReader(bytes.NewReader(data))

	// NonVerifiedDealSpace may not be included
	if maj == cbg.MajByteString {
		// t.NonVerifiedDealSpace (big.Int) (struct)

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

	// we need to read all bytes because checking the header reads bytes from the original buffer
	data, err = io.ReadAll(cr)
	if err != nil {
		return fmt.Errorf("reading t.UnsealedCid: %w", err)
	}

	cr = cbg.NewCborReader(bytes.NewReader(data))
	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj == cbg.MajOther && extra == 22 {
		return t.setRemainingBytes(cr)
	}

	// reset buffer and read the unsealed cid
	cr = cbg.NewCborReader(bytes.NewReader(data))
	// t.UnsealedCid (cid.Cid) (struct)
	{
		if err := t.UnsealedCid.UnmarshalCBOR(cr); err != nil {
			return fmt.Errorf("unmarshaling t.UnsealedCid: %w", err)
		}
	}

	return t.setRemainingBytes(cr)
}
func (t *SectorDealActivation) setRemainingBytes(cr *cbg.CborReader) error {
	// read remaining bytes
	data, err := io.ReadAll(cr)
	if err != nil {
		return fmt.Errorf("reading remaining bytes: %w", err)
	}
	t.nextSectorDealsBytes = data
	return nil
}

func (t *SectorDeals) UnmarshalCBOR(r io.Reader) (err error) {
	*t = SectorDeals{}
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

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.SectorNumber  (uint64)
	{
		maj, extra, err := cr.ReadHeader()
		if err != nil {
			return err
		}
		var extraI uint64
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = uint64(extra)

		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SectorNumber = extraI
	}

	// t.SectorType (abi.RegisteredSealProof) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		if err != nil {
			return err
		}
		var extraI int64
		switch maj {
		case cbg.MajUnsignedInt:
			// Check for positive overflow before conversion
			if extra > uint64(math.MaxInt64) {
				return fmt.Errorf("int64 positive overflow")
			}
			extraI = int64(extra)
		case cbg.MajNegativeInt:
			// Check for negative overflow before conversion
			// We need -1-extra >= MinInt64, which simplifies to extra <= MaxInt64
			if extra > uint64(math.MaxInt64) {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - int64(extra)
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SectorType = abi.RegisteredSealProof(extraI)
	}
	// t.SectorExpiry (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		if err != nil {
			return err
		}
		var extraI int64
		switch maj {
		case cbg.MajUnsignedInt:
			// Check for positive overflow before conversion
			if extra > uint64(math.MaxInt64) {
				return fmt.Errorf("int64 positive overflow")
			}
			extraI = int64(extra)
		case cbg.MajNegativeInt:
			// Check for negative overflow before conversion
			// We need -1-extra >= MinInt64, which simplifies to extra <= MaxInt64
			if extra > uint64(math.MaxInt64) {
				return fmt.Errorf("int64 negative overflow")
			}
			extraI = -1 - int64(extra)
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SectorExpiry = abi.ChainEpoch(extraI)
	}
	// t.DealIDs ([]abi.DealID) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > 8192 {
		return fmt.Errorf("t.DealIDs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.DealIDs = make([]abi.DealID, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				maj, extra, err = cr.ReadHeader()
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.DealIDs[i] = abi.DealID(extra)

			}

		}
	}

	return nil
}
