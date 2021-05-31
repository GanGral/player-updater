package updater

//updater.go
// contains handlers to process incoming requests.
import (
	"encoding/json"
	"fmt"
	"net/http"
	"player-updater/common"

	"github.com/gorilla/mux"
)

// hard-coded, we expect the files to be local to the updater
const destAddresses = "players.csv"
const currentProfilePath = "currentVersion.json"
const headerContentKey = "content-type"
const headerContentValue = "application/json"

var macAddresses []string

// HandleUpdate request handler to process music player update.
// 200 Success is returned together with new Player version profile.
// 404 If no such macaddress exists
// 409 If the update profile is not in expected format or empty

func HandleUpdate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(headerContentKey, headerContentValue)

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

	if len(macAddresses) == 0 {
		macAddresses = common.ReadAddresses(destAddresses)
	}

	//We should check for validity of macAddress. For now just verifying it exists in common csv file.
	_, macFound := find(macAddresses, params["macaddress"]) //check if macaddress is in the slice.

	if macFound {

		//need to verify body before processing
		err := verifyBody(r)

		if err != nil {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, `child \"profile\" fails because [child \"applications\" fails because [\"applications\" is required]]`)
		} else {

			//Returns the updated version of the player upon success. Currently using mock version json.
			json.NewEncoder(w).Encode(common.GetLatestVersion(currentProfilePath)) //return latest player version back in the response
		}

	} else {
		//no such mac address
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "profile of client %s does not exist", params["macaddress"])
	}
}

//helper funciton
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
