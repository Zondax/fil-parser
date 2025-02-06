# Actors V2

## Overview

ActorsV2 is a new implementation of the actor parsing logic that is more flexible and easier to maintain while properly handling all legacy spec-actors and new builtin-actors structs. This ensures that the parser can properly parse all actor messages for any network version.

## Design

The `ActorParser` struct is tasked with parsing actor messages and is composed of a `helper.Helper` and a `zap.Logger`.

Each actor implements the `Actor` interface, which includes the following methods:

- `Name() string`: Returns the actor's name.
- `Parse(network string, height int64, txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid, key filTypes.TipSetKey) (map[string]interface{}, *types.AddressInfo, error)`: Parses the actor message.
- `TransactionTypes() map[string]any`: Provides a map of all transaction types supported by the actor.

### Actor Folder Structure

Each actor is organized within its own directory under `actors/v2`. The directory contains the following files:

- `generic.go`: Defines generic functions, typed with specific builtin-actors/spec-actors version structs, for parsing actor messages.
- `parse.go`: Contains the switch case logic for parsing specific transaction types for the actor.

## Testing

There are different types of tests to ensure comprehensive coverage:

### Actor Support Verification

**Test Function:** `TestAllActorsSupported`

This test verifies that all actors in the latest builtin-actor release and all legacy spec-actors are supported by the parser.
It will fail if any actor is not supported.

> Note: No modification is needed to this test for new releases, the test automatically gets the latest builtin-actor release from github.

To add a new actor:

- Create a new package under `actors/v2`.
- Implement a struct that satisfies the `Actor` interface.
- Integrate the actor into the `ActorParser.GetActor` function.

### Method Coverage Verification

**Test Function:** `TestMethodCoverage`

This test verifies that all methods exposed by the actor in all builtin-actor and spec-actor releases are supported by the actors.
It will fail if any method is not covered.

> Note: No modification is needed to this test on new builtin-actor releases. The test will pass once support is added.

To add support for a new method:

- Implement the new method within the actor struct.
- Add the method to the `TransactionTypes` map.
- Incorporate the method into the `Parse` function.

### Network Version Coverage Verification

**Test Function:** `TestVersionCoverage`

This test verifies that all the actor methods can correctly handle all network versions ( decided by the height of the block ).
It will fail if any network version is not supported.

> Note: No modification is needed to this test on new builtin-actor releases. The test will pass once support is added.

To add support for a new network version:

- Include the network version in `tools/version_mapping.go`.
- Ensure the version is supported in the actor parsing methods (use existing switch cases as a reference).

These tests are designed to ensure that the actor parser accurately handles all releases of both builtin-actors and spec-actors. They are configured to automatically fail upon the release of any new builtin-actor version. This failure mechanism guides developers to the necessary modifications, thereby eliminating the need for manual verification of the parser with each new Filecoin upgrade.

### Actor Functionality Tests

**Test Location:** `actors/tests/{actor_name}_test.go`

These tests, originally developed for actors version 1 (v1), are designed to validate the functionality of actors by comparing their function outputs against a set of pre-calculated expected values. These expected values are stored in the data/actors/{actor_name} directory.

Each actor undergoes both v1 and v2 tests, and passing both test suites is a mandatory requirement. This dual testing approach ensures backward compatibility and adherence to established specifications.

**Important Note:** The pre-computed data currently stored within data/actors corresponds to network version **V20**. If testing against a different network version is required, the `cmd/tracedl` tool provides a mechanism for automatically updating the stored data to the desired network version. This ensures that tests are always executed against the correct expected values for the target network version.

These tests are designed to ensure that the actor parser accurately handles all releases of both builtin-actors and spec-actors. They are configured to automatically fail upon the release of any new builtin-actor version. This failure mechanism guides developers to the necessary modifications, thereby eliminating the need for manual verification of the parser with each new Filecoin upgrade.

## Compatibility

The `ActorParser` is designed to be backwards compatible and a drop in replacement for ActorsV1.

## Misc

### Filecoin Network Version - Actor Version Mapping

>
> [https://github.com/filecoin-project/builtin-actors/releases](https://github.com/filecoin-project/builtin-actors/releases?page=11)
>
> [https://github.com/filecoin-project/community/discussions/74](https://github.com/filecoin-project/community/discussions/74)

The following table shows the mapping of Filecoin network versions to actor versions:

| Network Version | Actor Version              | Height(Mainnet) | Height(Calibration) |
|-----------------|----------------------------|-----------------|---------------------|
| v8              | v2(spec-actors)            | 170000          | UNKNOWN             |
| v9              | v2(spec-actors)            | 265200          | UNKNOWN             |
| v10             | v3(spec-actors)            | 550321          | UNKNOWN             |
| v11             | v3(spec-actors)            | 665280          | UNKNOWN             |
| v12             | v4(spec-actors)            | 712320          | 193789              |
| v13             | v5(spec-actors)            | 892800          | 0 (RESET)           |
| v14             | v6(spec-actors)            | 1231620         | 312746              |
| v15             | v7(spec-actors)            | 1594680         | 682006              |
| v16             | v8(builtin-actors)         | 1960320         | 1044660             |
| v17             | v9(builtin-actors)         | 2383680         | 16800 (RESET)       |
| v18             | v10(builtin-actors)        | 2683348         | 322354              |
| v19             | v11(builtin-actors)        | 2809800         | 489094              |
| v20             | v11(builtin-actors)        | 2870280         | 492214              |
| v21             | v12(builtin-actors)        | 3469380         | 1108174             |
| v22             | v13(builtin-actors)        | 3817920         | 1427974             |
| v23             | v14(builtin-actors)        | 4154640         | 1779094             |
| v24             | v15(builtin-actors)        | 4461240         | 2081674             |
