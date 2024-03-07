package main

import (
	"github.com/pspiagicw/groom/pkg/argparse"
	"github.com/pspiagicw/groom/pkg/handler"
)

var VERSION string = "unversioned"

func main() {
	opts := argparse.ParseArguments(VERSION)
	handler.HandleArgs(opts)
}
