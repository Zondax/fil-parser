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
		return parseGeneric(rawParams, nil, false, &miner15.ExtendSectorExpiration2Params{}, &miner15.ExtendSectorExpiration2Params{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ExtendSectorExpiration2Params{}, &miner14.ExtendSectorExpiration2Params{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ExtendSectorExpiration2Params{}, &miner13.ExtendSectorExpiration2Params{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ExtendSectorExpiration2Params{}, &miner12.ExtendSectorExpiration2Params{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ExtendSectorExpiration2Params{}, &miner11.ExtendSectorExpiration2Params{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ExtendSectorExpiration2Params{}, &miner10.ExtendSectorExpiration2Params{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ExtendSectorExpiration2Params{}, &miner9.ExtendSectorExpiration2Params{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V16)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) PreCommitSector(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.PreCommitSectorParams{}, &miner15.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.PreCommitSectorParams{}, &miner14.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.PreCommitSectorParams{}, &miner13.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.PreCommitSectorParams{}, &miner12.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.PreCommitSectorParams{}, &miner11.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.PreCommitSectorParams{}, &miner10.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.PreCommitSectorParams{}, &miner9.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.PreCommitSectorParams{}, &miner8.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.PreCommitSectorParams{}, &legacyv7.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.PreCommitSectorParams{}, &legacyv6.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.PreCommitSectorParams{}, &legacyv5.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.PreCommitSectorParams{}, &legacyv4.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.PreCommitSectorParams{}, &legacyv3.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.PreCommitSectorParams{}, &legacyv2.PreCommitSectorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.SectorPreCommitInfo{}, &legacyv1.SectorPreCommitInfo{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ProveCommitSector(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ProveCommitSectorParams{}, &miner15.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ProveCommitSectorParams{}, &miner14.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ProveCommitSectorParams{}, &miner13.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ProveCommitSectorParams{}, &miner12.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ProveCommitSectorParams{}, &miner11.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ProveCommitSectorParams{}, &miner10.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ProveCommitSectorParams{}, &miner9.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ProveCommitSectorParams{}, &miner8.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ProveCommitSectorParams{}, &legacyv7.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ProveCommitSectorParams{}, &legacyv6.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ProveCommitSectorParams{}, &legacyv5.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ProveCommitSectorParams{}, &legacyv4.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ProveCommitSectorParams{}, &legacyv3.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.ProveCommitSectorParams{}, &legacyv2.ProveCommitSectorParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.ProveCommitSectorParams{}, &legacyv1.ProveCommitSectorParams{}, parser.ParamsKey)
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
		return parseGeneric(rawParams, nil, false, &miner15.SubmitWindowedPoStParams{}, &miner15.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.SubmitWindowedPoStParams{}, &miner14.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.SubmitWindowedPoStParams{}, &miner13.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.SubmitWindowedPoStParams{}, &miner12.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.SubmitWindowedPoStParams{}, &miner11.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.SubmitWindowedPoStParams{}, &miner10.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.SubmitWindowedPoStParams{}, &miner9.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.SubmitWindowedPoStParams{}, &miner8.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.SubmitWindowedPoStParams{}, &legacyv7.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.SubmitWindowedPoStParams{}, &legacyv6.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.SubmitWindowedPoStParams{}, &legacyv5.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.SubmitWindowedPoStParams{}, &legacyv4.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.SubmitWindowedPoStParams{}, &legacyv3.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.SubmitWindowedPoStParams{}, &legacyv2.SubmitWindowedPoStParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.SubmitWindowedPoStParams{}, &legacyv1.SubmitWindowedPoStParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ConfirmSectorProofsValid(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.AnyIsSupported(network, height, tools.VersionsAfter(tools.V23)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ConfirmSectorProofsParams{}, &miner13.ConfirmSectorProofsParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ConfirmSectorProofsParams{}, &miner12.ConfirmSectorProofsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ConfirmSectorProofsParams{}, &miner11.ConfirmSectorProofsParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ConfirmSectorProofsParams{}, &miner10.ConfirmSectorProofsParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ConfirmSectorProofsParams{}, &miner9.ConfirmSectorProofsParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ConfirmSectorProofsParams{}, &miner8.ConfirmSectorProofsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V15)...):
		return map[string]interface{}{}, fmt.Errorf("%w: %d", actors.ErrInvalidHeightForMethod, height)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) CheckSectorProven(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.CheckSectorProvenParams{}, &miner15.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.CheckSectorProvenParams{}, &miner14.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.CheckSectorProvenParams{}, &miner13.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.CheckSectorProvenParams{}, &miner12.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.CheckSectorProvenParams{}, &miner11.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.CheckSectorProvenParams{}, &miner10.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.CheckSectorProvenParams{}, &miner9.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.CheckSectorProvenParams{}, &miner8.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.CheckSectorProvenParams{}, &legacyv7.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.CheckSectorProvenParams{}, &legacyv6.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.CheckSectorProvenParams{}, &legacyv5.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.CheckSectorProvenParams{}, &legacyv4.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.CheckSectorProvenParams{}, &legacyv3.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.CheckSectorProvenParams{}, &legacyv2.CheckSectorProvenParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.CheckSectorProvenParams{}, &legacyv1.CheckSectorProvenParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) ExtendSectorExpiration(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.ExtendSectorExpirationParams{}, &miner15.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.ExtendSectorExpirationParams{}, &miner14.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.ExtendSectorExpirationParams{}, &miner13.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.ExtendSectorExpirationParams{}, &miner12.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.ExtendSectorExpirationParams{}, &miner11.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.ExtendSectorExpirationParams{}, &miner10.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.ExtendSectorExpirationParams{}, &miner9.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.ExtendSectorExpirationParams{}, &miner8.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.ExtendSectorExpirationParams{}, &legacyv7.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.ExtendSectorExpirationParams{}, &legacyv6.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.ExtendSectorExpirationParams{}, &legacyv5.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.ExtendSectorExpirationParams{}, &legacyv4.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.ExtendSectorExpirationParams{}, &legacyv3.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.ExtendSectorExpirationParams{}, &legacyv2.ExtendSectorExpirationParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.ExtendSectorExpirationParams{}, &legacyv1.ExtendSectorExpirationParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) CompactSectorNumbers(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.CompactSectorNumbersParams{}, &miner15.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.CompactSectorNumbersParams{}, &miner14.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.CompactSectorNumbersParams{}, &miner13.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.CompactSectorNumbersParams{}, &miner12.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.CompactSectorNumbersParams{}, &miner11.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.CompactSectorNumbersParams{}, &miner10.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.CompactSectorNumbersParams{}, &miner9.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.CompactSectorNumbersParams{}, &miner8.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.CompactSectorNumbersParams{}, &legacyv7.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.CompactSectorNumbersParams{}, &legacyv6.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.CompactSectorNumbersParams{}, &legacyv5.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.CompactSectorNumbersParams{}, &legacyv4.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.CompactSectorNumbersParams{}, &legacyv3.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.CompactSectorNumbersParams{}, &legacyv2.CompactSectorNumbersParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.CompactSectorNumbersParams{}, &legacyv1.CompactSectorNumbersParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) CompactPartitions(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.CompactPartitionsParams{}, &miner15.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.CompactPartitionsParams{}, &miner14.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.CompactPartitionsParams{}, &miner13.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.CompactPartitionsParams{}, &miner12.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.CompactPartitionsParams{}, &miner11.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.CompactPartitionsParams{}, &miner10.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.CompactPartitionsParams{}, &miner9.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.CompactPartitionsParams{}, &miner8.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.CompactPartitionsParams{}, &legacyv7.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.CompactPartitionsParams{}, &legacyv6.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.CompactPartitionsParams{}, &legacyv5.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.V12.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv4.CompactPartitionsParams{}, &legacyv4.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V11, tools.V10):
		return parseGeneric(rawParams, nil, false, &legacyv3.CompactPartitionsParams{}, &legacyv3.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V9, tools.V8):
		return parseGeneric(rawParams, nil, false, &legacyv2.CompactPartitionsParams{}, &legacyv2.CompactPartitionsParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.VersionsBefore(tools.V7)...):
		return parseGeneric(rawParams, nil, false, &legacyv1.CompactPartitionsParams{}, &legacyv1.CompactPartitionsParams{}, parser.ParamsKey)
	}
	return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
}

func (*Miner) PreCommitSectorBatch(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	switch {
	case tools.V24.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner15.PreCommitSectorBatchParams{}, &miner15.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.V23.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner14.PreCommitSectorBatchParams{}, &miner14.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.V22.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner13.PreCommitSectorBatchParams{}, &miner13.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.V21.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner12.PreCommitSectorBatchParams{}, &miner12.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.AnyIsSupported(network, height, tools.V19, tools.V20):
		return parseGeneric(rawParams, nil, false, &miner11.PreCommitSectorBatchParams{}, &miner11.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.V18.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner10.PreCommitSectorBatchParams{}, &miner10.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.V17.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner9.PreCommitSectorBatchParams{}, &miner9.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.V16.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &miner8.PreCommitSectorBatchParams{}, &miner8.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.V15.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv7.PreCommitSectorBatchParams{}, &legacyv7.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.V14.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv6.PreCommitSectorBatchParams{}, &legacyv6.PreCommitSectorBatchParams{}, parser.ParamsKey)
	case tools.V13.IsSupported(network, height):
		return parseGeneric(rawParams, nil, false, &legacyv5.PreCommitSectorBatchParams{}, &legacyv5.PreCommitSectorBatchParams{}, parser.ParamsKey)
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
