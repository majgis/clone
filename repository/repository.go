package repository

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

// Clone the given gitURI into GITTK_PATH
func Clone(repoURI string) (string, error) {
	repoDir, repoDirErr := GetDir(repoURI)
	if repoDirErr != nil {
		return repoDir, repoDirErr
	}
	err := os.MkdirAll(repoDir, os.ModePerm)
	if err != nil {
		return repoDir, fmt.Errorf("unable to create directory: %v", repoDir)
	}
	os.Chdir(repoDir)
	cmd := exec.Command("git", "clone", repoURI, ".")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	startErr := cmd.Start()
	if startErr != nil {
		return repoDir, fmt.Errorf("unable to execute git clone")
	}
	cmd.Wait()
	return repoDir, nil
}

// GetDir takes a git uri and returns the corresponding local folder
func GetDir(uri string) (string, error) {
	repoDir := os.Getenv("GITTK_PATH")

	// Get the project directory
	if repoDir == "" {
		user, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("unable to determine the home directory")
		}
		repoDir = filepath.Join(user.HomeDir, "projects")
	}

	// Parse git URI to get subtree
	subDir, err := GetSubDir(uri)
	if err != nil {
		return "", fmt.Errorf("unable to load project directory, try setting GITTK_PATH")
	}
	fullDir := filepath.Join(repoDir, subDir)
	return fullDir, nil
}

// Return subdirectory from different git URI types
// git@github.com:majgis/gittk.git
// https://github.com/majgis/gittk.git
func GetSubDir(uri string) (string, error) {
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
