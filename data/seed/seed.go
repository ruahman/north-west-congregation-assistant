package seed

import (
	"database/sql"
	"jw/data/sqlutils"
)

func Run(db *sql.DB) {
	sqlutils.ExecFile(db, "./seed-table_create.sql")
}

func Drop(db *sql.DB) {
	sqlutils.ExecFile(db, "./seed-table_drop.sql")
}
