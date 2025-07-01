package actors

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/zondax/fil-parser/parser"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v11/eam"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/types"
)

// TODO: do we need ethLogs?
func (p *ActorParser) ParseEam(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	var err error
	switch txType {
	case parser.MethodConstructor:
		metadata, err = p.emptyParamsAndReturn()
	case parser.MethodCreate:
		return p.parseCreate(msg, msgRct, mainMsgCid)
	case parser.MethodCreate2:
		return p.parseCreate2(msg, msgRct, mainMsgCid)
	case parser.MethodCreateExternal:
		return p.parseCreateExternal(msg, msgRct, mainMsgCid)
	case parser.UnknownStr:
		metadata, err = p.unknownMetadata(msg.Params, msgRct.Return)
	default:
		err = parser.ErrUnknownMethod
	}
	return metadata, nil, err
}

func (p *ActorParser) parseEamReturn(rawReturn []byte, method string) (cr eam.CreateReturn, err error) {
	reader := bytes.NewReader(rawReturn)
	err = cr.UnmarshalCBOR(reader)
	if err != nil {
		return cr, err
	}

	err = p.validateEamReturn(&cr)
	if err != nil {
		rawString := hex.EncodeToString(rawReturn)
		_ = p.metrics.UpdateActorMethodErrorMetric(manifest.EamKey, method)
		p.logger.Errorf("[parseEamReturn]- Detected invalid return bytes: %s. Raw: %s", err, rawString)
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

func (p *ActorParser) parseCreate(msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})

	reader := bytes.NewReader(msg.Params)
	var params eam.CreateParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ParamsKey] = params

	createReturn, err := p.parseEamReturn(msgRct.Return, parser.MethodCreate)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ReturnKey] = p.newEamCreate(createReturn)

	ethHash, err := ethtypes.EthHashFromCid(mainMsgCid)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.EthHashKey] = ethHash.String()

	r := eam.Return(createReturn)
	createdEvmActor := &types.AddressInfo{
		Short:         parser.FilPrefix + strconv.FormatUint(r.ActorID, 10),
		Robust:        r.RobustAddress.String(),
		EthAddress:    parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
		ActorType:     manifest.EvmKey,
		CreationTxCid: mainMsgCid.String(),
	}

	if msgRct.ExitCode.IsSuccess() {
		p.helper.GetActorsCache().StoreAddressInfo(*createdEvmActor)
	}

	return metadata, createdEvmActor, nil
}

func (p *ActorParser) parseCreate2(msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})

	reader := bytes.NewReader(msg.Params)
	var params eam.Create2Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ParamsKey] = params

	createReturn, err := p.parseEamReturn(msgRct.Return, parser.MethodCreate2)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ReturnKey] = p.newEamCreate(createReturn)

	ethHash, err := ethtypes.EthHashFromCid(mainMsgCid)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.EthHashKey] = ethHash.String()
	r := eam.Return(createReturn)
	createdEvmActor := &types.AddressInfo{
		Short:         parser.FilPrefix + strconv.FormatUint(r.ActorID, 10),
		Robust:        r.RobustAddress.String(),
		EthAddress:    parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
		ActorType:     manifest.EvmKey,
		CreationTxCid: mainMsgCid.String(),
	}

	if msgRct.ExitCode.IsSuccess() {
		p.helper.GetActorsCache().StoreAddressInfo(*createdEvmActor)
	}

	return metadata, createdEvmActor, nil
}

func (p *ActorParser) parseCreateExternal(msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, mainMsgCid cid.Cid) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(msg.Params)
	metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(msg.Params)

	var params abi.CborBytes
	if err := params.UnmarshalCBOR(reader); err != nil {
		_ = p.metrics.UpdateActorMethodErrorMetric(manifest.EamKey, parser.MethodCreateExternal)
		p.logger.Warn(fmt.Sprintf("error deserializing rawParams: %s - hex data: %s", err.Error(), hex.EncodeToString(msg.Params)))
	}

	if reader.Len() == 0 { // This means that the reader has processed all the bytes
		metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(params)
	}

	createExternalReturn, err := p.parseEamReturn(msgRct.Return, parser.MethodCreateExternal)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.ReturnKey] = p.newEamCreate(createExternalReturn)

	ethHash, err := ethtypes.EthHashFromCid(mainMsgCid)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.EthHashKey] = ethHash.String()
	r := eam.Return(createExternalReturn)
	createdEvmActor := &types.AddressInfo{
		Short:         parser.FilPrefix + strconv.FormatUint(r.ActorID, 10),
		Robust:        r.RobustAddress.String(),
		EthAddress:    parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
		ActorType:     manifest.EvmKey,
		CreationTxCid: mainMsgCid.String(),
	}

	if msgRct.ExitCode.IsSuccess() {
		p.helper.GetActorsCache().StoreAddressInfo(*createdEvmActor)
	}

	return metadata, createdEvmActor, nil
}
