package evm

import (
	"bytes"

	"github.com/zondax/fil-parser/parser"
)

func parse[T evmParams, R evmReturn](rawParams, rawReturn []byte, customReturn bool) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	if !customReturn {
		return metadata, nil
	}

	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func invokeContractDelegate[T evmParams, R evmReturn](rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func getByteCode[R evmReturn](key string, raw []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	var r R
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[key] = r
	return metadata, nil
}

func getStorageAt[T evmParams, R evmReturn](rawParams, rawReturn []byte) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawParams)
	var params T
	err := params.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ParamsKey] = params
	reader = bytes.NewReader(rawReturn)
	var r R
	err = r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}
