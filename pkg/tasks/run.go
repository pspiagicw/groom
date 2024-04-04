package tasks

import (
	"github.com/buildkite/shellwords"
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
	components, err := shellwords.Split(command)
	if err != nil {
		goreland.LogFatal("Error parsing command [%s] for task [%s]", command, name)
	}

	if len(components) == 0 {
		goreland.LogFatal("Command is not provided for task [%s]", name)
	}

	logCommand(environment, command, name)

	execute.Execute(components[0], components[1:], environment)
}
