package utils

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/goreland"
)

var LOG_PREFIX = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd93f9")).Render(" [make] ")

var TASK_FILE = "groom.toml"

func ConfigFilePath() string {
	if fileExists(TASK_FILE) {
		return TASK_FILE
	}
	return findConfig(getCurrentDir())
}
func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return true
}
func getCurrentDir() string {
	curDir, err := os.Getwd()
	if err != nil {
		goreland.LogFatal("Error getting current directory")
	}
	return curDir
}
func findConfig(dir string) string {
	parent := filepath.Dir(dir)

	if parent == xdg.Home {
		noConfigFound()
	}

	configFile := filepath.Join(parent, TASK_FILE)

	if fileExists(configFile) {
		changeDir(parent)
		return configFile
	}
	return findConfig(parent)
}
func changeDir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		goreland.LogFatal("Error changing working directory %v", err)
	}
}
func noConfigFound() {
	goreland.LogError("Couldn't find `groom.toml` in current or parent directories.")
	goreland.LogError("Search stopped at home directory.")
	goreland.LogInfo("If you need more information regarding the groom.toml file, run `groom help`")
	goreland.LogFatal("Make sure the `groom.toml` file is created.")
}
