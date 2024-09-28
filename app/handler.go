package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Anshualawa/go-microservice/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode int    `json:"zip" xml:"zip"`
}

// CustomerHandler struct
type CustomerHandler struct {
	service service.CustomerService
}

func (ch CustomerHandler) customers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()
	// customers := []Customer{
	// 	{"pappu", "Alawa", 454545},
	// 	{"pappu", "Alawa", 454545},
	// 	{"pappu", "Alawa", 454545},
	// }
	if r.Header.Get("Content-Type") == "xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

/*

func getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cid := params["customer_id"]
	fmt.Fprintf(w, cid)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "post request received for new customer.")
}
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go ....")
}
*/
