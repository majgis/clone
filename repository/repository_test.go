package repository

import (
	"fmt"
	"os/user"
	"path/filepath"
	"testing"
)

func TestGetSubDir(t *testing.T) {
	actual, err := GetSubDir("https://github.com/majgis/gittk.git")
	if err != nil {
		t.Fatalf("Error when getting directory: %v", err)
	}

	expected := filepath.Join("github.com", "majgis", "gittk")

	if actual != expected {
		t.Fatalf("'%v' does not equal '%v'", actual, expected)
	}
}

func TestGetDirGithubHTTPS(t *testing.T) {
	actual, err := GetDir("https://github.com/majgis/gittk.git")
	if err != nil {
		t.Fatalf("Error when getting directory: %v", err)
	}

	user, usrErr := user.Current()
	if usrErr != nil {
		t.Fatalf("Error retrieving user: %v", usrErr)
	}

	expected := filepath.Join(user.HomeDir, "projects", "github.com", "majgis", "gittk")
	if actual != expected {
		t.Fatalf("'%v' does not equal '%v'", actual, expected)
	}
}

func TestGetDirGithubHTTPSWithUser(t *testing.T) {
	actual, err := GetDir("https://github.com/majgis/gittk.git")
	if err != nil {
		t.Fatalf("Error when getting directory: %v", err)
	}

	user, usrErr := user.Current()
	if usrErr != nil {
		t.Fatalf("Error retrieving user: %v", usrErr)
	}

	expected := filepath.Join(user.HomeDir, "projects", "github.com", "majgis", "gittk")
	if actual != expected {
		t.Fatalf("'%v' does not equal '%v'", actual, expected)
	}
}

func TestGetDirGithubSSH(t *testing.T) {
	actual, err := GetDir("git@github.com:majgis/gittk.git")
	if err != nil {
		t.Fatalf("Error when getting directory: %v", err)
	}

	user, usrErr := user.Current()
	if usrErr != nil {
		t.Fatalf("Error retrieving user: %v", usrErr)
	}

	expected := filepath.Join(user.HomeDir, "projects", "github.com", "majgis", "gittk")
	if actual != expected {
		t.Fatalf("'%v' does not equal '%v'", actual, expected)
	}
}

func TestGetDirBitbucketServerHTTPS(t *testing.T) {
	actual, err := GetDir("https://git.somewhere.com/scm/teamid/appname.git")
	if err != nil {
		t.Fatalf("Error when getting directory: %v", err)
	}

	user, usrErr := user.Current()
	if usrErr != nil {
		t.Fatalf("Error retrieving user: %v", usrErr)
	}

	expected := filepath.Join(user.HomeDir, "projects", "git.somewhere.com", "teamid", "appname")
	if actual != expected {
		t.Fatalf("'%v' does not equal '%v'", actual, expected)
	}
}

func TestGetDirBitbucketServerHTTPSWithUser(t *testing.T) {
	actual, err := GetDir("https://user@git.somewhere.com/scm/teamid/appname.git")
	if err != nil {
		t.Fatalf("Error when getting directory: %v", err)
	}

	user, usrErr := user.Current()
	if usrErr != nil {
		t.Fatalf("Error retrieving user: %v", usrErr)
	}

	expected := filepath.Join(user.HomeDir, "projects", "git.somewhere.com", "teamid", "appname")
	if actual != expected {
		t.Fatalf("'%v' does not equal '%v'", actual, expected)
	}
}

func TestGetDirBitbucketServerSSH(t *testing.T) {
	actual, err := GetDir("ssh://git@git.somewhere.com:1111/teamid/appname.git")
	if err != nil {
		t.Fatalf("Error when getting directory: %v", err)
	}

	user, usrErr := user.Current()
	if usrErr != nil {
		t.Fatalf("Error retrieving user: %v", usrErr)
	}

	expected := filepath.Join(user.HomeDir, "projects", "git.somewhere.com", "teamid", "appname")
	if actual != expected {
		t.Fatalf("'%v' does not equal '%v'", actual, expected)
	}
}

func TestCloneWithBashOnly(t *testing.T) {
	actual, err := Clone("https://github.com/majgis/gittk.git", true)
	if err != nil {
		t.Fatalf("Error when getting directory: %v", err)
	}

	user, usrErr := user.Current()
	if usrErr != nil {
		t.Fatalf("Error retrieving user: %v", usrErr)
	}

	path := filepath.Join(user.HomeDir, "projects", "github.com", "majgis", "gittk")
	expected := fmt.Sprintf(`mkdir -p %v \
	&& cd %v \
	&& git clone https://github.com/majgis/gittk.git %v`,
		path, path, path)
	if actual != expected {
		t.Fatalf("'%v' does not equal '%v'", actual, expected)
	}
}
