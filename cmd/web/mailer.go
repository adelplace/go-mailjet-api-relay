package main

import (
	"github.com/mailjet/mailjet-apiv3-go"
)

func (app *application) sendMail(contact *contact) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: app.email,
				Name:  app.email,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: app.email,
					Name:  app.email,
				},
			},
			Subject: "Message from Mailjet API Relay",
			HTMLPart: `
			<ul>
				<li><b>Email:</b> ` + contact.email + `</li>
				<li><b>Name:</b> ` + contact.name + `</li>
				<li><b>Sujet:</b> ` + contact.subject + `</li>
				<li><b>Message:</b> ` + contact.message + `</li>
			</ul>`,
			CustomID: "DefaultAppId",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := app.mailjetClient.SendMailV31(&messages)

	if err != nil {
		app.errorLog.Printf("Error: %v\n", err)
	} else {
		app.infoLog.Printf("Data: %+v\n", res)
	}
	return err
}
