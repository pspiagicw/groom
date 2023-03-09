package main

import (
	"github.com/golang-groom/groom-make/pkg/argparse"
	"github.com/golang-groom/groom-make/pkg/controller"
)

var VERSION string

func main() {

	requestedTasks := argparse.ParseArguments(VERSION)
	controller.ExecuteTasks(requestedTasks)

}
