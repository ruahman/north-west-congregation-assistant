package data

import (
	"data/database"
	"data/migrations"
	"data/seed"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func DatabaseExec(args []string) {
	if len(args) > 0 {
		if args[0] == "migration" {
			if len(args) > 1 {
				if args[1] == "run" {
					db, err := database.Connect(
						os.Getenv("POSTGRES_USER"),
						os.Getenv("POSTGRES_PASSWORD"),
						os.Getenv("POSTGRES_HOST"),
						os.Getenv("DATABASE"),
					)
					if err != nil {
						log.Fatal(err)
					}
					defer db.Close()
					migrations.Run(db)
				} else if args[1] == "drop" {
					db, err := database.Connect(
						os.Getenv("POSTGRES_USER"),
						os.Getenv("POSTGRES_PASSWORD"),
						os.Getenv("POSTGRES_HOST"),
						os.Getenv("DATABASE"),
					)
					if err != nil {
						log.Fatal(err)
					}
					defer db.Close()
					migrations.Drop(db)
				} else if args[1] == "add" {
					migrations.Add(args[2])
				} else if args[1] == "up" {
					migrations.Up()
				} else if args[1] == "down" {
					migrations.Down()
				} else {
					fmt.Println("migration command", strings.Join(args, " "), "is not reconized.")
				}
			} else {
				fmt.Println("database migration needs more arguments")
			}
		} else if args[0] == "create" {
			db, err := database.Connect(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"))
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()
			database.CreateDB(db, os.Getenv("DATABASE"))
		} else if args[0] == "drop" {
			db, err := database.Connect(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"))
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()
			database.DropDB(db)
		} else if args[0] == "seed" {
			if args[1] == "run" {
				db, err := database.Connect(
					os.Getenv("POSTGRES_USER"),
					os.Getenv("POSTGRES_PASSWORD"),
					os.Getenv("POSTGRES_HOST"),
					os.Getenv("DATABASE"),
				)
				if err != nil {
					log.Fatal(err)
				}
				seed.Run(db)
			} else if args[1] == "drop" {
				db, err := database.Connect(
					os.Getenv("POSTGRES_USER"),
					os.Getenv("POSTGRES_PASSWORD"),
					os.Getenv("POSTGRES_HOST"),
					os.Getenv("DATABASE"),
				)
				if err != nil {
					log.Fatal(err)
				}
				seed.Drop(db)
			}
		} else {
			fmt.Println("command for database is not reconized.")
		}
	} else {
		fmt.Println("you need more arguments for database command.")
	}
}
