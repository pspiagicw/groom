package main

import (
	"github.com/pspiagicw/groom/pkg/argparse"
	"github.com/pspiagicw/groom/pkg/handle"
)

var VERSION string = "unversioned"

func main() {
	opts := argparse.ParseArguments(VERSION)
	handle.Handle(opts)
}
