package help

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const exampleConfig string = `
name = "example-project"

[variables]
version = "0.0.1"

# Tasks start with '[task.<task-name>]'
# They should contain, 'command' property.
# Other fields are optional.
[task.build]
description = "Build the project."
command = "go build ."

# Tasks can contain 'commands' as a list of commands.
[task.run]
commands = [
    "go run main.go",
    "python -m exaple-project",
]

# Tasks can contain dependencies, and environment variables defined
[task.test]
environment = [ "TESTS=1" ]
command = "python -m unittest"
depends = [
    "format"
]

[task.format]
command = "go fmt ./..."
`

func HandleHelp(args []string, version string) {
	if len(args) == 0 {
		PrintHelp(version)
	} else {
		arg := args[0]
		switch arg {
		case "tasks":
			taskHelp()
		}

	}
}
func printUsage() {
	fmt.Println("A simple task runner")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  groom [flags] [tasks/commands]")
	fmt.Println()
}

func printFlags() {
	fmt.Println("FLAGS")
	flags := `
--example-config:
--simple:
--dry-run:`
	descriptions := `
Dump example config.
List tasks without table
Dry run tasks`

	printAligned(flags, descriptions)
}
func printCommands() {
	fmt.Println("COMMANDS")
	commands := `
version:
help:`
	messages := `
Show version info
Show this message`
	printAligned(commands, messages)

}
func printFooter() {
	fmt.Println("MORE HELP")
	fmt.Println("  Use 'groom help tasks' for more info about the task format")
	fmt.Println("  For more information visit https://github.com/pspiagicw/groom")
}

func PrintHelp(version string) {
	PrintVersion(version)
	printUsage()
	printFlags()
	printCommands()
	printFooter()
}

func printAligned(left string, right string) {
	leftCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(left).MarginLeft(2).String()
	rightCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(right).MarginLeft(5).String()

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Bottom, leftCol, rightCol))
	fmt.Println()

}

func PrintVersion(version string) {
	fmt.Printf("groom version: %s\n", version)
}

func taskHelp() {
	fmt.Println("Task Format")
	fmt.Println("The task are difined in the `groom.toml` file")
	fmt.Println("Here is an example file.")
	fmt.Println()
	fmt.Println("[task.example]")
	fmt.Println("description=\"Do something simple\"")
	fmt.Println("command=\"echo Go rocks!\"")
	fmt.Println()
	fmt.Println("groom supports many more awesome features. Run `groom --example-config` to explore")
}
func PrintExampleConfig() {
	fmt.Println()
	fmt.Println(exampleConfig)
}
