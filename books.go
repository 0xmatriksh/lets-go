package main

type book struct {
	Id       int
	Title    string
	Author   string
	Quantity int
}

var books = []book{
	{Id: 1, Title: "A strange Loop", Author: "Alex Bradman", Quantity: 5},
	{Id: 2, Title: "Atomic Habits", Author: "James Clear", Quantity: 7},
	{Id: 3, Title: "Homo Sapiens", Author: "Yuval Noah Harari", Quantity: 2},
}
