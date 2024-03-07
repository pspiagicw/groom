package tasks

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/parse"
)

func PerformTasks(requests []string) {

	taskFile := parse.ParseTasks()

	executeTasks(requests, taskFile)
}

func checkTask(task *parse.Task, request string) {
	if task.Command == "" && len(task.Commands) == 0 {
		goreland.LogFatal("No command/commands specified for [%s]!", request)
	}

	if len(task.Commands) == 0 {
		task.Commands = []string{
			task.Command,
		}
	}
}

func getTask(request string, tasks map[string]*parse.Task) *parse.Task {

	task, ok := tasks[request]
	if !ok {
		goreland.LogFatal("No task named %s", request)
	}

	checkTask(task, request)
	return task

}

func runDependencies(request string, task *parse.Task, taskFile map[string]*parse.Task) {
	goreland.LogInfo("Executing dependencies for [%s]", request)
	executeTasks(task.Depends, taskFile)
}

func executeTasks(tasks []string, taskFile map[string]*parse.Task) {

	for _, name := range tasks {
		task := getTask(name, taskFile)
		runDependencies(name, task, taskFile)
		runCommands(task, name)
	}
}

func runCommands(task *parse.Task, name string) {
	for _, command := range task.Commands {
		runCommand(task.Environment, command, name)
	}

}
