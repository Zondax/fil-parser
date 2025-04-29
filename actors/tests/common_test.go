package actortest

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
)

func TestActorParserV1_CommonParseSend(t *testing.T) {
	msg, err := deserializeMessage(manifest.MultisigKey, parser.MethodSend)
	require.NoError(t, err)
	require.NotNil(t, msg)
	got := actors.ParseSend(msg)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
}

func TestActorParserV1_CommonParseConstructor(t *testing.T) {
	rawParams, err := loadFile(manifest.AccountKey, parser.MethodConstructor, parser.ParamsKey)
	require.NoError(t, err)
	require.NotNil(t, rawParams)
	got, err := actors.ParseConstructor(rawParams)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
}

func TestActorParserV1_CommonParseUnknown(t *testing.T) {
	rawParams, err := loadFile(manifest.AccountKey, parser.UnknownStr, parser.ParamsKey)
	require.NoError(t, err)
	require.NotNil(t, rawParams)
	got, err := actors.ParseUnknownMetadata(rawParams, nil)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
}
