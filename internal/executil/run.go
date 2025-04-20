package executil

import (
	"os"
	"os/exec"
)

// Run executes a command, streaming stdout/stderr
func Run(name string, args ...string) error {
	c := exec.Command(name, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}