package main

import (
	"fmt"
	"log"
	"os"

	"./repository"
)

const usage = `
Commands:

clone - Clone git repositories into consistent tree structure
        $ gittk clone <git URI>      
`

var lg = log.New(os.Stderr, "", 0)

func main() {

	// Print usage if no command was given
	if len(os.Args) == 1 {
		fmt.Print(usage)
		os.Exit(0)
	}

	command := os.Args[1]

	// Execute the given command
	switch command {
	case "clone":
		if len(os.Args) < 3 {
			lg.Fatalf("You must supply git URI to the clone command.\n%v", usage)
		}
		repoURI := os.Args[2]
		repoDir, err := repository.Clone(repoURI)
		if err != nil {
			lg.Fatalf("Unable to clone repository:  %v", err)
		}
		fmt.Printf("\n%v\n", repoDir)
	default:
		lg.Fatalf("The given command is unknown.\n%v", usage)
	}
}
