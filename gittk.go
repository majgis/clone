package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"./repository"
)

const usage = `
Commands:

    clone - Clone git repositories into consistent tree structure
			$ gittk clone <git URI>      

Options:
    -bash
`

var lg = log.New(os.Stderr, "", 0)

func main() {
	bash := flag.Bool("bash", false, "Output bash commands instead of executing.")
	flag.Parse()

	// Print usage if no command was given
	command := flag.Arg(0)
	if command == "" {
		fmt.Print(usage)
		os.Exit(0)
	}

	// Execute the given command
	switch command {
	case "clone":

		// Exit if URI argument was not given to clone command
		repoURI := flag.Arg(1)
		if repoURI == "" {
			lg.Fatalf("You must supply git URI to the clone command.\n%v", usage)
		}

		// Clone
		repoDir, err := repository.Clone(repoURI, *bash)

		// Exit if error during clone
		if err != nil {
			lg.Fatalf("Unable to clone repository:  %v", err)
		}

		// Print the cloned directory
		fmt.Printf("\n%v\n", repoDir)

	default:
		lg.Fatalf("The given command is unknown.\n%v", usage)
	}
}
