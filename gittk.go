package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
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
		clone(repoURI)
	default:
		lg.Fatalf("The given command is unknown.\n%v", usage)
	}
}

// Clone the given gitURI into GITTK_PATH
func clone(repoURI string) {
	repoDir := getRepoDir(repoURI)
	if repoDir == "" {
		lg.Fatalln("Unable to load project directory, try setting GITTK_PATH.")
	}
	err := os.MkdirAll(repoDir, os.ModePerm)
	if err != nil {
		lg.Fatalf("Unable to create directory: %v", repoDir)
	}
	os.Chdir(repoDir)
	cmd := exec.Command("git", "clone", repoURI, ".")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	startErr := cmd.Start()
	if startErr != nil {
		lg.Fatal("Unable to execute git clone")
	}
	cmd.Wait()
	fmt.Printf("\n%v\n", repoDir)
}

// Get the repository directory from git URI
func getRepoDir(uri string) string {
	repoDir := os.Getenv("GITTK_PATH")

	// Get the project directory
	if repoDir == "" {
		user, err := user.Current()
		if err != nil {
			return ""
		}
		repoDir = filepath.Join(user.HomeDir, "projects")
	}

	// Parse git URI to get subtree
	subDir, err := parseGitURI(uri)
	if err != nil {
		lg.Fatalln(err)
	}
	fullDir := filepath.Join(repoDir, subDir)
	return fullDir
}

// Return subdirectory from different git URI types
// git@github.com:majgis/gittk.git
// https://github.com/majgis/gittk.git
func parseGitURI(uri string) (string, error) {
	uriSplit := strings.Split(uri, "/")

	// github SSH
	isGithubSSH := strings.HasPrefix(uri, "git@github.com")
	if isGithubSSH {
		userName := strings.Split(uriSplit[0], ":")[1]
		projectName := strings.Split(uriSplit[1], ".")[0]
		result := filepath.Join("github.com", userName, projectName)
		return result, nil
	}

	// github HTTPS
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
