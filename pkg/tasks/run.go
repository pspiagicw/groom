package tasks

import (
	"strings"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/config"
	"github.com/pspiagicw/groom/pkg/execute"
)

func runDependencies(task *config.Task, taskList map[string]*config.Task) {
	goreland.LogInfo("Executing dependencies for [%s]", task.Name)
	executeTasks(task.Depends, taskList, false)
}

func runCommands(task *config.Task) {
	for _, command := range task.Commands {
		runCommand(task.Environment, command, task.Name)
	}

}
func runCommand(environment []string, command string, name string) {
	components := splitCommand(command)

	if len(components) == 0 {
		goreland.LogFatal("Command is not provided for task [%s]", name)
	}

	logCommand(environment, command, name)

	execute.Execute(components[0], components[1:], environment)
}

func splitCommand(command string) []string {

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
func cleanComponent(component string) string {
	component = strings.ReplaceAll(component, "\"", "")

	return component
}
