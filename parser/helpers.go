package parser

import (
	"encoding/hex"
	"errors"

	"github.com/filecoin-project/go-address"
	filTypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/zondax/rosetta-filecoin-lib/actors"
	"go.uber.org/zap"

	"github.com/zondax/fil-parser/database"
	"github.com/zondax/fil-parser/parser/methods"
	"github.com/zondax/fil-parser/types"
)

var allMethods = methods.All()["v11"]()

func (p *Parser) getActorAddressInfo(add address.Address, height int64, key filTypes.TipSetKey) types.AddressInfo {
	var (
		addInfo types.AddressInfo
		err     error
	)
	addInfo.Robust, err = database.ActorsDB.GetRobustAddress(add)
	if err != nil {
		zap.S().Errorf("could not get robust address for %s. Err: %v", add.String(), err)
	}

	addInfo.Short, err = database.ActorsDB.GetShortAddress(add)
	if err != nil {
		zap.S().Errorf("could not get short address for %s. Err: %v", add.String(), err)
	}

	addInfo.ActorCid, err = database.ActorsDB.GetActorCode(add, height, key)
	if err != nil {
		zap.S().Errorf("could not get actor code from address. Err:", err)
	} else {
		addInfo.ActorType, _ = p.lib.BuiltinActors.GetActorNameFromCid(addInfo.ActorCid)
	}

	return addInfo
}

func (p *Parser) getActorNameFromAddress(address address.Address, height int64, key filTypes.TipSetKey) string {
	var actorCode cid.Cid
	// Search for actor in cache
	var err error
	actorCode, err = database.ActorsDB.GetActorCode(address, height, key)
	if err != nil {
		return actors.UnknownStr
	}

	actorName, err := p.lib.BuiltinActors.GetActorNameFromCid(actorCode)
	if err != nil {
		return actors.UnknownStr
	}

	return actorName
}

func (p *Parser) GetMethodName(msg *filTypes.Message, height int64, key filTypes.TipSetKey) (string, error) {

	if msg == nil {
		return "", errors.New("malformed value")
	}

	// Shortcut 1 - Method "0" corresponds to "MethodSend"
	if msg.Method == 0 {
		return MethodSend, nil
	}

	// Shortcut 2 - Method "1" corresponds to "MethodConstructor"
	if msg.Method == 1 {
		return MethodConstructor, nil
	}

	actorName := p.getActorNameFromAddress(msg.To, height, key)

	actorMethods, ok := allMethods[actorName]
	if !ok {
		return "", errNotKnownActor
	}
	method, ok := actorMethods[msg.Method]
	if !ok {
		return UnknownStr, nil
	}
	return method.Name, nil
}

func (p *Parser) parseSend(msg *filTypes.Message) map[string]interface{} {
	metadata := make(map[string]interface{})
	metadata[ParamsKey] = msg.Params
	return metadata
}

func (p *Parser) unknownMetadata(msgParams, msgReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	if len(msgParams) > 0 {
		metadata[ParamsKey] = hex.EncodeToString(msgParams)
	}
	if len(msgReturn) > 0 {
		metadata[ReturnKey] = hex.EncodeToString(msgReturn)
	}
	return metadata, nil
}

func (p *Parser) emptyParamsAndReturn() (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}
