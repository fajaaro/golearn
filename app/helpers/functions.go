package helpers

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}

func UploadFile(c *gin.Context, file *multipart.FileHeader, folder string) (*string, error) {
	directory := fmt.Sprintf("storage/app/public/%s", folder)
	fileExtension := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d-%s%s", time.Now().Unix(), GenerateRandomString(20), fileExtension)
	filePath := fmt.Sprintf("%s/%s", directory, fileName)

	_ = os.MkdirAll(directory, os.ModePerm)
    
	err := c.SaveUploadedFile(file, filePath)
	if err == nil {
		filePathWithFolder := fmt.Sprintf("/%s/%s", folder, fileName)
		return &filePathWithFolder, nil
    }

	return nil, err
}

func SendEmail(to []string, subject, body string) error {
	mailHost := os.Getenv("MAIL_HOST")
	mailPort := os.Getenv("MAIL_PORT")
	mailUsername := os.Getenv("MAIL_USERNAME")
	mailPassword := os.Getenv("MAIL_PASSWORD")
	mailFrom := os.Getenv("MAIL_FROM")

	auth := smtp.PlainAuth("", mailUsername, mailPassword, mailHost)

	msg := "From: " + mailFrom + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body

	smtpAddr := fmt.Sprintf("%s:%s", mailHost, mailPort)
	err := smtp.SendMail(smtpAddr, auth, mailFrom, to, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}