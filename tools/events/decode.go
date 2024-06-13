package event_tools

import (
	"fmt"
	"github.com/filecoin-project/lotus/chain/types"

	"regexp"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
)

var (
	cidRegex    = regexp.MustCompile("cid")
	bigintRegex = regexp.MustCompile("balance")
)

// decode does an ipld decode of the entry.Value using dagcbor
func decode(entry types.EventEntry) (datamodel.Node, error) {
	n, err := ipld.Decode(entry.Value, dagcbor.Decode)
	if err != nil {
		return nil, fmt.Errorf("error ipld decode entry: %w ", err)
	}

	return n, nil
}

// parseBigInt uses the filecoin-project big package to decode a node into a big.Int
// required for the verifier_balance event
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
func parseCid(n datamodel.Node) (any, error) {
	if n.Kind() == datamodel.Kind_Null {
		// nullable CIDs that show up in unsealed_cid are represented as Null
		return nil, nil
	}
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
