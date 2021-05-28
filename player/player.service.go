package player

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetLatestVersion: returns latest player version
func GetLatestVersion() Player {
	CurrentProfile := Player{
		Profile: Profile{

			Applications: []Application{
				{ApplicationID: "555", Version: "53333=4"},
				{ApplicationID: "-rumba", Version: "534"},
			},
		},
	}
	return CurrentProfile
}

//PUT request handler to handle player update request
func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	//Setting response header
	w.Header().Set("content-type", "application/json")

	//reading input parameters
	params := mux.Vars(r)

	/*Required headers to verify:
	x-client-id: required
	x-authentication-token: required
	*/

	var playerProfile Player

	err := json.NewDecoder(r.Body).Decode(&playerProfile)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	canProceed := requestVerification(r.Header)

	if canProceed {

		if params["macaddress"] == "388" {

			json.NewEncoder(w).Encode(GetLatestVersion())

		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "profile of client %s does not exist", params["macaddress"])
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "invalid clientId or token supplied")
	}

}

func requestVerification(header http.Header) bool {
	okId := verifyClientID(header)
	okToken := verifyToken(header)

	if okId && okToken {
		return true
	}

	return false
}
func verifyClientID(header http.Header) bool {
	_, okId := header["X-Client-Id"] //using canonical here

	return okId

}

func verifyToken(header http.Header) bool {
	_, okToken := header["X-Authentication-Token"] //using canonical here

	return okToken

}
