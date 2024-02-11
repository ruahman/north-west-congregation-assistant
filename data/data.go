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

const (
	DATABASE_PATH   = "data/database"
	MIGRATIONS_PATH = "data/migrations"
	SEED_PATH       = "data/seed"
)

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

func exec_file(db *sql.DB, p string) {
	sql, err := utils.ReadFile(p)
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

func create_database(db *sql.DB) {
	exec_file(db, DATABASE_PATH+"/database_create.sql")
}

func drop_database(db *sql.DB) {
	exec_file(db, DATABASE_PATH+"/database_drop.sql")
}

func create_migrations(n string) {
	timeStamp := time.Now().Unix()

	u := fmt.Sprintf("%d--%s_up.sql", timeStamp, n)
	d := fmt.Sprintf("%d--%s_down.sql", timeStamp, n)

	fmt.Println("...Creating migration files", u, d)
	os.Create(MIGRATIONS_PATH + "/" + u)
	os.Create(MIGRATIONS_PATH + "/" + d)
}

func run_migrations(db *sql.DB) {
	fmt.Println("...run migrations")
	exec_file(db, MIGRATIONS_PATH+"/migrations-table_create.sql")
}

func drop_migrations(db *sql.DB) {
	fmt.Println("...drop migrations")
	exec_file(db, MIGRATIONS_PATH+"/migrations-table_drop.sql")
}

func up_migrations() {
	fmt.Println("...run migrations up")
}

func down_migrations() {
	fmt.Println("...run migrations down")
}

func run_seed(db *sql.DB) {
	exec_file(db, SEED_PATH+"/seed-table_create.sql")
}

func drop_seed(db *sql.DB) {
	exec_file(db, SEED_PATH+"/seed-table_drop.sql")
}

func Database(args []string) {
	if len(args) > 0 {
		if args[0] == "migration" {
			if len(args) > 1 {
				if args[1] == "run" {
					db := connectDB(
						os.Getenv("POSTGRES_USER"),
						os.Getenv("POSTGRES_PASSWORD"),
						os.Getenv("POSTGRES_HOST"),
						os.Getenv("DATABASE"),
					)
					run_migrations(db)
				} else if args[1] == "drop" {
					db := connectDB(
						os.Getenv("POSTGRES_USER"),
						os.Getenv("POSTGRES_PASSWORD"),
						os.Getenv("POSTGRES_HOST"),
						os.Getenv("DATABASE"),
					)
					drop_migrations(db)
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
			if args[1] == "run" {
				db := connectDB(
					os.Getenv("POSTGRES_USER"),
					os.Getenv("POSTGRES_PASSWORD"),
					os.Getenv("POSTGRES_HOST"),
					os.Getenv("DATABASE"),
				)
				run_seed(db)
			} else if args[1] == "drop" {
				db := connectDB(
					os.Getenv("POSTGRES_USER"),
					os.Getenv("POSTGRES_PASSWORD"),
					os.Getenv("POSTGRES_HOST"),
					os.Getenv("DATABASE"),
				)
				drop_seed(db)
			}
		} else {
			fmt.Println("command for database is not reconized.")
		}
	} else {
		fmt.Println("you need more arguments for database command.")
	}
}
