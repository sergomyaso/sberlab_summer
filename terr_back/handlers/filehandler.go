package handlers

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
)

func GetDataFromFile(file multipart.File) string {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return ""
	}
	return buf.String()
}

func CreteTempDir(path string, prefix string) string {
	dir, err := ioutil.TempDir(path, prefix)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func CreateTempFile(path string, prefix string) *os.File {
	file, err := ioutil.TempFile(path, prefix)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func InsertDataInFile(file *os.File, data string) {
	if _, err := file.WriteString(data); err != nil {
		log.Println(err)
	}
}

func CopyFile(src, dst string) error {
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
