package models

import (
	"database/sql"
	"golang/entities"
)

type BookModel struct {
	Db *sql.DB
}

func (bookModel BookModel) Find(Book_id int) (entities.Book, error) {
	rows, err := bookModel.Db.Query("select * from books where book_id = ?", Book_id)
	if err != nil {
		return entities.Book{}, err
	} else {
		book := entities.Book{}
		for rows.Next() {
			var Book_id int64
			var Book_title string
			var Book_publisher string
			var Book_author string
			var Book_language string
			err2 := rows.Scan(&Book_id, &Book_title, &Book_publisher, &Book_author, &Book_language)
			if err2 != nil {
				return entities.Book{}, err2
			} else {
				book = entities.Book{Book_id, Book_title, Book_publisher, Book_author, Book_language}
			}

		}
		return book, nil
	}
}
