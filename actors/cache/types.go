package cache

import (
	"context"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/go-resty/resty/v2"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/zondax/fil-parser/actors/cache/impl/common"
	"github.com/zondax/fil-parser/types"
	"go.uber.org/zap"
)

const SignatureDBURL = "https://www.4byte.directory/api/v1/event-signatures/"

type IActorsCache interface {
	NewImpl(source common.DataSource, logger *zap.Logger) error
	GetActorCode(add address.Address, key filTypes.TipSetKey) (string, error)
	GetRobustAddress(add address.Address) (string, error)
	GetShortAddress(add address.Address) (string, error)
	StoreAddressInfo(info types.AddressInfo)
	GetEVMSelectorSig(ctx context.Context, selectorHash string) (string, error)
	StoreEVMSelectorSig(ctx context.Context, selectorHash, selectorSig string) error
	BackFill() error
	ImplementationType() string
}

type ActorsCache struct {
	offChainCache IActorsCache
	onChainCache  IActorsCache
	badAddress    cmap.ConcurrentMap
	logger        *zap.Logger
	httpClient    *resty.Client
}

// FourBytesSignatureResult represents the response from SignatureDBURL
type FourBytesSignatureResult struct {
	Results []struct {
		HexSignature  string `json:"hex_signature"`
		TextSignature string `json:"text_signature"`
	} `json:"results"`
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}
