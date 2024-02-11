package data

import (
	"database/sql"
	"fmt"
	"jw/utils"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const MIGRATIONS_PATH = "data/database"

func connectDB(u string, p string, h string, d string) *sql.DB {
	c := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		u,
		p,
		h,
		d,
	)

	db, err := sql.Open("postgres", c)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

func create_database(db *sql.DB) {
	sql, err := utils.ReadFile(MIGRATIONS_PATH + "/database_create.sql")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func drop_database(db *sql.DB) {
	sql, err := utils.ReadFile(MIGRATIONS_PATH + "/database_drop.sql")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = db.Exec(sql)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func create_migrations(n string) {
	timeStamp := time.Now().Unix()

	u := fmt.Sprintf("%d--%s_up.sql", timeStamp, n)
	d := fmt.Sprintf("%d--%s_down.sql", timeStamp, n)

	fmt.Println("...Creating migration files", u, d)
	os.Create(MIGRATIONS_PATH + "/" + u)
	os.Create(MIGRATIONS_PATH + "/" + d)
}

func run_migrations() {
	fmt.Println("...run migrations")
}

func up_migrations() {
	fmt.Println("...run migrations up")
}

func down_migrations() {
	fmt.Println("...run migrations down")
}

func seed() {
	fmt.Println("...seed data")
}

func Database(args []string) {
	if len(args) > 0 {
		if args[0] == "migration" {
			if len(args) > 1 {
				if args[1] == "run" {
					run_migrations()
				} else if args[1] == "create" {
					create_migrations(args[2])
				} else if args[1] == "up" {
					up_migrations()
				} else if args[1] == "down" {
					down_migrations()
				} else {
					fmt.Println("migration command is not reconized.")
				}
			} else {
				fmt.Println("database migration needs more arguments")
			}
		} else if args[0] == "create" {
			db := connectDB(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"))
			defer db.Close()
			create_database(db)
		} else if args[0] == "drop" {
			db := connectDB(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"))
			defer db.Close()
			drop_database(db)
		} else if args[0] == "seed" {
			seed()
		} else {
			fmt.Println("command for database is not reconized.")
		}
	} else {
		fmt.Println("you need more arguments for database command.")
	}
}
