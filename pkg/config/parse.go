package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pspiagicw/demp"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/groom/pkg/utils"
)

type Task struct {
	Description string   `toml:"description"`
	Command     string   `toml:"command"`
	Commands    []string `toml:"commands"`
	Shell       string   `toml:"shell"`
	Environment []string `toml:"environment"`
	Depends     []string `toml:"depends"`
	Name        string
}

type Config struct {
	Name      string            `toml:"name"`
	Variables map[string]string `toml:"variables"`
	Tasks     map[string]*Task  `toml:"task"`
}

func readConf() *Config {

	config, err := os.Open(utils.ConfigFilePath())

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
func ParseConfig() *Config {

	config := readConf()

	resolveVariables(config)

	resolveTasks(config)

	return config
}

func resolveTasks(c *Config) {

	for name, task := range c.Tasks {
		task.Name = name
		newValue := demp.ResolveTemplate(task.Command, c.Variables)
		c.Tasks[name].Command = newValue

		for i, subtask := range task.Commands {
			newValue := demp.ResolveTemplate(subtask, c.Variables)
			c.Tasks[name].Commands[i] = newValue

		}
	}

}

func resolveVariables(c *Config) {

	// Add the name as a variable
	c.Variables["name"] = c.Name
	// Add other default variables
	addDefaultVariables(c)

	for name, value := range c.Variables {
		newValue := demp.ResolveTemplate(value, c.Variables)
		c.Variables[name] = newValue
	}
}

func addDefaultVariables(c *Config) {
	curDir, err := os.Getwd()
	if err != nil {
		goreland.LogFatal("Error getting current directory: %q", err)
	}
	c.Variables["pwd"] = curDir
}
