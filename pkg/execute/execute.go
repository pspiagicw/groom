package execute

import (
	"os"
	"os/exec"

	"github.com/pspiagicw/goreland"
)

func Execute(components []string, env []string) {

	command := components[0]
	args := components[1:]

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = append(cmd.Environ(), env...)

	err := cmd.Run()

	if err != nil {
		goreland.LogFatal("Proccess exited with a error: %v", err)
	}
}
