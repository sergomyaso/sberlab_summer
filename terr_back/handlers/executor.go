package handlers

import (
	"bytes"
	"log"
	"os/exec"
)

func ExecCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Start()
	if err != nil {
		log.Println(err)
		return "", err
	}
	err = cmd.Wait()
	log.Println(buf.String())
	return buf.String(), err
}
