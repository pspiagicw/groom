package handle

import (
	"os"

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

	checkExampleConfig(opts)
	checkArgLen(opts)

	cmd := opts.Args[0]

	handleFunc := handlers[cmd]

	if handleFunc == nil {
		tasks.PerformTasks(opts)
	} else {
		handleFunc(opts)
	}

}
func checkArgLen(opts *argparse.Opts) {
	if len(opts.Args) == 0 {
		tasks.ListTasks(opts)
		os.Exit(0)
	}
}
func checkExampleConfig(opts *argparse.Opts) {
	if opts.ExampleConfig {
		help.PrintExampleConfig()
		os.Exit(0)
	}
}
