package tasks

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/config"
)

func PerformTasks(tasks []string, dryRun bool) {

	groomConfig := config.ParseConfig()

	executeTasks(tasks, groomConfig.Tasks, dryRun)
}

func checkTask(task *config.Task) {
	if task.Command == "" && len(task.Commands) == 0 {
		goreland.LogFatal("No command/commands specified for [%s]!", task.Name)
	}

	if len(task.Commands) == 0 {
		task.Commands = []string{
			task.Command,
		}
	}
}

func getTask(name string, tasks map[string]*config.Task) *config.Task {

	task, ok := tasks[name]

	if !ok {
		goreland.LogFatal("No task named %s", name)
	}

	checkTask(task)
	return task

}

func executeTasks(tasks []string, taskList map[string]*config.Task, dryRun bool) {
	for _, name := range tasks {
		task := getTask(name, taskList)
		if !dryRun {
			runDependencies(task, taskList)
			runCommands(task)
		} else {
			logTask(task)
		}
	}
}
