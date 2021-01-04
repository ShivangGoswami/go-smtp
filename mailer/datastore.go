package mailer

import (
	"context"
	"fmt"
	"go-smtp/restapi/ops/smtp"
	"log"
	"strings"
)

//Datastore is a generic interface
type Datastore interface {
	SendMail(ctx context.Context, params smtp.PostSendmailParams) error
}

func init() {
	Register("mailjet", NewMailjetDatastore)
}

//DatastoreFactory is a type for datastore factory methods
type DatastoreFactory func(conf map[string]string) (Datastore, error)

// datastoreFactories maintains a list of supported factories.
var datastoreFactories = make(map[string]DatastoreFactory)

//Register adds a specific factory to the datastore factories list
func Register(name string, factory DatastoreFactory) {
	log.Println("System started registeration of the service:", name)
	if factory == nil {
		log.Fatal("Datastore factory does not exist")
	}
	_, registered := datastoreFactories[name]
	if registered {
		log.Println("Datastore factory already registered. Ignoring.")
		return
	}
	log.Println("System completed registeration")
	datastoreFactories[name] = factory
}

// CreateDatastore creates a new db.Datastore entity based on the configuration
// provided.
func CreateDatastore(conf map[string]string) (Datastore, error) {
	log.Println("CreateDatastore going to create data store")
	var engineName string
	if name, ok := conf["DATASTORE"]; ok {
		engineName = name
	}
	engineFactory, ok := datastoreFactories[engineName]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		availableDatastores := make([]string, len(datastoreFactories))
		for k := range datastoreFactories {
			availableDatastores = append(availableDatastores, k)
		}
		return nil, fmt.Errorf("Invalid Datastore name. Must be one of: %s", strings.Join(availableDatastores, ", "))
	}
	log.Println("CreateDatastore completed")
	return engineFactory(conf)
}
