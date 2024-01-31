package main

import (
	"github.com/pspiagicw/groom/pkg/argparse"
	"github.com/pspiagicw/groom/pkg/handler"
)

var VERSION string

func main() {
	args, opts := argparse.ParseArguments(VERSION)
	handler.HandleArgs(args, VERSION, opts)
}
