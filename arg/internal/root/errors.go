package argroot

import (
	"errors"
)

var (
	errNilDestination = errors.New("argroot: nil destination")
	errNotParsedYet   = errors.New("argroot: not parsed yet")
)
