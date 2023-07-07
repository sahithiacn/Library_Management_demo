package main

import (
	"fmt"
	"golang/config"
	"models"
	"os"
	"strconv"
)

func main() {
	argsWithProg := os.Args[1]
	argsWithout := os.Args[2]
	Demo1_CallFind(argsWithProg, argsWithout)

}
func Demo1_CallFind(ArgsWithProg, ArgsWithout string) {
	db, err := config.GetMySQLDB()
	id, err := strconv.Atoi(ArgsWithProg)
	fmt.Println("name:", ArgsWithout)
	if err != nil {
		fmt.Println(err)
	} else {
		bookModel := models.BookModel{
			Db: db,
		}

		book, err := bookModel.Find(id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(book)
			fmt.Println("Book info")
			fmt.Println("Id:", book.Book_id)
			fmt.Println("Name:", book.Book_title)
			fmt.Println("Publisher:", book.Book_publisher)
			fmt.Println("Author:", book.Book_author)
			fmt.Println("Author:", book.Book_language)
		}

	}
}
