package infrastructure

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

const (
	SmtpHost      = "smtp.gmail.com"        // Correct SMTP server for Gmail
	SmtpPort      = 465                     // Port for SMTPS (SSL/TLS)
	EmailFrom     = "adanemoges6@gmail.com" // Your Gmail address
	EmailPassword = "nqcbzwothtmiedus"      // Your app-specific password
	ServerHost    = "http://localhost:8080" // Change to your domain in production
	TokenTTlL     = time.Hour               // Token Time-To-Live
)

// Generates a secure random token
func GenerateResetToken() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Sends the password reset email
func SendResetEmail(to, token string) error {
	resetLink := fmt.Sprintf("%s/users/password-update?token=%s", ServerHost, token)
	body := fmt.Sprintf(`
        Hi,

        You requested a password reset. Click the link below to reset your password:

        %s

        If you did not request this, please ignore this email.
    `, resetLink)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", "Loan Services", EmailFrom))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Password Reset")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(SmtpHost, SmtpPort, EmailFrom, EmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
func SendVerificationEmail(to, token string) error {
	verifyLink := fmt.Sprintf("%s/users/verify-email?token=%s", ServerHost, token)
	body := fmt.Sprintf(`
		Hi,

		Welcome to Lone Tracker! Click the link below to verify your email address:

		%s

		If you did not sign up for an account, please ignore this email.
	`, verifyLink)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", "Lone Services", EmailFrom))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Email Verification")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(SmtpHost, SmtpPort, EmailFrom, EmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
