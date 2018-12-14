package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Settings struct {
	User         User         `json:"user"`
	Repositories Repositories `json:"repositories"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Repositories struct {
	List []string `json:"repositories"`
}

var CurrentSettings Settings
var settingsPath = "/tmp/gogit/settings.json"

func (s Settings) UpdateEmail(v string) {
	s.User.Email = v
	s.WriteSettingsFile()
}

func (s Settings) UpdateName(v string) {
	s.User.Name = v
	s.WriteSettingsFile()
}

func (s Settings) AddRepositories(v string) {
	var hasPath = false
	for _, t := range s.Repositories.List {
		if t == v {
			hasPath = true
		}
	}
	if hasPath != true {
		s.Repositories.List = append(s.Repositories.List, v)
		s.WriteSettingsFile()
	} else {
		fmt.Println("This repo already added")
	}
}

func (s Settings) RemoveRepositories(v string) {
	var pathIndex = -1
	for i, t := range s.Repositories.List {
		if t == v {
			pathIndex = i
		}
	}
	if pathIndex != -1 {
		s.Repositories.List = append(s.Repositories.List[:pathIndex], s.Repositories.List[pathIndex+1:]...)
		s.WriteSettingsFile()
	} else {
		fmt.Println("This repo already deleted")
	}
}

func (s Settings) WriteSettingsFile() {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(s)

	err := json.Unmarshal(reqBodyBytes.Bytes(), &s)
	if err != nil {
		fmt.Println(err)
	}

	emptyJSON, _ := json.Marshal(s)
	err = ioutil.WriteFile(settingsPath, emptyJSON, 0777)
}

func main() {
	// Read the settings file
	settingsFile, err := os.Open(settingsPath)

	if err != nil {
		fmt.Println("Settings file could not found.")

		// Create the gogit directory
		gogitPath := filepath.Join("/", "tmp", "gogit")
		os.MkdirAll(gogitPath, os.ModePerm)
		settingsFile, err = os.Create(settingsPath)
		fmt.Println("Settings folder created.")

		if err != nil {
			fmt.Println("Something went wrong, settings folder could not created.", err)
		}

		settingsBodyBytes := new(bytes.Buffer)
		json.NewEncoder(settingsBodyBytes).Encode(CurrentSettings)

		err = json.Unmarshal(settingsBodyBytes.Bytes(), &CurrentSettings)

		if err != nil {
			fmt.Println("Something went wrong, settings file could not created.", err)
		}

		settingsJSON, _ := json.Marshal(CurrentSettings)
		err = ioutil.WriteFile(settingsPath, settingsJSON, 0777)
	} else {
		// Read current settings
		settingsByteValue, _ := ioutil.ReadAll(settingsFile)

		var result map[string]interface{}
		json.Unmarshal([]byte(settingsByteValue), &result)

		json.Unmarshal(settingsByteValue, &CurrentSettings)

		// Set flags
		emailFlag := flag.String("e", "", "GitHub User Email")
		nameFlag := flag.String("n", "", "GitHub User Name")
		repostoriesAddFlag := flag.String("a", "", "Add the new directory")
		repostoriesRemoveFlag := flag.String("r", "", "Remove the directory")
		listFlag := flag.String("l", "s", "List the current settings")

		flag.Parse()
		if *emailFlag != "" {
			CurrentSettings.UpdateEmail(*emailFlag)
		}

		if *nameFlag != "" {
			CurrentSettings.UpdateName(*nameFlag)
		}
		if *repostoriesAddFlag != "" {
			CurrentSettings.AddRepositories(*repostoriesAddFlag)
		}
		if *repostoriesRemoveFlag != "" {
			CurrentSettings.RemoveRepositories(*repostoriesRemoveFlag)
		}
		if *listFlag == "s" {
			fmt.Println(CurrentSettings)
		}
	}
	defer settingsFile.Close()
}
