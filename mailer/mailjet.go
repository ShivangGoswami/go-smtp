package mailer

import (
	"log"

	"github.com/mailjet/mailjet-apiv3-go"
)

type mailjetStore struct {
	client *mailjet.Client
}

func NewMailjetDatastore(conf map[string]string) (Datastore, error) {
	log.Println("Creating mailjet Client")
	mj := mailjet.NewMailjetClient(conf["mailjetpublickey"], conf["mailjetsecretkey"])
	log.Println("Created mailjet Client")
	return mailjetStore{mj}, nil
}
