package execute

import (
	"os"
	"os/exec"
)

func Execute(command string, args []string, env []string) error {

	cmd := exec.Command(command, args...)
	// fmt.Println(arg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = append(cmd.Environ(), env...)

	err := cmd.Run()

	return err
}
