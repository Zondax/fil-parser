package actors

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/filecoin-project/lotus/api/client"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	"github.com/zondax/fil-parser/actors/database"
)

const testUrl = "https://api.zondax.ch/fil/node/mainnet/rpc/v1"

func loadFile(actor, txType, name string) ([]byte, error) {
	item, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/%s", testDataPath, actor, txType, name))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func getActorParser() *ActorParser {
	lotusClient, _, err := client.NewFullNodeRPCV1(context.Background(), testUrl, http.Header{})
	if err != nil {
		return nil
	}
	database.SetupActorsDatabase(&lotusClient)

	lib := rosettaFilecoinLib.NewRosettaConstructionFilecoin(lotusClient)
	return &ActorParser{
		lib: lib,
	}
}

func getParmasAndReturn(actor, txType string) ([]byte, []byte, error) {
	rawParams, err := loadFile(actor, txType, "params")
	if err != nil {
		return nil, nil, err
	}
	rawReturn, err := loadFile(actor, txType, "return")
	if err != nil {
		return nil, nil, err
	}
	return rawParams, rawReturn, nil
}
