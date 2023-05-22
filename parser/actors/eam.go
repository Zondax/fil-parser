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

func ParseEam(txType string, msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, msgCid cid.Cid, ethLogs []types.EthLog) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	switch txType {
	case parser.MethodConstructor:
		return emptyParamsAndReturn()
	case parser.MethodCreate:
		return parseCreate(msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreate2:
		return parseCreate2(msg.Params, msgRct.Return, msgCid)
	case parser.MethodCreateExternal:
		return parseCreateExternal(msg, msgRct, msgCid)
	case parser.UnknownStr:
		return unknownMetadata(msg.Params, msgRct.Return)
	}
	return metadata, nil
}

func parseEamReturn(rawReturn []byte) (cr eam.CreateReturn, err error) {
	reader := bytes.NewReader(rawReturn)
	err = cr.UnmarshalCBOR(reader)
	if err != nil {
		return cr, err
	}

	err = validateEamReturn(&cr)
	if err != nil {
		rawString := hex.EncodeToString(rawReturn)
		zap.S().Errorf("[parseEamReturn]- Detected invalid return bytes: %s. Raw: %s", err, rawString)
	}

	return cr, nil
}

func validateEamReturn(ret *eam.CreateReturn) error {
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

func newEamCreate(r eam.CreateReturn) parser.EamCreateReturn {
	return parser.EamCreateReturn{
		ActorId:       r.ActorID,
		RobustAddress: r.RobustAddress,
		EthAddress:    parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
	}
}

func parseCreate(rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})

	reader := bytes.NewReader(rawParams)
	var params eam.CreateParams
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	createReturn, err := parseEamReturn(rawReturn)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = newEamCreate(createReturn)
	appendCreatedEVMActor(eam.Return(createReturn), msgCid.String())

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, err
	}
	metadata[parser.EthHashKey] = ethHash.String()

	return metadata, nil
}

func parseCreate2(rawParams, rawReturn []byte, msgCid cid.Cid) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})

	reader := bytes.NewReader(rawParams)
	var params eam.Create2Params
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params

	createReturn, err := parseEamReturn(rawReturn)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = newEamCreate(createReturn)
	p.appendCreatedEVMActor(eam.Return(createReturn), msgCid.String())

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, err
	}
	metadata[parser.EthHashKey] = ethHash.String()

	return metadata, nil
}

func parseCreateExternal(msg *parser.LotusMessage, msgRct *parser.LotusMessageReceipt, msgCid cid.Cid) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(msg.Params[3:]) // TODO

	createExternalReturn, err := parseEamReturn(msgRct.Return)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = newEamCreate(createExternalReturn)
	appendCreatedEVMActor(eam.Return(createExternalReturn), msgCid.String())

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, err
	}
	metadata[parser.EthHashKey] = ethHash.String()

	return metadata, nil
}

func appendCreatedEVMActor(r eam.Return, msgCid string) {
	appendToAddresses(types.AddressInfo{
		Short:          parser.FilPrefix + strconv.FormatUint(r.ActorID, 10),
		Robust:         r.RobustAddress.String(),
		EthAddress:     parser.EthPrefix + hex.EncodeToString(r.EthAddress[:]),
		ActorType:      "evm",
		CreationTxHash: msgCid,
	})
}
