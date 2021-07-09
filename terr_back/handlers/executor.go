package handlers

import (
	"bytes"
	"log"
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
	log.Println("Your command:" + command)
	log.Println(buf.String())
	return buf.String(), nil
}

