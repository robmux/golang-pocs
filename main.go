package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Sender and recipient email addresses
	from := "sender@example.com"
	to := "robinsonmu232@gmail.com"

	// SMTP server configuration
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.ParseInt(os.Getenv("SMTP_PORT"), 10, 64)
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	// Email subject
	subject := "Email with images and attachments"

	// Set up HTML part with embedded image
	htmlTemplate := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Email Template</title>
	</head>
	<body>
		<h1>Hello,</h1>
		<p>This is a sample email with an image:</p>
		<img src="cid:logo.png" alt="Logo">
		<p>Thank you!</p>
	</body>
	</html>
	`
	tmpl, err := template.New("email").Parse(htmlTemplate)
	if err != nil {
		log.Fatal(err)
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, nil)
	if err != nil {
		log.Fatal(err)
	}

	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	message.SetBody("text/html", tpl.String())
	message.Attach("document.pdf")
	message.Embed("logo.png")

	// Connect to the SMTP server with TLS encryption
	d := gomail.NewDialer(smtpHost, int(smtpPort), smtpUsername, smtpPassword)
	if err := d.DialAndSend(message); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")

	time.Sleep(50 * time.Second)
}
