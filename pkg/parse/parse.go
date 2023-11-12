package parse

import (
	"bytes"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pspiagicw/groom/pkg/constants"
)

type Task struct {
	Description string   `toml:"description"`
	Command     string   `toml:"command"`
	Commands    []string `toml:"commands"`
	Shell       string   `toml:"shell"`
	Environment []string `toml:"environment"`
	Depends     []string `toml:"depends"`
}

type Config struct {
	Name      string            `toml:"name"`
	Variables map[string]string `toml:"variables"`
	Tasks     map[string]*Task  `toml:"task"`
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
func ParseTasks() map[string]*Task {

	config := ParseConf()

	resolveVariables(config)

	resolveTasks(config)

	return config.Tasks
}

func resolveTasks(c *Config) {

	for name, task := range c.Tasks {
		newValue := resolveString(task.Command, c)
		c.Tasks[name].Command = newValue

		for i, subtask := range task.Commands {
			newValue := resolveString(subtask, c)
			c.Tasks[name].Commands[i] = newValue

		}
	}

}

func resolveVariables(c *Config) {
	for name, value := range c.Variables {
		newValue := resolveString(value, c)
		// fmt.Printf("%s has value '%s'\n", name, newValue)
		c.Variables[name] = newValue
	}
}
func resolveString(content string, config *Config) string {

	var out bytes.Buffer

	for i := 0; i < len(content); i++ {
		if content[i] == '$' {
			startIndex := i + 1
			currentIndex := i + 1
			if content[currentIndex] == '{' {
				for currentIndex < len(content) {
					if content[currentIndex] != '}' {
						currentIndex += 1
					} else {
						break
					}
				}

				token := content[startIndex+1 : currentIndex]

				if token == "name" {
					out.WriteString(config.Name)
					break
				}

				value, exists := config.Variables[token]

				if !exists {
					log.Printf("Token '%s' does not exist within variables.", token)
				}
				out.WriteString(value)
				i = currentIndex + 1
			} else {
				for currentIndex < len(content) {
					if isLetter(content[currentIndex]) {
						currentIndex += 1
					} else {
						break
					}
				}
				token := content[startIndex:currentIndex]

				// fmt.Printf("Extracted token: %s\n", token)

				if token == "name" {
					out.WriteString(config.Name)
					i = currentIndex - 1
				} else {
					value, exists := config.Variables[token]

					if !exists {
						log.Printf("Token '%s' does not exist within variables.", token)
					}
					out.WriteString(value)
					i = currentIndex - 1
					// fmt.Printf("%d is the position, total length: %d\n", i, len(content))
				}

			}
		} else {
			out.WriteByte(content[i])
		}
	}

	return out.String()
}

func isLetter(element byte) bool {
	return ('a' <= element && element <= 'z') || ('A' <= element && element <= 'Z')
}
