package utils

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/goreland"
)

var LOG_PREFIX = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd93f9")).Render(" [make] ")

var TASK_FILE = "groom.toml"

func AssertFile() {
	_, err := os.Stat(TASK_FILE)
	if err != nil {
		goreland.LogError("Error while reading groom.toml: %v", err)
		goreland.LogInfo("If you need more information regarding the groom.toml file, run `groom help`")
		goreland.LogFatal("Make sure the current directory has the `groom.toml` file.")
	}
}
