package main

import "os"
import "fmt"

func printUsage() {
	fmt.Println(`
Commands:

clone - Clone git repositories into consistent tree structure
        $ gittk clone <git URI>      
`)
}

func main() {

	// Print usage if no command was given
	if len(os.Args) == 1 {
		printUsage()
		os.Exit(0)
	}

	command := os.Args[1]

	// Execute the given command
	switch command {
	case "clone":
		if len(os.Args) < 3 {
			fmt.Println("You must supply git URI to the clone command.")
			printUsage()
			os.Exit(1)
		}
		repo := os.Args[2]
		clone(repo)
	default:
		fmt.Println("The given command is unknown.")
		printUsage()
		os.Exit(1)
	}
}

func clone(repo string) {
	fmt.Printf("TODO: clone %v\n", repo)
}
