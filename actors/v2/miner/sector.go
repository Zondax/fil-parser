package miner

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	legacyv1 "github.com/filecoin-project/specs-actors/actors/builtin/miner"
	legacyv2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	legacyv3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/miner"
	legacyv4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/miner"
	legacyv5 "github.com/filecoin-project/specs-actors/v5/actors/builtin/miner"
	legacyv6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/miner"
	legacyv7 "github.com/filecoin-project/specs-actors/v7/actors/builtin/miner"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

func (*Miner) ExtendSectorExpiration2(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ExtendSectorExpiration2Params{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ExtendSectorExpiration2Params{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ExtendSectorExpiration2Params{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ExtendSectorExpiration2Params{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ExtendSectorExpiration2Params{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ExtendSectorExpiration2Params{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ExtendSectorExpiration2Params{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) PreCommitSector(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.PreCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.SectorPreCommitInfo{}, &abi.EmptyValue{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveCommitSector(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.ProveCommitSectorParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveCommitSectors3(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner15.ProveCommitSectors3Params{}, &miner15.ProveCommitSectors3Return{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner14.ProveCommitSectors3Params{}, &miner14.ProveCommitSectors3Return{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner13.ProveCommitSectors3Params{}, &miner13.ProveCommitSectors3Return{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V21)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) InternalSectorSetupForPreseal(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner15.InternalSectorSetupForPresealParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, rawReturn, true, &miner14.InternalSectorSetupForPresealParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V23)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) SubmitWindowedPoSt(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.SubmitWindowedPoStParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ConfirmSectorProofsValid(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsAfter(tools.V23)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ConfirmSectorProofsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ConfirmSectorProofsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ConfirmSectorProofsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ConfirmSectorProofsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ConfirmSectorProofsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ConfirmSectorProofsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) CheckSectorProven(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.CheckSectorProvenParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ExtendSectorExpiration(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.ExtendSectorExpirationParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) CompactSectorNumbers(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.CompactSectorNumbersParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) CompactPartitions(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.CompactPartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) PreCommitSectorBatch(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.PreCommitSectorBatchParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V12)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) GetSectorSize(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	return nil, fmt.Errorf("unsupported height: %d", height)
}

func (*Miner) ProveCommitSectorsNI(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ProveCommitSectorsNIParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ProveCommitSectorsNIParams{}, &abi.EmptyValue{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V22)...):
		return nil, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return parseGeneric(rawParams, nil, false, &abi.EmptyValue{}, &abi.EmptyValue{}, parser.ParamsKey)
}
