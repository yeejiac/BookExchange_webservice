package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var mariadb_conn *sql.DB

func MariaDBConnection() {
	db, err := sql.Open("mysql", "root:Aa1234@tcp(127.0.0.1:3306)/user")
	mariadb_conn = db
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}
