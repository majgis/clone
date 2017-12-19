package repository

import "testing"

func TestHelloWorld(t *testing.T) {
	actual, err := GetSubDir("https://github.com/majgis/gittk.git")
	if err != nil {
		t.Fatalf("Error when getting directory: %v", err)
	}

	expected := "github.com/majgis/gittk"

	if actual != expected {
		t.Fatalf("%v does not equal %v", actual, expected)
	}
}
