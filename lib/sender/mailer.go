package mailer

import (
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

	user := lib.GetEnv().MailUser
	password := lib.GetEnv().MailPassword

	// Set up the SMTP dialer
	_ = gomail.NewDialer("live.smtp.mailtrap.io", 587,
		user, password)

	// Send the email
	//if err := dialer.DialAndSend(message); err != nil {
	//	log.Get(ctx).Error("Error sending email: ", err)
	//	panic(err)
	//} else {
	//	log.Get(ctx).Info("Email sent successfully")
	//}

	return nil
}
