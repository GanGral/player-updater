//main.go

//Entry point for updater_service. Creates a HTTP listener on port 8487 to process PUT requests to "/profiles/clientId:{macaddress}".
//Updater service processes HTTP requests to upgrade Players profile using their macaddresses.
package main

import (
	"fmt"
	"log"
	"net/http"
	"player-updater/updater"

	"github.com/gorilla/mux"
)

func handleRequests() {

	router := mux.NewRouter()
	router.HandleFunc("/profiles/clientId:{macaddress}", updater.HandleUpdate).Methods("PUT")
	fmt.Printf("Starting Updater Service at port 8457\n")
	log.Fatal(http.ListenAndServe(":8457", router))

}
func main() {

	handleRequests()
}
