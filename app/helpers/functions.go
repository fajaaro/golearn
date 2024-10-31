package helpers

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
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

func UploadFile(c *gin.Context, file *multipart.FileHeader, permissionType, folder string) (*string, error) {
	if permissionType == "private" {
		permissionType = ""
	} else {
		permissionType = "public/"
	}
	directory := fmt.Sprintf("storage/app/%s%s", permissionType, folder)
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

func ReadExcel(file *multipart.FileHeader) ([][]string, error) {
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}

	xlsx, err := excelize.OpenReader(fileContent)
	if err != nil {
		return nil, err
	}

	rows := xlsx.GetRows("Sheet1")

	fileContent.Close()

	return rows, nil
}

func ExtractModelExcelColIndexes(structType interface{}) map[string]int {
	result := make(map[string]int)
	val := reflect.TypeOf(structType)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		excelTag := field.Tag.Get("excel-col-index")
		if excelTag != "" {
			if index, err := strconv.Atoi(excelTag); err == nil {
				result[field.Name] = index
			}
		}
	}
	return result
}