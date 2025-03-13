# Parsing Failures

## CBOR unmarshal parameter failures

> Note: The same parsing logic successfully processes identical message types at other heights, with successful parses significantly outnumbering failures.

The "got" and "expected" values shown below were derived by directly converting raw CBOR data to JSON arrays.
The discrepancies appear in parameter formatting and ordering rather than in the underlying data.
These parameter structures do not match any known builtin-actor or spec-actor implementations, requiring further investigation to determine their origin.

| method | count | got | expected |
|---------|------|---- |--------- |
| VerifyDealsForActivation | 59,684 | `[[[528, 8, 5211920, [199775, 199776]]]` | `[[[8, 4973322, [183029, 183030, 183031, 183032, 183033, 183034]]]]` |
| ActivateDeals | 11,055 | `[[[8,2553148,[141635]],[8,2553160,[141636]],[8,2553196,[141649]],[8,2553211,[141662]]],false]` | `[[141641],1620353]` |
| ClaimAllocations | 6,200 | `[[[654,2656054,[[33431,29115,"AAGB4gOSICAq+/vXPqVNpVihCnnIjdkcIP/GLIVfeUHIvoUXQ48cIA==",34359738368]]]],true]` | ``[[[1011,2,"AAGB4gOSICBH1Nv50zeXIfZ4TPoqlwRCzcrO+NEr4GEalxhV2hm2Lg==",34359738368,14,1509893]],true]`` |
| OnMinerSectorsTerminate | 160 | `[1435439,"GA=="]` |  `[27761,[2454]]` |
| GetStorageAt | 4 | `{}` | `multisig-inner-proposal` |
| Resurrect | 4 | `{}` | `multisig-inner-proposal` |

### Parameter Structure Analysis

#### VerifyDealsForActivation

**Expected format:** `[[[sector, epoch, [dealIds]]]]`  
**Observed format:** `[[[unknown, sector, epoch, [dealIds]]]]`  
The failing messages contain an additional unknown parameter at the beginning of each array.

#### ActivateDeals

**Expected format:** `[[dealIds], epoch]`  
**Observed format:** `[[[sector, epoch, [dealId]], [sector, epoch, [dealId]], ...], boolean]`  
The parameter structure differs significantly, with the observed format containing sector information and a boolean flag.

#### ClaimAllocations

**Expected format:** `[[[param1, param2, "base64string", number, param5, param6]], boolean]`  
**Observed format:** `[[[param1, param2, [[param3, param4, "base64string", number]]]], boolean]`  
The parameter nesting structure differs, with variations in the number and arrangement of parameters.

#### OnMinerSectorsTerminate

**Expected format:** `[epoch, [dealIds]]`  
**Observed format:** `[epoch, "base64string"]`  
The second parameter appears to be encoded differently.

#### Implementation Differences

There are notable differences between Golang and Rust implementations:

**Golang:**

```go
type OnMinerSectorsTerminate struct {
    abi.Epoch
    []Deals
}
```

**Rust:**

```rust
pub struct OnMinerSectorsTerminateParams {
    pub epoch: ChainEpoch,
    pub sectors: BitField,
}
```

The second parameter in failing messages likely represents a sector BitField rather than deal IDs.

#### GetStorageAt & Resurrect

These methods are executed by the multisig actor. Current implementation does not correctly handle inner proposal parameters.

## Method Number Resolution Issues

The following method numbers are unexpected for their respective actors. Research indicates that `account`, `evm`, and `ethaccount` actors implement Universal Receiver interfaces and accept any method number greater than FIRST_EXPORTED_METHOD_NUMBER (as defined in the library).

|height| actor | method number | calling-actor |
|------|-------|---------------|---------------|
|1793034|ethaccount|23|account|
|2152440|account|16|account|
|2129712|evm|28|account|
|2134311|account|28|account|
|2230551|account|5|account|
