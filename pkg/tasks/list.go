package tasks

import (
	"fmt"
	"strings"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/argparse"
	"github.com/pspiagicw/groom/pkg/config"
)

func ListTasks(opts *argparse.Opts) {

	groomConfig := config.ParseConfig()

	if len(groomConfig.Tasks) == 0 {
		goreland.LogFatal("No tasks declared.")
	}

	if !opts.SimpleListing {
		printTaskTable(groomConfig.Tasks)
	} else {
		printTaskList(groomConfig.Tasks)
	}
}

func printTaskList(taskList map[string]*config.Task) {
	for name, _ := range taskList {
		fmt.Println(name)
	}
}

func printTaskTable(taskList map[string]*config.Task) {
	fmt.Println("Tasks:")
	headers := []string{"Name", "Description", "Depends"}
	rows := buildRows(taskList)
	goreland.LogTable(headers, rows)

}

func buildRows(tasks map[string]*config.Task) [][]string {
	rows := [][]string{}

	for name, task := range tasks {
		description := task.Description

		if description == "" {
			description = "No description provided"
		}

		deps := strings.Join(task.Depends, ",")
		if deps == "" {
			deps = "No dependencies"
		}

		rows = append(rows, []string{name, description, deps})
	}

	return rows
}
