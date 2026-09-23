package event_tools

import (
	"encoding/hex"
	"encoding/json"
	"math"
	"testing"

	"github.com/filecoin-project/go-address"
	filBig "github.com/filecoin-project/go-state-types/big"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	cidLink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/zondax/fil-parser/types"
)

const (
	cborCodec = 0x51 // DAG-CBOR, used by builtin actors (runtime/src/util/events.rs)
	rawCodec  = 0x55 // raw, used by the EVM actor (evm/src/interpreter/instructions/log_event.rs)

	flagIndexedKey = 0x02 // Flags::FLAG_INDEXED_KEY, EventBuilder::field
	flagIndexedAll = 0x03 // Flags::FLAG_INDEXED_ALL, EventBuilder::typ / field_indexed
)

// rewardActor is f02, the emitter of the v19 reward events.
var rewardActor, _ = address.NewIDAddress(2)

func cborEncode(t *testing.T, data any) []byte {
	t.Helper()
	var (
		builder datamodel.NodeBuilder
		err     error
	)
	switch x := data.(type) {
	case string:
		builder = basicnode.Prototype.String.NewBuilder()
		err = builder.AssignString(x)
	case []byte:
		builder = basicnode.Prototype.Bytes.NewBuilder()
		err = builder.AssignBytes(x)
	case int64:
		builder = basicnode.Prototype.Int.NewBuilder()
		err = builder.AssignInt(x)
	case datamodel.Link:
		builder = basicnode.Prototype.Link.NewBuilder()
		err = builder.AssignLink(x)
	case nil:
		builder = basicnode.Prototype.Any.NewBuilder()
		err = builder.AssignNull()
	default:
		t.Fatalf("unsupported type %T", data)
	}
	require.NoError(t, err)
	encoded, err := ipld.Encode(builder.Build(), dagcbor.Encode)
	require.NoError(t, err)
	return encoded
}

// cborTokenAmount encodes a TokenAmount the way fvm_shared serializes it: a CBOR byte string
// holding the big-endian magnitude with a leading sign byte.
func cborTokenAmount(t *testing.T, amount string) []byte {
	t.Helper()
	v, err := filBig.FromString(amount)
	require.NoError(t, err)
	b, err := v.Bytes()
	require.NoError(t, err)
	return cborEncode(t, b)
}

func mustHex(t *testing.T, s string) []byte {
	t.Helper()
	b, err := hex.DecodeString(s)
	require.NoError(t, err)
	return b
}

func parseNative(t *testing.T, emitter address.Address, entries []filTypes.EventEntry) (*types.Event, error) {
	t.Helper()
	tipset := &types.ExtendedTipSet{TipSet: filTypes.TipSet{}}
	return ParseNativeLog(tipset, &filTypes.ActorEvent{Emitter: emitter, Entries: entries}, 0, nil)
}

func metadataValues(t *testing.T, event *types.Event) map[string]any {
	t.Helper()
	got := map[int]map[string]any{}
	require.NoError(t, json.Unmarshal([]byte(event.Metadata), &got))
	values := map[string]any{}
	for _, entry := range got {
		values[entry[parsedEntryKey].(string)] = entry[parsedEntryValue]
	}
	return values
}

