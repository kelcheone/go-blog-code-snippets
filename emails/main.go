package main

import (
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	from := "pinky@mail.com"
	to := []string{
		"frazer@mail.com",
	}

	message := []byte(`To: frazer@mail.com

			From: pinky@mail.com

			Subject: A Guide To SMTP

			Content-Type: text/plain; charset="us-ascii"

			This is a test email message.
		`)

	godotenv.Load()
	username := os.Getenv("EMAIL_HOST_USER")
	passwd := os.Getenv("EMAIL_HOST_PASSWORD")
	host := os.Getenv("EMAIL_HOST")

	auth := smtp.PlainAuth("", username, passwd, host)
	smtpHost := host + ":" + os.Getenv("EMAIL_PORT")

	if err := smtp.SendMail(smtpHost, auth, from, to, message); err != nil {
		panic(err)
	}

}
