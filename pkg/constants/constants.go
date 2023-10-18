package constants

import "github.com/charmbracelet/lipgloss"

var LOG_PREFIX = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd93f9")).Render(" [make] ")

var TASK_FILE = "groom.toml"
