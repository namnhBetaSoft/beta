package mail

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
	"strconv"
)

var (
	mailSmtpUsername = os.Getenv("MAIL_SMTP_USERNAME")
	mailSmtpPassword = os.Getenv("MAIL_SMTP_PASSWORD")
	mailSmtpHost     = os.Getenv("MAIL_SMTP_HOST")
	mailSmtpPort     = os.Getenv("MAIL_SMTP_PORT")
	mailEmailForm    = os.Getenv("MAIL_EMAIL_FORM")
)

func SendMailForgotPassword(name, email string, code string) error {
	auth := smtp.PlainAuth("", mailSmtpUsername,
		mailSmtpPassword, mailSmtpHost)

	templateData := struct {
		Name string
		Code string
	}{
		name,
		code,
	}

	body, err := parseTemplate("helper/mail/template_mail.html", templateData)
	if err != nil {
		return err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: Location Tracker forgot password!\n"
	msg := []byte(subject + mime + "\n" + body)

	smtpPort, err := strconv.Atoi(mailSmtpPort)
	if err != nil {
		return err
	}
	addr := mailSmtpHost + ":" + strconv.Itoa(smtpPort)

	if err = smtp.SendMail(addr, auth, mailEmailForm, []string{email}, msg); err != nil {
		return err
	}

	return nil
}

func parseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
