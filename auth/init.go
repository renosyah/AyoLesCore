package auth

import "database/sql"

var db *sql.DB

func Init(d *sql.DB) {
	db = d
}
