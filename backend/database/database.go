package database

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//DB Exported for access to mysql
type DB struct {
	*sql.DB
}

//var Database *sql.DB
/*
func NewOpen(dt string, c string) (DB, error) {
	db, err := sql.Open(dt, c)
	return DB{db}, err
}
*/

// NewOpen exported for creating a new MySQL instance
func NewOpen() (DB, error) {
	db, err := sql.Open("mysql", "root:kemar@tcp(localhost:3306)/reactheaddb")
	return DB{db}, err
}
