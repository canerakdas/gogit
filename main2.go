package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "status", "-s")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("go build failed with %s\n", err)
	}
	lines := strings.Split(string(out), "\n")
	var modifiedArray []string
	var newFileArray []string
	for i := 0; i < len(lines)-1; i++ {
		if strings.Fields(lines[i])[0] == "M" {
			modifiedArray = append(modifiedArray, lines[i])
		} else if strings.Fields(lines[i])[0] == "??" {
			newFileArray = append(newFileArray, lines[i])
		}
	}
	fmt.Println("Modified Files : ")
	fmt.Println(modifiedArray)

	fmt.Println("New Files : ")
	fmt.Println(newFileArray)
}
