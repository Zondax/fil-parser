# Deals Lifecycle

## Creation

storageMarket.PublishStorageDeals 

Validation: provider and client have enough locked balance (collateral + initial)

status: published

## Activation

miner ProRep ( proof of replication)
implicit storageMarket.ActivateStorageDeals by storageMiner.ProveCommit

status: active

## Maintenance

regular storageMiner.SubmitWindowedPoSt 

storageMarket.OnSuccesfulPost triggers deal payments from the clients locked collateral.


FaultySectors do not slash deal collateral but may slash the pledge collateral

## Termination

status: Deleted

### Expiration

When Deal EndEpoch is reached and sector containing the deal expires or reaches end of its term.
Power associated with the deal is lost.

### Early Termination / Faulty Sector Termination

storageMiner.TerminateSectors

Smart Contract termination

# Overview

## StorageMarketActor

- PublishStorageDeals
- ActivateStorageDeals
- AddBalance ( both client and provider)
- WithdrawBalance ( unlocked FIL from the escrow)
- (implicit) OnSuccessfulPost
- GetDealDataCommitment
- GetDealProvider
- GetDealClient
- GetDealTerm
- GetDealActivation

## StorageMinerActor

- ProveCommitSector / ProveCommitAggregate
- TerminateSectors
- ExtendSectorExpiration

# Misc

- VerifiedRegistryActor manages datacap allocations for verified clients. Deals made with datacap receive a storage power multiplier.
- RewardActor distributes block rewards to SPs based on proven storage power.
- FVM contracts can programatically create and manage deals by invoking the relevant builtin actors methods.




