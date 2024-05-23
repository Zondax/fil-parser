package event_tools

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/zondax/fil-parser/parser"
	"strings"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/multiformats/go-multicodec"
	"github.com/zondax/fil-parser/parser/helper"
	"github.com/zondax/fil-parser/tools"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
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

func ParseNativeLog(tipset *types.ExtendedTipSet, actorEvent *filTypes.ActorEvent, logIndex uint64) (*types.Event, error) {
	event := &types.Event{}
	event.TxCid = actorEvent.MsgCid.String()
	event.Height = uint64(tipset.Height())
	event.TipsetCid = tipset.GetCidString()
	event.Reverted = actorEvent.Reverted
	event.Emitter = actorEvent.Emitter.String()
	event.EventTimestamp = parser.GetTimestamp(tipset.MinTimestamp())
	addr, err := address.NewFromString(event.Emitter)
	if err != nil {
		return nil, err
	}
	var metaData string
	if addr.Protocol() == address.Delegated {
		// this is an evm compatible address
		event.Type = types.EventTypeEVM
		// if the native event is of type evm, the topics are encoded as entries with keys=t1..t4 ( topics ) and key=d ( data )
		parsedEntries, err := parseNativeEventEntry(event.Type, actorEvent.Entries)
		if err != nil {
			return nil, fmt.Errorf("error parsing native evm event entries: %w", err)
		}
		// the first item, t1 contains the selector_hash
		var selectorHash string
		if parsedEntries[0] != nil {
			var ok bool
			selectorHash, ok = parsedEntries[0][parsedEntryValue].(string)
			if !ok {
				return nil, fmt.Errorf("unable to retrieve %s from event entries", EVMTopic0EventEntryKey)
			}
		}
		event.SelectorID = selectorHash
		// retrieve the EVM topics and data to build the metadata object
		var (
			data   []byte
			topics []string
		)
		// maintain the order of topics
		for i := 0; i < len(parsedEntries); i++ {
			k, _ := parsedEntries[i][parsedEntryKey].(string)
			v := parsedEntries[i][parsedEntryValue]
			if strings.HasPrefix(k, EVMTopicPrefixEventEntryKey) { // topic
				val, _ := v.(string)
				topics = append(topics, val)
			}
			if k == EVMDataEventEntryKey { // data
				data, _ = v.([]byte)
			}
		}

		// we store the evm event metadata in the same format as if the event was parsed from an ethLog
		metaDataBytes, err := buildEVMEventMetaData[string](data, topics)
		if err != nil {
			return nil, fmt.Errorf("error building native evm event metadata %w", err)
		}
		metaData = string(metaDataBytes)

	} else {
		event.Type = types.EventTypeNative
		parsedEntries, err := parseNativeEventEntry(event.Type, actorEvent.Entries)
		if err != nil {
			return nil, fmt.Errorf("error parsing native event entries: %w", err)
		}

		var eventType datamodel.Node
		if parsedEntries[0] != nil {
			var ok bool
			eventType, ok = parsedEntries[0][parsedEntryValue].(datamodel.Node)
			if !ok {
				return nil, fmt.Errorf("unable to retrieve %s from event entries", NativeTypeEventEntryKey)
			}
		}

		metaDataBytes, err := json.Marshal(parsedEntries)
		if err != nil {
			return nil, fmt.Errorf("error marshalling parsedEntries to JSON: %w", err)
		}
		metaData = string(metaDataBytes)

		if eventType != nil {
			event.SelectorID, err = eventType.AsString()
			if err != nil {
				return nil, fmt.Errorf("error converting %s to string: %w", NativeTypeEventEntryKey, err)
			}
		}

	}

	event.Metadata = metaData
	event.SelectorSig = genFVMSelectorSig(actorEvent)
	event.LogIndex = logIndex
	event.ID = tools.BuildId(event.TipsetCid, event.TxCid, fmt.Sprint(event.LogIndex), event.Type)
	return event, nil
}

func ParseEthLog(tipset *types.ExtendedTipSet, ethLog types.EthLog, helper *helper.Helper) (*types.Event, error) {
	event := &types.Event{}
	event.TxCid = ethLog.TransactionCid
	event.Emitter = ethLog.Address.String()
	event.LogIndex = uint64(ethLog.LogIndex)
	event.Height = uint64(tipset.Height())
	event.TipsetCid = tipset.GetCidString()
	event.EventTimestamp = parser.GetTimestamp(tipset.MinTimestamp())
	event.SelectorID = extractSelectorIDFromTopics(ethLog.Topics)

	if event.SelectorID != "" {
		var err error
		event.SelectorSig, err = helper.GetEVMSelectorSig(context.Background(), event.SelectorID)
		if err != nil {
			zap.S().Errorf("error retrieving selector_sig for hash: %s err: %s", event.SelectorID, err)
		}
	} else {
		zap.S().Debugf("empty selector_id for event: %v", *event)
	}

	metaDataBytes, err := buildEVMEventMetaData[ethtypes.EthHash](ethLog.Data, ethLog.Topics)
	if err != nil {
		return nil, fmt.Errorf("error marshalling ethLog metadata: %w", err)
	}

	event.Metadata = string(metaDataBytes)
	event.Reverted = ethLog.Removed
	event.Type = types.EventTypeEVM
	event.ID = tools.BuildId(event.TipsetCid, event.TxCid, fmt.Sprint(event.LogIndex), event.Type)

	return event, nil

}

func genFVMSelectorSig(event *filTypes.ActorEvent) string {
	// TODO: generate signature
	return ""
}

// extractSelectorIDFromTopics extracts the selector_hash from a list of topics of an event.
func extractSelectorIDFromTopics(topics []ethtypes.EthHash) string {
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
