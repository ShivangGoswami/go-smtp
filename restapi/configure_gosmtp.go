// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	syserr "errors"

	pkgapi "go-smtp/api"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"go-smtp/models"
	"go-smtp/restapi/ops"
	"go-smtp/restapi/ops/smtp"
)

//go:generate swagger generate server --target ..\..\go-smtp --name Gosmtp --spec ..\api.yml --api-package ops
var statictoken string

func init() {
	if temp := os.Getenv("staticapikey"); temp == "" {
		log.Fatal("Static API KEY not present in environmental variables!")
	} else {
		statictoken = temp
	}
}

func configureFlags(api *ops.GosmtpAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *ops.GosmtpAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-api" header is set
	api.KeyAuth = func(token string) (interface{}, error) {
		//return nil, errors.NotImplemented("api key auth (key) x-api from header param [x-api] has not yet been implemented")
		if token == statictoken {
			return token, nil
		}
		return nil, syserr.New("Static API Key authentication failed")
	}

	apiService, err := setup()
	if err != nil {
		log.Fatal("unable to configure MicroService:", err)
	}
	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	//if api.SMTPPostSendmailHandler == nil {
	api.SMTPPostSendmailHandler = smtp.PostSendmailHandlerFunc(func(params smtp.PostSendmailParams, principal interface{}) middleware.Responder {
		log.Println("Send Mail Request Received", params.HTTPRequest)
		log.Println("Input Parameters", params.InputParam)
		if len(params.InputParam.To) == 0 && len(params.InputParam.Cc) == 0 && len(params.InputParam.Bcc) == 0 {
			return smtp.NewPostSendmailDefault(422).WithPayload(&models.Error{
				Message: &[]string{"Destination address not available in request"}[0],
				Code:    422,
			})
		}
		err := apiService.SendMail(params.HTTPRequest.Context(), params)
		if err != nil {
			errorObj := &models.Error{
				Message: &[]string{err.Error()}[0],
				Code:    500,
			}
			log.Println("Send Mail Request Error")
			return smtp.NewPostSendmailDefault(500).WithPayload(errorObj)
		}
		log.Println("Send Mail Request Completed")
		return smtp.NewPostSendmailOK().WithPayload(&models.Status{"Success"})
		//return middleware.NotImplemented("operation smtp.PostSendmail has not yet been implemented")
	})
	//}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

/*
=====================================================================================
  ---------------- Code changes started for the custom service call ------------
======================================================================================
*/

//setup is getting called to set up the api service
func setup() (*pkgapi.Service, error) {
	// Construct service configuration instructions
	log.Println("Started Creating Service Object")
	conf, err := buildServiceConfig()
	// If there's an error, fail out
	if err != nil {
		log.Println("Error in creating service object")
		return nil, err
	}
	log.Println("Completed Creating Service Object")
	return pkgapi.NewAPIService(conf)
}

func buildServiceConfig() (map[string]string, error) {
	config := make(map[string]string)
	log.Println("Started initialization of map from environment variables")
	//can be used from env but hardcoded for now
	config["DATASTORE"] = "mailjet"
	err := syserr.New("Configuration missing from environment variables")
	if publicKey := os.Getenv("MJ_APIKEY_PUBLIC"); publicKey != "" {
		config["mailjetpublickey"] = publicKey
	} else {
		log.Println("mailjet public key not present in environment variables")
		return nil, err
	}
	if secretKey := os.Getenv("MJ_APIKEY_PRIVATE"); secretKey != "" {
		config["mailjetsecretkey"] = secretKey
	} else {
		log.Println("mailjet secret key not present in environment variables")
		return nil, err
	}
	log.Println("Completed initialization of map from environment variables")
	return config, nil
}
