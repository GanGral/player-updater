//Updater Tool
// interactive command line tool to perform updates of multiple player's profile.
// sends PUT requests to updater service to initiate updates
// Usage:
//
//  -csvpath string
//        Specify the path to the player's Mac Addresses csv file (default "../players.csv")
//  -port int
//        specifies the updater_service port (default 8457)
//  -profile string
//        Specify the location of json file with up-to-date player profile (default "../tool/currentVersion.json")
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"player-updater/common"
	"strconv"
)

func main() {

	// Setting up flags for user
	macAddressPath := flag.String("csvpath", "../players.csv", "Specify the path to the player's Mac Addresses csv file")
	latestProfileJson := flag.String("profile", "../tool/currentVersion.json", "Specify the location of json file with up-to-date player profile")
	updaterPort := flag.Int("port", 8457, "specifies the updater_service port")
	flag.Parse()

	// Assemble mac address list, latest player profile version
	addressList := common.ReadAddresses(*macAddressPath)
	latestProfile := common.GetLatestVersion(*latestProfileJson)
	authParams := getAuthParams("../tool/authparams.json")

	for _, mac := range addressList {
		RunUpdate(latestProfile, mac, authParams, *updaterPort)
	}

}

// RunUpdate updates player profiles.
func RunUpdate(currentVersion common.Player, mac string, authParams AuthParams, updaterPort int) {

	payloadBytes, err := json.Marshal(currentVersion)
	if err != nil {
		fmt.Print(err)
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("PUT", "http://localhost:"+strconv.Itoa(updaterPort)+"/profiles/clientId:"+mac, body)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Client-Id", authParams.ClientId)
	if !authParams.ClientToken.Expired {
		req.Header.Set("X-Authentication-Token", authParams.ClientToken.Token)
	}

	// run assembled request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(respBody)
	log.Printf(sb, resp.Status)
	defer resp.Body.Close()
}
