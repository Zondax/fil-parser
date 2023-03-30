package parser

import "errors"

var (
	errUnknownMethod = errors.New("not known method")
	errBlockHash     = errors.New("unable to get block hash")
	errNotValidActor = errors.New("not a valid actor")
	errNotKnownActor = errors.New("actor is unknown")
)
