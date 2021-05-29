package updater

import (
	"encoding/json"
	"log"
	"net/http"
	"player-updater/common"
)

func verifyBody(r *http.Request) error {
	//Try decode the body into accepted Player object

	var playerProfile common.Player

	err := json.NewDecoder(r.Body).Decode(&playerProfile)
	if err != nil {

		log.Print(err)

	}

	//json.NewEncoder(w).Encode(playerProfile) //debugging purposes
	return err
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

}

func verifyToken(header http.Header) bool {
	_, okToken := header["X-Authentication-Token"] //using canonical here

	return okToken

}
