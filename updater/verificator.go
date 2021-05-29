package updater

import (
	"encoding/json"
	"errors"
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

	err = verifyProfile(&playerProfile)

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

	clientId, okId := header["X-Client-Id"] //using canonical here
	if clientId[0] == "" {
		okId = false
	}

	//fmt.Println(clientId, okId)
	return okId

}
func verifyProfile(player *common.Player) error {

	if player.Profile.Applications == nil {
		err := errors.New("invalid body")
		return err
	}
	return nil

}

func verifyToken(header http.Header) bool {
	token, okToken := header["X-Authentication-Token"] //using canonical here
	if token[0] == "" {
		okToken = false
	}

	return okToken

}
