package updater

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"player-updater/common"

	"github.com/gorilla/mux"
)

const filename = "players.csv"

var macAddresses []string

func Init() {
	readAddresses()
	//fmt.Println(macAddresses)
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

//reads the Csv file. File should exist in the root of the tool
func readCsv() *os.File {
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	//defer csvFile.Close()
	return csvFile
}

func readAddresses() {
	macAddresses = make([]string, 0, 10)
	macFile := readCsv()
	macFileReader, err := csv.NewReader(macFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i < len(macFileReader); i++ {
		macAddresses = append(macAddresses, macFileReader[i][0])
	}

	/* 	for i, mac := range macFileReader {
		fmt.Println(i)

	} */
	defer macFile.Close()

}

func processRequest(params map[string]string, r *http.Request, w http.ResponseWriter) {

	//We should check for validity of macAddress. For now just verifying it exists in common csv file.
	_, macFound := find(macAddresses, params["macaddress"]) //check if macaddress is in the slice.

	if macFound {

		verifyBody(r, w)
		//TODO:
		//shouldn't get version back if verification fails
		//json.NewEncoder(w).Encode(GetLatestVersion()) //return latest player version back in the response

	} else {
		//no such mac address
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "profile of client %s does not exist", params["macaddress"])
	}
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func verifyBody(r *http.Request, w http.ResponseWriter) {
	//Try decode the body into accepted Player object

	var playerProfile common.Player

	err := json.NewDecoder(r.Body).Decode(&playerProfile)
	if err != nil {

		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

	}
	//attempting to verify body for errors through reflect TODO:
	/* 	body, err := ioutil.ReadAll(r.Body)
	   	if err != nil {
	   		log.Print(err)
	   		//w.WriteHeader(http.StatusBadRequest)

	   	} */

	/* 	err = playerProfile.UnmarshalJSON(body)

	   	if err != nil {
	   		log.Print(err)
	   		w.WriteHeader(http.StatusBadRequest)

	   	} */

	//json.NewEncoder(w).Encode(playerProfile) //debugging purposes
}
func requestAuthentication(header http.Header) bool {
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
func verifyProfile(player *common.Player) {

}

func verifyToken(header http.Header) bool {
	_, okToken := header["X-Authentication-Token"] //using canonical here

	return okToken

}
