package power

import "io"

type powerParams interface {
	UnmarshalCBOR(io.Reader) error
}

type powerReturn interface {
	UnmarshalCBOR(io.Reader) error
}
