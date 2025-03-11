package verifiedRegistry

import "io"

type verifiedRegistryParams interface {
	UnmarshalCBOR(io.Reader) error
}

type verifiedRegistryReturn interface {
	UnmarshalCBOR(io.Reader) error
}
