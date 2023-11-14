package controller

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/constants"
	"github.com/pspiagicw/groom/pkg/execute"
	"github.com/pspiagicw/groom/pkg/parse"
)

func ExecuteTasks(requests []string) {

	assertFile()
	tasks := parse.ParseTasks()
	if len(requests) == 0 {
		listTasks(tasks)

	}
	executeTasks(requests, tasks)

}
func assertFile() {
	_, err := os.Stat(constants.TASK_FILE)
	if err != nil {
		goreland.LogError("Error while reading groom.toml: %v", err)
		goreland.LogFatal("Make sure the current directory has the `groom.toml` file.")
	}
}

func getEnvironmentString(env []string) string {

	var out bytes.Buffer
	for _, value := range env {

		out.WriteString(" ")
		out.WriteString(value)

	}

	return out.String()
}
func cleanComponent(component string) string {
	component = strings.ReplaceAll(component, "\"", "")

	return component

}
func splitCommandString(command string) []string {

	if len(command) == 0 {
		return []string{}
	}

	components := make([]string, 0)
	startIndex := 0
	currentIndex := 0

	parenStack := make([]byte, 0)

	for currentIndex < len(command) {
		// fmt.Println(parenStack, currentIndex, startIndex)
		if command[currentIndex] == ' ' && len(parenStack) == 0 {
			component := command[startIndex:currentIndex]
			startIndex = currentIndex + 1

			if component != "" {
				components = append(components, cleanComponent(component))
			}
		} else if command[currentIndex] == '\'' {
			if len(parenStack) == 0 {
				parenStack = append(parenStack, '\'')
			} else {
				lastElement := parenStack[len(parenStack)-1]

				if lastElement == command[currentIndex] {
					parenStack = parenStack[:len(parenStack)-1]
				} else {
					parenStack = append(parenStack, command[currentIndex])
				}
			}
		} else if command[currentIndex] == '"' {
			if len(parenStack) == 0 {
				parenStack = append(parenStack, '"')
			} else {
				lastElement := parenStack[len(parenStack)-1]

				if lastElement == command[currentIndex] {

					parenStack = parenStack[:len(parenStack)-1]

				} else {
					parenStack = append(parenStack, command[currentIndex])
				}
			}
		}
		currentIndex += 1
	}

	lastComponent := command[startIndex:currentIndex]

	components = append(components, lastComponent)

	return components

}
func executeTasks(requested []string, tasks map[string]*parse.Task) {
	for _, request := range requested {
		task, ok := tasks[request]

		if !ok {
			goreland.LogFatal("No task named %s", request)
		}

		if len(task.Depends) != 0 {
			goreland.LogInfo("Executing dependencies for [%s]", request)
			executeTasks(task.Depends, tasks)
		}

		if task.Command == "" && len(task.Commands) == 0 {
			goreland.LogFatal("No command/commands specified for [%s]!", request)
		}

		environmentString := getEnvironmentString(task.Environment)

		if len(task.Commands) != 0 {
			for _, subtask := range task.Commands {
				components := splitCommandString(subtask)

				if len(components) == 0 {
					goreland.LogFatal("Command is not provided for task %s", request)
				}
				fmt.Printf(constants.LOG_PREFIX+"%s =>"+environmentString+" %s\n", request, subtask)

				err := execute.Execute(components[0], components[1:], task.Environment)

				if err != nil {
					goreland.LogError("exited with a error: " + err.Error())
				}

			}

		} else {
			components := splitCommandString(task.Command)

			if len(components) == 0 {
				goreland.LogError("Command is not provided for task %s", request)
			}

			fmt.Printf(constants.LOG_PREFIX+"%s =>"+environmentString+" %s\n", request, task.Command)

			err := execute.Execute(components[0], components[1:], task.Environment)

			if err != nil {
				goreland.LogError("exited with a error:" + err.Error())
			}

		}
	}
}
func listTasks(tasks map[string]*parse.Task) {

	if len(tasks) == 0 {
		goreland.LogFatal("No tasks declared.")
	}

	fmt.Println("Tasks:")

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

	headers := []string{"Name", "Description", "Depends"}
	goreland.LogTable(headers, rows)
}
