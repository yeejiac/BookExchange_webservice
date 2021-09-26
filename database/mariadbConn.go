package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yeejiac/BookExchange_webservice/internal"
	"github.com/yeejiac/BookExchange_webservice/models"
)

func MariaDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:Aa1234@tcp(127.0.0.1:3306)/?charset=utf8")
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("mariadb conn failed")
		os.Exit(1)
	}
	fmt.Println("mariadb conn success")
	return db
}

func InsertBookInfo(bookinfo models.BookInfo, mc *sql.DB) {
	sql := fmt.Sprintf("INSERT INTO `Bookstore`.`BookInfo` (`book_id`, `bookname`, `book_catelog`) VALUES ('%s', '%s', %d);",
		internal.GenerateGroupid(5), bookinfo.BookName, bookinfo.BookCatelog)
	fmt.Println(sql)
	_, err := mc.Exec(sql)

	if err != nil {
		fmt.Println("Insert error")
		fmt.Println(err)
	}
}

type MariaDB interface {
	MariaDBConnection()
	Insert(sql string)
}
