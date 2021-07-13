package handlers

import (
	"io"
	"io/ioutil"
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
	trValidateCommand = " validate"
	trApplyCommand    = "apply"
	trInitCommand     = "init"
)

func creteTempDir(path string, prefix string) string {
	dir, err := ioutil.TempDir(path, prefix)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func createTempFile(path string, prefix string) *os.File {
	file, err := ioutil.TempFile(path, prefix)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func insertDataInFile(file *os.File, data string) {
	if _, err := file.WriteString(data); err != nil {
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
	err := copyFile(trClientPath, copyPath)
	if err != nil {
		log.Println(err)
	}
}

func RunEcsScript(configScript, ecsScript string) {
	tempDir := creteTempDir(tempPath, ecsDirPrefix)
	ecsFileScript := createTempFile(tempDir, ecsFilePrefix)
	insertDataInFile(ecsFileScript, ecsScript)
	configFileScript := createTempFile(tempDir, configFilePrefix)
	insertDataInFile(configFileScript, configScript)
	//copyClient(tempDir + "\\" + trClientTitle)
	ExecCommand(trClientPath, "-chdir="+tempDir, trInitCommand)
	ExecCommand(trClientPath, "-chdir="+tempDir, trApplyCommand, "-auto-approve ")
}
