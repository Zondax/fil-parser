# Parsing Failures

## CBOR unmarshal parameter failures

> Note: The same parsing logic successfully processes identical message types at other heights, with successful parses significantly outnumbering failures.

The "got" and "expected" values shown below were derived by directly converting raw CBOR data to JSON arrays.
The discrepancies appear in parameter formatting and ordering rather than in the underlying data.
These parameter structures do not match any known builtin-actor or spec-actor implementations, requiring further investigation to determine their origin.

| method | count | got | expected |
|---------|------|---- |--------- |
| VerifyDealsForActivation | 59,684 | `[[[528, 8, 5211920, [199775, 199776]]]` | `[[[8, 4973322, [183029, 183030, 183031, 183032, 183033, 183034]]]]` |


### Parameter Structure Analysis

#### VerifyDealsForActivation

**Expected format:** `[[[sector, epoch, [dealIds]]]]`  
**Observed format:** `[[[unknown, sector, epoch, [dealIds]]]]`  
The failing messages contain an additional unknown parameter at the beginning of each array.
