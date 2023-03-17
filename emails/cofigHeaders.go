// Configuring email headers and body
package main

import (
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
	headers.Set("Content-Type", "text/plain; charset=\"us-ascii\"")
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

	if err := smtp.SendMail(smtpHost, auth, from, to, []byte(message)); err != nil {
		panic(err)
	}
}
