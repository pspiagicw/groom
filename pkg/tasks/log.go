package tasks

import (
	"bytes"
	"fmt"

	"github.com/pspiagicw/groom/pkg/config"
	"github.com/pspiagicw/groom/pkg/utils"
)

func logTask(task *config.Task) {
	environmentString := getEnvPrefix(task.Environment)

	fmt.Printf(utils.LOG_PREFIX+"%s =>"+environmentString+" %s\n", task.Name, task.Command)
}

func getEnvPrefix(env []string) string {

	var out bytes.Buffer
	for _, value := range env {
		out.WriteString(" ")
		out.WriteString(value)
	}
	return out.String()
}
