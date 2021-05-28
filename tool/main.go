package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"player-updater/updater"
)

func main() {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer b7d03a6947b217efb6f3ec3bd3504582" -d '{"type":"A","name":"www","data":"162.10.66.0","priority":null,"port":null,"weight":null}' "https://api.digitalocean.com/v2/domains/example.com/records"
	addressList := updater.GetAddresses()
	currentVersion := updater.GetLatestVersion("../currentVersion.json")
	fmt.Print(addressList)

	for _, mac := range addressList {
		runUpdate(currentVersion, mac)

	}

}
func runUpdate(currentVersion updater.Player, mac string) {

	payloadBytes, err := json.Marshal(currentVersion)
	if err != nil {
		fmt.Print(err)
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("PUT", "http://localhost:8457/profiles/clientId/"+mac, body)

	if err != nil {
		fmt.Print(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Client-Id", "b7d03a6947b217efb6f3ec3bd3504582")
	req.Header.Set("X-Authentication-Token", "b7d03a6947b217efb6f3ec3bd3504582")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(respBody)
	log.Printf(sb, resp.Status)
	defer resp.Body.Close()
}
