package help

import (
	"github.com/pspiagicw/pelp"
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
	pelp.Print("A simple task runner")
	pelp.HeaderWithDescription("usage", []string{"groom [flags] [tasks/commands]"})
}

func printFlags() {
	flags := []string{"example-config", "simple", "dry-run"}
	descriptions := []string{"Dump example config", "List tasks without table", "Dry run tasks"}
	pelp.Flags("flags", flags, descriptions)
}
func printCommands() {
	pelp.Aligned("commands", []string{"version:", "help:"}, []string{"Show version info", "Show this message"})

}
func printFooter() {
	pelp.HeaderWithDescription(
		"more help",
		[]string{
			"Use 'groom help tasks' for more info about the task format",
			"For more information visit https://github.com/pspiagicw/groom",
		})
}

func PrintHelp(version string) {
	PrintVersion(version)
	printUsage()
	printFlags()
	printCommands()
	printFooter()
}

func PrintVersion(version string) {
	pelp.Version("groom", version)
}

func taskHelp() {
	pelp.HeaderWithDescription(
		"Task Format",
		[]string{
			"The task are difined in the `groom.toml` file",
			"Here is an example file.",
		})
	pelp.HeaderWithDescription(
		"example",
		[]string{
			"[task.example]",
			`description="Do something simple"`,
			`command="echo Go rocks!"`,
		})
	pelp.HeaderWithDescription("more help",
		[]string{
			"groom supports many more awesome features. Run `groom --example-config` to explore",
		})

}
func PrintExampleConfig() {
	pelp.Print(exampleConfig)
}
