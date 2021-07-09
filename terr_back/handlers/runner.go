package handlers

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	tempPath      = ""
	configFilePrefix = "conf.*.tf"
	ecsFilePrefix = "ecs.*.tf"
	ecsDirPrefix  = "ecs"
)

const (
	trClientTitle     = "terraform.exe"
	trClientPath      = "/utilities/" + trClientTitle
	trValidateCommand = trClientPath + " validate"
	trApplyCommand    = trClientPath + " apply"
	trInitCommand     = trClientPath + " init"
)

func creteTempDir(path string, prefix string) string {
	dir, err := ioutil.TempDir(path, prefix)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func createTempFile(path string, prefix string) string {
	dir, err := ioutil.TempDir(path, prefix)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func insertDataInFile(path string, data string) {
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		log.Println(err)
	}
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func copyClient(copyPath string) {
	copyFile(trClientPath, copyPath)
}

func RunEcsScript(configScript, ecsScript string) {
	tempDir := creteTempDir(tempPath, ecsDirPrefix)
	ecsScriptPath := createTempFile(tempDir, ecsFilePrefix)
	insertDataInFile(ecsScriptPath, ecsScript)
	configScriptPath := createTempFile(tempDir, configFilePrefix)
	insertDataInFile(ecsScriptPath, configScriptPath)
	copyClient(tempPath + "/" + trClientTitle)
	ExecCommand(trInitCommand)
	ExecCommand(trApplyCommand)
}
