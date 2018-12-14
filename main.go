package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "status")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("go build failed with %s\n", err)
	}
	fmt.Printf("%s\n", string(out))
}
