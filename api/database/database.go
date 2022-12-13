package database

import (
	"api/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// OpenConnection open the connection to the database.
func OpenConnection() (*sql.DB, error) {

	db, err := sql.Open("mysql", config.ConnectionStringDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
