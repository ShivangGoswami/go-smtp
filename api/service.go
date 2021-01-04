package api

import (
	"go-smtp/mailer"
	"log"
)

// Service manages interaction with the datastore and other services
type Service struct {
	mailClient mailer.Datastore
}

// NewAPIService constructs an APIService
func NewAPIService(conf map[string]string) (*Service, error) {
	log.Println("Started injection of mailer client into service object")
	client, err := mailer.CreateDatastore(conf)
	if err != nil {
		return nil, err
	}
	service := &Service{client}
	log.Println("Completed injection of mailer client into service object")
	return service, nil
}

// ClearAll is a utility function that allows for the clearing of all
// persisted data in the store
func (svc *Service) ClearAll() error {
	return nil
}
