package argroot

import (
	"github.com/reiver/hydrargyrum/lib/opt/str"

	"flag"
)

var (
	value optstr.String
)

func init() {
	flag.Var(&value, "root", "path to docoument root")
}

func Receive(dst *optstr.String) error {
	if !flag.Parsed() {
		return errNotParsedYet
	}

	if nil == dst {
		return errNilDestination
	}

	optstr.Push(dst, value)

	return nil
}
