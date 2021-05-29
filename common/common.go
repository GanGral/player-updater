package common

import (
	"encoding/csv"
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

	return CurrentProfile
}

//reads the Csv file. File should exist in the root of the tool
func readCsv(path string) *os.File {
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	//defer csvFile.Close()
	return csvFile
}

func ReadAddresses(path string) []string {
	macAddresses := make([]string, 0, 10)
	macFile := readCsv(path)
	macFileReader, err := csv.NewReader(macFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i < len(macFileReader); i++ {
		macAddresses = append(macAddresses, macFileReader[i][0])
	}

	defer macFile.Close()
	return macAddresses

}
