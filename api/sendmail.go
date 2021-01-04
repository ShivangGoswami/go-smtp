package api

import (
	"context"
	"go-smtp/restapi/ops/smtp"
	"log"
)

func (svc *Service) SendMail(ctx context.Context, params smtp.PostSendmailParams) error {
	log.Println("Send Mail Request Started in API layer")
	err := svc.mailClient.SendMail(ctx, params)
	if err != nil {
		log.Println("Send Mail Request Error in API layer")
		return err
	}
	log.Println("Send Mail Request Completed in API layer")
	return nil
}
