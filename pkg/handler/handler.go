package handler

import (
	"os"

	"github.com/pspiagicw/groom/pkg/controller"
	"github.com/pspiagicw/groom/pkg/helper"
)

var handlers = map[string]func(string){
	"version": helper.PrintVersion,
	"help":    helper.PrintHelp,
}

func HandleArgs(args []string, version string) {

	if len(args) == 0 {
		controller.ListTasks()
		os.Exit(0)
	}

	cmd := args[0]

	handlerFunc, exists := handlers[cmd]

	if !exists {
		controller.PerformTasks(args)
	} else {
		handlerFunc(version)
	}
}
