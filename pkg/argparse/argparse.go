package argparse

import (
	"flag"
	"github.com/pspiagicw/groom/pkg/help"
)

type Opts struct {
	SimpleListing bool
	Args          []string
	Version       string
	ExampleConfig bool
	DryRun        bool
}

func ParseArguments(version string) *Opts {

	PrintHelp := func() {
		help.PrintHelp(version)
	}
	opts := new(Opts)

	flag.BoolVar(&opts.SimpleListing, "simple", false, "Print simple listing")
	flag.BoolVar(&opts.ExampleConfig, "example-config", false, "Print example config")
	flag.BoolVar(&opts.DryRun, "dry-run", false, "Dry run comands.")
	flag.Usage = PrintHelp
	flag.Parse()
	opts.Version = version
	opts.Args = flag.Args()
	return opts
}
