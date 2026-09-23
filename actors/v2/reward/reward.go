package reward

import (
	"context"
	"fmt"

	"github.com/zondax/golem/pkg/logger"

	"github.com/filecoin-project/go-state-types/abi"
	nonLegacyBuiltin "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/manifest"

	rewardv10 "github.com/filecoin-project/go-state-types/builtin/v10/reward"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv12 "github.com/filecoin-project/go-state-types/builtin/v12/reward"
	rewardv13 "github.com/filecoin-project/go-state-types/builtin/v13/reward"
	rewardv14 "github.com/filecoin-project/go-state-types/builtin/v14/reward"
	rewardv15 "github.com/filecoin-project/go-state-types/builtin/v15/reward"
	rewardv16 "github.com/filecoin-project/go-state-types/builtin/v16/reward"
	rewardv17 "github.com/filecoin-project/go-state-types/builtin/v17/reward"
	rewardv18 "github.com/filecoin-project/go-state-types/builtin/v18/reward"
	rewardv19 "github.com/filecoin-project/go-state-types/builtin/v19/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	rewardv9 "github.com/filecoin-project/go-state-types/builtin/v9/reward"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

type Reward struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Reward {
	return &Reward{
		logger: logger,
	}
}

func (r *Reward) Name() string {
	return manifest.RewardKey
}

func (*Reward) StartNetworkHeight() int64 {
	return tools.V1.Height()
}

var methods = map[string]map[abi.MethodNum]nonLegacyBuiltin.MethodMeta{
	tools.V0.String(): v1Methods(),
	tools.V1.String(): v1Methods(),
	tools.V2.String(): v1Methods(),
	tools.V3.String(): v1Methods(),

	tools.V4.String(): v2Methods(),
	tools.V5.String(): v2Methods(),
	tools.V6.String(): v2Methods(),
	tools.V7.String(): v2Methods(),
	tools.V8.String(): v2Methods(),
	tools.V9.String(): v2Methods(),

	tools.V10.String(): v3Methods(),
	tools.V11.String(): v3Methods(),

	tools.V12.String(): v4Methods(),
	tools.V13.String(): v5Methods(),
	tools.V14.String(): v6Methods(),
	tools.V15.String(): v7Methods(),
	tools.V16.String(): actors.CopyMethods(rewardv8.Methods),
	tools.V17.String(): actors.CopyMethods(rewardv9.Methods),
	tools.V18.String(): actors.CopyMethods(rewardv10.Methods),
	tools.V19.String(): actors.CopyMethods(rewardv11.Methods),
	tools.V20.String(): actors.CopyMethods(rewardv11.Methods),
	tools.V21.String(): actors.CopyMethods(rewardv12.Methods),
	tools.V22.String(): actors.CopyMethods(rewardv13.Methods),
	tools.V23.String(): actors.CopyMethods(rewardv14.Methods),
	tools.V24.String(): actors.CopyMethods(rewardv15.Methods),
	tools.V25.String(): actors.CopyMethods(rewardv16.Methods),
	tools.V26.String(): actors.CopyMethods(rewardv16.Methods),
	tools.V27.String(): actors.CopyMethods(rewardv17.Methods),
	tools.V28.String(): actors.CopyMethods(rewardv18.Methods),
	tools.V29.String(): actors.CopyMethods(rewardv19.Methods),
}

func (r *Reward) Methods(_ context.Context, network string, height int64) (map[abi.MethodNum]nonLegacyBuiltin.MethodMeta, error) {
	version := tools.VersionFromHeight(network, height)
	methods, ok := methods[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return methods, nil
}

func (*Reward) Constructor(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := constructorParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	metadata, err := parse(raw, params(), parser.ParamsKey)
	if err != nil {
		versions := tools.GetSupportedVersions(network)
		for _, v := range versions {
			params, ok = constructorParams[v.String()]
			if !ok {
				continue
			}
			metadata, err = parse(raw, params(), parser.ParamsKey)
			if err == nil {
				break
			}
		}
	}
	return metadata, err
}

func (*Reward) AwardBlockReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := awardBlockRewardParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	return parse(raw, params(), parser.ParamsKey)
}

func (*Reward) UpdateNetworkKPI(network string, height int64, raw []byte) (map[string]interface{}, error) {
	return parse(raw, &abi.StoragePower{}, parser.ParamsKey)
}

func (*Reward) ThisEpochReward(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	returns, ok := thisEpochRewardReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}

	metadata, err := parse(raw, returns(), parser.ReturnKey)
	if err != nil {
		versions := tools.GetSupportedVersions(network)
		for _, v := range versions {
			returns, ok = thisEpochRewardReturn[v.String()]
			if !ok {
				continue
			}
			metadata, err = parse(raw, returns(), parser.ReturnKey)
			if err == nil {
				break
			}
		}
	}
	return metadata, err
}

// The nine methods below were added to the reward actor in builtin-actors v19 (NV29,
// Solstice / FIP-0118), which turns the reward actor into a stream allocator. They are
// FRC-42 dispatched, so their method numbers are name hashes rather than small integers;
// go-state-types v19 already carries them in rewardv19.Methods, so no customMethods()
// entry is needed. Params/returns resolve through per-version maps in params.go, matching
// the dispatch pattern used by the miner actor's v18 FRC-42 methods.

func (*Reward) SetWeightRecordsExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := setWeightRecordsExportedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, params(), parser.ParamsKey)
}

func (*Reward) StepWeightRecordsExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := stepWeightRecordsExportedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, params(), parser.ParamsKey)
}

func (*Reward) RegisterStreamExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := registerStreamExportedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, params(), parser.ParamsKey)
}

func (*Reward) RemoveStreamExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := removeStreamExportedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, params(), parser.ParamsKey)
}

func (*Reward) SetDistributionExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := setDistributionExportedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, params(), parser.ParamsKey)
}

func (*Reward) SetSharesExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := setSharesExportedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, params(), parser.ParamsKey)
}

// ReplaceAddressExported returns ReplaceAddressReturn since builtin-actors v19.0.1: whether the
// old recipient's share moved (AddressReplaced) or the old address had no share
// (OldAddressNotInLedger, in which case no address-replaced event is emitted).
func (*Reward) ReplaceAddressExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := replaceAddressExportedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := replaceAddressExportedReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseWithReturn(rawParams, rawReturn, params(), returnValue())
}

func (*Reward) CancelPendingExported(network string, height int64, raw []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := cancelPendingExportedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parse(raw, params(), parser.ParamsKey)
}

func (*Reward) ClaimExported(network string, height int64, rawParams, rawReturn []byte) (map[string]interface{}, error) {
	version := tools.VersionFromHeight(network, height)
	params, ok := claimExportedParams[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	returnValue, ok := claimExportedReturn[version.String()]
	if !ok {
		return nil, fmt.Errorf("%w: %d", actors.ErrUnsupportedHeight, height)
	}
	return parseWithReturn(rawParams, rawReturn, params(), returnValue())
}
