package controller

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/execute"
	"github.com/pspiagicw/groom/pkg/parse"
	"github.com/pspiagicw/groom/pkg/utils"
	"github.com/samber/lo"
)

func PerformTasks(requests []string) {

	utils.AssertFile()

	if len(requests) == 0 {
		ListTasks()
	}

	executeTasks(requests)

}

func getEnvironmentString(env []string) string {

	var out bytes.Buffer
	lo.ForEach(env, func(value string, index int) {
		out.WriteString(" ")
		out.WriteString(value)
	})
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

func runTask(environment []string, task string, name string) {
	components := splitCommandString(task)

	if len(components) == 0 {
		goreland.LogFatal("Command is not provided for task [%s]", name)
	}

	logTask(environment, task, name)

	execute.Execute(components[0], components[1:], environment)

}

func logTask(environment []string, task string, name string) {
	environmentString := getEnvironmentString(environment)

	fmt.Printf(utils.LOG_PREFIX+"%s =>"+environmentString+" %s\n", name, task)
}

func checkTask(request string, tasks map[string]*parse.Task) *parse.Task {
	task, ok := tasks[request]

	if !ok {
		goreland.LogFatal("No task named %s", request)
	}

	if task.Command == "" && len(task.Commands) == 0 {
		goreland.LogFatal("No command/commands specified for [%s]!", request)
	}
	return task

}
func executeTasks(requested []string) {

	tasks := parse.ParseTasks()

	lo.ForEach(requested, func(request string, index int) {

		task := checkTask(request, tasks)

		goreland.LogInfo("Executing dependencies for [%s]", request)

		executeTasks(task.Depends)

		if len(task.Commands) != 0 {
			lo.ForEach(task.Commands, func(item string, index int) {
				runTask(task.Environment, item, request)
			})
		} else {
			runTask(task.Environment, task.Command, request)
		}
	})
}
