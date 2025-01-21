package eam

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/ipfs/go-cid"
	"github.com/zondax/fil-parser/parser"
	"github.com/zondax/fil-parser/types"
)

func parseEamReturn[R createReturn](rawReturn []byte) (R, error) {
	var cr R

	reader := bytes.NewReader(rawReturn)
	err := cr.UnmarshalCBOR(reader)
	if err != nil {
		return cr, err
	}

	err = validateEamReturn(cr)
	if err != nil {
		rawString := hex.EncodeToString(rawReturn)
		return cr, fmt.Errorf("[parseEamReturn]- Detected invalid return bytes: %s. Raw: %s", err, rawString)
	}

	return cr, nil
}

func parseCreate[T createReturn](rawParams, rawReturn []byte, msgCid cid.Cid, isExternal bool) (map[string]interface{}, *types.AddressInfo, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)

	if isExternal {
		metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(rawParams)
		var params abi.CborBytes
		if err := params.UnmarshalCBOR(reader); err != nil {
			return metadata, nil, fmt.Errorf("error deserializing rawParams: %s - hex data: %s", err.Error(), hex.EncodeToString(rawParams))
		}

		if reader.Len() == 0 { // This means that the reader has processed all the bytes
			metadata[parser.ParamsKey] = parser.EthPrefix + hex.EncodeToString(params)
		}
	} else {
		var params T
		err := params.UnmarshalCBOR(reader)
		if err != nil {
			return metadata, nil, err
		}
	}

	createReturn, err := parseEamReturn[T](rawReturn)
	if err != nil {
		return metadata, nil, err
	}

	metadata[parser.ReturnKey] = newEamCreate(createReturn)

	ethHash, err := ethtypes.EthHashFromCid(msgCid)
	if err != nil {
		return metadata, nil, err
	}
	metadata[parser.EthHashKey] = ethHash.String()

	r := newEamCreate(createReturn)
	createdEvmActor := &types.AddressInfo{
		Short:         parser.FilPrefix + strconv.FormatUint(r.ActorId, 10),
		Robust:        r.RobustAddress.String(),
		EthAddress:    parser.EthPrefix + r.EthAddress,
		ActorType:     "evm",
		CreationTxCid: msgCid.String(),
	}
	return metadata, createdEvmActor, nil
}
