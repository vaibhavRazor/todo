package database

import (
	"database/sql"

	"todo/lib"

	"github.com/go-gorp/gorp"
)

func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:Root@123@tcp(host.docker.internal:3306)/test")
	lib.CheckErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	lib.CheckErr(err, "Create tables failed")
	return dbmap
}
