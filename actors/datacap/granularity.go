package datacap

import (
	"bytes"
	"fmt"

	datacapv10 "github.com/filecoin-project/go-state-types/builtin/v10/datacap"
	datacapv11 "github.com/filecoin-project/go-state-types/builtin/v11/datacap"
	datacapv12 "github.com/filecoin-project/go-state-types/builtin/v12/datacap"
	datacapv13 "github.com/filecoin-project/go-state-types/builtin/v13/datacap"
	datacapv14 "github.com/filecoin-project/go-state-types/builtin/v14/datacap"
	datacapv15 "github.com/filecoin-project/go-state-types/builtin/v15/datacap"
	"github.com/zondax/fil-parser/parser"
)

type granularityReturn = unmarshalCBOR

func GranularityExported(height uint64, rawReturn []byte) (map[string]interface{}, error) {
	switch height {
	case 9:
		return nil, fmt.Errorf("not supported")
	case 10:
		return granularityExportedv10(rawReturn)
	case 11:
		return granularityExportedv11(rawReturn)
	case 12:
		return granularityExportedv12(rawReturn)
	case 13:
		return granularityExportedv13(rawReturn)
	case 14:
		return granularityExportedv14(rawReturn)
	case 15:
		return granularityExportedv15(rawReturn)
	}
	return nil, fmt.Errorf("not supported")
}

func granularityGeneric[T granularityReturn](rawReturn []byte, r T) (map[string]interface{}, error) {
	metadata := make(map[string]interface{})
	reader := bytes.NewReader(rawReturn)
	err := r.UnmarshalCBOR(reader)
	if err != nil {
		return metadata, err
	}
	metadata[parser.ReturnKey] = r
	return metadata, nil
}

// Granularity Exported

func granularityExportedv10(rawReturn []byte) (map[string]interface{}, error) {
	var r datacapv10.GranularityReturn
	return granularityGeneric[*datacapv10.GranularityReturn](rawReturn, &r)
}

func granularityExportedv11(rawReturn []byte) (map[string]interface{}, error) {
	var r datacapv11.GranularityReturn
	return granularityGeneric[*datacapv11.GranularityReturn](rawReturn, &r)
}

func granularityExportedv12(rawReturn []byte) (map[string]interface{}, error) {
	var r datacapv12.GranularityReturn
	return granularityGeneric[*datacapv12.GranularityReturn](rawReturn, &r)
}

func granularityExportedv13(rawReturn []byte) (map[string]interface{}, error) {
	var r datacapv13.GranularityReturn
	return granularityGeneric[*datacapv13.GranularityReturn](rawReturn, &r)
}

func granularityExportedv14(rawReturn []byte) (map[string]interface{}, error) {
	var r datacapv14.GranularityReturn
	return granularityGeneric[*datacapv14.GranularityReturn](rawReturn, &r)
}

func granularityExportedv15(rawReturn []byte) (map[string]interface{}, error) {
	var r datacapv15.GranularityReturn
	return granularityGeneric[*datacapv15.GranularityReturn](rawReturn, &r)
}
