package miner

import "io"

type minerParam interface {
	UnmarshalCBOR(io.Reader) error
}

type minerReturn interface {
	UnmarshalCBOR(io.Reader) error
}
