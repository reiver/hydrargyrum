package argaddress

import (
	"errors"
)

var (
	errNilDestination = errors.New("argaddress: nil destination")
	errNotParsedYet   = errors.New("argaddress: not parsed yet")
)
