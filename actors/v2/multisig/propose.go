package multisig

import (
	"context"
	"strings"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/actors/v2/internal"

	"github.com/filecoin-project/go-state-types/abi"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
)

// parseInnerProposeMsg processes the parameters for a multisig proposal by:
// 1. Creating a new LotusMessage with the proposal details
// 2. Getting the actor and method name for the proposal
// 3. Parsing the proposal parameters using the actor's Parse method
func (m *Msig) parseInnerProposeMsg(
	msg *parser.LotusMessage, to address.Address, network string, height int64, method abi.MethodNum,
	proposeParams, proposeReturn []byte, key filTypes.TipSetKey,
) (string, map[string]interface{}, error) {
	proposeMsg := &parser.LotusMessage{
		To:     to,
		From:   msg.From,
		Method: method,
		Cid:    msg.Cid,
		Params: proposeParams,
	}

	proposeMsgRct := &parser.LotusMessageReceipt{ExitCode: exitcode.Ok, Return: proposeReturn}

	actor, proposedMethod, err := m.innerProposeMethod(proposeMsg, network, height, key)
	if err != nil {
		return "", nil, err
	}

	metadata, _, err := actor.Parse(context.Background(), network, height, proposedMethod, proposeMsg, proposeMsgRct, msg.Cid, key)
	if err != nil {
		return "", nil, err
	}

	return proposedMethod, metadata, nil
}

// innerProposeMethod determines the actor and method name for a multisig proposal by:
// 1. Getting the actor name from the target address
// 2. Using the methodNameFn to get the methodName from the methodNum for the actor.
func (m *Msig) innerProposeMethod(
	msg *parser.LotusMessage, network string, height int64, key filTypes.TipSetKey,
) (actors.Actor, string, error) {
	_, actorName, err := m.helper.GetActorNameFromAddress(msg.To, height, key)
	if err != nil {
		return nil, "", err
	}
	var actor actors.Actor
	if strings.Contains(actorName, manifest.MultisigKey) {
		actor = m
	} else {
		actor, err = internal.GetActor(actorName, m.logger, m.helper, m.metrics)
		if err != nil {
			return nil, "", err
		}
	}

	methodName, err := m.methodNameFn(context.Background(), msg.Method, actorName, height, network, m.helper, m.logger)
	if err != nil {
		return nil, "", err
	}
	return actor, methodName, nil
}
