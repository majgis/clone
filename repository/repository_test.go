package repository

import "testing"
import "os/user"
import "path/filepath"

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

func TestGetDir(t *testing.T) {
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
