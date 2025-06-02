package multisig

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-state-types/manifest"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/zondax/fil-parser/actors/v2/internal"

	"github.com/filecoin-project/go-state-types/abi"
	filTypes "github.com/filecoin-project/lotus/chain/types"

	"github.com/zondax/fil-parser/actors"
	"github.com/zondax/fil-parser/parser"
)

// innerProposeParams processes the parameters for a multisig proposal by:
// 1. Creating a new LotusMessage with the proposal details
// 2. Getting the actor and method name for the proposal
// 3. Parsing the proposal parameters using the actor's Parse method
func (m *Msig) innerProposeParams(
	msg *parser.LotusMessage, to address.Address, network string, height int64, method abi.MethodNum,
	proposeParams []byte, key filTypes.TipSetKey,
) (string, map[string]interface{}, error) {
	proposeMsg := &parser.LotusMessage{
		To:     to,
		From:   msg.From,
		Method: method,
		Cid:    msg.Cid,
		Params: proposeParams,
	}

	actor, proposedMethod, err := m.innerProposeMethod(proposeMsg, network, height, key)
	if err != nil {
		return "", nil, err
	}

	metadata, _, err := actor.Parse(context.Background(), network, height, proposedMethod, proposeMsg, &parser.LotusMessageReceipt{ExitCode: exitcode.Ok, Return: []byte{}}, msg.Cid, key)
	if err != nil {
		return "", nil, err
	}

	return proposedMethod, metadata, nil
}

// innerProposeMethod determines the actor and method name for a multisig proposal by:
// 1. Getting the actor name from the target address
// 2. Getting the appropriate actor implementation
// 3. Checking for common methods
// 4. Looking up the method name in the actor's method list
func (m *Msig) innerProposeMethod(
	msg *parser.LotusMessage, network string, height int64, key filTypes.TipSetKey,
) (actors.Actor, string, error) {
	_, actorName, err := m.helper.GetActorNameFromAddress(msg.To, height, key)
	if err != nil {
		return nil, "", err
	}
	var actor actors.Actor
	actor = m
	if actorName != manifest.MultisigKey {
		actor, err = internal.GetActor(actorName, m.logger, m.helper, m.metrics)
		if err != nil {
			return nil, "", err
		}
	}

	method, err := m.helper.CheckCommonMethods(msg, height, key)
	if err != nil {
		return nil, "", err
	}
	if method != "" {
		return actor, method, nil
	}

	actorMethods, err := actor.Methods(context.Background(), network, height)
	if err != nil {
		return nil, "", err
	}

	proposeMethod, ok := actorMethods[msg.Method]
	if !ok {
		return nil, "", fmt.Errorf("unrecognized propose method: %s for actor %s", method, actorName)
	}

	return actor, proposeMethod.Name, nil
}
