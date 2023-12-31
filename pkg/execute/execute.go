package execute

import (
	"os"
	"os/exec"

	"github.com/pspiagicw/goreland"
)

func Execute(command string, args []string, env []string) {

	cmd := exec.Command(command, args...)
	// fmt.Println(arg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = append(cmd.Environ(), env...)

	err := cmd.Run()

	if err != nil {
		goreland.LogFatal("Proccess exited with a error: %v", err)
	}
}
