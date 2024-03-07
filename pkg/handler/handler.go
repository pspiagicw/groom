package handler

import (
	"github.com/pspiagicw/groom/pkg/argparse"
	"github.com/pspiagicw/groom/pkg/helper"
	"github.com/pspiagicw/groom/pkg/tasks"
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
		tasks.ListTasks(opts)
	} else {
		handleCmd(opts)
	}
}
func handleCmd(opts *argparse.Opts) {
	cmd := opts.Args[0]
	handleCmd, ok := handlers[cmd]
	if !ok {
		tasks.PerformTasks(opts.Args, opts.DryRun)
	} else {
		handleCmd(opts)
	}

}
