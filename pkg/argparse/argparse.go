package argparse

import (
	"flag"
	"github.com/pspiagicw/groom/pkg/helper"
)

func ParseArguments(version string) []string {

	PrintHelp := func() {
		helper.PrintHelp(version)
	}

	flag.Usage = PrintHelp
	flag.Parse()
	return flag.Args()
}
