package account

import (
	"io"
)

type authenticateMessageParams interface {
	UnmarshalCBOR(r io.Reader) error
}

type authenticateMessageReturn interface {
	UnmarshalCBOR(r io.Reader) error
}
