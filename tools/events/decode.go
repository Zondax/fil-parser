package event_tools

import (
	"errors"
	"fmt"
	"strings"

	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/datamodel"
)

const (
	cidEntryValues = "unsealed-cid,piece-cid"
)

// decode does an ipld decode of the entry.Value using dagcbor
// special cases include entries that have a CID as a value.
// CIDs are represented as an ipld.Link which needs an extra step of decoding the CID
// to get the correct JSON representation.
func decode(entry filTypes.EventEntry) (any, error) {
	n, err := ipld.Decode(entry.Value, dagcbor.Decode)
	if err != nil {
		return nil, fmt.Errorf("error ipld decode entry: %w ", err)
	}

	if strings.Contains(entry.Key, cidEntryValues) {
		return parseNullableCid(n)
	}

	return n, nil
}

// parseNullableCid handles the edge-case of event entries with
// nullable CIDs e.g unsealed-cid
func parseNullableCid(n datamodel.Node) (any, error) {
	switch n.Kind() {
	case datamodel.Kind_Null:
		return nil, nil
	case datamodel.Kind_Invalid:
		return nil, errors.New("invalid datamodel kind for ipld node")
	}
	return parseCid(n)
}

func parseCid(n datamodel.Node) (any, error) {
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
