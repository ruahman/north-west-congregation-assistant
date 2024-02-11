package main

import (
	"fmt"
	"jw/data"
	"jw/service"
	"jw/utils"
	"log"
	"os"
	"strings"
)

func main() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	fmt.Println("=== Running command:", strings.Join(os.Args[1:], " "), "===")

	command := os.Args[1]
	if command == "database" {
		data.Database(os.Args[2:])
	} else if command == "server" {
		service.Server()
	} else {
		fmt.Println("command not found")
	}
}
