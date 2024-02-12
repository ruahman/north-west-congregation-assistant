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
