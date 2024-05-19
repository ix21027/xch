package service

import (
	"fmt"
	"net/smtp"
	"os"
	"xch/db"
)

func NotifySubscribers() {
	emails, err := db.GetSubscribers()
	if err != nil {
		fmt.Println(err)
	}
	rate, err := GetRate()
	if err != nil {
		fmt.Println(err)
	}

	for _, email := range emails {
		err := sendEmail(email, rate)
		if err != nil {
			fmt.Println("sending email error:", err)
		}
	}
}

func sendEmail(to, rate string) error {
	var (
		server   = os.Getenv("SMTP_SERVER")
		port     = os.Getenv("PORT_EMAIL")
		email    = os.Getenv("EMAIL_ADDR")
		password = os.Getenv("EMAIL_PASS")
	)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: XCH Rate (BTCUAH)" + "\r\n" +
		"\r\n" +
		rate)

	auth := smtp.PlainAuth("", email, password, server)

	err := smtp.SendMail(server+":"+port, auth, email, []string{to}, msg)
	if err != nil {
		return err
	}

	return nil
}
