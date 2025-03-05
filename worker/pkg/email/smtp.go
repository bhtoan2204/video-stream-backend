package email

import (
	"log"
	"net/smtp"
	"os"

	"github.com/aymerick/raymond"
	"github.com/bhtoan2204/worker/internal/payload"
)

func SendEmail(email payload.EmailPayload) error {
	// TODO: replace this by config viper
	smtpHost := "smtp.example.com"
	smtpPort := "587"
	senderEmail := "your-email@example.com"
	password := "your-password"

	templateContent, err := os.ReadFile("templates/" + email.Template)
	if err != nil {
		log.Printf("error reading email template: %v", err)
		return err
	}
	renderedBody, err := raymond.Render(string(templateContent), email.Data)
	if err != nil {
		log.Printf("error rendering email template: %v", err)
	}

	msg := []byte("MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"From: " + senderEmail + "\r\n" +
		"To: " + email.To + "\r\n\r\n" +
		renderedBody)

	auth := smtp.PlainAuth("", senderEmail, password, smtpHost)
	addr := smtpHost + ":" + smtpPort

	err = smtp.SendMail(addr, auth, senderEmail, []string{email.To}, msg)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}

	log.Printf("Email sent successfully to %s", email.To)
	return nil
}
