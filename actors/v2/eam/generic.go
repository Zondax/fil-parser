package eam

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func parseEamReturn[R createReturn](rawReturn []byte, r R) (R, error) {
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

func parseCreate[T createParams, R createReturn](rawParams, rawReturn []byte, msgCid cid.Cid, params T, r R) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	if len(rawParams) > 0 {
		reader := bytes.NewReader(rawParams)

		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, nil, err
		}
		metadata[parser.ParamsKey] = params
	}

	if len(rawReturn) > 0 {
		return handleReturnValue(rawReturn, metadata, msgCid, r)
	}
	return metadata, nil, nil
}

func parseCreateExternal[T createReturn](rawParams, rawReturn []byte, msgCid cid.Cid, params abi.CborBytes, r T) (map[string]interface{}, *types.AddressInfo, error) {
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
		return handleReturnValue(rawReturn, metadata, msgCid, r)
	}
	return metadata, nil, nil
}

func handleReturnValue[R createReturn](rawReturn []byte, metadata map[string]interface{}, msgCid cid.Cid, r R) (map[string]interface{}, *types.AddressInfo, error) {
	createReturn, err := parseEamReturn[R](rawReturn, r)
	if err != nil {
		return nil, nil, err
	}

	ethHash, createdEvmActor, cr, err := newEamCreate(createReturn, msgCid)
	metadata[parser.ReturnKey] = cr
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.EthHashKey] = ethHash

	return metadata, createdEvmActor, nil
}
