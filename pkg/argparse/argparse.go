package argparse

import (
	"flag"
)

func ParseArguments(VERSION string) []string {

	flag.Parse()

	return flag.Args()
}
