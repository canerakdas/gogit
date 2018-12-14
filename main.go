package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Settings struct {
	UserName     string   `json:"user.name"`
	UserEmail    string   `json:"user.email"`
	Repositories []string `json:"repositories"`
}

func main() {
	// Read the settings file

	var settingsPath = "/tmp/gogit/settings.json"
	settingsFile, err := os.Open(settingsPath)

	if err != nil {
		newpath := filepath.Join("/", "tmp", "gogit")
		os.MkdirAll(newpath, os.ModePerm)
		fmt.Println("Settings file not found, settings file created")
		settingsFile, err = os.Create(settingsPath)

		if err != nil {
			fmt.Println("Settings file can't added :(", err)
		}

		settingsFile, err = os.Open(settingsPath)

		if err != nil {
			fmt.Print("Whoops")
		}
		byteValue, _ := ioutil.ReadAll(settingsFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		fmt.Println(result["repositories"])
		fmt.Println(result["user.name"])
	}

	byteValue, _ := ioutil.ReadAll(settingsFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["repositories"])
	fmt.Println(result["user.name"])

	settings := Settings{
		UserName:     "caner",
		UserEmail:    "gmail.com",
		Repositories: []string{"hello"},
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(settings)

	err = json.Unmarshal(reqBodyBytes.Bytes(), &settings)
	if err != nil {
		fmt.Println(err)
		// nozzle.printError("opening config file", err.Error())
	}

	rankingsJson, _ := json.Marshal(settings)
	err = ioutil.WriteFile(settingsPath, rankingsJson, 0777)
	fmt.Printf("%+v", settings)
	// defer the closing of our jsonFile so that we can parse it later on
	defer settingsFile.Close()
}
