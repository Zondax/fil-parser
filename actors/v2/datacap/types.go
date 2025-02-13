package datacap

import "io"

type unmarshalCBOR interface {
	UnmarshalCBOR(io.Reader) error
}

type datacapParams interface {
	unmarshalCBOR
}

type datacapReturn interface {
	unmarshalCBOR
}
