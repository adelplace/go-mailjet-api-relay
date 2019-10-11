package main

import (
	"github.com/mailjet/mailjet-apiv3-go"
)

const privateKey = ""
const publicKey = ""

func (app *application) sendMail(contact *contact) error {
	mailjetClient := mailjet.NewMailjetClient(publicKey, privateKey)
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "alexandre.delplacemille@gmail.com",
				Name:  "Alexandre",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "alexandre.delplacemille@gmail.com",
					Name:  "Alexandre",
				},
			},
			Subject: "Message from adelplace.fr",
			HTMLPart: `
			<ul>
				<li><b>Email :</b> ` + contact.email + `</li>
				<li><b>Sujet :</b> ` + contact.subject + `</li>
				<li><b>Message :</b> ` + contact.message + `</li>
			</ul>`,
			CustomID: "DefaultAppId",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)

	if err != nil {
		app.errorLog.Printf("Error: %v\n", err)
	} else {
		app.infoLog.Printf("Data: %+v\n", res)
	}
	return err
}
