package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//go get -u "github.com/gorilla/mux"
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/customer/{custid}/address/{addrid}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		custid := vars["custid"]
		addrid := vars["addrid"]

		fmt.Fprintf(w, "You've requested the cust id: %s address %s\n", custid, addrid)
	})

	http.ListenAndServe(":8080", r)
}
