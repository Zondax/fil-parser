package evm

import "io"

type evmParams interface {
	UnmarshalCBOR(io.Reader) error
}

type evmReturn interface {
	UnmarshalCBOR(io.Reader) error
}

type createReturn interface {
	UnmarshalCBOR(io.Reader) error
}

type createReturnStruct[T any] struct {
	CreateReturn T
}
