package v2

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/multiformats/go-multicodec"
	"github.com/zondax/fil-parser/types"
)

const (
	EVMTopicPrefixEventEntryKey = "t"     // check the lotus function ethLogFromEvent
	EVMTopic0EventEntryKey      = "t1"    // check the lotus function ethLogFromEvent
	EVMDataEventEntryKey        = "d"     // check the lotus function ethLogFromEvent
	NativeTypeEventEntryKey     = "$type" // check lotus actor event documentation

	parsedEntryKey   = "key"
	parsedEntryValue = "value"
	parsedEntryFlags = "flags"
)

// parseNativeEventEntry parses event entries which are represented as [{key: x, value: y}...] into
//
//	{
//		logIndex: { entryKey: x,entryValue: y,entryFlags: 0x03}...
//
// }
func parseNativeEventEntry(eventType string, entries []filTypes.EventEntry) (map[int]map[string]any, error) {
	parsedEntries := map[int]map[string]any{}

	for idx, entry := range entries {
		parsedEntry := map[string]any{
			parsedEntryKey: entry.Key,
		}

		if entry.Flags&0x03 != 0 { // parse the values if both key and value are indexable
			switch eventType {
			case types.EventTypeEVM:
				if entry.Codec != cid.Raw {
					// TODO: log
					// "Built-in actors emit CBOR, and anything else would be invalid anyway" - Lotus Docs
					break
				}
				if entry.Key == EVMDataEventEntryKey {
					parsedEntry[parsedEntryValue] = entry.Value
					break
				}
				var selectorHash ethtypes.EthHash
				if copy(selectorHash[:], entry.Value) == 0 {
					return nil, fmt.Errorf("unable to retrieve %s from evm event entry", entry.Key)
				}
				parsedEntry["value"] = selectorHash.String()
			case types.EventTypeNative:
				if entry.Codec != uint64(multicodec.Cbor) {
					break
				}
				n, err := ipld.Decode(entry.Value, dagcbor.Decode)
				if err != nil {
					return nil, fmt.Errorf("error ipld decode native event: %w ", err)
				}
				parsedEntry[parsedEntryValue] = n
			}
		}
		parsedEntry[parsedEntryFlags] = entry.Flags
		if parsedEntry[parsedEntryValue] == nil {
			parsedEntry[parsedEntryValue] = entry.Value
		}

		parsedEntries[idx] = parsedEntry
	}
	return parsedEntries, nil
}

func genFVMSelectorSig(event *filTypes.ActorEvent) string {
	// TODO: generate signature
	return ""
}

// extractSigFromTopics extracts the selector_hash from a list of topics of an event.
func extractSigFromTopics(topics []ethtypes.EthHash) string {
	if len(topics) > 0 {
		return topics[0].String()
	}
	return ""
}

// buildEVMEventMetaData marshals the data and topics into a JSON object
// the type parameter constraint:  when ethtypes.EthHash is marshalled to JSON, it's String() method is called
func buildEVMEventMetaData[T interface{ string | ethtypes.EthHash }](data []byte, topics []T) ([]byte, error) {
	metaDataBytes, err := json.Marshal(map[string]any{
		"data":   hex.EncodeToString(data),
		"topics": topics,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshalling evm event metadata: %w", err)
	}
	return metaDataBytes, nil
}
