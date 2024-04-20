package dbtools

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	driverName     string
	dataSourceName string
)

func DBInitializer(dn, dsn string) {
	driverName = dn
	dataSourceName = dsn
}

func connect() *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
