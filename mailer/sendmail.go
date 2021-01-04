package mailer

import (
	"context"
	"go-smtp/models"
	"go-smtp/restapi/ops/smtp"
	"log"

	"github.com/mailjet/mailjet-apiv3-go"
)

func parseEmails(persons []*models.Person) *mailjet.RecipientsV31 {
	var emails mailjet.RecipientsV31
	if len(persons) != 0 {
		for _, val := range persons {
			emails = append(emails, mailjet.RecipientV31{
				Email: val.Email.String(),
				Name:  *val.Name,
			})
		}
		return &emails
	}
	return nil
}
func (mj mailjetStore) SendMail(ctx context.Context, params smtp.PostSendmailParams) error {
	log.Println("Send Mail Request Started in Mailer layer")
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				//Email: "shivang.goswami.here@gmail.com",
				//Name:  "Shivang",
				Email: params.InputParam.From.Email.String(),
				Name:  *params.InputParam.From.Name,
			},
			To:  parseEmails(params.InputParam.To),
			Cc:  parseEmails(params.InputParam.Cc),
			Bcc: parseEmails(params.InputParam.Bcc),
			// To: &mailjet.RecipientsV31{
			// 	mailjet.RecipientV31{
			// 		Email: "shivang.goswami.here@gmail.com",
			// 		Name:  "Shivang",
			// 	},
			// },
			Subject: *params.InputParam.Subject,
			//Subject:  "Greetings from Mailjet.",
			TextPart: params.InputParam.Text,
			//TextPart: "My first Mailjet email",
			HTMLPart: params.InputParam.HTML,
			//HTMLPart: "<h3>Dear passenger 1, welcome to <a href='https://www.mailjet.com/'>Mailjet</a>!</h3><br />May the delivery force be with you!",
			CustomID: *params.InputParam.CustomID,
			// CustomID: "AppGettingStartedTest",
		},
	}
	res, err := mj.client.SendMailV31(&mailjet.MessagesV31{Info: messagesInfo})
	if err != nil {
		log.Println("Send Mail Error:", err.Error())
		return err
	}
	log.Printf("SendMailResponse: %+v\n", res)
	log.Println("Send Mail Request Completed in Mailer layer")
	return nil
}
