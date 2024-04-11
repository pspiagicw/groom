package handle

import (
	"github.com/pspiagicw/groom/pkg/argparse"
	"github.com/pspiagicw/groom/pkg/help"
	"github.com/pspiagicw/groom/pkg/tasks"
)

var handlers = map[string]func(*argparse.Opts){
	"version": func(opts *argparse.Opts) {
		help.PrintVersion(opts.Version)
	},
	"help": func(opts *argparse.Opts) {
		help.HandleHelp(opts.Args[1:], opts.Version)
	},
}

func Handle(opts *argparse.Opts) {

	if opts.ExampleConfig {
		help.PrintExampleConfig()
	} else if len(opts.Args) == 0 {
		tasks.ListTasks(opts)
	} else {
		cmd := opts.Args[0]
		handleFunc, ok := handlers[cmd]
		if !ok {
			tasks.PerformTasks(opts.Args, opts.DryRun)
		} else {
			handleFunc(opts)
		}
	}

}
