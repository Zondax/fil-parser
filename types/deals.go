package types

import (
	"time"
)

type DealsEvents struct {
	DealsMessages    []*DealsMessages
	DealsProposals   []*DealsProposals
	DealsActivations []*DealsActivations
	DealsSpaceInfo   []*DealsSpaceInfo
}

type DealsMessages struct {
	ID           string    `json:"id"`
	ActorAddress string    `json:"actor_address"`
	Height       uint64    `json:"height"`
	TxCid        string    `json:"tx_cid"`
	ActionType   string    `json:"action_type"`
	Data         string    `json:"data"`
	TxTimestamp  time.Time `json:"tx_timestamp"`
}

type DealsProposals struct {
	ID           string `json:"id"`
	Height       uint64 `json:"height"`
	ActorAddress string `json:"actor_address"`
	DealID       uint64 `json:"deal_id"`
	TxCid        string `json:"tx_cid"`

	// proposal details
	ClientSignature string `json:"client_signature"`
	ProviderAddress string `json:"provider_address"`
	ClientAddress   string `json:"client_address"`
	PieceCid        string `json:"piece_cid"`
	PieceSize       uint64 `json:"piece_size"`
	Verified        bool   `json:"verified"`

	// Arbitrary client chosen label to apply to the deal
	Label string `json:"label"`

	// Nominal start epoch. Deal payment is linear between StartEpoch and EndEpoch,
	// with total amount StoragePricePerEpoch * (EndEpoch - StartEpoch).
	// Storage deal must appear in a sealed (proven) sector no later than StartEpoch,
	// otherwise it is invalid.
	StartEpoch    int64  `json:"start_epoch"`
	EndEpoch      int64  `json:"end_epoch"`
	PricePerEpoch uint64 `json:"price_per_epoch"`

	ProviderCollateral uint64 `json:"provider_collateral"`
	ClientCollateral   uint64 `json:"client_collateral"`

	TxTimestamp time.Time `json:"tx_timestamp"`
}

type DealsActivations struct {
	ID           string    `json:"id"`
	Height       uint64    `json:"height"`
	ActorAddress string    `json:"actor_address"`
	TxCid        string    `json:"tx_cid"`
	DealID       uint64    `json:"deal_id"`
	SectorExpiry int64     `json:"sector_expiry"`
	ActionType   string    `json:"action_type"`
	TxTimestamp  time.Time `json:"tx_timestamp"`
}

type DealsSpaceInfo struct {
	ID           string   `json:"id"`
	Height       uint64   `json:"height"`
	ActorAddress string   `json:"actor_address"`
	TxCid        string   `json:"tx_cid"`
	DealID       uint64   `json:"deal_id"`
	GroupDealIDs []uint64 `json:"group_deal_ids"`
	// NonVerifiedDealWeight is the sum(piece_size * deal_duration) of all the non-verified deals
	// VerifiedDealWeight is the sum(piece_size * deal_duration) of all the verified deals
	// NonVerifiedDealSpace is the sum(piece_size) of all the deals
	// VerifiedDealSpace is the sum(piece_size) of all the verified deals
	NonVerifiedDealSpace uint64 `json:"non_verified_deal_space"`
	VerifiedDealSpace    uint64 `json:"verified_deal_space"`
	// SpaceAsWeight is true if the deal space is expressed as a weight. Retrieve the deal space by dividing the VerifiedDealSpace/NonVerifiedDealSpace by the DealDuration for each dealId in the DealIDs slice.
	SpaceAsWeight bool      `json:"space_as_weight"`
	ActionType    string    `json:"action_type"`
	TxTimestamp   time.Time `json:"tx_timestamp"`
}
