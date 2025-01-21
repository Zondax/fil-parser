package reward

import (
	"bytes"
	"io"

	"github.com/filecoin-project/go-state-types/abi"
	rewardv11 "github.com/filecoin-project/go-state-types/builtin/v11/reward"
	rewardv8 "github.com/filecoin-project/go-state-types/builtin/v8/reward"
	"github.com/zondax/fil-parser/parser"
)

type rewardParams interface {
	UnmarshalCBOR(io.Reader) error
}

func parse[T rewardParams](raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	return metadata, nil
}

func RewardConstructor(height int64, raw []byte) (map[string]interface{}, error) {
	return parse[*abi.StoragePower](raw)
}

func AwardBlockReward(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*rewardv8.AwardBlockRewardParams](raw)
	case 11:
		return parse[*rewardv11.AwardBlockRewardParams](raw)
	}
	return nil, nil
}

func UpdateNetworkKPI(height int64, raw []byte) (map[string]interface{}, error) {
	return parse[*abi.StoragePower](raw)
}

func ThisEpochReward(height int64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return parse[*rewardv8.ThisEpochRewardReturn](raw)
	case 11:
		return parse[*rewardv11.ThisEpochRewardReturn](raw)
	}
	return nil, nil
}
