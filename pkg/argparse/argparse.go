package argparse

import (
	"flag"
	"os"

	"github.com/pspiagicw/goreland"
)

func ParseArguments(VERSION string) []string {

	version := flag.Bool("version", false, "Print version info.")

	flag.Parse()

	if *version {
		goreland.LogSuccess("groom-make, Version: %s", VERSION)
		os.Exit(0)
	}

	requested := flag.Args()

	return requested
}
