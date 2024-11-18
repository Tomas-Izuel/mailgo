package mailer

import (
	"fmt"
	gomail "gopkg.in/mail.v2"
	"mailgo/lib"
	"mailgo/modules/template"
)

func SendEmail(mail template.MailNotificationTemplate, recipient string) error {
	// Create a new message
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", "tomasizuel@gmail.com")
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", mail.Subject)

	// Set email body
	message.SetBody("text/html", mail.BodyHTML)
	message.AddAlternative("text/plain", mail.BodyText)

	// Set up the SMTP dialer
	dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587,
		lib.GetEnv().MailUser, lib.GetEnv().MailPassword)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	} else {
		fmt.Println("Email sent successfully!")
	}

	return nil
}
