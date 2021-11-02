package argaddress

import (
	"github.com/reiver/hydrargyrum/lib/opt/str"

	"flag"
)

var (
	value optstr.String
)

func init() {
	flag.Var(&value, "address", "The address (host & port) to make the request from  — ex: “example.com:1961”, “192.0.2.1:8008”")
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
