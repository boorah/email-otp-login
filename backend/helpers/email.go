package helpers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"

	"github.com/resend/resend-go/v2"

	"boorah/email-otp-login-backend/config"
)

var resendClient *resend.Client

func InitEmailClient() {
	if resendClient == nil {
		resendClient = resend.NewClient(config.ConfigData.RESEND_API_KEY)
	}
}

func renderOTPEmailBody(otp string) (string, error) {
	tmpl, err := template.ParseFiles("templates/otp/otp.html")
	if err != nil {
		return "", err
	}

	var renderedTemplate bytes.Buffer
	if err := tmpl.Execute(&renderedTemplate, map[string]string{
		"OTP":       otp,
		"ExpiresIn": strconv.Itoa(config.ConfigData.OTP_VALIDITY_MINUTES),
	}); err != nil {
		return "", err
	}

	return renderedTemplate.String(), nil
}

func SendOTPEmail(to, subject, otp string) error {
	fromAddress := fmt.Sprintf("%s <%s>", config.ConfigData.RESEND_FROM_NAME, config.ConfigData.RESEND_FROM_EMAIL)

	emailBody, err := renderOTPEmailBody(otp)
	if err != nil {
		return fmt.Errorf("error rendering email template: %v", err)
	}

	sendEmailParams := &resend.SendEmailRequest{
		From:    fromAddress,
		To:      []string{to},
		Subject: subject,
		Html:    emailBody,
	}

	sent, err := resendClient.Emails.Send(sendEmailParams)
	if err != nil {
		return err
	}

	log.Println("Email sent successfully:", sent.Id)

	return nil
}
