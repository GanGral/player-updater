package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Structure to construct for authentication
type AuthParams struct {
	ClientToken Token  `json:"token"`
	ClientId    string `json:"x-client-id"`
}
type Token struct {
	Token   string `json:"x-authentication-token"`
	Expired bool   `json:"expired"`
}

/*Authentication parameters.
Ideally should be retrieved through OAuth and not exposed, but for purpose of assignement this is a plain text JSON.
*/

func getAuthParams(path string) AuthParams {

	var authParams AuthParams

	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened currentVersion.json")

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &authParams)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	return authParams

}
