package parser

import (
	"encoding/json"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/types"
	"strings"
	"time"
)

func GetExitCodeStatus(exitCode exitcode.ExitCode) string {
	code := exitCode.String()
	status := strings.Split(code, "(")
	if len(status) == 2 {
		return status[0]
	}
	return CheckExitCodeCommonError(code)
}

func parseMetadata(key string, metadata map[string]interface{}) string {
	params, ok := metadata[key].(string)
	if ok && params != "" {
		return params
	}
	jsonMetadata, err := json.Marshal(metadata[key])
	if err == nil && string(jsonMetadata) != "null" && string(jsonMetadata) != "\"\"" {
		return string(jsonMetadata)
	}
	return ""
}

func ParseParams(metadata map[string]interface{}) string {
	return parseMetadata(ParamsKey, metadata)
}

func ParseReturn(metadata map[string]interface{}) string {
	return parseMetadata(ReturnKey, metadata)
}

func GetTimestamp(timestamp uint64) time.Time {
	blockTimeStamp := int64(timestamp) * 1000
	return time.Unix(blockTimeStamp/1000, blockTimeStamp%1000)
}

func AppendToAddressesMap(addressMap types.AddressInfoMap, info ...*types.AddressInfo) {
	if addressMap == nil {
		return
	}

	for _, i := range info {
		switch i.ActorType {
		case manifest.MultisigKey:
			// with multisig accounts we can skip checking for robust addresses because some
			// addresses do not have a robust address (genesis addresses)
			if i.Short != "" {
				if _, ok := addressMap[i.Short]; !ok {
					addressMap[i.Short] = i
				}
			}
		default:
			if i.Robust != "" && i.Short != "" && i.Robust != i.Short {
				if _, ok := addressMap[i.Short]; !ok {
					addressMap[i.Short] = i
				}
			}
		}
	}
}
