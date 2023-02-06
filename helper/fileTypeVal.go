package helper

import (
	"io"
	"mime/multipart"
	"net/http"
)

func TypeFile(test multipart.File) bool {
	fileByte, _ := io.ReadAll(test)
	fileType := http.DetectContentType(fileByte)
	if fileType == "image/png" || fileType == "image/jpeg" {
		return true
	}
	return false
}