// TestParseNativeLog_ScalarOutputUnchanged pins the exact metadata JSON for scalar, CID and
// bigint entries so that the list/map and token-amount handling cannot alter it.
func TestParseNativeLog_ScalarOutputUnchanged(t *testing.T) {
	pieceCid, err := cid.Decode("baga6ea4seaqeyz6zikyr2bqbhy6mrocoqwagx45vlbpsbem7euqv5mf3hrvn2fy")
	require.NoError(t, err)

	tb := []struct {
		name    string
		entries []filTypes.EventEntry
		want    string
	}{
		{
			name: "sector-activated",
			entries: []filTypes.EventEntry{
				{Flags: flagIndexedAll, Key: "$type", Codec: cborCodec, Value: cborEncode(t, "sector-activated")},
				{Flags: flagIndexedAll, Key: "sector", Codec: cborCodec, Value: cborEncode(t, int64(10))},
				{Flags: flagIndexedAll, Key: "unsealed-cid", Codec: cborCodec, Value: cborEncode(t, nil)},
				{Flags: flagIndexedAll, Key: "piece-cid", Codec: cborCodec, Value: cborEncode(t, cidLink.Link{Cid: pieceCid})},
				{Flags: flagIndexedKey, Key: "piece-size", Codec: cborCodec, Value: cborEncode(t, int64(math.MaxInt64))},
				{Flags: flagIndexedKey, Key: "neg", Codec: cborCodec, Value: cborEncode(t, int64(-5))},
			},
			want: `{"0":{"flags":3,"key":"$type","value":"sector-activated"},"1":{"flags":3,"key":"sector","value":10},"2":{"flags":3,"key":"unsealed-cid","value":null},"3":{"flags":3,"key":"piece-cid","value":{"/":"baga6ea4seaqeyz6zikyr2bqbhy6mrocoqwagx45vlbpsbem7euqv5mf3hrvn2fy"}},"4":{"flags":2,"key":"piece-size","value":"9223372036854775807"},"5":{"flags":2,"key":"neg","value":-5}}`,
		},
		{
			name: "verifier-balance",
			entries: []filTypes.EventEntry{
				{Flags: flagIndexedAll, Key: "$type", Codec: cborCodec, Value: cborEncode(t, "verifier-balance")},
				{Flags: flagIndexedAll, Key: "verifier", Codec: cborCodec, Value: cborEncode(t, int64(1234))},
				{Flags: flagIndexedKey, Key: "balance", Codec: cborCodec, Value: cborTokenAmount(t, "12345678901234567891234567890")},
			},
			want: `{"0":{"flags":3,"key":"$type","value":"verifier-balance"},"1":{"flags":3,"key":"verifier","value":1234},"2":{"flags":2,"key":"balance","value":"12345678901234567891234567890"}}`,
		},
		{
			name: "write-queued with bytes payload",
			entries: []filTypes.EventEntry{
				{Flags: flagIndexedAll, Key: "$type", Codec: cborCodec, Value: cborEncode(t, "write-queued")},
				{Flags: flagIndexedAll, Key: "op", Codec: cborCodec, Value: cborEncode(t, int64(0))},
				{Flags: flagIndexedKey, Key: "effective-epoch", Codec: cborCodec, Value: cborEncode(t, int64(5_000_000))},
				{Flags: flagIndexedKey, Key: "payload", Codec: cborCodec, Value: cborEncode(t, []byte("test_data"))},
			},
			want: `{"0":{"flags":3,"key":"$type","value":"write-queued"},"1":{"flags":3,"key":"op","value":0},"2":{"flags":2,"key":"effective-epoch","value":5000000},"3":{"flags":2,"key":"payload","value":"dGVzdF9kYXRh"}}`,
		},
	}

	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			event, err := parseNative(t, rewardActor, tt.entries)
			require.NoError(t, err)
			assert.Equal(t, tt.want, event.Metadata)
			assert.Equal(t, types.EventTypeNative, event.Type)
		})
	}
}

// TestParseNativeLog_SharesSetKeepsList covers the v19 reward `shares-set` event, whose
// `shares` value is a CBOR array of (recipient, share) tuples
// (builtin-actors v19.0.1 actors/reward/src/emit.rs:121).
func TestParseNativeLog_SharesSetKeepsList(t *testing.T) {
	// [[1000, 5000], [1001, 4294967296]] exactly as Serialize_tuple(ShareRow) produces it.
	shares := mustHex(t, "82"+"82"+"1903e8"+"191388"+"82"+"1903e9"+"1b0000000100000000")

	event, err := parseNative(t, rewardActor, []filTypes.EventEntry{
		{Flags: flagIndexedAll, Key: "$type", Codec: cborCodec, Value: cborEncode(t, "shares-set")},
		{Flags: flagIndexedAll, Key: "stream-id", Codec: cborCodec, Value: cborEncode(t, int64(7))},
		{Flags: flagIndexedKey, Key: "shares", Codec: cborCodec, Value: shares},
	})
	require.NoError(t, err)
	assert.Equal(t, "shares-set", event.SelectorID)
	assert.Equal(t,
		`{"0":{"flags":3,"key":"$type","value":"shares-set"},"1":{"flags":3,"key":"stream-id","value":7},"2":{"flags":2,"key":"shares","value":[[1000,5000],[1001,"4294967296"]]}}`,
		event.Metadata)
}

// TestParseNativeLog_NestedValues checks that maps, bytes, links, nulls and bools nested in a
// list or map survive, using the same JSON encoding as top-level values.
func TestParseNativeLog_NestedValues(t *testing.T) {
	pieceCid, err := cid.Decode("baga6ea4seaqeyz6zikyr2bqbhy6mrocoqwagx45vlbpsbem7euqv5mf3hrvn2fy")
	require.NoError(t, err)
	cidBytes := append([]byte{0x00}, pieceCid.Bytes()...)
	require.Less(t, len(cidBytes), 256)
	link := append(mustHex(t, "d82a58"), byte(len(cidBytes)))
	link = append(link, cidBytes...)
	// {"a": [h'746573745f64617461', <link>, null, true], "b": {"c": -1}}
	value := append(append(mustHex(t, "a2"+"6161"+"84"+"49746573745f64617461"), link...), mustHex(t, "f6"+"f5"+"6162"+"a1"+"6163"+"20")...)

	event, err := parseNative(t, rewardActor, []filTypes.EventEntry{
		{Flags: flagIndexedAll, Key: "$type", Codec: cborCodec, Value: cborEncode(t, "nested")},
		{Flags: flagIndexedKey, Key: "value", Codec: cborCodec, Value: value},
	})
	require.NoError(t, err)
	assert.Equal(t,
		`{"0":{"flags":3,"key":"$type","value":"nested"},"1":{"flags":2,"key":"value","value":{"a":["dGVzdF9kYXRh",{"/":"baga6ea4seaqeyz6zikyr2bqbhy6mrocoqwagx45vlbpsbem7euqv5mf3hrvn2fy"},null,true],"b":{"c":-1}}}}`,
		event.Metadata)
}

