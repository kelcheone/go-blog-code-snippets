package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	from := "pinky@mail.com"
	to := []string{
		"frazer@mail.com",
	}

	message := &bytes.Buffer{}
	writer := multipart.NewWriter(message)
	writer.WriteField("To", to[0])
	writer.WriteField("From", from)
	writer.WriteField("Subject", "A Guide To SMTP")
	writer.WriteField("Content-Type", "text/plain; charset=\"us-ascii\"")

	body := []byte("This is a test email message.")
	writer.WriteField("Content-Type", "text/plain; charset=\"us-ascii\"")
	writer.CreatePart(body)

	file, err := os.Open("attachment.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("attachment", file.Name())
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	godotenv.Load()
	username := os.Getenv("EMAIL_HOST_USER")
	passwd := os.Getenv("EMAIL_HOST_PASSWORD")
	host := os.Getenv("EMAIL_HOST")

	auth := smtp.PlainAuth("", username, passwd, host)
	smtpHost := host + ":" + os.Getenv("EMAIL_PORT")

	if err := smtp.SendMail(smtpHost, auth, from, to, message.Bytes()); err != nil {
		panic(err)
	}
}
