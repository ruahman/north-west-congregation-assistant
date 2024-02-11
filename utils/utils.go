package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(p string) (string, error) {
	content, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}
	r := string(content)
	r = strings.TrimSpace(r)
	return r, nil
}

func LoadEnv() error {
	fmt.Println("--- Loading .env file ---")
	f, err := os.Open(".env")
	defer f.Close()
	if err != nil {
		fmt.Println("No .env file found")
		return err
	}

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
