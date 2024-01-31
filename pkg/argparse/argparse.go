package argparse

import (
	"flag"
	"github.com/pspiagicw/groom/pkg/helper"
)

type Opts struct {
	SimpleListing bool
}

func ParseArguments(version string) ([]string, *Opts) {

	PrintHelp := func() {
		helper.PrintHelp(version)
	}
	opts := new(Opts)

	flag.BoolVar(&opts.SimpleListing, "simple", false, "Print simple listing")
	flag.Usage = PrintHelp
	flag.Parse()
	return flag.Args(), opts
}
