package helper

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func PrintHelp(version string) {
	// goreland.LogError("Help printing not implemented yet!")
	PrintVersion(version)
	fmt.Println("A simple task runner")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  groom [flags] [tasks]")
	fmt.Println()
	fmt.Println("COMMANDS")
	commands := `
version:
help:`
	messages := `
Show version info
Show this message`

	commandCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(commands).MarginLeft(2).String()
	messageCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(messages).MarginLeft(5).String()

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Bottom, commandCol, messageCol))
	fmt.Println()
	// taskHelp()
	fmt.Println("MORE HELP")
	fmt.Println("  Use 'groom help tasks' for more info about the task format")
	fmt.Println("  For more information visit https://github.com/pspiagicw/groom")
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
