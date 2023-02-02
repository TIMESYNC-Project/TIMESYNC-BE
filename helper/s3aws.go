package helper

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"timesync-be/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

var theSession *session.Session

// GetConfig Initiatilize config in singleton way
func GetSession() *session.Session {
	if theSession == nil {
		theSession = initSession()
	}
	return theSession
}
func initSession() *session.Session {
	newSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(config.AWS_REGION),
		Credentials: credentials.NewStaticCredentials(config.ACCESS_KEY_ID, config.ACCESS_KEY_SECRET, ""),
	}))
	return newSession
}

type UploadResult struct {
	Path string `json:"path" xml:"path"`
}

// Helper
func UploadToS3(fileName string, src multipart.File) (string, error) {
	// The session the S3 Uploader will use
	sess := GetSession()
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String("fauziawsbucket"),
		Key:         aws.String(fileName),
		Body:        src,
		ContentType: aws.String("image/png"),
	})
	// content type penting agar saat link dibuka file tidak langsung auto download
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	return result.Location, nil
}

// Handler Controller
func UploadController(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	// karena saat upload file aws tidak generate nama file secara manual, sehingga harus generate nama filenya secara manual
	// gunakan package github.com/satori/go.uuid lalu panggil fungsinya uuid.NewV4().String()
	fileName := uuid.NewV4().String()
	file.Filename = fileName + file.Filename[(len(file.Filename)-5):len(file.Filename)]
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	uploadURL, err := UploadToS3(file.Filename, src)
	if err != nil {
		return err
	}
	responseJson := &UploadResult{
		Path: uploadURL,
	}
	return c.JSON(http.StatusOK, responseJson)
}
