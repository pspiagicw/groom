package argparse

import (
	"flag"
	"os"

	"github.com/pspiagicw/colorlog"
)

func ParseArguments(VERSION string) []string {
	version := flag.Bool("version", false, "Print version info.")

	flag.Parse()

	if *version {
		colorlog.LogSuccess("groom-make, Version: %s", VERSION)
		os.Exit(0)
	}

	requested := flag.Args()

	return requested
}
