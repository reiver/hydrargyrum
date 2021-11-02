package arg

import (
	"github.com/reiver/hydrargyrum/lib/opt/str"

	"github.com/reiver/hydrargyrum/arg/internal/address"
	"github.com/reiver/hydrargyrum/arg/internal/help"
	"github.com/reiver/hydrargyrum/arg/internal/loglevel"
	"github.com/reiver/hydrargyrum/arg/internal/root"

	"flag"
)

var (
	Address optstr.String
	Help bool
	LogLevel uint8 = 0
	Root optstr.String
)

func init() {

	flag.Parse()

	{
		var err error

		// -- help
		err = arghelp.Receive(&Help)
		if nil != err {
			panic(err)
		}

		// -v -vv -vvv -vvvv -vvvvv -vvvvvv
		err = argloglevel.Receive(&LogLevel)
		if nil != err {
			panic(err)
		}

		// --address
		err = argaddress.Receive(&Address)
		if nil != err {
			panic(err)
		}

		// --root
		err = argroot.Receive(&Root)
		if nil != err {
			panic(err)
		}
	}
}
