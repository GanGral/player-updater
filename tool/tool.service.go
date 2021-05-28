package tool

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const filename = "players.csv"

var macAddresses []string

func Init() {
	readAddresses()
	fmt.Println(macAddresses)
	fmt.Println(macAddresses[0])
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

//GetLatestVersion: returns latest player version
func GetLatestVersion() Player {

	var CurrentProfile Player
	// Open our jsonFile
	jsonFile, err := os.Open("currentVersion.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened currentVersion.json")

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &CurrentProfile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	/*
		CurrentProfile := Player{
			Profile: Profile{

				Applications: []Application{
					{ApplicationID: "555", Version: "53333=4"},
					{ApplicationID: "-rumba", Version: "534"},
				},
			},
		} */
	return CurrentProfile
}

//PUT request handler to handle player update request
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

	_, macFound := find(macAddresses, params["macaddress"]) //check if macaddress is in the slice.

	if macFound {

		verifyBody(r, w)

		json.NewEncoder(w).Encode(GetLatestVersion()) //return latest player version back in the response

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

	var playerProfile Player

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
func verifyProfile(player *Player) {

}

func verifyToken(header http.Header) bool {
	_, okToken := header["X-Authentication-Token"] //using canonical here

	return okToken

}
