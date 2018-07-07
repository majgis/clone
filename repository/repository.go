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
		return "", fmt.Errorf("unable to load project directory, only GitHub and Bitbucket Server URI's are supported")
	}
	fullDir := filepath.Join(repoDir, subDir)
	return fullDir, nil
}

// Return subdirectory from different git URI types
//
// github:
//   https://github.com/majgis/gittk.git
//   https://user@github.com/majgis/gittk.git
//   git@github.com:majgis/gittk.git
//
//
// Bitbucket Server:
//   https://xxx.somewhere.com/scm/teamid/appname.git
//   https://user@xxx.somewhere.com/scm/teamid/appname.git
//   ssh://git@xxx.somewhere.com:1111/teamid/appname.git
func GetSubDir(uri string) (string, error) {
	uriSplit := strings.Split(uri, "/")

	// GitHub SSH
	if strings.HasPrefix(uri, "git@github.com") {
		userName := strings.Split(uriSplit[0], ":")[1]
		projectName := strings.Split(uriSplit[1], ".")[0]
		result := filepath.Join("github.com", userName, projectName)
		return result, nil
	}

	// GitHub HTTPS
	if strings.HasPrefix(uri, "https://github.com") {
		userName := uriSplit[3]
		projectName := strings.Split(uriSplit[4], ".")[0]
		result := filepath.Join("github.com", userName, projectName)
		return result, nil
	}

	// GitHub HTTPS with user
	if strings.Contains(uri, "@github.com/") {
		userName := uriSplit[3]
		projectName := strings.Split(uriSplit[4], ".")[0]
		result := filepath.Join("github.com", userName, projectName)
		return result, nil
	}

	// Bitbucket Server HTTPS
	if strings.HasPrefix(uri, "https://") && !strings.Contains(uriSplit[2], "@") {
		userName := uriSplit[4]
		projectName := strings.Split(uriSplit[5], ".")[0]
		result := filepath.Join(uriSplit[2], userName, projectName)
		return result, nil
	}

	// Bitbucket Server HTTPS with user
	if strings.HasPrefix(uri, "https://") && strings.Contains(uriSplit[2], "@") {
		userName := uriSplit[4]
		projectName := strings.Split(uriSplit[5], ".")[0]
		domain := strings.Split(uriSplit[2], "@")[1]
		result := filepath.Join(domain, userName, projectName)
		return result, nil
	}

	// Bitbucket Server SSH
	if strings.HasPrefix(uri, "ssh://") {
		userName := uriSplit[3]
		projectName := strings.Split(uriSplit[4], ".")[0]
		result := filepath.Join(strings.Split(strings.Split(uriSplit[2], ":")[0], "@")[1], userName, projectName)
		return result, nil
	}

	// Unknown type
	return "", errors.New("unknown URI type")
}
