package context

import (
	"context"
	"fmt"

	"github.com/filecoin-project/lotus/api"
)

type nodesContextKey string

const (
	key nodesContextKey = "nodes"
)

func GetNodes(ctx context.Context) ([]api.FullNode, error) {
	nodesVal := ctx.Value(key)
	if nodesVal != nil {
		nodes, ok := nodesVal.([]api.FullNode)
		if ok {
			return nodes, nil
		}
	}
	return nil, fmt.Errorf("nodes missing in context")
}

func SetNodes(parent context.Context, nodes []api.FullNode) context.Context {
	return context.WithValue(parent, key, nodes)
}
