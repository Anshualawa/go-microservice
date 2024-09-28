package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Anshualawa/go-microservice/domain"
	"github.com/Anshualawa/go-microservice/service"
	"github.com/gorilla/mux"
)

func Start() {

	SanityCheck()

	// Defined custom mux and register route
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wire
	// ch := CustomerHandler{service.NewCustomerService(domain.NewCutomerRepositoryStub())}
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositiryDb())}

	router.HandleFunc("/customers", ch.customers).Methods(http.MethodGet)
	/*
		mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
		mux.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
		mux.HandleFunc("/createCustomer/", createCustomer).Methods(http.MethodPost)
	*/

	serverAdd := os.Getenv("SERVER_ADDRESS")
	serverpORT := os.Getenv("SERVER_PORT")

	srString := fmt.Sprintf("%s:%s", serverAdd, serverpORT)
	log.Panic(http.ListenAndServe(srString, router))
}

func SanityCheck() {
	// TODO - get all environment variables
	envVar := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, k := range envVar {
		if os.Getenv(k) == "" {
			log.Fatal(fmt.Sprintf("Environment Variable is missing %s", k))
		}
	}
}
