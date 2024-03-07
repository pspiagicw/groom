package handler

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/argparse"
	"github.com/pspiagicw/groom/pkg/controller"
	"github.com/pspiagicw/groom/pkg/helper"
)

var handlers = map[string]func(*argparse.Opts){
	"version": func(opts *argparse.Opts) {
		helper.PrintVersion(opts.Version)
	},
	"help": func(opts *argparse.Opts) {
		helper.HandleHelp(opts.Args[1:], opts.Version)
	},
}

func HandleArgs(opts *argparse.Opts) {

	if opts.ExampleConfig {
		helper.PrintExampleConfig()

	} else if len(opts.Args) == 0 {
		controller.ListTasks(opts)
	} else {
		cmd := opts.Args[0]

		handleCmd, ok := handlers[cmd]

		if !ok {
			goreland.LogFatal("No such command")
		}

		handleCmd(opts)
	}
}
