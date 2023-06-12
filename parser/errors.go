package parser

import "errors"

var (
	ErrUnknownMethod  = errors.New("not known method")
	ErrBlockHash      = errors.New("unable to get block hash")
	ErrNotValidActor  = errors.New("not a valid actor")
	ErrNotKnownActor  = errors.New("actor is unknown")
	ErrInvalidType    = errors.New("invalid trace version")
	ErrInvalidVersion = errors.New("invalid trace version")
)
