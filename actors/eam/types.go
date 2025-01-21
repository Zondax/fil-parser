package eam

import "io"

type createReturn interface {
	UnmarshalCBOR(io.Reader) error
}
