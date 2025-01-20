package datacap

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	"github.com/zondax/fil-parser/parser"
)

type granularityReturn = unmarshalCBOR

func granularityExported(rawReturn []byte, r granularityReturn) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

func granularityExportedv11(rawReturn []byte) (map[string]interface{}, error) {
	return granularityExported(rawReturn, datacap.GranularityReturn{})
}

func granularityExportedv14(rawReturn []byte) (map[string]interface{}, error) {
	return granularityExported(rawReturn, datacap.GranularityReturn{})
}
