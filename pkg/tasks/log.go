package tasks

import (
	"bytes"
	"fmt"

	"github.com/pspiagicw/groom/pkg/utils"
	"github.com/samber/lo"
)

func logTask(environment []string, task string, name string) {
	environmentString := getEnvPrefix(environment)

	fmt.Printf(utils.LOG_PREFIX+"%s =>"+environmentString+" %s\n", name, task)
}

func getEnvPrefix(env []string) string {

	var out bytes.Buffer
	lo.ForEach(env, func(value string, index int) {
		out.WriteString(" ")
		out.WriteString(value)
	})
	return out.String()
}
