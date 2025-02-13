package actors

import "errors"

var ErrUnsupportedHeight = errors.New("unsupported height")
var ErrInvalidHeightForMethod = errors.New("invalid height for method")
