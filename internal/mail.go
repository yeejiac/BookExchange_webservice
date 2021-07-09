package internal

import (
	"fmt"
	"net/smtp"

	"github.com/yeejiac/BookExchange_webservice/models"
	"gopkg.in/ini.v1"
)

func SendRegisterMail(userinfo models.User) {
	cfg, err := ini.Load("./config/setting.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}
	from := cfg.Section("gmail").Key("username").String()
	password := cfg.Section("gmail").Key("password").String()
	// Receiver email address.
	to := []string{
		userinfo.Email,
	}

	// smtp server configuration.
	smtpHost := cfg.Section("gmail").Key("host").String()
	smtpPort := cfg.Section("gmail").Key("port").String()

	// Message.
	message := []byte("Account verification\n" + "http://192.168.56.105:8080/validation?Name=" + userinfo.Name)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	senderr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if senderr != nil {
		fmt.Println(senderr)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
