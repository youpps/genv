package genv

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Clearenv()

	err := Load()
	if err != nil {
		t.Fatal(err)
	}

	if os.Getenv("KEY") != "1234567890" {
		t.Fatal("No valid os.Getenv result.")
	}

}

func TestLoadCustomFile(t *testing.T) {
	os.Clearenv()

	err := Load(".env-custom")
	if err != nil {
		t.Fatal(err)
	}

	if os.Getenv("KEY") != "0987654321" {
		t.Fatal("No valid os.Getenv result.")
	}
}

func TestGetFilenames(t *testing.T) {
	if filenames := getFileNames([]string{}); filenames[0] != ".env" {
		t.Fatal("Zero filenames lenght.")
	}
}
