package controller

import (
	"fmt"
	"strings"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/parse"
)

func ListTasks() {

	tasks := parse.ParseTasks()

	if len(tasks) == 0 {
		goreland.LogFatal("No tasks declared.")
	}

	printTaskTable(tasks)
}

func printTaskTable(tasks map[string]*parse.Task) {
	fmt.Println("Tasks:")
	headers := []string{"Name", "Description", "Depends"}
	rows := buildRows(tasks)
	goreland.LogTable(headers, rows)

}

func buildRows(tasks map[string]*parse.Task) [][]string {
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
