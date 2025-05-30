package multisig

import (
	"io"
)

type multisigParams interface {
	UnmarshalCBOR(io.Reader) error
}
