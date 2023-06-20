package actors

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
	"net/http"
	"os"

	"github.com/filecoin-project/lotus/api/client"
	rosettaFilecoinLib "github.com/zondax/rosetta-filecoin-lib"

	"github.com/zondax/fil-parser/actors/database"
)

const (
	testUrl  = "https://api.zondax.ch/fil/node/mainnet/rpc/v1"
	dataPath = "../data/actors"
)

func loadFile(actor, txType, name string) ([]byte, error) {
	item, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/%s", dataPath, actor, txType, name))
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
	rawParams, err := loadFile(actor, txType, parser.ParamsKey)
	if err != nil {
		return nil, nil, err
	}
	rawReturn, err := loadFile(actor, txType, parser.ReturnKey)
	if err != nil {
		return nil, nil, err
	}
	return rawParams, rawReturn, nil
}

func deserializeMessage(actor, txType string) (*parser.LotusMessage, error) {
	file, err := os.Open(fmt.Sprintf("%s/%s/%s/%s", dataPath, actor, txType, "Message"))
	if err != nil {
		return nil, err
	}
	decoder := gob.NewDecoder(file)
	message := &parser.LotusMessage{}
	err = decoder.Decode(message)
	return message, err
}

func deserializeTipset(actor, txType string) (t filTypes.TipSet, err error) {
	file, err := os.Open(fmt.Sprintf("%s/%s/%s/%s", dataPath, actor, txType, "Tipset"))
	if err != nil {
		return
	}
	err = t.UnmarshalCBOR(file)
	return
}

func getEthLogs(actor, txType string) ([]types.EthLog, error) {
	file, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/%s", dataPath, actor, txType, parser.EthLogsKey))
	if err != nil {
		return nil, err
	}
	var ethLogs []types.EthLog
	err = json.Unmarshal(file, &ethLogs)
	return ethLogs, err
}
