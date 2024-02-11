package sqlutils

import (
	"database/sql"
	"jw/utils"
	"log"
)

func ExecFile(db *sql.DB, p string) {
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
