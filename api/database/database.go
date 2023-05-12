package database

import (
	"api/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// OpenConnection open the connection to the database.
func main() {
	db, err := sql.Open("mysql", config.GetConfig().DbConn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to MySQL database!")
}
