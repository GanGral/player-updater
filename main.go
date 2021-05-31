package main

import (
	"fmt"
	"log"
	"net/http"
	"player-updater/updater"

	//"player-updater/updater"

	"github.com/gorilla/mux"
)

//TODO: Create a list of macaddresses
//TODO: https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj we should g with server to set the response headers

//Simple GET request handler

func init() {

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to main page!")
}

func handleRequests() {

	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/profiles/clientId:{macaddress}", updater.HandleUpdate).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8457", router))

}
func main() {
	//updater.Init()
	handleRequests()
}

/* func handleProductReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var productFilter ProductReportFilter
		err := json.NewDecoder(r.Body).Decode(&productFilter)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
} */
