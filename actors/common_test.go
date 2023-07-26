package actors

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/stretchr/testify/require"
	"github.com/zondax/fil-parser/parser"
	"testing"
)

func TestActorParser_parseSend(t *testing.T) {
	p := getActorParser()
	msg, err := deserializeMessage(manifest.MultisigKey, parser.MethodSend)
	require.NoError(t, err)
	require.NotNil(t, msg)
	got := p.parseSend(msg)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
}

func TestActorParser_parseConstructor(t *testing.T) {
	p := getActorParser()
	rawParams, err := loadFile(manifest.AccountKey, parser.MethodConstructor, parser.ParamsKey)
	require.NoError(t, err)
	require.NotNil(t, rawParams)
	got, err := p.parseConstructor(rawParams)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
}

func TestActorParser_parseUnknown(t *testing.T) {
	p := getActorParser()
	rawParams, err := loadFile(manifest.AccountKey, parser.UnknownStr, parser.ParamsKey)
	require.NoError(t, err)
	require.NotNil(t, rawParams)
	got, err := p.unknownMetadata(rawParams, nil)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Contains(t, got, parser.ParamsKey, fmt.Sprintf("%s could no be found in metadata", parser.ParamsKey))
}
