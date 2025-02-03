package eam

import "io"

type createReturn interface {
	UnmarshalCBOR(io.Reader) error
}

type createParams interface {
	UnmarshalCBOR(io.Reader) error
}
