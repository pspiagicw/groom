package controller

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/golang-groom/groom-make/pkg/constants"
	"github.com/golang-groom/groom-make/pkg/execute"
	"github.com/golang-groom/groom-make/pkg/parse"
	"github.com/pspiagicw/colorlog"
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
		colorlog.LogError("Error while reading groom.toml: %v\n", err)
		colorlog.LogError("Make sure the current directory has the `groom.toml` file.\n")
		os.Exit(1)
	}
}

func getEnvironmentString(env []string) string {

    var out bytes.Buffer
    for _ , value := range env {

        out.WriteString(" ")
        out.WriteString(value)


    }

    return out.String()
}
func executeTasks(requested []string, tasks map[string]parse.Task) {
	for _, request := range requested {
		task, ok := tasks[request]

		if !ok {
			colorlog.LogError(constants.LOG_PREFIX+"No task named %s", request)
			os.Exit(1)
		}
        environmentString := getEnvironmentString(task.Environment)

		fmt.Printf(constants.LOG_PREFIX + "%s =>" + environmentString + " %s\n", request , task.Command)

		components := strings.Split(task.Command, " ")

		if len(components) == 0 {
			colorlog.LogError(constants.LOG_PREFIX+" Command is not provided for task %s", request)
		}

		err := execute.Execute(components[0], components[1:], task.Environment)

		if err != nil {
            colorlog.LogError(constants.LOG_PREFIX + " exited with a error:" + err.Error())
		}
	}
}
func listTasks(tasks map[string]parse.Task) {

	if len(tasks) == 0 {
		colorlog.LogError("No tasks declared.")
		return
	}

	fmt.Println()
	fmt.Println(lipgloss.NewStyle().MarginLeft(1).Background(lipgloss.Color("#7e56f4")).Foreground(lipgloss.Color("#ffffff")).Render(" tasks "))
	fmt.Println()
	taskStyle := lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#50fa7b"))
	descriptionStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffb86c"))

	for name, task := range tasks {
		description := task.Description

		if description != "" {
			description = "No description provided"
		}
		fmt.Println("-" + taskStyle.Render(name) + " : " + descriptionStyle.Render(task.Description))
	}

}
