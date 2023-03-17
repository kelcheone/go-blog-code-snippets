// send attachments

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	from := "pinky@mail.com"
	to := []string{
		"frazer@mail.com",
	}

	headers := make(textproto.MIMEHeader)
	headers.Set("From", from)
	headers.Set("To", strings.Join(to, ","))
	headers.Set("Subject", "A Guide To SMTP")
	headers.Set("Content-Type", "multipart/mixed; boundary=\"----=_NextPart_01C7B6B1.6E7B6AF0\"")
	headers.Set("MIME-Version", "1.0")
	headers.Set("Content-Transfer-Encoding", "7bit")

	body := `This is a test email message.`

	godotenv.Load()
	username := os.Getenv("EMAIL_HOST_USER")
	passwd := os.Getenv("EMAIL_HOST_PASSWORD")
	host := os.Getenv("EMAIL_HOST")

	auth := smtp.PlainAuth("", username, passwd, host)
	smtpHost := host + ":" + os.Getenv("EMAIL_PORT")

	message := ""

	for k, v := range headers {
		message = message + k + ": " + strings.Join(v, "") + "\r "
	}

	message = message + "\r " + body

	// Create a buffer to write our email to.
	buf := bytes.NewBuffer(nil)

	// Create a multipart writer to create the message.
	writer := multipart.NewWriter(buf)

	// Create a new text part.
	textPart, err := writer.CreatePart(textproto.MIMEHeader{"Content-Type": {"text/plain"}})
	if err != nil {
		log.Fatal(err)
	}

	// Write the plain text to the text part.
	if _, err := textPart.Write([]byte(body)); err != nil {
		log.Fatal(err)
	}

	// Create a new file part.
	filePart, err := writer.CreatePart(textproto.MIMEHeader{"Content-Type": {"image/png"}, "Content-Disposition": {`attachment; filename="image1.png"`}})
	if err != nil {
		log.Fatal(err)
	}

	// Open the file to be attached.
	file, err := os.Open("image1.png")
	if err != nil {
		log.Fatal(err)
	}

	// Copy the file data to the file part.
	if _, err := io.Copy(filePart, file); err != nil {
		log.Fatal(err)
	}

	// Close the file.
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	// Create another new file part.
	filePart, err = writer.CreatePart(textproto.MIMEHeader{"Content-Type": {"image/png"}, "Content-Disposition": {`attachment; filename="image2.png"`}})
	if err != nil {
		log.Fatal(err)
	}

	// Open the file to be attached.
	file, err = os.Open("image2.png")
	if err != nil {
		log.Fatal(err)
	}

	// Copy the file data to the file part.
	if _, err := io.Copy(filePart, file); err != nil {
		log.Fatal(err)
	}

	// Close the file.
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	// Close the writer.
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}

	// Send the email.
	if err := smtp.SendMail(smtpHost, auth, from, to, buf.Bytes()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email Sent!")
}
