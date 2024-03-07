package argparse

import (
	"flag"
	"github.com/pspiagicw/groom/pkg/helper"
)

type Opts struct {
	SimpleListing bool
	Args          []string
	Version       string
	ExampleConfig bool
}

func ParseArguments(version string) *Opts {

	PrintHelp := func() {
		helper.PrintHelp(version)
	}
	opts := new(Opts)

	flag.BoolVar(&opts.SimpleListing, "simple", false, "Print simple listing")
	flag.BoolVar(&opts.ExampleConfig, "example-config", false, "Print example config")
	flag.Usage = PrintHelp
	flag.Parse()
	opts.Version = version
	opts.Args = flag.Args()
	return opts
}
