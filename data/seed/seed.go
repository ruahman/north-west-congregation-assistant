package seed

import (
	"database/sql"
	"jw/data/database"
)

func Run(db *sql.DB) {
	_, _ = database.ExecFile(db, "./seed-table_create.sql")
}

func Drop(db *sql.DB) {
	_, _ = database.ExecFile(db, "./seed-table_drop.sql")
}
