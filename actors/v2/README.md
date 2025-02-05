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

To add a new actor:

- Create a new package under `actors/v2`.
- Implement a struct that satisfies the `Actor` interface.
- Integrate the actor into the `ActorParser.GetActor` function.

### Method Coverage Verification

**Test Function:** `TestMethodCoverage`

To add support for a new method:

- Implement the new method within the actor struct.
- Add the method to the `TransactionTypes` map.
- Incorporate the method into the `Parse` function.

### Network Version Coverage Verification

**Test Function:** `TestVersionCoverage`

To add support for a new network version:

- Include the network version in `tools/version_mapping.go`.
- Ensure the version is supported in the actor parsing methods (use existing switch cases as a reference).

## Compatibility

The `ActorParser` is designed to be backwards compatible and a drop in replacement for ActorsV1.
