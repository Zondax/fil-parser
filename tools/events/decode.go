package event_tools

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"math"
	"regexp"
)

const (
	maxJSONNumber = math.MaxInt32
	minJSONNumber = math.MinInt32
)

var (
	// - https://github.com/filecoin-project/lotus/blob/6e13eac5d51f08d964f1338d9fab7cca42014e5c/documentation/en/actor-events-api.md?plain=1#L365
	cidRegex = regexp.MustCompile("cid")
	// https://github.com/filecoin-project/lotus/blob/6e13eac5d51f08d964f1338d9fab7cca42014e5c/documentation/en/actor-events-api.md?plain=1#L112
	bigintRegex = regexp.MustCompile("balance")

	// tokenAmountKeys lists, per event $type, the entry keys whose value is a TokenAmount
	// (a CBOR byte string holding a big.Int). They are scoped by $type because the keys are
	// generic and must not match unrelated events.
	// - builtin-actors v19.0.1 actors/reward/src/emit.rs:82 (claim-payout: amount)
	// - builtin-actors v19.0.1 actors/reward/src/emit.rs:103-104 (period-folded: accrued, dust)
	tokenAmountKeys = map[string]map[string]struct{}{
		"claim-payout":  {"amount": {}},
		"period-folded": {"accrued": {}, "dust": {}},
	}
)

// isBigIntEntry reports whether the entry value must be decoded as a big.Int.
func isBigIntEntry(eventType, key string) bool {
	if bigintRegex.MatchString(key) {
		return true
	}
	_, ok := tokenAmountKeys[eventType][key]
	return ok
}

// decode does an ipld decode of the entry.Value using dagcbor
func decode(entry types.EventEntry) (datamodel.Node, error) {
	n, err := ipld.Decode(entry.Value, dagcbor.Decode)
	if err != nil {
		return nil, fmt.Errorf("error ipld decode entry: %w ", err)
	}

	if n.Kind() == datamodel.Kind_Int {
		val, err := n.AsInt()
		if err != nil {
			return nil, fmt.Errorf("error ipld node to int : %w ", err)
		}
		if val > maxJSONNumber || val < minJSONNumber {
			return basicnode.NewString(fmt.Sprint(val)), nil
		}
	}

	return n, nil
}

// nodeToAny converts a list or map ipld node into plain Go values so that it can be marshalled to
// JSON; go-ipld-prime nodes of these kinds marshal to {}. Nested values use the same JSON encoding
// as top-level entry values: ints outside the int32 range become strings (see decode), bytes are
// base64 encoded, and links use the {"/": cid} form produced by parseCid.
func nodeToAny(n datamodel.Node) (any, error) {
	switch n.Kind() {
	case datamodel.Kind_Null:
		return nil, nil
	case datamodel.Kind_Bool:
		return n.AsBool()
	case datamodel.Kind_Int:
		if un, ok := n.(datamodel.UintNode); ok {
			// values above math.MaxInt64 are only readable as uint64
			val, err := un.AsUint()
			if err != nil {
				return nil, fmt.Errorf("error ipld node to uint: %w", err)
			}
			if val > maxJSONNumber {
				return fmt.Sprint(val), nil
			}
			return int64(val), nil // #nosec G115 -- val <= maxJSONNumber
		}
		val, err := n.AsInt()
		if err != nil {
			return nil, fmt.Errorf("error ipld node to int: %w", err)
		}
		if val > maxJSONNumber || val < minJSONNumber {
			return fmt.Sprint(val), nil
		}
		return val, nil
	case datamodel.Kind_Float:
		return n.AsFloat()
	case datamodel.Kind_String:
		return n.AsString()
	case datamodel.Kind_Bytes:
		return n.AsBytes()
	case datamodel.Kind_Link:
		return parseCid(n)
	case datamodel.Kind_List:
		list := make([]any, 0, n.Length())
		it := n.ListIterator()
		for !it.Done() {
			_, v, err := it.Next()
			if err != nil {
				return nil, fmt.Errorf("error iterating ipld list: %w", err)
			}
			item, err := nodeToAny(v)
			if err != nil {
				return nil, err
			}
			list = append(list, item)
		}
		return list, nil
	case datamodel.Kind_Map:
		m := make(map[string]any, n.Length())
		it := n.MapIterator()
		for !it.Done() {
			k, v, err := it.Next()
			if err != nil {
				return nil, fmt.Errorf("error iterating ipld map: %w", err)
			}
			key, err := k.AsString()
			if err != nil {
				return nil, fmt.Errorf("error reading ipld map key: %w", err)
			}
			item, err := nodeToAny(v)
			if err != nil {
				return nil, err
			}
			m[key] = item
		}
		return m, nil
	default:
		return nil, fmt.Errorf("unsupported ipld node kind: %s", n.Kind())
	}
}

// parseBigInt uses the filecoin-project big package to decode a node into a big.Int
// required for the verifier_balance event
// https://github.com/filecoin-project/lotus/blob/6e13eac5d51f08d964f1338d9fab7cca42014e5c/documentation/en/actor-events-api.md?plain=1#L112
func parseBigInt(n datamodel.Node) (any, error) {
	hexEncodedInt, err := n.AsBytes()
	if err != nil {
		return nil, fmt.Errorf("error converting ipld node to string: %w", err)
	}

	bigInt, err := big.FromBytes(hexEncodedInt)
	if err != nil {
		return nil, fmt.Errorf("error converting hex encoded bigint to big.Int: %w", err)
	}
	return bigInt.String(), nil
}

// parseCid parses an ipld node into the correct cid implementation.
// special cases include entries that have a CID as a value.
// CIDs are represented as an ipld.Link which needs an extra step of decoding the CID
// to get the correct JSON representation.
// Current edge case entry keys: unsealed-cid,piece-cid
// References:
// - https://github.com/filecoin-project/lotus/blob/6e13eac5d51f08d964f1338d9fab7cca42014e5c/documentation/en/actor-events-api.md?plain=1#L365
func parseCid(n datamodel.Node) (any, error) {
	if n.Kind() == datamodel.Kind_Null {
		// nullable CIDs that show up in unsealed_cid are represented as Null
		// - https://github.com/filecoin-project/lotus/blob/5dffc05a30894283287345d61b6578be7897ee4b/itests/direct_data_onboard_verified_test.go#L194
		return nil, nil
	}
	// - https://github.com/filecoin-project/lotus/blob/5dffc05a30894283287345d61b6578be7897ee4b/itests/direct_data_onboard_verified_test.go#L201
	if n.Kind() != datamodel.Kind_Link {
		return nil, fmt.Errorf("unexpected datamodel kind for cid: %s ,expected: link", n.Kind())
	}

	link, err := n.AsLink()
	if err != nil {
		return nil, fmt.Errorf("error converting cid ipld node to link : %s : %w", n.Kind(), err)
	}

	c, err := cid.Decode(link.String())
	if err != nil {
		return nil, fmt.Errorf("error decoding %s to cid: %w", link.String(), err)
	}

	return c, nil
}
