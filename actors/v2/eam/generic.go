package eam

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/zondax/fil-parser/parser/helper"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"

	typegen "github.com/whyrusleeping/cbor-gen"
)

func parseEamReturn[R typegen.CBORUnmarshaler](rawReturn []byte, r R) (R, error) {
	reader := bytes.NewReader(rawReturn)
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return r, err
	}

	err = validateEamReturn(r)
	if err != nil {
		rawString := hex.EncodeToString(rawReturn)
		return r, fmt.Errorf("[parseEamReturn]- Detected invalid return bytes: %s. Raw: %s", err, rawString)
	}

	return r, nil
}

func parseCreate[T typegen.CBORUnmarshaler, R typegen.CBORUnmarshaler](e *Eam, rawParams, rawReturn []byte, ec exitcode.ExitCode, msgCid cid.Cid, params T, r R, h *helper.Helper) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	if len(rawParams) > 0 {
		reader := bytes.NewReader(rawParams)

		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, nil, fmt.Errorf("error deserializing rawParams: %s - hex data: %s", err.Error(), hex.EncodeToString(rawParams))
		}
		metadata[parser.ParamsKey] = params
	}

	if len(rawReturn) > 0 {
		return handleReturnValue(e, rawReturn, ec, metadata, msgCid, r, h)
	}

	return metadata, nil, nil
}

func parseCreateExternal[T typegen.CBORUnmarshaler](e *Eam, rawParams, rawReturn []byte, ec exitcode.ExitCode, msgCid cid.Cid, params abi.CborBytes, r T, h *helper.Helper) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	if len(rawParams) > 0 {
		reader := bytes.NewReader(rawParams)
		metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(rawParams)
		if err := params.UnmarshalCBOR(reader); err != nil {
			return metadata, nil, fmt.Errorf("error deserializing rawParams: %s - hex data: %s", err.Error(), hex.EncodeToString(rawParams))
		}

		if reader.Len() == 0 { // This means that the reader has processed all the bytes
			metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(params)
		}

	}

	if len(rawReturn) > 0 {
		return handleReturnValue(e, rawReturn, ec, metadata, msgCid, r, h)
	}
	return metadata, nil, nil
}

func handleReturnValue[R typegen.CBORUnmarshaler](e *Eam, rawReturn []byte, ec exitcode.ExitCode, metadata map[string]interface{}, msgCid cid.Cid, r R, h *helper.Helper) (map[string]interface{}, *types.AddressInfo, error) {
	createReturn, err := parseEamReturn(rawReturn, r)
	if err != nil {
		return nil, nil, err
	}

	ethHash, createdEvmActor, cr, err := e.newEamCreate(createReturn, msgCid)
	metadata[parser.ReturnKey] = cr
	if err != nil {
		return metadata, nil, fmt.Errorf("error parsing createReturn: %s", err)
	}
	metadata[parser.EthHashKey] = ethHash

	if ec.IsSuccess() {
		h.GetActorsCache().StoreAddressInfo(*createdEvmActor)
	}

	return metadata, createdEvmActor, nil
}
