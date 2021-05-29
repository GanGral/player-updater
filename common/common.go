package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//GetLatestVersion: returns latest player version
func GetLatestVersion(path string) Player {

	var CurrentProfile Player
	// Open our jsonFile
	jsonFile, err := os.Open(path)
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
