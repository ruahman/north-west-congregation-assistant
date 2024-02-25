package main

import (
	"fmt"
	app "framework"
	"log"
	"os"
	"strings"
	"utils"
)

func main() {
	err := utils.LoadEnv(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	fmt.Println("=== Running command:", strings.Join(os.Args[1:], " "), "===")

	command := os.Args[1]
	if command == "database" {
		// data.DatabaseExec(os.Args[2:])
	} else if command == "server" {
		app.Server()
	} else {
		fmt.Println("command not found")
	}
}
