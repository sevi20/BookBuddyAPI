package models

import (
	"BOOKBUDDYAPI/database"
)

type Book struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Isbn   string `json:"isbn"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var bookCollection []Book

func (b *Book) Save() error {
	statement, err := database.GetDb().Prepare(`
		INSERT INTO 
		books
		    (title, isbn, author, year)
		VALUES
		    (?, ?, ?, ?)
	`)
	defer statement.Close()

	if err != nil {
		return err
	}

	result, err := statement.Exec(b.Title, b.Isbn, b.Author, b.Year)
	if err != nil {
		return err
	}

	bookId, err := result.LastInsertId()
	b.Id = bookId

	return err
}

func GetAllBooks() ([]Book, error) {
	dbCursor, err := database.GetDb().Query(`SELECT * FROM books`)
	if err != nil {
		return nil, err
	}
	defer dbCursor.Close()

	bookCollection = []Book{}
	for dbCursor.Next() {
		var bookObject Book
		err := dbCursor.Scan(
			&bookObject.Id,
			&bookObject.Title,
			&bookObject.Isbn,
			&bookObject.Author,
			&bookObject.Year,
		)
		if err != nil {
			return nil, err
		}

		bookCollection = append(bookCollection, bookObject)
	}

	return bookCollection, nil
}

func GetBookByID(id int) (Book, error) {
	var book Book
	row := database.GetDb().QueryRow(`SELECT * FROM books WHERE id = ?`, id)
	err := row.Scan(&book.Id, &book.Title, &book.Isbn, &book.Author, &book.Year)
	if err != nil {
		return book, err
	}
	return book, nil
}

func UpdateBookByID(id int, updatedBook *Book) error {
	statement, err := database.GetDb().Prepare(`
        UPDATE books SET title = ?, isbn = ?, author = ?, year = ? WHERE id = ?
    `)
	defer statement.Close()

	if err != nil {
		return err
	}

	_, err = statement.Exec(updatedBook.Title, updatedBook.Isbn, updatedBook.Author, updatedBook.Year, id)
	return err
}

func DeleteBookByID(id int) error {
	statement, err := database.GetDb().Prepare(`DELETE FROM books WHERE id = ?`)
	defer statement.Close()

	if err != nil {
		return err
	}

	_, err = statement.Exec(id)
	return err
}
