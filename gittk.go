package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

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
		projectDir := clone(repo)
		if projectDir == "" {
			fmt.Println("Unable to load project directory, try setting GITTK_PATH.")
		}
		fmt.Println(projectDir)
	default:
		fmt.Println("The given command is unknown.")
		printUsage()
		os.Exit(1)
	}
}

func clone(repo string) string {
	projectDir := os.Getenv("GITTK_PATH")

	// Get the project directory
	if projectDir == "" {
		user, err := user.Current()
		if err != nil {
			return ""
		}
		projectDir = filepath.Join(user.HomeDir, "projects")
	}

	// Parse git URI to get subtree
	fmt.Printf("TODO: clone %v into %v\n", repo, projectDir)
	return projectDir
}
