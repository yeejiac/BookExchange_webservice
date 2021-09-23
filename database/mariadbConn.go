package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var mariadb_conn *sql.DB

func MariaDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:Aa1234@tcp(127.0.0.1:3306)/user")
	mariadb_conn = db
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("mariadb conn success")
	return db
}

func Insert(sql string) {
	statement, err := mariadb_conn.Prepare(sql)

	if err != nil {
		fmt.Println(err)
	}
	defer statement.Close()
}

type MariaDB interface {
	MariaDBConnection()
	Insert(sql string)
}
