package Capabilities

import (
	"User_system/Error_Handling"
	"bytes"
	"fmt"
	"html/template"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/gomail.v2"
)

func SendEmailOtp(nameFileEmail string, otpCod string, SetHeaderServer string, SetHeaderClient string, Subject string, emailBody1 string) string {
	Error_Handling.InitLogger(zapcore.InfoLevel)
	data := struct {
		Code string
	}{
		Code: otpCod,
	}

	htmlData, _ := template.ParseFiles(nameFileEmail)

	var htmlContent bytes.Buffer
	if err := htmlData.Execute(&htmlContent, data); err != nil {
		zap.L().Error("err in template execution", zap.String("err", err.Error()))
		return "error execution"
	}

	m := gomail.NewMessage()
	Email := Error_Handling.LoadConfigFile().Email.Email
	EmailPass := Error_Handling.LoadConfigFile().Email.EmailPassword
	EmailSmtpServer := Error_Handling.LoadConfigFile().Email.EmailSmtpServer
	Port := Error_Handling.LoadConfigFile().Email.EmailPort

	m.SetHeader("From", m.FormatAddress(Email, SetHeaderServer))

	m.SetHeader("To", SetHeaderClient)

	m.SetHeader("Subject", Subject)
	emailBody := fmt.Sprintf(emailBody1, otpCod)
	m.SetBody("text/plain", emailBody)

	m.AddAlternative("text/html", htmlContent.String())

	d := gomail.NewDialer(EmailSmtpServer, Port, Email, EmailPass)

	if err := d.DialAndSend(m); err != nil {
		zap.L().Error("Failed to send email", zap.String("err", err.Error()))
		return "error send"
	}
	zap.L().Info("Email sent successfully")
	return "sent"
}
