package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5431
	user     = "admin"
	password = "admin123"
	dbname   = "booksDB"
)

func DBConnect() *sql.DB {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	mydb, err := sql.Open("postgres", connInfo)
	if err != nil {
		fmt.Println(err)
	}
	err = mydb.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to db!")
	return mydb
}

func DBCloseConnection(mydb *sql.DB) {
	defer mydb.Close()
}
