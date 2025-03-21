package v2_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/zondax/fil-parser/actors/metrics"
	metrics2 "github.com/zondax/fil-parser/metrics"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	builtinActors "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	"github.com/zondax/fil-parser/actors"
	v2 "github.com/zondax/fil-parser/actors/v2"
	logger2 "github.com/zondax/fil-parser/logger"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/tools"
)

var latestBuiltinActorVersion uint64
var l *zap.Logger

func TestMain(m *testing.M) {
	version, err := v2.LatestBuiltinActorVersion()
	if err != nil {
		panic(fmt.Sprintf("failed to get latest builtin actor version: %v", err))
	}
	fmt.Printf("latestBuiltinActorVersion: %d\n", version)
	latestBuiltinActorVersion = version

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("failed to create logger: %v", err))
	}
	l = logger2.GetSafeLogger(logger)
	os.Exit(m.Run())
}

// TestVersionCoverage tests that all actor methods are supported for all supported network versions
func TestVersionCoverage(t *testing.T) {
	tests := []struct {
		name    string
		network string
	}{
		{name: "mainnet", network: "mainnet"},
		{name: "calibration", network: "calibration"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			versions := tools.GetSupportedVersions(tt.network)
			actorParsers := getActors(t)
			ctx := context.Background()
			for _, version := range versions {
				height := tools.DeterministicTestHeight(version)
				for _, actor := range actorParsers {
					methods, err := actor.Methods(ctx, tt.network, height)
					if err != nil {
						continue
					}
					if actor.Name() != manifest.PlaceholderKey {
						require.NotEmptyf(t, methods, "Methods are empty for actor: %s version: %s height: %d", actor.Name(), version, height)
					}
					for _, method := range methods {
						if method.Method == nil {
							// depracated
							continue
						}
						_, _, err := actor.Parse(context.Background(), tt.network, height, method.Name, &parser.LotusMessage{}, &parser.LotusMessageReceipt{}, cid.Undef, filTypes.TipSetKey{})
						require.Falsef(t, errors.Is(err, actors.ErrUnsupportedHeight), "Missing support for txType: %s, actor: %s version: %s height: %d", method.Name, actor.Name(), version, height)
						require.Falsef(t, errors.Is(err, parser.ErrUnknownMethod), "Method missing in actor.Parse: %s, actor: %s version: %s height: %d", method.Name, actor.Name(), version, height)
					}
				}
			}
		})

	}

}

// TestAllActorsSupported tests that all actors are supported for the latest actor version
func TestAllActorsSupported(t *testing.T) {
	getActors(t)
}

func getActors(t *testing.T) []v2.Actor {
	actorParser := v2.NewActorParser("mainnet", nil, l, metrics2.NewNoopMetricsClient()).(*v2.ActorParser)
	// #nosec G115
	filActors := manifest.GetBuiltinActorsKeys(builtinActors.Version(latestBuiltinActorVersion))
	actors := []v2.Actor{}
	for _, filActor := range filActors {
		actor, err := actorParser.GetActor(filActor, &metrics.ActorsMetricsClient{MetricsClient: metrics2.NewNoopMetricsClient()})
		require.NoErrorf(t, err, "Actor %s is not supported", filActor)
		actors = append(actors, actor)
	}
	return actors
}

// TestABIMethodNumberToMethodName tests that the method number is mapped to the correct method name for every version
func TestABIMethodNumberToMethodName(t *testing.T) {
	network := "mainnet"

	versions := tools.GetSupportedVersions(network)
	require.NotEmpty(t, versions)

	actorParsers := getActors(t)
	require.NotEmpty(t, actorParsers)

	for _, version := range versions {
		height := tools.DeterministicTestHeight(version)
		for _, actor := range actorParsers {
			transactionTypes := actor.TransactionTypes()
			// Placeholder actor has no methods
			if actor.Name() == manifest.PlaceholderKey {
				continue
			}
			require.NotEmptyf(t, transactionTypes, "Transaction types are empty for actor: %s version: %s height: %d", actor.Name(), version, height)
			methods, err := actor.Methods(context.Background(), network, height)
			if actor.StartNetworkHeight() > height {
				continue
			}
			require.NoErrorf(t, err, "Failed to get methods for actor: %s version: %s height: %d", actor.Name(), version, height)
			require.NotEmptyf(t, methods, "Methods are empty for actor: %s version: %s height: %d", actor.Name(), version, height)

			for methodNum := range methods {
				methodName := methods[methodNum].Name
				assert.Containsf(t, transactionTypes, methodName, "Method name: %s is not in transaction types for actor: %s version: %s height: %d", methodName, actor.Name(), version, height)
			}
		}
	}
}
