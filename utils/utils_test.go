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

func TestPrettyPrint(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 30}

	PrettyPrint(person)
}

func TestPrettyJSON(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 30}

	PrettyJSON(person)
}

func TestCheckAddDeleteFile(t *testing.T) {
	if CheckFile("./this-file-do-not-exist.txt") {
		t.Error("File does exist")
	}
	if !CheckFile("./utils_test.go") {
		t.Error("File does not exist")
	}
	if !CheckFilePattern("utils_test") {
		t.Error("File does not exist")
	}
	err := AddFile("./now-this-file-do-not-exist.txt")
	if err != nil {
		t.Error("AddFile failed")
	}
	if !CheckFile("./now-this-file-do-not-exist.txt") {
		t.Error("File does not exist")
	}
	err = DeleteFile("./now-this-file-do-not-exist.txt")
	if err != nil {
		t.Error("DeleteFile failed")
	}
	if CheckFile("./now-this-file-do-not-exist.txt") {
		t.Error("File does exist")
	}
}
