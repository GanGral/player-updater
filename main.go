package main

//main.go

//entry point for updater_service.Creates a HTTP listener on port 8487 to process PUT requests to "/profiles/clientId:{macaddress}".
//Updater service processes HTTP requests to upgrade Players profile using their macaddresses.

import (
	"log"
	"net/http"
	"player-updater/updater"

	"github.com/gorilla/mux"
)

func handleRequests() {

	router := mux.NewRouter()
	router.HandleFunc("/profiles/clientId:{macaddress}", updater.HandleUpdate).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8457", router))

}
func main() {

	handleRequests()
}
