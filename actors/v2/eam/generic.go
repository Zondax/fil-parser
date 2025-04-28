package eam

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/zondax/fil-parser/parser/helper"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
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

func parseCreate[T typegen.CBORUnmarshaler, R typegen.CBORUnmarshaler](rawParams, rawReturn []byte, msgCid cid.Cid, params T, r R, h *helper.Helper) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	if len(rawParams) > 0 {
		reader := bytes.NewReader(rawParams)

		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, nil, fmt.Errorf("error deserializing rawParams: %s - hex data: %s", err.Error(), hex.EncodeToString(rawParams))
		}
		metadata[parser.ParamsKey] = params
	}
	var createdEvmActor *types.AddressInfo

	if len(rawReturn) > 0 {
		var err error
		metadata, createdEvmActor, err = handleReturnValue(rawReturn, metadata, msgCid, r, h)
		if err != nil {
			return metadata, nil, err
		}
	}

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, createdEvmActor, fmt.Errorf("error getting ethHash: %s", err)
	}
	metadata[parser.EthHashKey] = ethHash.String()

	return metadata, createdEvmActor, nil
}

func parseCreateExternal[T typegen.CBORUnmarshaler](rawParams, rawReturn []byte, msgCid cid.Cid, params abi.CborBytes, r T, h *helper.Helper) (map[string]interface{}, *types.AddressInfo, error) {
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
	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, nil, fmt.Errorf("error getting ethHash: %s", err)
	}
	metadata[parser.EthHashKey] = ethHash.String()

	if len(rawReturn) > 0 {
		return handleReturnValue(rawReturn, metadata, msgCid, r, h)
	}
	return metadata, nil, nil
}

func handleReturnValue[R typegen.CBORUnmarshaler](rawReturn []byte, metadata map[string]interface{}, msgCid cid.Cid, r R, h *helper.Helper) (map[string]interface{}, *types.AddressInfo, error) {
	createReturn, err := parseEamReturn(rawReturn, r)
	if err != nil {
		return nil, nil, err
	}
	createdEvmActor, cr := newEamCreate(createReturn, msgCid)
	metadata[parser.ReturnKey] = cr

	h.GetActorsCache().StoreAddressInfoAddress(*createdEvmActor)

	return metadata, createdEvmActor, nil
}
