package miner

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"

	miner10 "github.com/filecoin-project/go-state-types/builtin/v10/miner"
	miner11 "github.com/filecoin-project/go-state-types/builtin/v11/miner"
	miner12 "github.com/filecoin-project/go-state-types/builtin/v12/miner"
	miner13 "github.com/filecoin-project/go-state-types/builtin/v13/miner"
	miner14 "github.com/filecoin-project/go-state-types/builtin/v14/miner"
	miner15 "github.com/filecoin-project/go-state-types/builtin/v15/miner"
	miner16 "github.com/filecoin-project/go-state-types/builtin/v16/miner"
	miner8 "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/actors/v2/miner/types"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Miner struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Miner {
	return &Miner{
		logger: logger,
	}
}

func (m *Miner) Name() string {
	return manifest.MinerKey
}

func (*Miner) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

// Implemented in the rust builtin-actors but not the golang version
var initialPledgeMethodNum = abi.MethodNum(nonLegacyBuiltin.MustGenerateFRCMethodNum(parser.MethodInitialPledge))
var maxTerminationFeeMethodNum = abi.MethodNum(nonLegacyBuiltin.MustGenerateFRCMethodNum(parser.MethodMaxTerminationFee))

// Implemented in a fork https://github.com/ipfs-force-community/builtin-actors/blob/99642572098400e6bbdff27c5126714781350fce/actors/miner/src/lib.rs#L131
var movePartitionsMethodNum = abi.MethodNum(33)

func customMethods() map[abi.MethodNum]nonLegacyBuiltin.MethodMeta {
	m := &Miner{}
	return map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
		initialPledgeMethodNum: {
			Name:   parser.MethodInitialPledge,
			Method: m.InitialPledgeExported,
		},
		maxTerminationFeeMethodNum: {
			Name:   parser.MethodMaxTerminationFee,
			Method: m.MaxTerminationFeeExported,
		},
		// missing in go-state-types
		nonLegacyBuiltin.MustGenerateFRCMethodNum(parser.MethodGetBeneficiary): {
			Name:   parser.MethodGetBeneficiary,
			Method: m.GetBeneficiary,
		},
		movePartitionsMethodNum: {
			Name:   parser.MethodMovePartitions,
			Method: m.MovePartitions,
		},

		// these methods appear in unexpected versions
		25: {
			Name:   parser.MethodPreCommitSectorBatch,
			Method: m.PreCommitSectorBatch,
		},
		26: {
			Name:   parser.MethodProveCommitAggregate,
			Method: m.ProveCommitAggregate,
		},
		27: {
			Name:   parser.MethodProveReplicaUpdates,
			Method: m.ProveReplicaUpdates,
		},
	}
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V0.String(): actors.CopyMethods(customMethods(), v1Methods()),
	tools.V1.String(): actors.CopyMethods(customMethods(), v1Methods()),
	tools.V2.String(): actors.CopyMethods(customMethods(), v1Methods()),
	tools.V3.String(): actors.CopyMethods(customMethods(), v1Methods()),

	tools.V4.String(): actors.CopyMethods(customMethods(), v2Methods()),
	tools.V5.String(): actors.CopyMethods(customMethods(), v2Methods()),
	tools.V6.String(): actors.CopyMethods(customMethods(), v2Methods()),
	tools.V7.String(): actors.CopyMethods(customMethods(), v2Methods()),
	tools.V8.String(): actors.CopyMethods(customMethods(), v2Methods()),
	tools.V9.String(): actors.CopyMethods(customMethods(), v2Methods()),

	tools.V10.String(): actors.CopyMethods(customMethods(), v3Methods()),
	tools.V11.String(): actors.CopyMethods(customMethods(), v3Methods()),

	tools.V12.String(): actors.CopyMethods(customMethods(), v4Methods()),
	tools.V13.String(): actors.CopyMethods(customMethods(), v5Methods()),
	tools.V14.String(): actors.CopyMethods(customMethods(), v6Methods()),
	tools.V15.String(): actors.CopyMethods(customMethods(), v7Methods()),
	tools.V16.String(): actors.CopyMethods(customMethods(), miner8.Methods),
	tools.V17.String(): actors.CopyMethods(customMethods(), miner9.Methods),
	tools.V18.String(): actors.CopyMethods(customMethods(), miner10.Methods),
	tools.V19.String(): actors.CopyMethods(customMethods(), miner11.Methods),
	tools.V20.String(): actors.CopyMethods(customMethods(), miner11.Methods),
	tools.V21.String(): actors.CopyMethods(customMethods(), miner12.Methods),
	tools.V22.String(): actors.CopyMethods(customMethods(), miner13.Methods),
	tools.V23.String(): actors.CopyMethods(customMethods(), miner14.Methods),
	tools.V24.String(): actors.CopyMethods(customMethods(), miner15.Methods),
	tools.V25.String(): actors.CopyMethods(customMethods(), miner16.Methods),
}

