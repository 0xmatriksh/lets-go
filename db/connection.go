package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5431
)

var dbname string = os.Getenv("POSTGRES_DB")
var user string = os.Getenv("POSTGRES_USER")
var password string = os.Getenv("POSTGRES_PASSWORD")

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
