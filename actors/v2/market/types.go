package market

import "io"

type marketParam interface {
	UnmarshalCBOR(io.Reader) error
}

type marketReturn interface {
	UnmarshalCBOR(io.Reader) error
}
