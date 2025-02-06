package init

import "io"

type constructorParams interface {
	UnmarshalCBOR(io.Reader) error
}

type execReturn interface {
	UnmarshalCBOR(io.Reader) error
}
