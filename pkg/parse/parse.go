package parse

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/golang-groom/groom-make/pkg/constants"
)

type Task struct {
	Description string   `toml:"description"`
	Command     string   `toml:"command"`
	Shell       string   `toml:"shell"`
	Depends     []string `toml:"depends"`
}

type Config struct {
    Name string `toml:"name"`
	Section map[string]Task `toml:"task"`
}

func ParseConf() *Config {

	config, err := os.Open(constants.TASK_FILE)

	if err != nil {
		log.Fatalf("Error reading goproject.toml: %q", err)

	}

	defer config.Close()

	decoder := toml.NewDecoder(config)

	var read Config

	_, err = decoder.Decode(&read)

	if err != nil {
		log.Fatalf("Error parsing toml: %v", err)
	}

    return &read

}
func ParseTasks() map[string]Task {

    config := ParseConf()

	return config.Section

}
