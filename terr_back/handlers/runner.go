package handlers

import (
	"log"
	"os"
)

const (
	tempPath         = "./"
	configFilePrefix = "conf.*.tf"
	ecsFilePrefix    = "ecs.*.tf"
	ecsDirPrefix     = "ecs"
)

const (
	trClientTitle     = "terraform"
	trClientPath      = "./" + trClientTitle
	trValidateCommand = "validate"
	trApplyCommand    = "apply"
	trInitCommand     = "init"
)

func copyClient(copyPath string) {
	err := CopyFile(trClientPath, copyPath)
	if err != nil {
		log.Println(err)
	}
}

func RunUserScript(script string) (string, error) {
	tempDir := CreteTempDir(tempPath, ecsDirPrefix)
	ecsFileScript := CreateTempFile(tempDir, ecsFilePrefix)
	InsertDataInFile(ecsFileScript, script)
	result, err := ExecCommand(trClientPath, "-chdir="+tempDir, trInitCommand)
	if err != nil {
		os.RemoveAll(tempPath + tempDir)
		log.Printf("Exec error:%v\n", err)
		return "", err
	}
	result, err = ExecCommand(trClientPath, trValidateCommand)
	if err != nil {
		os.RemoveAll(tempPath + tempDir)
		log.Printf("Exec error:%v\n", err)
		return "", err
	}
	result, err = ExecCommand(trClientPath, "-chdir="+tempDir, trApplyCommand, "-auto-approve")
	if err != nil {
		os.RemoveAll(tempPath + tempDir)
		log.Printf("Exec error:%v\n", err)
		return "", err
	}
	log.Println(result)
	return result, nil
}
