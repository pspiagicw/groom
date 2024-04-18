package tasks

import (
	"bytes"
	"fmt"

	"github.com/pspiagicw/groom/pkg/config"
	"github.com/pspiagicw/groom/pkg/utils"
)

func logTask(task *config.Task) {
	environmentString := getEnvPrefix(task.Environment)

	for _, command := range task.Commands {
		fmt.Printf(utils.LOG_PREFIX+"%s =>"+environmentString+" %s\n", task.Name, command)
	}
}

func getEnvPrefix(env []string) string {

	var out bytes.Buffer
	for _, value := range env {
		out.WriteString(" ")
		out.WriteString(value)
	}
	return out.String()
}
