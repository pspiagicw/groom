package execute

import (
	"os"
	"os/exec"
)

func Execute(command string, args []string) (error) {

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

    return err
}
