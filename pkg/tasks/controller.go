package tasks

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

	taskFile := parse.ParseTasks()

	executeTasks(requests, taskFile)
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

func getTask(request string, tasks map[string]*parse.Task) *parse.Task {

	task, ok := tasks[request]

	if !ok {
		goreland.LogFatal("No task named %s", request)
	}

	if task.Command == "" && len(task.Commands) == 0 {
		goreland.LogFatal("No command/commands specified for [%s]!", request)
	}

	if len(task.Commands) == 0 {
		task.Commands = []string{
			task.Command,
		}
	}

	return task

}

func runDependencies(request string, task *parse.Task, taskFile map[string]*parse.Task) {
	goreland.LogInfo("Executing dependencies for [%s]", request)

	executeTasks(task.Depends, taskFile)
}

func executeTasks(requested []string, taskFile map[string]*parse.Task) {

	lo.ForEach(requested, func(request string, _ int) {

		task := getTask(request, taskFile)

		runDependencies(request, task, taskFile)

		lo.ForEach(task.Commands, func(item string, _ int) {
			runTask(task.Environment, item, request)
		})
	})
}