func (m *Miner) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	actorMethods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return actorMethods, nil
}

func (*Miner) ConfirmUpdateWorkerKey(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &abi.EmptyValue{}, &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) TerminateSectors(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := terminateSectorsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := terminateSectorsReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue(), parser.ParamsKey)
}

func (*Miner) DeclareFaults(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)

	params, ok := declareFaultsParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) DeclareFaultsRecovered(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := declareFaultsRecoveredParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ProveReplicaUpdates(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := proveReplicaUpdatesParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) PreCommitSectorBatch2(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := preCommitSectorBatchParams2[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ProveReplicaUpdates2(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := proveReplicaUpdatesParams2[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), &bitfield.BitField{}, parser.ParamsKey)
}

func (*Miner) ProveReplicaUpdates3(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := proveReplicaUpdates3Params[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := proveReplicaUpdates3Return[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, rawReturn, true, params(), returnValue(), parser.ParamsKey)
}

func (*Miner) ProveCommitAggregate(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := proveCommitAggregateParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) DisputeWindowedPoSt(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := disputeWindowedPoStParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ReportConsensusFault(network string, height int64, rawParams []byte) (map[string]interface{}, error) {

	version := tools.VersionFromHeight(network, height)
	params, ok := reportConsensusFaultParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ChangeBeneficiaryExported(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := changeBeneficiaryParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) GetBeneficiary(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if rawParams != nil {
		metadata[parser.ParamsKey] = base64.StdEncoding.EncodeToString(rawParams)
	}
	beneficiaryReturn, err := getBeneficiaryReturn(network, height, rawReturn)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = beneficiaryReturn
	return metadata, nil
}

func (*Miner) Constructor(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := minerConstructorParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
}

func (*Miner) ApplyRewards(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := applyRewardParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	metadata, err := parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
	if err != nil {
		versions := tools.GetSupportedVersions(network)
		for _, v := range versions {
			params, ok := applyRewardParams[v.String()]
			if !ok {
				continue
			}
			metadata, err = parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
			if err != nil {
				continue
			}
			break
		}
	}
	return metadata, err
}

func (*Miner) OnDeferredCronEvent(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := deferredCronEventParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	metadata, err := parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
	if err != nil {
		versions := tools.GetSupportedVersions(network)
		for _, v := range versions {
			params, ok := deferredCronEventParams[v.String()]
			if !ok {
				continue
			}
			metadata, err = parseGeneric(rawParams, nil, false, params(), &abi.EmptyValue{}, parser.ParamsKey)
			if err != nil {
				continue
			}
			break
		}
	}
	return metadata, err
}

func (*Miner) InitialPledgeExported(network string, height int64, rawReturn []byte) (map[string]interface{}, error) {
	return parseGeneric(rawReturn, nil, true, &types.InitialPledgeReturn{}, &types.InitialPledgeReturn{}, parser.ReturnKey)
}

func (*Miner) MaxTerminationFeeExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, rawReturn, true, &types.MaxTerminationFeeParams{}, &types.MaxTerminationFeeReturn{}, parser.ParamsKey)
}

func (*Miner) MovePartitions(network string, height int64, rawParams []byte) (map[string]interface{}, error) {
	return parseGeneric(rawParams, nil, false, &types.MovePartitionsParams{}, &abi.EmptyValue{}, parser.ParamsKey)
}
