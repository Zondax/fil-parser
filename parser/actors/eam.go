package actors

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/zondax/fil-parser/parser"
	"strconv"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v11/eam"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"go.uber.org/zap"

	"github.com/zondax/fil-parser/types"
)

// TODO: do we need ethLogs?
func (p *ActorParser) ParseEam(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	var err error
	switch txType {
	case parser.MethodConstructor:
		metadata, err = p.emptyParamsAndReturn()
	case parser.MethodCreate:
		return p.parseCreate(msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreate2:
		return p.parseCreate2(msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreateExternal:
		return p.parseCreateExternal(msg, msgRct, msgCid)
	case parser.UnknownStr:
		metadata, err = p.unknownMetadata(msg.Params, msgRct.Return)
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, nil, err
}

func (p *ActorParser) parseEamReturn(rawReturn []byte) (cr eam.CreateReturn, err error) {
	reader := bytes.NewReader(rawReturn)
	err = cr.UnmarshalCBOR(reader)
	if err != nil {
		return cr, err
	}

	err = p.validateEamReturn(&cr)
	if err != nil {
		rawString := hex.EncodeToString(rawReturn)
		zap.S().Errorf("[parseEamReturn]- Detected invalid return bytes: %s. Raw: %s", err, rawString)
	}

	return cr, nil
}

func (p *ActorParser) validateEamReturn(ret *eam.CreateReturn) error {
	if ret == nil {
		return fmt.Errorf("input is nil")
	}

	if ret.RobustAddress == nil {
		emptyAdd, _ := address.NewFromString("")
		ret.RobustAddress = &emptyAdd
		return fmt.Errorf("RobustAddress field is nil. Replaced with empty address")
	}

	return nil
}

func (p *ActorParser) newEamCreate(r eam.CreateReturn) parser.EamCreateReturn {
	return parser.EamCreateReturn{
		ActorId:       r.ActorID,
		RobustAddress: r.RobustAddress,
		EthAddress:    parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
	}
}

func (p *ActorParser) parseCreate(rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})

	reader := bytes.NewReader(rawParams)
	var params eam.CreateParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ParamsKey] = params

	createReturn, err := p.parseEamReturn(rawReturn)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ReturnKey] = p.newEamCreate(createReturn)
	//p.appendCreatedEVMActor(eam.Return(createReturn), msgCid.String())

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.EthHashKey] = ethHash.String()

	r := eam.Return(createReturn)
	createdEvmActor := &types.AddressInfo{
		Short:          parser.FilPrefix + strconv.FormatUint(r.ActorID, 10),
		Robust:         r.RobustAddress.String(),
		EthAddress:     parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
		ActorType:      "evm",
		CreationTxHash: msgCid.String(),
	}
	return metadata, createdEvmActor, nil
}

func (p *ActorParser) parseCreate2(rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})

	reader := bytes.NewReader(rawParams)
	var params eam.Create2Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ParamsKey] = params

	createReturn, err := p.parseEamReturn(rawReturn)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ReturnKey] = p.newEamCreate(createReturn)
	//p.appendCreatedEVMActor(eam.Return(createReturn), msgCid.String())

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.EthHashKey] = ethHash.String()
	r := eam.Return(createReturn)
	createdEvmActor := &types.AddressInfo{
		Short:          parser.FilPrefix + strconv.FormatUint(r.ActorID, 10),
		Robust:         r.RobustAddress.String(),
		EthAddress:     parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
		ActorType:      "evm",
		CreationTxHash: msgCid.String(),
	}
	return metadata, createdEvmActor, nil
}

func (p *ActorParser) parseCreateExternal(msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, msgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(msg.Params[3:]) // TODO

	createExternalReturn, err := p.parseEamReturn(msgRct.Return)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ReturnKey] = p.newEamCreate(createExternalReturn)
	//p.appendCreatedEVMActor(eam.Return(createExternalReturn), msgCid.String())

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.EthHashKey] = ethHash.String()
	r := eam.Return(createExternalReturn)
	createdEvmActor := &types.AddressInfo{
		Short:          parser.FilPrefix + strconv.FormatUint(r.ActorID, 10),
		Robust:         r.RobustAddress.String(),
		EthAddress:     parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
		ActorType:      "evm",
		CreationTxHash: msgCid.String(),
	}
	return metadata, createdEvmActor, nil
}

//func (p *ActorParser) appendCreatedEVMActor(r eam.Return, msgCid string) {
//	appendToAddresses(types.AddressInfo{
//		Short:          parser.FilPrefix + strconv.FormatUint(r.ActorID, 10),
//		Robust:         r.RobustAddress.String(),
//		EthAddress:     parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
//		ActorType:      "evm",
//		CreationTxHash: msgCid,
//	})
//}
