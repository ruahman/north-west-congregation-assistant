package utils

import (
	"fmt"
	"testing"
)

func TestReadDir(t *testing.T) {
	files, err := ReadDir("../data/migrations")
	if err != nil {
		t.Error(err)
	}
	if len(files) == 0 {
		t.Error("No files found")
	}
	fmt.Println(files)
}
