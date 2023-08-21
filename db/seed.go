package db

import (
	"database/sql"
	"fmt"
	model "lets-go/models"

	_ "github.com/lib/pq"
)

var seed_books = []model.Book{
	{Id: "1", Title: "A strange Loop", Author: "Alex Bradman", Quantity: 5},
	{Id: "2", Title: "Atomic Habits", Author: "James Clear", Quantity: 7},
	{Id: "3", Title: "Homo Sapiens", Author: "Yuval Noah Harari", Quantity: 2},
}

func InitializeTable(mydb *sql.DB) {
	var exists bool
	if err := mydb.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE  schemaname = 'public' AND tablename = 'articles' );").Scan(&exists); err != nil {
		fmt.Println("EXISTS: ", err)
		return
	}
	if !exists {
		var createTableQuery string = "CREATE TABLE books (id VARCHAR(36) PRIMARY KEY, title VARCHAR(100) NOT NULL, author VARCHAR(100) NOT NULL, quantity INT NOT NULL);"
		_, err := mydb.Query(createTableQuery)
		// res.Close()
		if err != nil {
			fmt.Println("failed to execute query", err)
			return
		}
		fmt.Println("Table Created Successfully!!!")

		for _, book := range seed_books {
			insertBookQuery := `INSERT INTO books (id,title,author,quantity) VALUES ($1, $2, $3, $4);`

			err := mydb.QueryRow(insertBookQuery, book.Id, book.Title, book.Author, book.Quantity)
			if err != nil {
				fmt.Println("failed to execute query", err)
				return
			}
		}
		fmt.Println("Books Inserted Successfully!!!")
	} else {
		fmt.Println("Table 'articles' already exists ")
	}
}
