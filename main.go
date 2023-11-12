package main

import (
	"github.com/pspiagicw/groom/pkg/argparse"
	"github.com/pspiagicw/groom/pkg/controller"
)

var VERSION string

func main() {

	requestedTasks := argparse.ParseArguments(VERSION)
	controller.ExecuteTasks(requestedTasks)

}
