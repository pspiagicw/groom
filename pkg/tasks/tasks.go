package tasks

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/config"
)

func PerformTasks(tasks []string, dryRun bool) {

	groomConfig := config.ParseConfig()

	checkTasks(tasks, groomConfig)

	executeTasks(tasks, groomConfig.Tasks, dryRun)
}

func checkTasks(tasks []string, groomConfig *config.Config) {
	for _, name := range tasks {
		task, ok := groomConfig.Tasks[name]
		if !ok {
			goreland.LogFatal("No task named %s", name)
		}

		for _, dep := range task.Depends {
			if _, ok := groomConfig.Tasks[dep]; !ok {
				goreland.LogFatal("Task %s depends on %s, which does not exist", name, dep)
			}
		}

		if task.Command == "" && len(task.Commands) == 0 {
			goreland.LogFatal("No command/commands specified for [%s]!", task.Name)
		}

	}
}

func sanitizeTask(task *config.Task) {

	if len(task.Commands) == 0 {
		task.Commands = []string{
			task.Command,
		}
	}
}

func getTask(name string, tasks map[string]*config.Task) *config.Task {

	task := tasks[name]

	sanitizeTask(task)

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
