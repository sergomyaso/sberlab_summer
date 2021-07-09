package ScriptRunner

import (
	"bytes"
	"os/exec"
)

func ExecCommand(command string) (string, error) {
	cmd := exec.Command(command)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	return buf.String(), nil
}

