package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
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
		projectDir := getProjectDir(repo)
		if projectDir == "" {
			fmt.Println("Unable to load project directory, try setting GITTK_PATH.")
		}
		os.MkdirAll(projectDir)
		fmt.Println(projectDir)
	default:
		fmt.Println("The given command is unknown.")
		printUsage()
		os.Exit(1)
	}
}

func getProjectDir(uri string) string {
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
	subDir, err := parseGitURI(uri)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fullDir := filepath.Join(projectDir, subDir)
	return fullDir
}

// git@github.com:majgis/gittk.git
// https://github.com/majgis/gittk.git
//
func parseGitURI(uri string) (string, error) {
	uriSplit := strings.Split(uri, "/")

	// SSH
	isGithubSSH := strings.HasPrefix(uri, "git@github.com")
	if isGithubSSH {
		userName := strings.Split(uriSplit[0], ":")[1]
		projectName := strings.Split(uriSplit[1], ".")[0]
		result := filepath.Join("github.com", userName, projectName)
		return result, nil
	}

	// HTTPS
	isGithubHTTPS := strings.HasPrefix(uri, "https://github.com")
	if isGithubHTTPS {
		userName := uriSplit[3]
		projectName := strings.Split(uriSplit[4], ".")[0]
		result := filepath.Join("github.com", userName, projectName)
		return result, nil
	}

	// Unknown
	return "", errors.New("unknown URI type")
}
