package helper

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"time"
	"timesync-be/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var theSession *session.Session

// ======================================================================
// INIT SESSION KEY CREDENTIAL S3 BUCKET AWS
// ======================================================================
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

// ======================================================================
// UPLOAD IMAGE PROGRESS
// ======================================================================
func GetUrlImagesFromAWS(fileData multipart.FileHeader) (string, error) {
	if fileData.Filename == "" {
		return "", errors.New("file name error")
	}
	if fileData.Size == 0 {
		return "", errors.New("file size error")
	}
	if fileData.Size > 500000 {
		return "", errors.New("size error")
	}
	file, err := fileData.Open()
	if err != nil {
		return "", errors.New("error open fileData")
	}
	// Validasi Type
	tipeNameFile, err := TypeFile(file)
	if err != nil {
		return "", errors.New("file type error only jpg or png file can be upload")
	}
	defer file.Close()

	log.Println("size:", fileData.Filename, file)
	namaFile := GenerateRandomString()
	namaFile = namaFile + tipeNameFile
	fileData.Filename = namaFile
	log.Println(namaFile)
	file2, _ := fileData.Open()
	defer file2.Close()
	uploadURL, err := UploadToS3(fileData.Filename, file2)
	if err != nil {
		return "", errors.New("cannot upload to s3 server error")
	}
	return uploadURL, nil
}

// ======================================================================
// UPLOAD TO S3
// ======================================================================
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

func GenerateRandomString() string {
	rand.Seed(time.Now().Unix())
	str := "AsDfzGhBvCX123456MnBp"
	shuff := []rune(str)
	// Shuffling the string
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	// Displaying the random string
	// fmt.Println(string(shuff))
	return string(shuff)
}

// Handler Controller
// func UploadController(c echo.Context) error {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return err
// 	}
// 	// karena saat upload file aws tidak generate nama file secara manual, sehingga harus generate nama filenya secara manual
// 	// gunakan package github.com/satori/go.uuid lalu panggil fungsinya uuid.NewV4().String()
// 	fileName := uuid.NewV4().String()
// 	file.Filename = fileName + file.Filename[(len(file.Filename)-5):len(file.Filename)]
// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()
// 	uploadURL, err := UploadToS3(file.Filename, src)
// 	if err != nil {
// 		return err
// 	}
// 	responseJson := &UploadResult{
// 		Path: uploadURL,
// 	}
// 	return c.JSON(http.StatusOK, responseJson)
// }
