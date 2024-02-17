package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func PrettyPrint(x interface{}) {
	fmt.Printf("%+v\n", x)
}

func PrettyJSON(x interface{}) {
	prettyJSON, _ := json.MarshalIndent(x, "", "  ")
	fmt.Println(string(prettyJSON))
}

func ReadDir(p string) ([]string, error) {
	// absolutePath := filepath.Join(".")
	// fmt.Println("Reading directory:", absolutePath)

	d, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer d.Close()

	files, err := d.Readdir(0)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var r []string
	for _, f := range files {
		if !f.IsDir() {
			r = append(r, f.Name())
		}
	}

	return r, nil
}

func ReadFile(p string) (string, error) {
	content, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	r := string(content)
	r = strings.TrimSpace(r)
	return r, nil
}

func AddFile(p string) error {
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func CheckFile(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

func CheckFilePattern(p string) bool {
	files, err := ReadDir(".")
	if err != nil {
		return false
	}
	for _, f := range files {
		if strings.Contains(f, p) {
			return true
		}
	}
	return false
}

func DeleteFile(p string) error {
	err := os.Remove(p)
	if err != nil {
		return err
	}
	return nil
}

func LoadEnv(p string) error {
	fmt.Println("--- Loading dotenv file ---")
	f, err := os.Open(p)
	if err != nil {
		fmt.Println("No dotenv file found")
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	n := 1
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			n++
			continue
		}
		key, value := parts[0], parts[1]
		fmt.Println(key, "=", value)
		os.Setenv(key, value)
		n++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func Search[T comparable](a []T, x T) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}
