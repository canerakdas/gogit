package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var repostories []string

func main() {
	repostories = append(repostories, "/Users/segmentify/accounts")
	repostories = append(repostories, "/Users/segmentify/go/src/github.com/canerakdas/gogit")
	for i := 0; i < len(repostories); i++ {
		goGitStatus(repostories[i])
	}
}
func goGitStatus(repo string) {
	cmd := exec.Command("git", "-C", repo, "status", "-s")
	out, err := cmd.Output()
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
	color.Blue("\n Repo : %s \n\n", repo)
	if len(modifiedArray) != 0 {
		fmt.Println("Modified Files")
		for i := 0; i < len(modifiedArray); i++ {
			color.Yellow("  ->  %s", strings.Fields(modifiedArray[i])[1])
		}
	}

	if len(newFileArray) != 0 {
		fmt.Println("New Files")
		for i := 0; i < len(newFileArray); i++ {
			color.Green("  ->  %s", strings.Fields(newFileArray[i])[1])
		}
	}
	color.White("\n--------------------------------- \n\n")
}
