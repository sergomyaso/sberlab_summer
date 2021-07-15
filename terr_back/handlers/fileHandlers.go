package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
)

func GetDataFromFile(file multipart.File) string {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return ""
	}
	return buf.String()
}