// TestParseNativeLog_RewardTokenAmounts covers the v19 reward events carrying TokenAmounts
// (builtin-actors v19.0.1 actors/reward/src/emit.rs:82 `amount`, :103 `accrued`, :104 `dust`).
func TestParseNativeLog_RewardTokenAmounts(t *testing.T) {
	tb := []struct {
		name    string
		entries []filTypes.EventEntry
		want    map[string]any
	}{
		{
			name: "claim-payout",
			entries: []filTypes.EventEntry{
				{Flags: flagIndexedAll, Key: "$type", Codec: cborCodec, Value: cborEncode(t, "claim-payout")},
				{Flags: flagIndexedAll, Key: "stream-id", Codec: cborCodec, Value: cborEncode(t, int64(7))},
				{Flags: flagIndexedAll, Key: "recipient", Codec: cborCodec, Value: cborEncode(t, int64(1000))},
				{Flags: flagIndexedKey, Key: "amount", Codec: cborCodec, Value: cborTokenAmount(t, "1000000000000000000")},
			},
			want: map[string]any{"amount": "1000000000000000000", "recipient": float64(1000), "stream-id": float64(7)},
		},
		{
			name: "period-folded",
			entries: []filTypes.EventEntry{
				{Flags: flagIndexedAll, Key: "$type", Codec: cborCodec, Value: cborEncode(t, "period-folded")},
				{Flags: flagIndexedAll, Key: "stream-id", Codec: cborCodec, Value: cborEncode(t, int64(7))},
				{Flags: flagIndexedKey, Key: "cause", Codec: cborCodec, Value: cborEncode(t, "set-shares")},
				{Flags: flagIndexedKey, Key: "accrued", Codec: cborCodec, Value: cborTokenAmount(t, "123456789012345678901234")},
				{Flags: flagIndexedKey, Key: "dust", Codec: cborCodec, Value: cborTokenAmount(t, "0")},
			},
			want: map[string]any{"accrued": "123456789012345678901234", "dust": "0", "cause": "set-shares"},
		},
		{
			// the rule is scoped by $type: an `amount` key on another event keeps its raw value
			name: "amount on an unrelated event is untouched",
			entries: []filTypes.EventEntry{
				{Flags: flagIndexedAll, Key: "$type", Codec: cborCodec, Value: cborEncode(t, "other")},
				{Flags: flagIndexedKey, Key: "amount", Codec: cborCodec, Value: cborEncode(t, int64(3))},
			},
			want: map[string]any{"amount": float64(3)},
		},
	}

	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			event, err := parseNative(t, rewardActor, tt.entries)
			require.NoError(t, err)
			got := metadataValues(t, event)
			for k, v := range tt.want {
				assert.Equal(t, v, got[k], k)
			}
		})
	}
}

// TestParseNativeLog_EVMLog0WithData covers an EVM LOG0 with data: the EVM actor emits only a
// `d` entry (builtin-actors v19.0.1 actors/evm/src/interpreter/instructions/log_event.rs:52-58).
func TestParseNativeLog_EVMLog0WithData(t *testing.T) {
	emitter, err := address.NewDelegatedAddress(10, mustHex(t, "d4c5fb16488aa48081296299d54b0c648c9333da"))
	require.NoError(t, err)

	data := mustHex(t, "00000000000000000000000000000000000000000000000000000000000000ff")
	event, err := parseNative(t, emitter, []filTypes.EventEntry{
		{Flags: flagIndexedAll, Key: EVMDataEventEntryKey, Codec: rawCodec, Value: data},
	})
	require.NoError(t, err)
	assert.Equal(t, types.EventTypeEVM, event.Type)
	assert.Equal(t, "", event.SelectorID)
	assert.Equal(t, `{"data":"00000000000000000000000000000000000000000000000000000000000000ff","topics":null}`, event.Metadata)
}

// TestParseNativeLog_EVMTopicStillValidated keeps the existing behaviour for a t1 entry that
// cannot be read as a topic.
func TestParseNativeLog_EVMTopicStillValidated(t *testing.T) {
	emitter, err := address.NewDelegatedAddress(10, mustHex(t, "d4c5fb16488aa48081296299d54b0c648c9333da"))
	require.NoError(t, err)

	topic := ethtypes.EthHash{0x01}
	event, err := parseNative(t, emitter, []filTypes.EventEntry{
		{Flags: flagIndexedAll, Key: EVMTopic0EventEntryKey, Codec: rawCodec, Value: topic[:]},
		{Flags: flagIndexedAll, Key: EVMDataEventEntryKey, Codec: rawCodec, Value: []byte{0xff}},
	})
	require.NoError(t, err)
	assert.Equal(t, topic.String(), event.SelectorID)

	_, err = parseNative(t, emitter, []filTypes.EventEntry{
		{Flags: flagIndexedAll, Key: EVMTopic0EventEntryKey, Codec: 0x52, Value: []byte{}},
	})
	assert.Error(t, err)
}
