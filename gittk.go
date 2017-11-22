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

var stderr = log.New(os.Stderr, "", 0)

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
			stderr.Fatalf("You must supply git URI to the clone command.\n%v", usage)
		}
		repoURI := os.Args[2]
		projectDir := getProjectDir(repoURI)
		if projectDir == "" {
			stderr.Fatalln("Unable to load project directory, try setting GITTK_PATH.")
		}
		err := os.MkdirAll(projectDir, os.ModePerm)
		if err != nil {
			stderr.Fatalf("Unable to create directory: %v", projectDir)
		}
		os.Chdir(projectDir)
		cmd := exec.Command("git", "clone", repoURI, ".")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Start()
		cmd.Wait()
		fmt.Printf("\n%v\n", projectDir)
	default:
		stderr.Fatalf("The given command is unknown.\n%v", usage)
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
		stderr.Fatalln(err)
	}
	fullDir := filepath.Join(projectDir, subDir)
	return fullDir
}

// git@github.com:majgis/gittk.git
// https://github.com/majgis/gittk.git
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
