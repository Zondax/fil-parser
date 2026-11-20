package types

import (
	"math/big"
	"time"
)

type DealsEvents struct {
	DealsMessages    []*DealsMessages
	DealsProposals   []*DealsProposals
	DealsActivations []*DealsActivations
	DealsSpaceInfo   []*DealsSpaceInfo
	// DealsTerminations are the market OnMinerSectorsTerminate calls, see DealsTerminations.
	DealsTerminations []*DealsTerminations
	// DealsSectorMappings map deals to the sectors that hold them, see DealsSectorMappings.
	DealsSectorMappings []*DealsSectorMappings
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
	StartEpoch    int64    `json:"start_epoch"`
	EndEpoch      int64    `json:"end_epoch"`
	PricePerEpoch *big.Int `json:"price_per_epoch" gorm:"column:price_per_epoch;type:UInt256"`

	ProviderCollateral *big.Int `json:"provider_collateral" gorm:"column:provider_collateral;type:UInt256"`
	ClientCollateral   *big.Int `json:"client_collateral" gorm:"column:client_collateral;type:UInt256"`

	TxTimestamp time.Time `json:"tx_timestamp"`
}

type DealsActivations struct {
	ID           string `json:"id"`
	Height       uint64 `json:"height"`
	ActorAddress string `json:"actor_address"`
	TxCid        string `json:"tx_cid"`
	DealID       uint64 `json:"deal_id"`
	SectorExpiry int64  `json:"sector_expiry"`
	// SectorNumber is the sector the deal was activated in. It is known for SectorContentChanged
	// and for BatchActivateDeals from actors v13 (NV22) on; nil when the params do not carry it.
	SectorNumber *uint64   `json:"sector_number" gorm:"column:sector_number"`
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
	NonVerifiedDealSpace *big.Int `json:"non_verified_deal_space" gorm:"column:non_verified_deal_space;type:UInt256"`
	VerifiedDealSpace    *big.Int `json:"verified_deal_space" gorm:"column:verified_deal_space;type:UInt256"`
	// SpaceAsWeight is true if the deal space is expressed as a weight. Retrieve the deal space by dividing the VerifiedDealSpace/NonVerifiedDealSpace by the DealDuration for each dealId in the DealIDs slice.
	SpaceAsWeight bool      `json:"space_as_weight"`
	ActionType    string    `json:"action_type"`
	TxTimestamp   time.Time `json:"tx_timestamp"`
}

// DealsTerminations records a successful market OnMinerSectorsTerminate call, made by a miner when
// sectors are terminated early (TerminateSectors, faults expired for too long, or the remainder of
// a termination processed later under cron).
//
// Up to actors v12 (NV21) the call carries deal ids: there is one row per deal id, with DealID set
// and SectorNumber nil. The market skips deals whose end_epoch <= TerminationEpoch (expiries, not
// early terminations), deals already slashed and deals whose proposal is gone, so downstream must
// keep only deals with end_epoch > TerminationEpoch (and the first termination per deal).
//
// From actors v13 (NV22) the call carries the terminated sectors instead: there is one row per
// sector, with SectorNumber set and DealID nil. Downstream maps (ActorAddress, SectorNumber) to
// deals through DealsSectorMappings, and applies the same end_epoch filter. Actors v13 executes
// from the tipset after the NV22 upgrade height; the upgrade tipset itself still runs v12 and
// produces deal id rows.
type DealsTerminations struct {
	ID     string `json:"id"`
	Height uint64 `json:"height"`
	TxCid  string `json:"tx_cid"`
	// ActorAddress is the miner (the caller of the market), in the same form as
	// DealsActivations.ActorAddress.
	ActorAddress string  `json:"actor_address"`
	DealID       *uint64 `json:"deal_id" gorm:"column:deal_id"`
	SectorNumber *uint64 `json:"sector_number" gorm:"column:sector_number"`
	// TerminationEpoch is the epoch in the params: the epoch the sectors were terminated at up to
	// actors v12, the current epoch from actors v13 on.
	TerminationEpoch int64     `json:"termination_epoch"`
	ActionType       string    `json:"action_type"`
	TxTimestamp      time.Time `json:"tx_timestamp"`
}

// DealsSectorMappings records that a miner's sector holds a deal. It is the single source for
// joining terminations by sector (DealsTerminations from actors v13 on) to deals:
//
//	deals_terminations (actor_address, sector_number) -> deals_sector_mappings -> deal_id
//
// Rows come from:
//   - PreCommitSector / PreCommitSectorBatch / PreCommitSectorBatch2 before actors v13 (NV22): the
//     sector's deal ids. A pre-commit is all-or-nothing, a sector number is never reused, and a
//     sector whose deals fail to activate is never proven, so a sector that is ever terminated
//     holds exactly its pre-commit's deals;
//   - ProveReplicaUpdates / ProveReplicaUpdates2 (SnapDeals) before actors v13: the deals of each
//     update the return reports as successful (the sector's deal ids are replaced by them);
//   - every DealsActivations row that has a SectorNumber (SectorContentChanged, and
//     BatchActivateDeals from actors v13), copied so downstream needs no second source.
//
// "Before actors v13" means tipsets executed by actors v12 or older, which includes the tipset at
// the NV22 upgrade height (the migration runs before the next tipset's messages).
//
// These rows are not activations: they are never counted as deal space. The same
// (actor_address, sector_number, deal_id) can appear more than once (e.g. pre-committed before
// NV22 and activated after it), so downstream joins on distinct deal ids.
type DealsSectorMappings struct {
	ID     string `json:"id"`
	Height uint64 `json:"height"`
	TxCid  string `json:"tx_cid"`
	// ActorAddress is the miner, in the same form as DealsActivations.ActorAddress and
	// DealsTerminations.ActorAddress.
	ActorAddress string `json:"actor_address"`
	SectorNumber uint64 `json:"sector_number"`
	DealID       uint64 `json:"deal_id"`
	// ActionType is the method the mapping comes from.
	ActionType  string    `json:"action_type"`
	TxTimestamp time.Time `json:"tx_timestamp"`
}
