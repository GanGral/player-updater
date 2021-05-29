package updater

import (
	"encoding/json"
	"fmt"
	"net/http"
	"player-updater/common"

	"github.com/gorilla/mux"
)

const playersPath = "players.csv"
const currentVersionPath = "currentVersion.json"

var macAddresses []string

func Init() {

	macAddresses = common.ReadAddresses(playersPath)
	fmt.Println(macAddresses)
	//fmt.Println(macAddresses[0])
}

//HandleUpdate: PUT request handler to process music player update
func HandleUpdate(w http.ResponseWriter, r *http.Request) {

	/*Required request headers to verify:
	x-client-id: required
	x-authentication-token: required

	Requied response headers:
	content-type: application/json
	*/

	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)                         //reading input parameters
	authorized := requestAuthentication(r.Header) //First authorization check

	if authorized {

		processRequest(params, r, w)

	} else {
		//no token/clientId provided
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "invalid clientId or token supplied")
	}

}

func processRequest(params map[string]string, r *http.Request, w http.ResponseWriter) {

	//We should check for validity of macAddress. For now just verifying it exists in common csv file.
	_, macFound := find(macAddresses, params["macaddress"]) //check if macaddress is in the slice.

	if macFound {

		err := verifyBody(r)
		//TODO:
		//shouldn't get version back if verification fails
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, `child \"profile\" fails because [child \"applications\" fails because [\"applications\" is required]]`)
		} else {
			json.NewEncoder(w).Encode(common.GetLatestVersion(currentVersionPath)) //return latest player version back in the response
		}

	} else {
		//no such mac address
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "profile of client %s does not exist", params["macaddress"])
	}
}

//helper funcitons
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
