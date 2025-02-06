package cron

import (
	"io"
)

type cronConstructorParams interface {
	UnmarshalCBOR(r io.Reader) error
}
