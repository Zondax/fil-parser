package reward

import "io"

type rewardParams interface {
	UnmarshalCBOR(io.Reader) error
}
