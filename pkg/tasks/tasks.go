package tasks

import (
	"os"
	"slices"

	"github.com/buildkite/shellwords"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/argparse"
	"github.com/pspiagicw/groom/pkg/config"
	"github.com/pspiagicw/groom/pkg/execute"
)

func PerformTasks(opts *argparse.Opts) {

	groomConfig := config.ParseConfig()

	taskList := getTaskList(opts, groomConfig)

	taskList = sort(taskList, groomConfig)

	executeTasks(taskList, opts)
}
func sort(tasks []*config.Task, groomConfig *config.Config) []*config.Task {

	seen := make(map[string]bool)

	stack := tasks

	order := []*config.Task{}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if seen[current.Name] {
			continue
		}
		seen[current.Name] = true

		for _, dep := range current.Depends {
			dependency := getTask(dep, groomConfig)

			if dependency == nil {
				goreland.LogFatal("Task [%s] depends on [%s], which does not exist", current.Name, dep)
			}

			if !seen[dependency.Name] {
				stack = append(stack, dependency)
			}
		}
		order = append(order, current)
	}

	slices.Reverse(order)

	return order
}

func getTaskList(opts *argparse.Opts, groomConfig *config.Config) []*config.Task {
	requestedTasks := opts.Args

	currentTasks := []*config.Task{}

	for _, name := range requestedTasks {

		task := getTask(name, groomConfig)

		if task == nil {
			goreland.LogFatal("No task named %s", name)
		}

		currentTasks = append(currentTasks, task)

	}

	return currentTasks
}

func getTask(name string, groomConfig *config.Config) *config.Task {

	task, ok := groomConfig.Tasks[name]

	if !ok {
		return nil
	}

	sanitizeTask(task)

	return task
}

func sanitizeTask(task *config.Task) {

	if task.Command == "" && len(task.Commands) == 0 {
		goreland.LogFatal("No command/commands specified for [%s]!", task.Name)
	}

	if len(task.Commands) == 0 {
		task.Commands = []string{
			task.Command,
		}
	}

	task.Command = ""
}

func executeTasks(taskList []*config.Task, opts *argparse.Opts) {
	for _, task := range taskList {
		logTask(task)
		if !opts.DryRun {
			runTask(task)
		}
	}
}
func pushDirectory(task *config.Task) {
	if task.Directory != "" {
		curdir, err := os.Getwd()
		if err != nil {
			goreland.LogFatal("Error changing directory to %s: %s", task.Directory, err)
		}
		err = os.Chdir(task.Directory)
		if err != nil {
			goreland.LogFatal("Error changing directory to %s: %s", task.Directory, err)
		}
		task.Directory = curdir
	}
}
func popDirectory(task *config.Task) {
	if task.Directory != "" {
		err := os.Chdir(task.Directory)
		if err != nil {
			goreland.LogFatal("Error changing directory to %s: %s", task.Directory, err)
		}
	}
}
func runTask(task *config.Task) {

	pushDirectory(task)

	for _, command := range task.Commands {
		run(task, command)
	}

	popDirectory(task)
}
func run(task *config.Task, command string) {
	components, err := shellwords.Split(command)

	if err != nil {
		goreland.LogFatal("Error parsing command [%s] for task [%s]", task.Command, task.Name)
	}

	if len(components) == 0 {
		goreland.LogFatal("No command specified for task [%s]", task.Name)
	}

	execute.Execute(components, task.Environment)
}
